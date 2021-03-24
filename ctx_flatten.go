package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"go/types"
	"log"
	"reflect"
)

type flattenSlot struct {
	f      *File
	define Slot
	assign Slot
}

var zeroValues = map[types.BasicKind]*Statement{
	types.Bool:    Lit(false),
	types.String:  Lit(""),
	types.Int:     Lit(0),
	types.Float64: Lit(0.0),
}

var typeCovertFunc = map[types.BasicKind]*Statement{
	types.Bool:    Bool(),
	types.String:  String(),
	types.Int:     Int(),
	types.Float64: Float64(),
}

func (ctx *Ctx) flattenType(t types.Type, varHint *string, ref bool, input *Statement, slot flattenSlot) {
	switch t := t.(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal("Can't flatten basic type")
		}
		ctx.flattenBasic(t, varHint, ref, input, slot)
	case *types.Pointer:
		ctx.flattenType(t.Elem(), varHint, true, input, slot)
	case *types.Slice:
		ctx.flattenSlice(t, varHint, ref, input, slot)
	case *types.Map:
		ctx.flattenMap(t, varHint, ref, input, slot)
	case *types.Named:
		ctx.flattenNamedType(t, varHint, ref, input, slot)
	default:
		// Ignore: Array, Chan, Tuple, Signature, Struct, Interface
		return
	}
}

func (ctx *Ctx) flattenNamedType(t *types.Named, varHint *string, ref bool, input *Statement, slot flattenSlot) {
	switch t.Underlying().(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal("Can't flatten basic type")
		}
		ctx.flattenNamedBasic(t, varHint, ref, input, slot)
	case *types.Pointer:
		ctx.flattenNamedPointer(t, varHint, ref, input, slot)
	case *types.Slice:
		ctx.flattenNamedSlice(t, varHint, ref, input, slot)
	case *types.Map:
		ctx.flattenNamedMap(t, varHint, ref, input, slot)
	case *types.Struct:
		ctx.flattenNamedStruct(t, varHint, ref, input, slot)
	case *types.Interface:
		panic("TODO")
	default:
		// Array, Chan, Tuple, Signature
		return
	}
}

func (ctx *Ctx) flattenBasic(t *types.Basic, hint *string, ref bool, input *Statement, slot flattenSlot) {
	if !ref {
		slot.assign.Add(input)
		return
	}

	localVar := _idLocalVar
	if hint != nil {
		localVar = *hint
	}

	slot.assign.Add(Id(localVar))
	slot.define.Add(
		Id(localVar).Op(":=").Add(zeroValues[t.Kind()].Clone()),
	)
	slot.define.Add(
		If(input.Clone().Op("!=").Nil()).Block(
			Id(localVar).Op("=").Op("*").Add(input),
		),
	)
}

func (ctx *Ctx) flattenNamedBasic(t *types.Named, hint *string, ref bool, input *Statement, slot flattenSlot) {
	ut := t.Underlying().(*types.Basic)

	if !ref {
		slot.assign.Add(typeCovertFunc[ut.Kind()].Clone().Call(input))
		return
	}

	localVar := _idLocalVar
	if hint != nil {
		localVar = *hint
	}

	slot.assign.Add(Id(localVar))
	slot.define.Add(
		Id(localVar).Op(":=").Add(zeroValues[ut.Kind()].Clone()),
	)
	slot.define.Add(
		If(input.Clone().Op("!=").Nil()).Block(
			Id(localVar).Op("=").Add(typeCovertFunc[ut.Kind()].Clone().Call(Op("*").Add(input))),
		),
	)
}

func (ctx *Ctx) flattenNamedPointer(t *types.Named, hint *string, ref bool, input *Statement, slot flattenSlot) {
	if input == nil {
		ctx.flattenType(t.Underlying().(*types.Pointer).Elem(), hint, true, nil, slot)
		return
	}

	assignSlot := &Statement{}
	newSlot := flattenSlot{
		f:      slot.f,
		define: slot.define,
		assign: assignSlot,
	}
	ctx.flattenType(t.Underlying().(*types.Pointer).Elem(), hint, true, input, newSlot)
	slot.assign.Add(assignSlot)
}

func (ctx *Ctx) flattenSlice(t *types.Slice, hint *string, ref bool, input *Statement, slot flattenSlot) {
	etName, et, isPtr := elemType(t.Elem())
	if isPtr {
		etName += "Ptr"
	}
	if isPtr {
		et = Op("*").Add(et)
	}

	flattenFuncName := fmt.Sprintf("flatten%sSlice", strcase.ToCamel(etName))
	if ref {
		flattenFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(flattenFuncName).Call(input))
	}

	// Create a flatten function for the given type
	if ctx.existFuncs[flattenFuncName] {
		return
	}
	ctx.existFuncs[flattenFuncName] = true

	slot.f.Func().Id(flattenFuncName).Params(
		Id(_idInput).Do(func(stmt *Statement) {
			if ref {
				stmt.Op("*")
			}
		}).Add(Index().Add(et)),
	).Add(Index().Interface()).

		// Function block
		BlockFunc(func(g *Group) {
			if ref {
				// Nil check on the input, e.g.:
				//
				// if input == nil {
				//     return []interface{}{}
				// }
				g.If(Id(_idInput).Op("==").Nil()).BlockFunc(func(g *Group) {
					g.Return(Index().Interface().Values())
				})
			}

			// Initialize the output array, e.g.
			//
			// output := make([]interface{}, 0)
			g.Id(_idOutput).Op(":=").Make(Index().Interface(), Lit(0))

			localVar := _idLocalVar

			g.For(List(Id("_"), Id(_idSliceElem)).Op(":=").Range().Do(func(stmt *Statement) {
				if ref {
					stmt.Op("*")
				}
			}).Id(_idInput)).BlockFunc(func(ig *Group) {
				// Prepare the slots
				assignSlot := &Statement{}
				newSlot := flattenSlot{
					f:      slot.f,
					assign: assignSlot,
					define: ig,
				}
				ctx.flattenType(t.Elem(), &localVar, false, Id(_idSliceElem), newSlot)

				ig.Add(Id(_idOutput).Op("=").Append(Id(_idOutput), assignSlot))
			})

			g.Return(Id(_idOutput))
		})
}

func (ctx *Ctx) flattenNamedSlice(t *types.Named, hint *string, ref bool, input *Statement, slot flattenSlot) {
	flattenFuncName := fmt.Sprintf("flatten%sSlice", strcase.ToCamel(t.Obj().Name()))
	if ref {
		flattenFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(flattenFuncName).Call(input))
	}

	// Create a flatten function for the given type
	if ctx.existFuncs[flattenFuncName] {
		return
	}
	ctx.existFuncs[flattenFuncName] = true

	slot.f.Func().Id(flattenFuncName).Params(
		Id(_idInput).Do(func(stmt *Statement) {
			if ref {
				stmt.Op("*")
			}
		}).Add(qualifiedNamedType(t)),
	).Add(Index().Interface()).

		// Function block
		BlockFunc(func(g *Group) {
			if ref {
				// Nil check on the input, e.g.:
				//
				// if input == nil {
				//     return []interface{}{}
				// }
				g.If(Id(_idInput).Op("==").Nil()).BlockFunc(func(g *Group) {
					g.Return(Index().Interface().Values())
				})
			}

			// Initialize the output array, e.g.
			//
			// output := make([]interface{}, 0)
			g.Id(_idOutput).Op(":=").Make(Index().Interface(), Lit(0))

			localVar := _idLocalVar

			g.For(List(Id("_"), Id(_idSliceElem)).Op(":=").Range().Do(func(stmt *Statement) {
				if ref {
					stmt.Op("*")
				}
			}).Id(_idInput)).BlockFunc(func(ig *Group) {
				// Prepare the slots
				assignSlot := &Statement{}
				newSlot := flattenSlot{
					f:      slot.f,
					assign: assignSlot,
					define: ig,
				}
				ctx.flattenType(t.Underlying().(*types.Slice).Elem(), &localVar, false, Id(_idSliceElem), newSlot)

				ig.Add(Id(_idOutput).Op("=").Append(Id(_idOutput), assignSlot))
			})

			g.Return(Id(_idOutput))
		})
}

func (ctx *Ctx) flattenMap(t *types.Map, hint *string, ref bool, input *Statement, slot flattenSlot) {
	// Type guard, Terraform only support map[string]interface{}, it is non-trivial to flatten to a non-string keyed map from Terraform.
	kt, ok := t.Key().(*types.Basic)
	if !ok || kt.Kind() != types.String {
		log.Fatalf("Only support flatten Map with String key (%v has key type %v)", t, kt)
	}

	// Construct the flatten function name.
	etName, et, isPtr := elemType(t.Elem())
	if isPtr {
		etName += "Ptr"
	}
	if isPtr {
		et = Op("*").Add(et)
	}

	flattenFuncName := fmt.Sprintf("flatten%sMap", strcase.ToCamel(etName))
	if ref {
		flattenFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(flattenFuncName).Call(input))
	}

	// Create a flatten function for the given type
	if ctx.existFuncs[flattenFuncName] {
		return
	}
	ctx.existFuncs[flattenFuncName] = true

	slot.f.Func().Id(flattenFuncName).Params(
		Id(_idInput).Do(func(stmt *Statement) {
			if ref {
				stmt.Op("*")
			}
		}).Map(String()).Add(et),
	).Map(String()).Interface().

		// Function block
		BlockFunc(func(g *Group) {
			if ref {
				// Nil check, e.g.
				// if input == nil {
				//   return map[string]interface{}{}
				// }
				g.If(Id(_idInput)).Op("==").Nil().Block(
					Return(Map(String()).Interface().Values()),
				)
			}

			// Initialize the output map, e.g.
			//
			// output := make(map[string]interface{})
			g.Id(_idOutput).Op(":=").Make(Map(String()).Interface())

			g.For(List(Id("k"), Id("v")).Op(":=").Range().Do(func(stmt *Statement) {
				if ref {
					stmt.Op("*")
				}
			}).Id(_idInput)).BlockFunc(func(ig *Group) {
				assignSlot := &Statement{}
				newSlot := flattenSlot{
					f:      slot.f,
					assign: assignSlot,
					define: ig,
				}
				localVar := _idLocalVar
				ctx.flattenType(t.Elem(), &localVar, false, Id("v"), newSlot)

				ig.Id(_idOutput).Index(Id("k")).Op("=").Add(assignSlot)
			})

			g.Return(Id(_idOutput))
		})
}

func (ctx *Ctx) flattenNamedMap(t *types.Named, hint *string, ref bool, input *Statement, slot flattenSlot) {
	// Type guard, Terraform only support map[string]interface{}, it is non-trivial to flatten to a non-string keyed map from Terraform.
	ut := t.Underlying().(*types.Map)
	kt, ok := ut.Key().(*types.Basic)
	if !ok || kt.Kind() != types.String {
		log.Fatalf("Only support flatten Map with String key (%v has key type %v)", t, kt)
	}

	// Construct the flatten function name.
	flattenFuncName := fmt.Sprintf("flatten%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		flattenFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(flattenFuncName).Call(input))
	}

	// Create a flatten function for the given type
	if ctx.existFuncs[flattenFuncName] {
		return
	}
	ctx.existFuncs[flattenFuncName] = true

	slot.f.Func().Id(flattenFuncName).Params(
		Id(_idInput).Add(qualifiedNamedType(t)),
	).Map(String()).Interface().

		// Function block
		BlockFunc(func(g *Group) {
			if ref {
				// Nil check, e.g.
				// if input == nil {
				//   return map[string]interface{}{}
				// }
				g.If(Id(_idInput)).Op("==").Nil().Block(
					Return(Map(String()).Interface().Values()),
				)
			}

			// Initialize the output map, e.g.
			//
			// output := make(map[string]interface{})
			g.Id(_idOutput).Op(":=").Make(Map(String()).Interface())

			g.For(List(Id("k"), Id("v")).Op(":=").Range().Do(func(stmt *Statement) {
				if ref {
					stmt.Op("*")
				}
			}).Id(_idInput)).BlockFunc(func(ig *Group) {
				assignSlot := &Statement{}
				newSlot := flattenSlot{
					f:      slot.f,
					assign: assignSlot,
					define: ig,
				}
				localVar := _idLocalVar
				ctx.flattenType(ut.Elem(), &localVar, false, Id("v"), newSlot)

				ig.Id(_idOutput).Index(Id("k")).Op("=").Add(assignSlot)
			})

			g.Return(Id(_idOutput))
		})
}
func (ctx *Ctx) flattenNamedStruct(t *types.Named, hint *string, ref bool, input *Statement, slot flattenSlot) {
	flattenFuncName := fmt.Sprintf("flatten%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		flattenFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(flattenFuncName).Call(input))
	}

	// Create a flatten function for the given type
	if ctx.existFuncs[flattenFuncName] {
		return
	}
	ctx.existFuncs[flattenFuncName] = true

	slot.f.Func().Id(flattenFuncName).Params(
		Id(_idInput).Do(func(stmt *Statement) {
			if ref {
				stmt.Op("*")
			}
		}).Add(qualifiedNamedType(t)),
	).Index().Interface().

		// Function block
		BlockFunc(func(g *Group) {
			if ref {
				// Nil check on the input, e.g.:
				//
				// if input == nil {
				//     return []interface{}{}
				// }
				g.If(Id(_idInput).Op("==").Nil()).BlockFunc(func(g *Group) {
					g.Return(Index().Interface().Values())
				})
			}

			// Loop over the struct fields and get their "slots"
			type slotCtx struct {
				field    *types.Var
				slot     flattenSlot
				input    *Statement
				localVar string
			}

			var slotCtxList []slotCtx

			var assignSlots []*Statement
			ut := t.Underlying().(*types.Struct)
			for i := 0; i < ut.NumFields(); i++ {
				v := ut.Field(i)
				if !v.Exported() {
					continue
				}

				if ctx.honorJSONIgnore {
					if reflect.StructTag(ut.Tag(i)).Get("json") == "-" {
						continue
					}
				}

				defineSlot := g
				assignSlot := &Statement{}
				assignSlots = append(assignSlots, assignSlot)
				sctx := slotCtx{
					field: v,
					slot: flattenSlot{
						f:      slot.f,
						define: defineSlot,
						assign: assignSlot,
					},
					input:    Id(_idInput).Dot(v.Name()),
					localVar: strcase.ToLowerCamel(v.Name()),
				}
				slotCtxList = append(slotCtxList, sctx)
			}

			for _, sctx := range slotCtxList {
				ctx.flattenType(sctx.field.Type(), &sctx.localVar, false, sctx.input, sctx.slot)
			}

			g.Return(Index().Interface().Values(
				Map(String()).Interface().Values(DictFunc(func(d Dict) {
					for idx, sctx := range slotCtxList {
						d[Lit(strcase.ToSnake(sctx.field.Name()))] = assignSlots[idx]
					}
				})),
			))
		})
}

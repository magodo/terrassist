package main

import (
	"fmt"
	"go/types"
	"log"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type expandSlot struct {
	f      *File
	define Slot
	assign Slot
}

func (ctx *Ctx) expandType(t types.Type, varHint *string, ref bool, input *Statement, slot expandSlot) {
	switch t := t.(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal("Can't expand basic type")
		}
		ctx.expandBasic(t, varHint, ref, input, slot)
	case *types.Pointer:
		ctx.expandType(t.Elem(), varHint, true, input, slot)
	case *types.Slice:
		ctx.expandSlice(t, varHint, ref, input, slot)
	case *types.Map:
		ctx.expandMap(t, varHint, ref, input, slot)
	case *types.Named:
		ctx.expandNamedType(t, varHint, ref, input, slot)
	default:
		// Ignore: Array, Chan, Tuple, Signature, Struct, Interface
		return
	}
}

func (ctx *Ctx) expandNamedType(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
	switch t.Underlying().(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal("Can't expand basic type")
		}
		ctx.expandNamedBasic(t, varHint, ref, input, slot)
	case *types.Pointer:
		ctx.expandNamedPointer(t, varHint, ref, input, slot)
	case *types.Slice:
		ctx.expandNamedSlice(t, varHint, ref, input, slot)
	case *types.Map:
		ctx.expandNamedMap(t, varHint, ref, input, slot)
	case *types.Struct:
		ctx.expandNamedStruct(t, varHint, ref, input, slot)
	case *types.Interface:
		panic("TODO")
	default:
		// Array, Chan, Tuple, Signature
		return
	}
}

func (ctx *Ctx) expandBasic(t *types.Basic, varHint *string, ref bool, input *Statement, slot expandSlot) {
	cs := input.Assert(Id(t.Name()))
	if ref {
		cs = ptrUtils[t.Name()].Clone().Call(cs)
	}
	slot.assign.Add(cs)
}

func (ctx *Ctx) expandNamedBasic(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
	cs := input.Assert(qualifiedNamedType(t))
	if !ref {
		slot.assign.Add(cs)
		return
	}

	localVar := _idLocalVar
	if varHint != nil {
		localVar = *varHint
	}
	slot.define.Add(Id(localVar).Op(":=").Add(cs))

	slot.assign.Add(Op("&").Id(localVar))
}

func (ctx *Ctx) expandNamedPointer(t *types.Named, hint *string, ref bool, input *Statement, slot expandSlot) {
	if input == nil {
		ctx.expandType(t.Underlying().(*types.Pointer).Elem(), hint, true, nil, slot)
		return
	}

	assignSlot := &Statement{}
	newSlot := expandSlot{
		f:      slot.f,
		define: slot.define,
		assign: assignSlot,
	}
	ctx.expandType(t.Underlying().(*types.Pointer).Elem(), hint, true, input, newSlot)
	slot.assign.Add(qualifiedNamedType(t).Call(assignSlot))
}

func (ctx *Ctx) expandSlice(t *types.Slice, varHint *string, ref bool, input *Statement, slot expandSlot) {
	etName, et, isPtr := elemType(t.Elem())
	if isPtr {
		etName += "Ptr"
	}
	if isPtr {
		et = Op("*").Add(et)
	}

	expandFuncName := fmt.Sprintf("expand%sSlice", strcase.ToCamel(etName))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Index().Interface())))
	}

	// Create an expand function for the given type
	if ctx.existFuncs[expandFuncName] {
		return
	}
	defer func() { ctx.existFuncs[expandFuncName] = true }()

	slot.f.Func().Id(expandFuncName).Params(
		Id(_idInput).Index().Interface(),
	).Do(func(stmt *Statement) {
		if ref {
			stmt.Op("*")
		}
	}).Add(Index().Add(et)).

		// Function block
		BlockFunc(func(g *Group) {
			// Nil check on the input, e.g.:
			//
			// if len(input) == 0 {
			//     return nil
			// }
			g.If(Len(Id(_idInput)).Op("==").Lit(0)).BlockFunc(func(g *Group) {
				g.Return(Nil())
			})

			// Initialize the output array, e.g.
			//
			// output := make([]*int, 0)
			g.Id(_idOutput).Op(":=").Make(Index().Add(et), Lit(0))

			// Prepare the slots
			assignSlot, defineSlot := &Statement{}, &Statement{}
			newSlot := expandSlot{
				f:      slot.f,
				assign: assignSlot,
				define: defineSlot,
			}

			localVar := _idLocalVar
			ctx.expandType(t.Elem(), &localVar, false, Id(_idSliceElem), newSlot)

			g.For(List(Id("_"), Id(_idSliceElem)).Op(":=").Range().Id(_idInput)).Block(
				defineSlot,
				Id(_idOutput).Op("=").Append(Id(_idOutput), assignSlot),
			)

			g.Return(Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Id(_idOutput))
		})
}

func (ctx *Ctx) expandNamedSlice(t *types.Named, hint *string, ref bool, input *Statement, slot expandSlot) {
	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Index().Interface())))
	}

	// Create an expand function for the given type
	if ctx.existFuncs[expandFuncName] {
		return
	}
	defer func() { ctx.existFuncs[expandFuncName] = true }()

	slot.f.Func().Id(expandFuncName).Params(
		Id(_idInput).Index().Interface(),
	).Do(func(stmt *Statement) {
		if ref {
			stmt.Op("*")
		}
	}).Add(qualifiedNamedType(t)).

		// Function block
		BlockFunc(func(g *Group) {
			// Nil check on the input, e.g.:
			//
			// if len(input) == 0 {
			//     return nil
			// }
			g.If(Len(Id(_idInput)).Op("==").Lit(0)).BlockFunc(func(g *Group) {
				g.Return(Nil())
			})

			// Initialize the output array, e.g.
			//
			// output := make([]*int, 0)
			g.Id(_idOutput).Op(":=").Make(qualifiedNamedType(t), Lit(0))

			// Prepare the slots
			assignSlot, defineSlot := &Statement{}, &Statement{}
			newSlot := expandSlot{
				f:      slot.f,
				assign: assignSlot,
				define: defineSlot,
			}

			localVar := _idLocalVar
			ctx.expandType(t.Underlying().(*types.Slice).Elem(), &localVar, false, Id(_idSliceElem), newSlot)

			g.For(List(Id("_"), Id(_idSliceElem)).Op(":=").Range().Id(_idInput)).Block(
				defineSlot,
				Id(_idOutput).Op("=").Append(Id(_idOutput), assignSlot),
			)

			g.Return(Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Id(_idOutput))
		})
}

func (ctx *Ctx) expandMap(t *types.Map, hint *string, ref bool, input *Statement, slot expandSlot) {
	// Type guard, Terraform only support map[string]interface{}, it is non-trivial to expand from a non-string keyed map to Terraform.
	ut := t.Underlying().(*types.Map)
	kt, ok := ut.Key().(*types.Basic)
	if !ok || kt.Kind() != types.String {
		log.Fatalf("Only support expanding Map with String key (%v has key type %v)", t, kt)
	}

	// Construct the expand function name.
	etName, et, isPtr := elemType(ut.Elem())
	if isPtr {
		etName += "Ptr"
	}
	if isPtr {
		et = Op("*").Add(et)
	}

	expandFuncName := fmt.Sprintf("expand%sMap", strcase.ToCamel(etName))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Map(String()).Interface())))
	}

	// Create an expand function for the given type
	if ctx.existFuncs[expandFuncName] {
		return
	}
	defer func() { ctx.existFuncs[expandFuncName] = true }()

	slot.f.Func().Id(expandFuncName).Params(
		Id(_idInput).Map(String()).Interface(),
	).Do(func(stmt *Statement) {
		if ref {
			stmt.Op("*")
		}
	}).Add(Map(String()).Add(et)).

		// Function block
		BlockFunc(func(g *Group) {
			// Initialize the output map, e.g.
			//
			// output := make(map[string]*string)
			g.Id(_idOutput).Op(":=").Make(Map(String()).Add(et))

			// Prepare the slots
			assignSlot, defineSlot := &Statement{}, &Statement{}
			newSlot := expandSlot{
				f:      slot.f,
				assign: assignSlot,
				define: defineSlot,
			}

			localVar := _idLocalVar
			ctx.expandType(t.Elem(), &localVar, false, Id("v"), newSlot)

			g.For(List(Id("k"), Id("v")).Op(":=").Range().Id(_idInput)).Block(
				defineSlot,
				Id(_idOutput).Index(Id("k")).Op("=").Add(assignSlot),
			)

			g.Return(Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Id(_idOutput))
		})
}

func (ctx *Ctx) expandNamedMap(t *types.Named, hint *string, ref bool, input *Statement, slot expandSlot) {
	// Type guard, Terraform only support map[string]interface{}, it is non-trivial to expand from a non-string keyed map to Terraform.
	ut := t.Underlying().(*types.Map)
	kt, ok := ut.Key().(*types.Basic)
	if !ok || kt.Kind() != types.String {
		log.Fatalf("Only support expanding Map with String key (%v has key type %v)", t, kt)
	}

	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Map(String()).Interface())))
	}

	// Create an expand function for the given type
	if ctx.existFuncs[expandFuncName] {
		return
	}
	defer func() { ctx.existFuncs[expandFuncName] = true }()

	slot.f.Func().Id(expandFuncName).Params(
		Id(_idInput).Map(String()).Interface(),
	).Do(func(stmt *Statement) {
		if ref {
			stmt.Op("*")
		}
	}).Add(qualifiedNamedType(t)).

		// Function block
		BlockFunc(func(g *Group) {
			// Initialize the output map, e.g.
			//
			// output := make(map[string]*string)
			g.Id(_idOutput).Op(":=").Make(qualifiedNamedType(t))

			// Prepare the slots
			assignSlot, defineSlot := &Statement{}, &Statement{}
			newSlot := expandSlot{
				f:      slot.f,
				assign: assignSlot,
				define: defineSlot,
			}

			localVar := _idLocalVar
			ctx.expandType(ut.Elem(), &localVar, false, Id("v"), newSlot)

			g.For(List(Id("k"), Id("v")).Op(":=").Range().Id(_idInput)).Block(
				defineSlot,
				Id(_idOutput).Index(Id("k")).Op("=").Add(assignSlot),
			)

			g.Return(Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Id(_idOutput))
		})
}

func (ctx *Ctx) expandNamedStruct(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Index().Interface())))
	}

	// Create an expand function for the given type
	if ctx.existFuncs[expandFuncName] {
		return
	}
	defer func() { ctx.existFuncs[expandFuncName] = true }()

	slot.f.Func().Id(expandFuncName).Params(
		Id(_idInput).Index().Interface(),
	).Do(func(stmt *Statement) {
		if ref {
			stmt.Op("*")
		}
	}).Add(qualifiedNamedType(t)).

		// Function block
		BlockFunc(func(g *Group) {
			// Nil check on the input, e.g.:
			//
			// if len(input == 0 || input[0] == nil) {
			//     return foo.TypeFoo{}
			// }
			g.If(Len(Id(_idInput)).Op("==").Lit(0).Op("||").Id(_idInput).Index(Lit(0)).Op("==").Nil()).BlockFunc(func(g *Group) {
				if ref {
					g.Return(Nil())
				} else {
					g.Return(qualifiedNamedType(t).Values())
				}
			})

			// Get the nested block map, e.g.:
			//
			// b := input[0].(map[string]interface{})
			g.Id(_idEncloseBlock).Op(":=").Id(_idInput).Index(Lit(0)).Assert(Map(String()).Interface())

			// Loop over the struct fields and get their "slots"
			type slotCtx struct {
				field    *types.Var
				slot     expandSlot
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

				defineSlot := g
				assignSlot := &Statement{}
				assignSlots = append(assignSlots, assignSlot)
				sctx := slotCtx{
					field: v,
					slot: expandSlot{
						f:      slot.f,
						define: defineSlot,
						assign: assignSlot,
					},
					input:    Id(_idEncloseBlock).Index(Lit(strcase.ToSnake(v.Name()))),
					localVar: strcase.ToLowerCamel(v.Name()),
				}
				slotCtxList = append(slotCtxList, sctx)
			}

			for _, sctx := range slotCtxList {
				ctx.expandType(sctx.field.Type(), &sctx.localVar, false, sctx.input, sctx.slot)
			}

			g.Id(_idOutput).Op(":=").Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Add(qualifiedNamedType(t)).Values(DictFunc(func(d Dict) {
				for idx, sctx := range slotCtxList {
					d[Id(sctx.field.Name())] = assignSlots[idx]
				}
			}))
			g.Return(Id(_idOutput))
		})
}

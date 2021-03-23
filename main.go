package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"os"

	"github.com/dave/dst/decorator"
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

const usage = "terialize <pkg> <type>"

const _idEncloseBlock = "b"
const _idLocalVar = "e"
const _idOutput = "output"
const _idSliceElem = "elem"
const _idInput = "input"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	pkgName, typeName := flag.Args()[0], flag.Args()[1]

	f := run(".", pkgName, typeName)

	if err := f.Render(os.Stdout); err != nil {
		log.Fatalf("failed to render: %v", err)
	}
}

func run(dir string, pkgName string, typeName string) *File {
	pkgs, err := decorator.Load(&packages.Config{Dir: dir, Mode: packages.LoadSyntax}, pkgName)
	if err != nil {
		log.Fatal(err)
	}

	pkg := pkgs[0]

	var targetIdent *ast.Ident
	for ident := range pkg.TypesInfo.Defs {
		if ident.Name == typeName {
			targetIdent = ident
			break
		}
	}
	if targetIdent == nil {
		log.Fatalf("no type named %q found in package %q", typeName, pkgName)
	}

	typeObj := pkg.TypesInfo.Defs[targetIdent]

	f := NewFile("main")
	expandType(typeObj.Type(), nil, false, nil, expandSlot{f: f})
	return f
}

type Slot interface {
	Add(code ...Code) *Statement
}

type expandSlot struct {
	f      *File
	define Slot
	assign Slot
}

func expandType(t types.Type, varHint *string, ref bool, input *Statement, slot expandSlot) {
	switch t := t.(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal( "Can't expand basic type")
		}
		expandBasic(t, ref, input, slot)
	case *types.Pointer:
		expandType(t.Elem(), varHint, true, input, slot)
	case *types.Slice:
		expandSlice(t, varHint, ref, input, slot)
	case *types.Map:
		panic("TODO")
	case *types.Named:
		expandNamedType(t, varHint, ref, input, slot)
	default:
		// Ignore: Array, Chan, Tuple, Signature, Struct, Interface
		return
	}
}

func expandNamedType(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
	switch t.Underlying().(type) {
	case *types.Basic:
		if input == nil {
			log.Fatal( "Can't expand basic type")
		}
		expandNamedBasic(t, varHint, ref, input, slot)
	case *types.Pointer:
		expandNamedPointer(t, varHint, ref, input, slot)
	case *types.Slice:
		expandNamedSlice(t, varHint, ref, input, slot)
	case *types.Map:
		panic("TODO")
	case *types.Struct:
		expandNamedStruct(t, varHint, ref, input, slot)
	case *types.Interface:
		panic("TODO")
	default:
		// Array, Chan, Tuple, Signature
		return
	}
}

func expandBasic(t *types.Basic, ref bool, input *Statement, slot expandSlot) {
	cs := input.Assert(Id(t.Name()))
	if ref {
		cs = ptrUtils[t.Name()].Clone().Call(cs)
	}
	slot.assign.Add(cs)
}

func expandNamedBasic(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
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

func expandSlice(t *types.Slice, varHint *string, ref bool, input *Statement, slot expandSlot) {
	etName, et, isPtr := sliceElemType(t)
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
			expandType(t.Elem(), &localVar, false, Id(_idSliceElem), newSlot)

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

func expandNamedSlice(t *types.Named, hint *string, ref bool, input *Statement, slot expandSlot) {
	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Index().Interface())))
	}

	// Create an expand function for the given type
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
			expandType(t.Underlying().(*types.Slice).Elem(), &localVar, false, Id(_idSliceElem), newSlot)

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

func expandNamedPointer(t *types.Named, hint *string, ref bool, input *Statement, slot expandSlot) {
	if input == nil {
		expandType(t.Underlying().(*types.Pointer).Elem(), hint, true, nil, slot)
		return
	}

	assignSlot := &Statement{}
	newSlot := expandSlot{
		f:      slot.f,
		define: slot.define,
		assign: assignSlot,
	}
	expandType(t.Underlying().(*types.Pointer).Elem(), hint, true, input, newSlot)
	slot.assign.Add(Id(t.Obj().Name()).Call(assignSlot))
}

func expandNamedStruct(t *types.Named, varHint *string, ref bool, input *Statement, slot expandSlot) {
	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))
	if ref {
		expandFuncName += "Ptr"
	}

	// Fill in the assign slot of the invoker.
	if slot.assign != nil {
		slot.assign.Add(Id(expandFuncName).Call(input.Assert(Index().Interface())))
	}

	// Create an expand function for the given type
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

			var defineSlots []*Statement
			var assignSlots []*Statement
			ut := t.Underlying().(*types.Struct)
			for i := 0; i < ut.NumFields(); i++ {
				v := ut.Field(i)
				if !v.Exported() {
					continue
				}

				defineSlot := &Statement{}
				assignSlot := &Statement{}
				defineSlots = append(defineSlots, defineSlot)
				assignSlots = append(assignSlots, assignSlot)
				ctx := slotCtx{
					field: v,
					slot: expandSlot{
						f:      slot.f,
						define: defineSlot,
						assign: assignSlot,
					},
					input:    Id(_idEncloseBlock).Index(Lit(strcase.ToSnake(v.Name()))),
					localVar: strcase.ToLowerCamel(v.Name()),
				}
				slotCtxList = append(slotCtxList, ctx)
			}

			for i, ctx := range slotCtxList {
				expandType(ctx.field.Type(), &ctx.localVar, false, ctx.input, ctx.slot)
				g.Add(defineSlots[i])
			}

			g.Id(_idOutput).Op(":=").Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Add(qualifiedNamedType(t)).Values(DictFunc(func(d Dict) {
				for idx, ctx := range slotCtxList {
					d[Id(ctx.field.Name())] = assignSlots[idx]
				}
			}))
			g.Return(Id(_idOutput))
		})

}

// TODO: make this configurable
var ptrUtils = map[string]*Statement{
	"bool":    Qual("types/utils", "Bool"),
	"int":     Qual("types/utils", "Int"),
	"float64": Qual("types/utils", "Float64"),
	"string":  Qual("types/utils", "String"),
}

func qualifiedNamedType(t *types.Named) *Statement {
	tPkgPath, tName := t.Obj().Pkg().Path(), t.Obj().Name()
	return Qual(tPkgPath, tName)
}

func sliceElemType(t *types.Slice) (name string, stmt *Statement, ref bool) {
	et := t.Elem()
	if pt, ok := t.Elem().(*types.Pointer); ok {
		ref = true
		et = pt.Elem()
	}

	switch et := et.(type) {
	case *types.Named:
		return et.Obj().Name(), qualifiedNamedType(et), ref
	case *types.Basic:
		return et.Name(), Id(et.Name()), ref
	}
	panic(fmt.Sprintf("Currently, we only support slice of named type, primary type or pointer of those types. But got: %v (elem: %v)", t, et))
}

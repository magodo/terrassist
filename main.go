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
	expandType(f, typeObj.Type(), false)
	return f
}

var ptrUtils = map[string]*Statement{
	"bool":    Id("utils").Dot("Bool"),
	"int":     Id("utils").Dot("Int"),
	"float64": Id("utils").Dot("Float64"),
	"string":  Id("utils").Dot("String"),
}

func expandType(f *File, t types.Type, ref bool) (cs expandCallSite, inline bool) {
	switch t := t.(type) {
	case *types.Basic:
		return expandBasic(t, ref)
	case *types.Pointer:
		return expandType(f, t.Elem(), true)
	case *types.Slice:
		panic("TODO")
	case *types.Map:
		panic("TODO")
	case *types.Named:
		return expandNamedType(f, t, ref)
	default:
		// Ignore: Array, Chan, Tuple, Signature, Struct, Interface
		return nil, false
	}
}

func expandBasic(t *types.Basic, ref bool) (expandCallSite, bool) {
	return func(input *Statement) func(*Statement) {
		return func(s *Statement) {
			cs := input.Assert(Id(t.Name()))
			if ref {
				cs = ptrUtils[t.Name()].Clone().Call(cs)
			}
			s.Add(cs)
		}
	}, true
}

func expandNamedType(f *File, t *types.Named, ref bool) (expandCallSite, bool) {
	switch t.Underlying().(type) {
	case *types.Basic:
		return expandNamedBasic(t, ref)
	case *types.Pointer:
		panic("TODO")
	case *types.Slice:
		panic("TODO")
	case *types.Map:
		panic("TODO")
	case *types.Struct:
		return expandNamedStruct(f, t, false)
	case *types.Interface:
		panic("TODO")
	default:
		// Array, Chan, Tuple, Signature
		return nil, false
	}
}

type expandCallSite func(input *Statement) func(s *Statement)

func expandNamedBasic(t *types.Named, ref bool) (expandCallSite, bool) {
	return func(input *Statement) func(*Statement) {
		return func(s *Statement) {
			s.Add(input.Assert(qualifiedNamedType(t)))
		}
	}, !ref
}

func expandNamedStruct(f *File, t *types.Named, ref bool) (expandCallSite, bool) {
	expandFuncName := fmt.Sprintf("expand%s", strcase.ToCamel(t.Obj().Name()))

	// Create an expand function for the given type
	f.Func().Id(expandFuncName).Params(
		Id("input").Index().Interface(),
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
			g.If(Len(Id("input").Op("==").Lit(0).Op("||").Id("input").Index(Lit(0)).Op("==").Nil())).BlockFunc(func(g *Group) {
				if ref {
					g.Return(Nil())
				} else {
					g.Return(qualifiedNamedType(t).Values())
				}
			})

			// Get the nested block map, e.g.:
			//
			// b := input[0].(map[string]interface{})
			g.Id(_idEncloseBlock).Op(":=").Id("input").Index(Lit(0)).Assert(Map(String()).Interface())

			// Loop over the struct fields and get their "expandCallsite" callbacks and the location indicators.
			type litCompCtx struct {
				field  *types.Var
				doFunc func(*Statement)
			}

			var litCompCtxes []litCompCtx

			ut := t.Underlying().(*types.Struct)
			for i := 0; i < ut.NumFields(); i++ {
				v := ut.Field(i)
				if !v.Exported() {
					continue
				}

				input := Id(_idEncloseBlock).Index(Lit(strcase.ToSnake(v.Name())))
				cs, inline := expandType(f, v.Type(), false)

				ctx := litCompCtx{
					field:  v,
					doFunc: cs(input.Clone()),
				}
				if !inline {
					localVar := Id(strcase.ToLowerCamel(v.Name()))

					// Add expand call site before constructing struct.
					g.Add(localVar.Clone().Op(":=").Do(ctx.doFunc))

					// Modify the context so that the field assignment uses the local variable that are expanded into prior.
					ctx.doFunc = func(s *Statement) {
						s.Op("&").Add(localVar)
					}
				}
				litCompCtxes = append(litCompCtxes, ctx)
			}

			g.Id("output").Op(":=").Do(func(stmt *Statement) {
				if ref {
					stmt.Op("&")
				}
			}).Add(qualifiedNamedType(t)).Values(DictFunc(func(d Dict) {
				for _, ctx := range litCompCtxes {
					s := Statement{}
					d[Id(ctx.field.Name())] = s.Do(ctx.doFunc)
				}
			}))
			g.Return(Id("output"))
		})

	return func(input *Statement) func(*Statement) {
		return func(s *Statement) {
			s.Id(expandFuncName).Call(input.Assert(Index().Interface()))
		}
	}, true
}

func qualifiedNamedType(t *types.Named) *Statement {
	tPkgPath, tName := t.Obj().Pkg().Path(), t.Obj().Name()
	return Qual(tPkgPath, tName)
}

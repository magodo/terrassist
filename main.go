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
	"golang.org/x/tools/go/packages"
)

const usage = "terialize <pkg> <type>"

const _idEncloseBlock = "b"
const _idLocalVar = "e"
const _idOutput = "output"
const _idSliceElem = "elem"
const _idInput = "input"

type Slot interface {
	Add(code ...Code) *Statement
}

type Ctx struct {
	existFuncs map[string]bool
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

func elemType(et types.Type) (name string, stmt *Statement, ref bool) {
	if pt, ok := et.(*types.Pointer); ok {
		ref = true
		et = pt.Elem()
	}

	switch et := et.(type) {
	case *types.Named:
		return et.Obj().Name(), qualifiedNamedType(et), ref
	case *types.Basic:
		return et.Name(), Id(et.Name()), ref
	}
	log.Fatalf("Currently, we only support element of named type, primary type or pointer of those types (got: %v)", et)
	return
}

func (ctx *Ctx) run(dir string, pkgName string, typeName string) *File {
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

	f := NewFile("output")
	ctx.expandType(typeObj.Type(), nil, false, nil, expandSlot{f: f})
	ctx.flattenType(typeObj.Type(), nil, false, nil, flattenSlot{f: f})
	return f
}

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

	ctx := Ctx{existFuncs: map[string]bool{}}
	f := ctx.run(".", pkgName, typeName)

	if err := f.Render(os.Stdout); err != nil {
		log.Fatalf("failed to render: %v", err)
	}
}

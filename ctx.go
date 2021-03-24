package main

import (
	"github.com/dave/dst/decorator"
	. "github.com/dave/jennifer/jen"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
)

const _idEncloseBlock = "b"
const _idLocalVar = "e"
const _idOutput = "output"
const _idSliceElem = "elem"
const _idInput = "input"

type Slot interface {
	Add(code ...Code) *Statement
}

type options struct {
	honorJSONIgnore bool
	forPointer      bool
}

type Ctx struct {
	existFuncs map[string]bool
	options
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
	ctx.expandType(typeObj.Type(), nil, ctx.forPointer, nil, expandSlot{f: f})
	ctx.flattenType(typeObj.Type(), nil, ctx.forPointer, nil, flattenSlot{f: f})
	return f
}

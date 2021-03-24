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

type basicTypeInfo struct {
	PtrHelperFunc *Statement // TODO: make this configurable
	Type          *Statement
}

var basicTypeInfoMap = map[types.BasicKind]basicTypeInfo{
	types.Bool: {
		PtrHelperFunc: Qual("types/utils", "Bool"),
		Type:          Bool(),
	},
	types.Int: {
		PtrHelperFunc: Qual("types/utils", "Int"),
		Type:          Int(),
	},
	types.Int8: {
		PtrHelperFunc: Qual("types/utils", "Int8"),
		Type:          Int8(),
	},
	types.Int16: {
		PtrHelperFunc: Qual("types/utils", "Int16"),
		Type:          Int16(),
	},
	types.Int32: {
		PtrHelperFunc: Qual("types/utils", "Int32"),
		Type:          Int32(),
	},
	types.Int64: {
		PtrHelperFunc: Qual("types/utils", "Int64"),
		Type:          Int64(),
	},
	types.Uint: {
		PtrHelperFunc: Qual("types/utils", "Uint"),
		Type:          Uint(),
	},
	types.Uint8: {
		PtrHelperFunc: Qual("types/utils", "Uint8"),
		Type:          Uint8(),
	},
	types.Uint16: {
		PtrHelperFunc: Qual("types/utils", "Uint16"),
		Type:          Uint16(),
	},
	types.Uint32: {
		PtrHelperFunc: Qual("types/utils", "Uint32"),
		Type:          Uint32(),
	},
	types.Uint64: {
		PtrHelperFunc: Qual("types/utils", "Uint64"),
		Type:          Uint64(),
	},
	types.Float32: {
		PtrHelperFunc: Qual("types/utils", "Float32"),
		Type:          Float32(),
	},
	types.Float64: {
		PtrHelperFunc: Qual("types/utils", "Float64"),
		Type:          Float64(),
	},
	types.Complex64: {
		PtrHelperFunc: Qual("types/utils", "Complex64"),
		Type:          Complex64(),
	},
	types.Complex128: {
		PtrHelperFunc: Qual("types/utils", "Complex128"),
		Type:          Complex128(),
	},
	types.String: {
		PtrHelperFunc: Qual("types/utils", "String"),
		Type:          String(),
	},
}

var reservedWordAlternatives = map[string]string{
	"break":       "brk",
	"default":     "dflt",
	"func":        "fun",
	"interface":   "itf",
	"select":      "sel",
	"case":        "cs",
	"defer":       "df",
	"go":          "g",
	"map":         "m",
	"struct":      "strct",
	"chan":        "ch",
	"else":        "eLse",
	"goto":        "gt",
	"package":     "pkg",
	"switch":      "swtch",
	"const":       "cst",
	"fallthrough": "fth",
	"if":          "iF",
	"range":       "rng",
	"type":        "typ",
	"continue":    "cont",
	"for":         "fOr",
	"import":      "impt",
	"return":      "ret",
	"var":         "vAr",
}

// newIdent returns the input id, except the id is a Go reserved word, in which case we will use an alternative word instead.
func newIdent(id string) string {
	newid := reservedWordAlternatives[id]
	if newid != "" {
		return newid
	}
	return id
}

// qualifiedNamedType return the Qual() of the specified named type
func qualifiedNamedType(t *types.Named) *Statement {
	tPkgPath, tName := t.Obj().Pkg().Path(), t.Obj().Name()
	return Qual(tPkgPath, tName)
}

// elemType type return the name of the element type of the input type, together with its *Statement representation, and a flag indicating the element is a pointer or not.
// NOTE: currently it only supports element of named type, primary type or pointer of those types.
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

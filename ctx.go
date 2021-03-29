package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"strings"
)

const _idEncloseBlock = "b"
const _idLocalVar = "e"
const _idOutput = "output"
const _idSliceElem = "elem"
const _idInput = "input"

type Slot interface {
	Add(code ...Code) *Statement
}

type flags struct {
	honorJSONIgnore bool
}

type CtxOptions struct {
	Dir      string
	PkgName  string
	TypeExpr string
}

type Ctx struct {
	flags
	thisPkg    *packages.Package
	targetPkg  *packages.Package
	existFuncs map[string]bool
	t          types.Type
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

func (ctx *Ctx) run() *File {
	f := NewFile(ctx.thisPkg.Name)
	ctx.expandType(ctx.t, nil, false, nil, expandSlot{f: f})
	ctx.flattenType(ctx.t, nil, false, nil, flattenSlot{f: f})
	return f
}

func packagesErrors(pkgs []*packages.Package) error {
	errors := []interface{}{}
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for _, err := range pkg.Errors {
			errors = append(errors, err)
		}
	})

	if len(errors) == 0 {
		return nil
	}

	tpl := strings.Repeat("%w\n", len(errors))
	return fmt.Errorf(tpl, errors...)
}

func NewCtx(opts CtxOptions, flags flags) (*Ctx, error) {
	thisPkgs, err := packages.Load(&packages.Config{Dir: opts.Dir})
	if err != nil {
		return nil, err
	}
	if err := packagesErrors(thisPkgs); err != nil {
		return nil, err
	}
	thisPkg := thisPkgs[0]

	pkgs, err := packages.Load(&packages.Config{Dir: opts.Dir, Mode: packages.LoadSyntax}, opts.PkgName)
	if err != nil {
		return nil, err
	}
	if err := packagesErrors(thisPkgs); err != nil {
		return nil, err
	}
	pkg := pkgs[0]

	buildType, typeExpr := processTypeExpr(opts.TypeExpr)

	var t types.Type
	for _, obj := range pkg.TypesInfo.Defs {
		if _, ok := obj.(*types.TypeName); ok && obj.Name() == typeExpr {
			t = obj.Type()
			break
		}
	}
	if t == nil {
		return nil, fmt.Errorf("no type named %q found in package %q", typeExpr, opts.PkgName)
	}

	t = buildType(t)

	return &Ctx{
		flags:      flags,
		thisPkg:    thisPkg,
		targetPkg:  pkg,
		existFuncs: map[string]bool{},
		t:          t,
	}, nil
}

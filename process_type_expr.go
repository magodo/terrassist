package main

import (
	"fmt"
	"go/types"
	"log"
	"regexp"
)

var (
	idPattern = `[a-zA-Z_]\w*`

	tPattern            = regexp.MustCompile(fmt.Sprintf(`^(%s)$`, idPattern))
	tPtrPattern         = regexp.MustCompile(fmt.Sprintf(`^\*(%s)$`, idPattern))
	tMapPattern         = regexp.MustCompile(fmt.Sprintf(`^map\[string\](%s)$`, idPattern))
	tMapPtrPattern      = regexp.MustCompile(fmt.Sprintf(`^\*map\[string\](%s)$`, idPattern))
	tPtrMapPattern      = regexp.MustCompile(fmt.Sprintf(`^map\[string\]\*(%s)$`, idPattern))
	tPtrMapPtrPattern   = regexp.MustCompile(fmt.Sprintf(`^\*map\[string\]\*(%s)$`, idPattern))
	tSlicePattern       = regexp.MustCompile(fmt.Sprintf(`^\[\](%s)$`, idPattern))
	tSlicePtrPattern    = regexp.MustCompile(fmt.Sprintf(`^\*\[\](%s)$`, idPattern))
	tPtrSlicePattern    = regexp.MustCompile(fmt.Sprintf(`^\[\]\*(%s)$`, idPattern))
	tPtrSlicePtrPattern = regexp.MustCompile(fmt.Sprintf(`^\*\[\]\*(%s)$`, idPattern))
)

func processTypeExpr(typeExpr string) (func(t types.Type) types.Type, string) {
	var p *regexp.Regexp
	var buildType func(t types.Type) types.Type

	switch {
	case tPattern.MatchString(typeExpr):
		p = tPattern
		buildType = func(t types.Type) types.Type {
			return t
		}
	case tPtrPattern.MatchString(typeExpr):
		p = tPtrPattern
		buildType = func(t types.Type) types.Type {
			return types.NewPointer(t)
		}
	case tMapPattern.MatchString(typeExpr):
		p = tMapPattern
		buildType = func(t types.Type) types.Type {
			return types.NewMap(types.Typ[types.String], t)
		}
	case tMapPtrPattern.MatchString(typeExpr):
		p = tMapPtrPattern
		buildType = func(t types.Type) types.Type {
			return types.NewPointer(types.NewMap(types.Typ[types.String], t))
		}
	case tPtrMapPattern.MatchString(typeExpr):
		p = tPtrMapPattern
		buildType = func(t types.Type) types.Type {
			return types.NewMap(types.Typ[types.String], types.NewPointer(t))
		}
	case tPtrMapPtrPattern.MatchString(typeExpr):
		p = tPtrMapPtrPattern
		buildType = func(t types.Type) types.Type {
			return types.NewPointer(types.NewMap(types.Typ[types.String], types.NewPointer(t)))
		}
	case tSlicePattern.MatchString(typeExpr):
		p = tSlicePattern
		buildType = func(t types.Type) types.Type {
			return types.NewSlice(t)
		}
	case tSlicePtrPattern.MatchString(typeExpr):
		p = tSlicePtrPattern
		buildType = func(t types.Type) types.Type {
			return types.NewPointer(types.NewSlice(t))
		}
	case tPtrSlicePattern.MatchString(typeExpr):
		p = tPtrSlicePattern
		buildType = func(t types.Type) types.Type {
			return types.NewSlice(types.NewPointer(t))
		}
	case tPtrSlicePtrPattern.MatchString(typeExpr):
		p = tPtrSlicePtrPattern
		buildType = func(t types.Type) types.Type {
			return types.NewPointer(types.NewSlice(types.NewPointer(t)))
		}
	default:
		log.Fatalf("invalid type expression specified: %q", typeExpr)
	}

	typeName := p.FindStringSubmatch(typeExpr)[1]
	return buildType, typeName
}

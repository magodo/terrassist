package main

import (
	"github.com/stretchr/testify/require"
	"go/types"
	"testing"
)

func TestProcessTypeExpr(t *testing.T) {
	var tFoo types.Type

	cases := []struct {
		typeExpr           string
		expectMainTypeName string
		expectFinalType    types.Type
	}{
		{
			typeExpr:           "Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    tFoo,
		},
		{
			typeExpr:           "*Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewPointer(tFoo),
		},
		{
			typeExpr:           "map[string]Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewMap(types.Typ[types.String], tFoo),
		},
		{
			typeExpr:           "map[string]*Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewMap(types.Typ[types.String], types.NewPointer(tFoo)),
		},
		{
			typeExpr:           "*map[string]Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewPointer(types.NewMap(types.Typ[types.String], tFoo)),
		},
		{
			typeExpr:           "*map[string]*Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewPointer(types.NewMap(types.Typ[types.String], types.NewPointer(tFoo))),
		},
		{
			typeExpr:           "[]Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewSlice(tFoo),
		},
		{
			typeExpr:           "[]*Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewSlice(types.NewPointer(tFoo)),
		},
		{
			typeExpr:           "*[]Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewPointer(types.NewSlice(tFoo)),
		},
		{
			typeExpr:           "*[]*Foo",
			expectMainTypeName: "Foo",
			expectFinalType:    types.NewPointer(types.NewSlice(types.NewPointer(tFoo))),
		},
	}

	for _, c := range cases {
		b, mt := processTypeExpr(c.typeExpr)
		require.Equal(t, c.expectMainTypeName, mt, c.typeExpr)
		require.True(t, types.Identical(c.expectFinalType, b(tFoo)), c.typeExpr)
	}
}

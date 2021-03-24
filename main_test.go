package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	pwd, _ := os.Getwd()
	typesDir := filepath.Join(pwd, "testdata", "types")
	resultBaseDir := filepath.Join(pwd, "testdata", "results")

	cases := []struct {
		typ     string
		options options
	}{
		{
			typ: "TypePrimaryCollection",
		},
		{
			typ: "TypePrimaryPtrCollection",
		},
		{
			typ: "TypePrimarySliceCollection",
		},
		{
			typ: "TypePrimaryPtrSliceCollection",
		},
		{
			typ: "TypePrimarySlicePtrCollection",
		},
		{
			typ: "TypePrimaryPtrSlicePtrCollection",
		},
		{
			typ: "TypePrimaryMapCollection",
		},
		{
			typ: "TypePrimaryPtrMapCollection",
		},
		{
			typ: "TypePrimaryMapPtrCollection",
		},
		{
			typ: "TypePrimaryPtrMapPtrCollection",
		},
		{
			typ: "TypePrimaryAliasCollection",
		},
		{
			typ: "TypePrimaryAliasPtrCollection",
		},
		{
			typ: "TypePrimaryAliasSliceCollection",
		},
		{
			typ: "TypePrimaryAliasPtrSliceCollection",
		},
		{
			typ: "TypePrimaryAliasSlicePtrCollection",
		},
		{
			typ: "TypePrimaryAliasPtrSlicePtrCollection",
		},
		{
			typ: "TypePrimaryAliasMapCollection",
		},
		{
			typ: "TypePrimaryAliasPtrMapCollection",
		},
		{
			typ: "TypePrimaryAliasMapPtrCollection",
		},
		{
			typ: "TypePrimaryAliasPtrMapPtrCollection",
		},
		{
			typ: "TypePrimaryPtrAliasCollection",
		},
		{
			typ: "TypePrimaryPtrAliasSliceCollection",
		},
		{
			typ: "TypePrimaryPtrAliasSlicePtrCollection",
		},
		{
			typ: "TypePrimaryPtrAliasMapCollection",
		},
		{
			typ: "TypePrimaryPtrAliasMapPtrCollection",
		},
		{
			typ: "TypeNamedStructCollection",
		},
		{
			typ: "TypeNamedStructPointerCollection",
		},
		{
			typ: "TypeNamedStructSliceCollection",
		},
		{
			typ: "TypeNamedStructMapCollection",
		},
		{
			typ: "TypeIntSliceAlias",
		},
		{
			typ: "TypeIntPtrSliceAlias",
		},
		{
			typ: "TypeNamedBoolSliceAlias",
		},
		{
			typ: "TypeNamedStructSliceAlias",
		},
		{
			typ: "TypeIntMapAlias",
		},
		{
			typ: "TypeIntPtrMapAlias",
		},
		{
			typ: "TypeNamedBoolMapAlias",
		},
		{
			typ: "TypeNamedStructMapAlias",
		},
		{
			typ: "TypeNamedStructSliceAlias",
		},
		{
			typ: "TypeNamedStructAlias",
		},
		{
			typ: "TypeNamedStructPtrAlias",
		},
		{
			typ:     "TypeNamedStructWithJSONIgnore",
			options: options{honorJSONIgnore: true},
		},
		{
			typ:     "TypeCyclicRefStruct",
			options: options{forPointer: true},
		},
	}

	for _, c := range cases {
		buf := bytes.NewBufferString("")
		ctx := Ctx{existFuncs: map[string]bool{}, options: c.options}
		f := ctx.run(typesDir, "types", c.typ)
		require.NoError(t, f.Render(buf), c.typ)

		expectFile := filepath.Join(resultBaseDir, c.typ, "main.go")
		expectByte, _ := os.ReadFile(expectFile)
		require.Equal(t, string(expectByte), buf.String(), c.typ)
	}
}

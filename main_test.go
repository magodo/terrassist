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
		typ          string
		expectFolder string
		flags        Flags
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
			typ:   "TypeNamedStructWithJSONIgnore",
			flags: Flags{HonorJSONIgnore: true},
		},
		{
			typ: "TypeCyclicRefStruct",
		},
		{
			typ: "TypeNamedStructWithGoReservedWord",
		},
		{
			typ: "TypeNamedInterface",
		},
		{
			typ:          "[]TypeS1",
			expectFolder: "TypeS1_Slice",
		},
		{
			typ:          "[]TypePrimaryPtrCollection",
			expectFolder: "TypePrimaryPtrCollection_Slice",
		},
		{
			typ:          "[]TypeNamedStructCollection",
			expectFolder: "TypeNamedStructCollection_Slice",
		},
		{
			typ: "TypeNonNativePrimaryCollection",
		},
		{
			typ: "TypeNonNativePrimaryPtrCollection",
		},
		{
			typ: "TypeNonNativePrimaryAliasCollection",
		},
		{
			typ: "TypeNonNativePrimaryAliasPtrCollection",
		},
	}

	for _, c := range cases {
		buf := bytes.NewBufferString("")
		ctx, err := NewCtx(CtxOptions{
			Dir:      typesDir,
			PkgName:  "types",
			TypeExpr: c.typ,
			Flags:    c.flags,
		})
		require.NoError(t, err, c.typ)

		f := ctx.run()
		require.NoError(t, f.Render(buf), c.typ)

		expectFolder := c.expectFolder
		if expectFolder == "" {
			expectFolder = c.typ
		}
		expectFile := filepath.Join(resultBaseDir, expectFolder, "main.go")
		expectByte, _ := os.ReadFile(expectFile)
		require.Equal(t, string(expectByte), buf.String(), c.typ)
	}
}

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpand(t *testing.T) {
	pwd, _ := os.Getwd()
	typesDir := filepath.Join(pwd, "testdata", "types")
	resultBaseDir := filepath.Join(pwd, "testdata", "results")

	cases := []struct {
		typ string
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
			typ: "TypeIntSliceAlias",
		},
		{
			typ: "TypeIntPtrSliceAlias",
		},
		{
			typ: "TypeNamedBoolSliceAlias",
		},
		{
			typ: "TypeNamedBoolSliceAlias",
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
	}

	for _, c := range cases {
		buf := bytes.NewBufferString("")
		f := run(typesDir, "types", c.typ)
		require.NoError(t, f.Render(buf), c.typ)

		expectFile := filepath.Join(resultBaseDir, c.typ, "main.go")
		expectByte, _ := os.ReadFile(expectFile)
		require.Equal(t, string(expectByte), buf.String(), c.typ)
	}
}

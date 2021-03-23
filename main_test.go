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
		// 0
		{
			typ: "TypePrimaryCollection",
		},
		// 1
		{
			typ: "TypePrimaryPtrCollection",
		},
		// 2
		{
			typ: "TypePrimarySliceCollection",
		},
		// 3
		{
			typ: "TypePrimaryPtrSliceCollection",
		},
		// 4
		{
			typ: "TypePrimaryPtrSlicePtrCollection",
		},
		// 5
		{
			typ: "TypePrimaryAliasCollection",
		},
		// 6
		{
			typ: "TypePrimaryAliasPtrCollection",
		},
		// 7
		{
			typ: "TypePrimaryAliasSliceCollection",
		},
		// 8
		{
			typ: "TypePrimaryAliasPtrSliceCollection",
		},
		// 9
		{
			typ: "TypePrimaryAliasPtrSlicePtrCollection",
		},
		// 10
		{
			typ: "TypePrimaryPtrAliasCollection",
		},
		// 11
		{
			typ: "TypePrimaryPtrAliasSliceCollection",
		},
		// 12
		{
			typ: "TypePrimaryPtrAliasSlicePtrCollection",
		},
		// 13
		{
			typ: "TypeNamedStructCollection",
		},
		// 14
		{
			typ: "TypeNamedStructPointerCollection",
		},
		// 15
		{
			typ: "TypeNamedStructSliceCollection",
		},
		// 16
		{
			typ: "TypeIntSliceAlias",
		},
		// 17
		{
			typ: "TypeIntPtrSliceAlias",
		},
		// 18
		{
			typ: "TypeNamedBoolSliceAlias",
		},
		// 19
		{
			typ: "TypeNamedBoolSliceAlias",
		},
		// 20
		{
			typ: "TypeNamedStructSliceAlias",
		},
		// 21
		{
			typ: "TypeNamedStructAlias",
		},
		// 22
		{
			typ: "TypeNamedStructPtrAlias",
		},
	}

	for idx, c := range cases {
		buf := bytes.NewBufferString("")
		f := run(typesDir, "types", c.typ)
		require.NoError(t, f.Render(buf), idx)

		expectFile := filepath.Join(resultBaseDir, c.typ, "main.go")
		expectByte, _ := os.ReadFile(expectFile)
		require.Equal(t, string(expectByte), buf.String(), idx)
	}
}

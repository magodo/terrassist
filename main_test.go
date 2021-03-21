package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpand(t *testing.T) {
	cases := []struct {
		dir    string
		pkg    string
		typ    string
		expect string
	}{
		{
			dir:    "testdata",
			pkg:    "foo",
			typ:    "TypePrimaryCollection",
			expect: `package main

import foo "foo"

func expandTypePrimaryCollection(input []interface{}) foo.TypePrimaryCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypePrimaryCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypePrimaryCollection{
		Bool:     b["bool"].(bool),
		Float:    b["float"].(float64),
		Int:      b["int"].(int),
		PtrBool:  utils.Bool(b["ptr_bool"].(bool)),
		PtrFloat: utils.Float64(b["ptr_float"].(float64)),
		PtrInt:   utils.Int(b["ptr_int"].(int)),
		PtrStr:   utils.String(b["ptr_str"].(string)),
		Str:      b["str"].(string),
	}
	return output
}
`,

		},
		{
			dir:    "testdata",
			pkg:    "foo",
			typ:    "TypePrimaryAliasCollection",
			expect: `package main

import foo "foo"

func expandTypePrimaryAliasCollection(input []interface{}) foo.TypePrimaryAliasCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypePrimaryAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	ptrBool := b["ptr_bool"].(foo.Bool)
	ptrInt := b["ptr_int"].(foo.Int)
	ptrStr := b["ptr_str"].(foo.Str)
	ptrFloat := b["ptr_float"].(foo.Float)
	output := foo.TypePrimaryAliasCollection{
		Bool:     b["bool"].(foo.Bool),
		Float:    b["float"].(foo.Float),
		Int:      b["int"].(foo.Int),
		PtrBool:  &ptrBool,
		PtrFloat: &ptrFloat,
		PtrInt:   &ptrInt,
		PtrStr:   &ptrStr,
		Str:      b["str"].(foo.Str),
	}
	return output
}
`,

		},
	}

	for idx, c := range cases {
		buf := bytes.NewBufferString("")
		f := run(c.dir, c.pkg, c.typ)
		require.NoError(t, f.Render(buf), idx)
		require.Equal(t, c.expect, buf.String(), idx)
	}
}

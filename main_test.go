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
		// 0
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypePrimaryCollection",
			expect: `package main

import foo "foo"

func expandTypePrimaryCollection(input []interface{}) foo.TypePrimaryCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypePrimaryCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypePrimaryCollection{
		Bool:             b["bool"].(bool),
		Float:            b["float"].(float64),
		Int:              b["int"].(int),
		PtrBool:          utils.Bool(b["ptr_bool"].(bool)),
		PtrFloat:         utils.Float64(b["ptr_float"].(float64)),
		PtrInt:           utils.Int(b["ptr_int"].(int)),
		PtrSliceBool:     expandBoolSlicePtr(b["ptr_slice_bool"].([]interface{})),
		PtrSliceFloat:    expandFloat64SlicePtr(b["ptr_slice_float"].([]interface{})),
		PtrSliceInt:      expandIntSlicePtr(b["ptr_slice_int"].([]interface{})),
		PtrSlicePtrBool:  expandBoolPtrSlicePtr(b["ptr_slice_ptr_bool"].([]interface{})),
		PtrSlicePtrFloat: expandFloat64PtrSlicePtr(b["ptr_slice_ptr_float"].([]interface{})),
		PtrSlicePtrInt:   expandIntPtrSlicePtr(b["ptr_slice_ptr_int"].([]interface{})),
		PtrSlicePtrStr:   expandStringPtrSlicePtr(b["ptr_slice_ptr_str"].([]interface{})),
		PtrSliceStr:      expandStringSlicePtr(b["ptr_slice_str"].([]interface{})),
		PtrStr:           utils.String(b["ptr_str"].(string)),
		SliceBool:        expandBoolSlice(b["slice_bool"].([]interface{})),
		SliceFloat:       expandFloat64Slice(b["slice_float"].([]interface{})),
		SliceInt:         expandIntSlice(b["slice_int"].([]interface{})),
		SlicePtrBool:     expandBoolPtrSlice(b["slice_ptr_bool"].([]interface{})),
		SlicePtrFloat:    expandFloat64PtrSlice(b["slice_ptr_float"].([]interface{})),
		SlicePtrInt:      expandIntPtrSlice(b["slice_ptr_int"].([]interface{})),
		SlicePtrStr:      expandStringPtrSlice(b["slice_ptr_str"].([]interface{})),
		SliceStr:         expandStringSlice(b["slice_str"].([]interface{})),
		Str:              b["str"].(string),
	}
	return output
}
func expandBoolSlice(input []interface{}) []bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]bool, 0)
	for _, elem := range input {
		output = append(output, elem.(bool))
	}
	return output
}
func expandIntSlice(input []interface{}) []int {
	if len(input) == 0 {
		return nil
	}
	output := make([]int, 0)
	for _, elem := range input {
		output = append(output, elem.(int))
	}
	return output
}
func expandStringSlice(input []interface{}) []string {
	if len(input) == 0 {
		return nil
	}
	output := make([]string, 0)
	for _, elem := range input {
		output = append(output, elem.(string))
	}
	return output
}
func expandFloat64Slice(input []interface{}) []float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]float64, 0)
	for _, elem := range input {
		output = append(output, elem.(float64))
	}
	return output
}
func expandBoolPtrSlice(input []interface{}) []*bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*bool, 0)
	for _, elem := range input {
		output = append(output, utils.Bool(elem.(bool)))
	}
	return output
}
func expandIntPtrSlice(input []interface{}) []*int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*int, 0)
	for _, elem := range input {
		output = append(output, utils.Int(elem.(int)))
	}
	return output
}
func expandStringPtrSlice(input []interface{}) []*string {
	if len(input) == 0 {
		return nil
	}
	output := make([]*string, 0)
	for _, elem := range input {
		output = append(output, utils.String(elem.(string)))
	}
	return output
}
func expandFloat64PtrSlice(input []interface{}) []*float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*float64, 0)
	for _, elem := range input {
		output = append(output, utils.Float64(elem.(float64)))
	}
	return output
}
func expandBoolSlicePtr(input []interface{}) *[]bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]bool, 0)
	for _, elem := range input {
		output = append(output, elem.(bool))
	}
	return &output
}
func expandIntSlicePtr(input []interface{}) *[]int {
	if len(input) == 0 {
		return nil
	}
	output := make([]int, 0)
	for _, elem := range input {
		output = append(output, elem.(int))
	}
	return &output
}
func expandStringSlicePtr(input []interface{}) *[]string {
	if len(input) == 0 {
		return nil
	}
	output := make([]string, 0)
	for _, elem := range input {
		output = append(output, elem.(string))
	}
	return &output
}
func expandFloat64SlicePtr(input []interface{}) *[]float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]float64, 0)
	for _, elem := range input {
		output = append(output, elem.(float64))
	}
	return &output
}
func expandBoolPtrSlicePtr(input []interface{}) *[]*bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*bool, 0)
	for _, elem := range input {
		output = append(output, utils.Bool(elem.(bool)))
	}
	return &output
}
func expandIntPtrSlicePtr(input []interface{}) *[]*int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*int, 0)
	for _, elem := range input {
		output = append(output, utils.Int(elem.(int)))
	}
	return &output
}
func expandStringPtrSlicePtr(input []interface{}) *[]*string {
	if len(input) == 0 {
		return nil
	}
	output := make([]*string, 0)
	for _, elem := range input {
		output = append(output, utils.String(elem.(string)))
	}
	return &output
}
func expandFloat64PtrSlicePtr(input []interface{}) *[]*float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*float64, 0)
	for _, elem := range input {
		output = append(output, utils.Float64(elem.(float64)))
	}
	return &output
}
`,
		},
		// 1
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypePrimaryAliasCollection",
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
		Bool:             b["bool"].(foo.Bool),
		Float:            b["float"].(foo.Float),
		Int:              b["int"].(foo.Int),
		PtrBool:          &ptrBool,
		PtrFloat:         &ptrFloat,
		PtrInt:           &ptrInt,
		PtrSliceBool:     expandBoolSlicePtr(b["ptr_slice_bool"].([]interface{})),
		PtrSliceFloat:    expandFloatSlicePtr(b["ptr_slice_float"].([]interface{})),
		PtrSliceInt:      expandIntSlicePtr(b["ptr_slice_int"].([]interface{})),
		PtrSlicePtrBool:  expandBoolPtrSlicePtr(b["ptr_slice_ptr_bool"].([]interface{})),
		PtrSlicePtrFloat: expandFloatPtrSlicePtr(b["ptr_slice_ptr_float"].([]interface{})),
		PtrSlicePtrInt:   expandIntPtrSlicePtr(b["ptr_slice_ptr_int"].([]interface{})),
		PtrSlicePtrStr:   expandStrPtrSlicePtr(b["ptr_slice_ptr_str"].([]interface{})),
		PtrSliceStr:      expandStrSlicePtr(b["ptr_slice_str"].([]interface{})),
		PtrStr:           &ptrStr,
		SliceBool:        expandBoolSlice(b["slice_bool"].([]interface{})),
		SliceFloat:       expandFloatSlice(b["slice_float"].([]interface{})),
		SliceInt:         expandIntSlice(b["slice_int"].([]interface{})),
		SlicePtrBool:     expandBoolPtrSlice(b["slice_ptr_bool"].([]interface{})),
		SlicePtrFloat:    expandFloatPtrSlice(b["slice_ptr_float"].([]interface{})),
		SlicePtrInt:      expandIntPtrSlice(b["slice_ptr_int"].([]interface{})),
		SlicePtrStr:      expandStrPtrSlice(b["slice_ptr_str"].([]interface{})),
		SliceStr:         expandStrSlice(b["slice_str"].([]interface{})),
		Str:              b["str"].(foo.Str),
	}
	return output
}
func expandBoolSlice(input []interface{}) []foo.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Bool, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Bool))
	}
	return output
}
func expandIntSlice(input []interface{}) []foo.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Int, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Int))
	}
	return output
}
func expandStrSlice(input []interface{}) []foo.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Str, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Str))
	}
	return output
}
func expandFloatSlice(input []interface{}) []foo.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Float, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Float))
	}
	return output
}
func expandBoolPtrSlice(input []interface{}) []*foo.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Bool, 0)
	for _, elem := range input {
		e := elem.(foo.Bool)
		output = append(output, &e)
	}
	return output
}
func expandIntPtrSlice(input []interface{}) []*foo.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Int, 0)
	for _, elem := range input {
		e := elem.(foo.Int)
		output = append(output, &e)
	}
	return output
}
func expandStrPtrSlice(input []interface{}) []*foo.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Str, 0)
	for _, elem := range input {
		e := elem.(foo.Str)
		output = append(output, &e)
	}
	return output
}
func expandFloatPtrSlice(input []interface{}) []*foo.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Float, 0)
	for _, elem := range input {
		e := elem.(foo.Float)
		output = append(output, &e)
	}
	return output
}
func expandBoolSlicePtr(input []interface{}) *[]foo.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Bool, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Bool))
	}
	return &output
}
func expandIntSlicePtr(input []interface{}) *[]foo.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Int, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Int))
	}
	return &output
}
func expandStrSlicePtr(input []interface{}) *[]foo.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Str, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Str))
	}
	return &output
}
func expandFloatSlicePtr(input []interface{}) *[]foo.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.Float, 0)
	for _, elem := range input {
		output = append(output, elem.(foo.Float))
	}
	return &output
}
func expandBoolPtrSlicePtr(input []interface{}) *[]*foo.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Bool, 0)
	for _, elem := range input {
		e := elem.(foo.Bool)
		output = append(output, &e)
	}
	return &output
}
func expandIntPtrSlicePtr(input []interface{}) *[]*foo.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Int, 0)
	for _, elem := range input {
		e := elem.(foo.Int)
		output = append(output, &e)
	}
	return &output
}
func expandStrPtrSlicePtr(input []interface{}) *[]*foo.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Str, 0)
	for _, elem := range input {
		e := elem.(foo.Str)
		output = append(output, &e)
	}
	return &output
}
func expandFloatPtrSlicePtr(input []interface{}) *[]*foo.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.Float, 0)
	for _, elem := range input {
		e := elem.(foo.Float)
		output = append(output, &e)
	}
	return &output
}
`,
		},
		// 2
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypePrimaryPointerAliasCollection",
			expect: `package main

import foo "foo"

func expandTypePrimaryPointerAliasCollection(input []interface{}) foo.TypePrimaryPointerAliasCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypePrimaryPointerAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypePrimaryPointerAliasCollection{
		PtrBool:          PtrBool(utils.Bool(b["ptr_bool"].(bool))),
		PtrFloat:         PtrFloat(utils.Float64(b["ptr_float"].(float64))),
		PtrInt:           PtrInt(utils.Int(b["ptr_int"].(int))),
		PtrSlicePtrBool:  expandPtrBoolSlicePtr(b["ptr_slice_ptr_bool"].([]interface{})),
		PtrSlicePtrFloat: expandPtrFloatSlicePtr(b["ptr_slice_ptr_float"].([]interface{})),
		PtrSlicePtrInt:   expandPtrIntSlicePtr(b["ptr_slice_ptr_int"].([]interface{})),
		PtrSlicePtrStr:   expandPtrStrSlicePtr(b["ptr_slice_ptr_str"].([]interface{})),
		PtrStr:           PtrStr(utils.String(b["ptr_str"].(string))),
		SlicePtrBool:     expandPtrBoolSlice(b["slice_ptr_bool"].([]interface{})),
		SlicePtrFloat:    expandPtrFloatSlice(b["slice_ptr_float"].([]interface{})),
		SlicePtrInt:      expandPtrIntSlice(b["slice_ptr_int"].([]interface{})),
		SlicePtrStr:      expandPtrStrSlice(b["slice_ptr_str"].([]interface{})),
	}
	return output
}
func expandPtrBoolSlice(input []interface{}) []foo.PtrBool {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrBool, 0)
	for _, elem := range input {
		output = append(output, PtrBool(utils.Bool(elem.(bool))))
	}
	return output
}
func expandPtrIntSlice(input []interface{}) []foo.PtrInt {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrInt, 0)
	for _, elem := range input {
		output = append(output, PtrInt(utils.Int(elem.(int))))
	}
	return output
}
func expandPtrStrSlice(input []interface{}) []foo.PtrStr {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrStr, 0)
	for _, elem := range input {
		output = append(output, PtrStr(utils.String(elem.(string))))
	}
	return output
}
func expandPtrFloatSlice(input []interface{}) []foo.PtrFloat {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrFloat, 0)
	for _, elem := range input {
		output = append(output, PtrFloat(utils.Float64(elem.(float64))))
	}
	return output
}
func expandPtrBoolSlicePtr(input []interface{}) *[]foo.PtrBool {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrBool, 0)
	for _, elem := range input {
		output = append(output, PtrBool(utils.Bool(elem.(bool))))
	}
	return &output
}
func expandPtrIntSlicePtr(input []interface{}) *[]foo.PtrInt {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrInt, 0)
	for _, elem := range input {
		output = append(output, PtrInt(utils.Int(elem.(int))))
	}
	return &output
}
func expandPtrStrSlicePtr(input []interface{}) *[]foo.PtrStr {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrStr, 0)
	for _, elem := range input {
		output = append(output, PtrStr(utils.String(elem.(string))))
	}
	return &output
}
func expandPtrFloatSlicePtr(input []interface{}) *[]foo.PtrFloat {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.PtrFloat, 0)
	for _, elem := range input {
		output = append(output, PtrFloat(utils.Float64(elem.(float64))))
	}
	return &output
}
`,
		},
		// 3
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypeNamedStructCollection",
			expect: `package main

import foo "foo"

func expandTypeNamedStructCollection(input []interface{}) foo.TypeNamedStructCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeNamedStructCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeNamedStructCollection{
		Bar: expandTypeBar(b["bar"].([]interface{})),
		Foo: expandTypeFoo(b["foo"].([]interface{})),
	}
	return output
}
func expandTypeFoo(input []interface{}) foo.TypeFoo {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeFoo{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeFoo{I: b["i"].(int)}
	return output
}
func expandTypeBar(input []interface{}) foo.TypeBar {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeBar{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeBar{S: b["s"].(string)}
	return output
}
`,
		},
		// 4
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypeNamedStructPointerCollection",
			expect: `package main

import foo "foo"

func expandTypeNamedStructPointerCollection(input []interface{}) foo.TypeNamedStructPointerCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeNamedStructPointerCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeNamedStructPointerCollection{
		Bar: expandTypeBarPtr(b["bar"].([]interface{})),
		Foo: expandTypeFooPtr(b["foo"].([]interface{})),
	}
	return output
}
func expandTypeFooPtr(input []interface{}) *foo.TypeFoo {
	if len(input == 0 || input[0] == nil) {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &foo.TypeFoo{I: b["i"].(int)}
	return output
}
func expandTypeBarPtr(input []interface{}) *foo.TypeBar {
	if len(input == 0 || input[0] == nil) {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &foo.TypeBar{S: b["s"].(string)}
	return output
}
`,
		},
		// 5
		{
			dir: "testdata",
			pkg: "foo",
			typ: "TypeNamedStructSliceCollection",
			expect: `package main

import foo "foo"

func expandTypeNamedStructSliceCollection(input []interface{}) foo.TypeNamedStructSliceCollection {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeNamedStructSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeNamedStructSliceCollection{
		FooPtrs: expandTypeFooPtrSlice(b["foo_ptrs"].([]interface{})),
		Foos:    expandTypeFooSlice(b["foos"].([]interface{})),
		FoosPtr: expandTypeFooSlicePtr(b["foos_ptr"].([]interface{})),
	}
	return output
}
func expandTypeFooSlice(input []interface{}) []foo.TypeFoo {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.TypeFoo, 0)
	for _, elem := range input {
		output = append(output, expandTypeFoo(elem.([]interface{})))
	}
	return output
}
func expandTypeFoo(input []interface{}) foo.TypeFoo {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeFoo{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeFoo{I: b["i"].(int)}
	return output
}
func expandTypeFooSlicePtr(input []interface{}) *[]foo.TypeFoo {
	if len(input) == 0 {
		return nil
	}
	output := make([]foo.TypeFoo, 0)
	for _, elem := range input {
		output = append(output, expandTypeFoo(elem.([]interface{})))
	}
	return &output
}
func expandTypeFoo(input []interface{}) foo.TypeFoo {
	if len(input == 0 || input[0] == nil) {
		return foo.TypeFoo{}
	}
	b := input[0].(map[string]interface{})
	output := foo.TypeFoo{I: b["i"].(int)}
	return output
}
func expandTypeFooPtrSlice(input []interface{}) []*foo.TypeFoo {
	if len(input) == 0 {
		return nil
	}
	output := make([]*foo.TypeFoo, 0)
	for _, elem := range input {
		output = append(output, expandTypeFooPtr(elem.([]interface{})))
	}
	return output
}
func expandTypeFooPtr(input []interface{}) *foo.TypeFoo {
	if len(input == 0 || input[0] == nil) {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &foo.TypeFoo{I: b["i"].(int)}
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

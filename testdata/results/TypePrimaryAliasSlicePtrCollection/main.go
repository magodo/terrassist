package main

import types "types"

func expandTypePrimaryAliasSlicePtrCollection(input []interface{}) types.TypePrimaryAliasSlicePtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasSlicePtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasSlicePtrCollection{
		BoolSlicePtr:   expandBoolSlicePtr(b["bool_slice_ptr"].([]interface{})),
		FloatSlicePtr:  expandFloatSlicePtr(b["float_slice_ptr"].([]interface{})),
		IntSlicePtr:    expandIntSlicePtr(b["int_slice_ptr"].([]interface{})),
		StringSlicePtr: expandStrSlicePtr(b["string_slice_ptr"].([]interface{})),
	}
	return output
}
func expandBoolSlicePtr(input []interface{}) *[]types.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Bool, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Bool))
	}
	return &output
}
func expandIntSlicePtr(input []interface{}) *[]types.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Int, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Int))
	}
	return &output
}
func expandStrSlicePtr(input []interface{}) *[]types.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Str, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Str))
	}
	return &output
}
func expandFloatSlicePtr(input []interface{}) *[]types.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Float, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Float))
	}
	return &output
}

package main

import types "types"

func expandTypePrimarySlicePtrCollection(input []interface{}) types.TypePrimarySlicePtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimarySlicePtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimarySlicePtrCollection{
		BoolSlicePtr:  expandBoolSlicePtr(b["bool_slice_ptr"].([]interface{})),
		FloatSlicePtr: expandFloat64SlicePtr(b["float_slice_ptr"].([]interface{})),
		IntSlicePtr:   expandIntSlicePtr(b["int_slice_ptr"].([]interface{})),
		StrSlicePtr:   expandStringSlicePtr(b["str_slice_ptr"].([]interface{})),
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

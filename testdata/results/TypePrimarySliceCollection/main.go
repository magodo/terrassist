package main

import types "types"

func expandTypePrimarySliceCollection(input []interface{}) types.TypePrimarySliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimarySliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimarySliceCollection{
		BoolSlice:  expandBoolSlice(b["bool_slice"].([]interface{})),
		FloatSlice: expandFloat64Slice(b["float_slice"].([]interface{})),
		IntSlice:   expandIntSlice(b["int_slice"].([]interface{})),
		StrSlice:   expandStringSlice(b["str_slice"].([]interface{})),
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

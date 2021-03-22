package main

import types "types"

func expandTypePrimaryAliasSliceCollection(input []interface{}) types.TypePrimaryAliasSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasSliceCollection{
		BoolSlice:  expandBoolSlice(b["bool_slice"].([]interface{})),
		FloatSlice: expandFloatSlice(b["float_slice"].([]interface{})),
		IntSlice:   expandIntSlice(b["int_slice"].([]interface{})),
		StrSlice:   expandStrSlice(b["str_slice"].([]interface{})),
	}
	return output
}
func expandBoolSlice(input []interface{}) []types.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Bool, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Bool))
	}
	return output
}
func expandIntSlice(input []interface{}) []types.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Int, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Int))
	}
	return output
}
func expandStrSlice(input []interface{}) []types.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Str, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Str))
	}
	return output
}
func expandFloatSlice(input []interface{}) []types.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Float, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Float))
	}
	return output
}

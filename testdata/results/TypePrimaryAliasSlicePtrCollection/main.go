package types

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
		output = append(output, types.Bool(elem.(bool)))
	}
	return &output
}
func expandIntSlicePtr(input []interface{}) *[]types.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Int, 0)
	for _, elem := range input {
		output = append(output, types.Int(elem.(int)))
	}
	return &output
}
func expandStrSlicePtr(input []interface{}) *[]types.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Str, 0)
	for _, elem := range input {
		output = append(output, types.Str(elem.(string)))
	}
	return &output
}
func expandFloatSlicePtr(input []interface{}) *[]types.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.Float, 0)
	for _, elem := range input {
		output = append(output, types.Float(elem.(float64)))
	}
	return &output
}
func flattenTypePrimaryAliasSlicePtrCollection(input types.TypePrimaryAliasSlicePtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_slice_ptr":   flattenBoolSlicePtr(input.BoolSlicePtr),
		"float_slice_ptr":  flattenFloatSlicePtr(input.FloatSlicePtr),
		"int_slice_ptr":    flattenIntSlicePtr(input.IntSlicePtr),
		"string_slice_ptr": flattenStrSlicePtr(input.StringSlicePtr),
	}}
}
func flattenBoolSlicePtr(input *[]types.Bool) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		output = append(output, bool(elem))
	}
	return output
}
func flattenIntSlicePtr(input *[]types.Int) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		output = append(output, int(elem))
	}
	return output
}
func flattenStrSlicePtr(input *[]types.Str) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		output = append(output, string(elem))
	}
	return output
}
func flattenFloatSlicePtr(input *[]types.Float) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		output = append(output, float64(elem))
	}
	return output
}

package output

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
func flattenTypePrimaryAliasSliceCollection(input types.TypePrimaryAliasSliceCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_slice":  flattenBoolSlice(input.BoolSlice),
		"float_slice": flattenFloatSlice(input.FloatSlice),
		"int_slice":   flattenIntSlice(input.IntSlice),
		"str_slice":   flattenStrSlice(input.StrSlice),
	}}
}
func flattenBoolSlice(input []types.Bool) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, bool(elem))
	}
	return output
}
func flattenIntSlice(input []types.Int) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, int(elem))
	}
	return output
}
func flattenStrSlice(input []types.Str) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, string(elem))
	}
	return output
}
func flattenFloatSlice(input []types.Float) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, float64(elem))
	}
	return output
}

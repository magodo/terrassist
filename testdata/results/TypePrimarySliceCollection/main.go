package output

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
func flattenTypePrimarySliceCollection(input types.TypePrimarySliceCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_slice":  flattenBoolSlice(input.BoolSlice),
		"float_slice": flattenFloat64Slice(input.FloatSlice),
		"int_slice":   flattenIntSlice(input.IntSlice),
		"str_slice":   flattenStringSlice(input.StrSlice),
	}}
}
func flattenBoolSlice(input []bool) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, elem)
	}
	return output
}
func flattenIntSlice(input []int) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, elem)
	}
	return output
}
func flattenStringSlice(input []string) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, elem)
	}
	return output
}
func flattenFloat64Slice(input []float64) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, elem)
	}
	return output
}

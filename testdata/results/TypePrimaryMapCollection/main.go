package types

import types "types"

func expandTypePrimaryMapCollection(input []interface{}) types.TypePrimaryMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryMapCollection{
		BoolMap:  expandBoolMap(b["bool_map"].(map[string]interface{})),
		FloatMap: expandFloat64Map(b["float_map"].(map[string]interface{})),
		IntMap:   expandIntMap(b["int_map"].(map[string]interface{})),
		StrMap:   expandStringMap(b["str_map"].(map[string]interface{})),
	}
	return output
}
func expandBoolMap(input map[string]interface{}) map[string]bool {
	output := make(map[string]bool)
	for k, v := range input {
		output[k] = v.(bool)
	}
	return output
}
func expandIntMap(input map[string]interface{}) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		output[k] = v.(int)
	}
	return output
}
func expandStringMap(input map[string]interface{}) map[string]string {
	output := make(map[string]string)
	for k, v := range input {
		output[k] = v.(string)
	}
	return output
}
func expandFloat64Map(input map[string]interface{}) map[string]float64 {
	output := make(map[string]float64)
	for k, v := range input {
		output[k] = v.(float64)
	}
	return output
}
func flattenTypePrimaryMapCollection(input types.TypePrimaryMapCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_map":  flattenBoolMap(input.BoolMap),
		"float_map": flattenFloat64Map(input.FloatMap),
		"int_map":   flattenIntMap(input.IntMap),
		"str_map":   flattenStringMap(input.StrMap),
	}}
}
func flattenBoolMap(input map[string]bool) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}
func flattenIntMap(input map[string]int) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}
func flattenStringMap(input map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}
func flattenFloat64Map(input map[string]float64) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}

package output

import types "types"

func expandTypePrimaryMapPtrCollection(input []interface{}) types.TypePrimaryMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryMapPtrCollection{
		BoolMapPtr:  expandBoolMapPtr(b["bool_map_ptr"].(map[string]interface{})),
		FloatMapPtr: expandFloat64MapPtr(b["float_map_ptr"].(map[string]interface{})),
		IntMapPtr:   expandIntMapPtr(b["int_map_ptr"].(map[string]interface{})),
		StrMapPtr:   expandStringMapPtr(b["str_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolMapPtr(input map[string]interface{}) *map[string]bool {
	output := make(map[string]bool)
	for k, v := range input {
		output[k] = v.(bool)
	}
	return &output
}
func expandIntMapPtr(input map[string]interface{}) *map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		output[k] = v.(int)
	}
	return &output
}
func expandStringMapPtr(input map[string]interface{}) *map[string]string {
	output := make(map[string]string)
	for k, v := range input {
		output[k] = v.(string)
	}
	return &output
}
func expandFloat64MapPtr(input map[string]interface{}) *map[string]float64 {
	output := make(map[string]float64)
	for k, v := range input {
		output[k] = v.(float64)
	}
	return &output
}

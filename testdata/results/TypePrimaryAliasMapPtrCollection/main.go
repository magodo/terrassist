package main

import types "types"

func expandTypePrimaryAliasMapPtrCollection(input []interface{}) types.TypePrimaryAliasMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasMapPtrCollection{
		BoolMapPtr:   expandBoolMapPtr(b["bool_map_ptr"].(map[string]interface{})),
		FloatMapPtr:  expandFloatMapPtr(b["float_map_ptr"].(map[string]interface{})),
		IntMapPtr:    expandIntMapPtr(b["int_map_ptr"].(map[string]interface{})),
		StringMapPtr: expandStrMapPtr(b["string_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolMapPtr(input map[string]interface{}) *map[string]types.Bool {
	output := make(map[string]types.Bool)
	for k, v := range input {
		output[k] = v.(types.Bool)
	}
	return &output
}
func expandIntMapPtr(input map[string]interface{}) *map[string]types.Int {
	output := make(map[string]types.Int)
	for k, v := range input {
		output[k] = v.(types.Int)
	}
	return &output
}
func expandStrMapPtr(input map[string]interface{}) *map[string]types.Str {
	output := make(map[string]types.Str)
	for k, v := range input {
		output[k] = v.(types.Str)
	}
	return &output
}
func expandFloatMapPtr(input map[string]interface{}) *map[string]types.Float {
	output := make(map[string]types.Float)
	for k, v := range input {
		output[k] = v.(types.Float)
	}
	return &output
}

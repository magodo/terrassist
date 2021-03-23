package main

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasMapPtrCollection(input []interface{}) types.TypePrimaryPtrAliasMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasMapPtrCollection{
		BoolPtrMapPtr: expandBoolPtrMapPtr(b["bool_ptr_map_ptr"].(map[string]interface{})),
		FlatPtrMapPtr: expandFloatPtrMapPtr(b["flat_ptr_map_ptr"].(map[string]interface{})),
		IntPtrMapPtr:  expandIntPtrMapPtr(b["int_ptr_map_ptr"].(map[string]interface{})),
		StrPtrMapPtr:  expandStrPtrMapPtr(b["str_ptr_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMapPtr(input map[string]interface{}) *map[string]types.BoolPtr {
	output := make(map[string]types.BoolPtr)
	for k, v := range input {
		output[k] = BoolPtr(utils.Bool(v.(bool)))
	}
	return &output
}
func expandIntPtrMapPtr(input map[string]interface{}) *map[string]types.IntPtr {
	output := make(map[string]types.IntPtr)
	for k, v := range input {
		output[k] = IntPtr(utils.Int(v.(int)))
	}
	return &output
}
func expandStrPtrMapPtr(input map[string]interface{}) *map[string]types.StrPtr {
	output := make(map[string]types.StrPtr)
	for k, v := range input {
		output[k] = StrPtr(utils.String(v.(string)))
	}
	return &output
}
func expandFloatPtrMapPtr(input map[string]interface{}) *map[string]types.FloatPtr {
	output := make(map[string]types.FloatPtr)
	for k, v := range input {
		output[k] = FloatPtr(utils.Float64(v.(float64)))
	}
	return &output
}

package main

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrMapPtrCollection(input []interface{}) types.TypePrimaryPtrMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrMapPtrCollection{
		BoolPtrMapPtr:  expandBoolPtrMapPtr(b["bool_ptr_map_ptr"].(map[string]interface{})),
		FloatPtrMapPtr: expandFloat64PtrMapPtr(b["float_ptr_map_ptr"].(map[string]interface{})),
		IntPtrMapPtr:   expandIntPtrMapPtr(b["int_ptr_map_ptr"].(map[string]interface{})),
		StrPtrMapPtr:   expandStringPtrMapPtr(b["str_ptr_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMapPtr(input map[string]interface{}) *map[string]*bool {
	output := make(map[string]*bool)
	for k, v := range input {
		output[k] = utils.Bool(v.(bool))
	}
	return &output
}
func expandIntPtrMapPtr(input map[string]interface{}) *map[string]*int {
	output := make(map[string]*int)
	for k, v := range input {
		output[k] = utils.Int(v.(int))
	}
	return &output
}
func expandStringPtrMapPtr(input map[string]interface{}) *map[string]*string {
	output := make(map[string]*string)
	for k, v := range input {
		output[k] = utils.String(v.(string))
	}
	return &output
}
func expandFloat64PtrMapPtr(input map[string]interface{}) *map[string]*float64 {
	output := make(map[string]*float64)
	for k, v := range input {
		output[k] = utils.Float64(v.(float64))
	}
	return &output
}

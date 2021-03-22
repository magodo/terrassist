package main

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrSlicePtrCollection(input []interface{}) types.TypePrimaryPtrSlicePtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrSlicePtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrSlicePtrCollection{
		PtrBool:  expandBoolPtrSlicePtr(b["ptr_bool"].([]interface{})),
		PtrFloat: expandFloat64PtrSlicePtr(b["ptr_float"].([]interface{})),
		PtrInt:   expandIntPtrSlicePtr(b["ptr_int"].([]interface{})),
		PtrStr:   expandStringPtrSlicePtr(b["ptr_str"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlicePtr(input []interface{}) *[]*bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*bool, 0)
	for _, elem := range input {
		output = append(output, utils.Bool(elem.(bool)))
	}
	return &output
}
func expandIntPtrSlicePtr(input []interface{}) *[]*int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*int, 0)
	for _, elem := range input {
		output = append(output, utils.Int(elem.(int)))
	}
	return &output
}
func expandStringPtrSlicePtr(input []interface{}) *[]*string {
	if len(input) == 0 {
		return nil
	}
	output := make([]*string, 0)
	for _, elem := range input {
		output = append(output, utils.String(elem.(string)))
	}
	return &output
}
func expandFloat64PtrSlicePtr(input []interface{}) *[]*float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*float64, 0)
	for _, elem := range input {
		output = append(output, utils.Float64(elem.(float64)))
	}
	return &output
}

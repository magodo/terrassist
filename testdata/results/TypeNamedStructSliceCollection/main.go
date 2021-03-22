package main

import types "types"

func expandTypeNamedStructSliceCollection(input []interface{}) types.TypeNamedStructSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructSliceCollection{
		S1PtrSlice:    expandTypeS1PtrSlice(b["s_1_ptr_slice"].([]interface{})),
		S1PtrSlicePtr: expandTypeS1PtrSlicePtr(b["s_1_ptr_slice_ptr"].([]interface{})),
		S1Slice:       expandTypeS1Slice(b["s_1_slice"].([]interface{})),
		S1SlicePtr:    expandTypeS1SlicePtr(b["s_1_slice_ptr"].([]interface{})),
	}
	return output
}
func expandTypeS1Slice(input []interface{}) []types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypeS1, 0)
	for _, elem := range input {
		output = append(output, expandTypeS1(elem.([]interface{})))
	}
	return output
}
func expandTypeS1(input []interface{}) types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return types.TypeS1{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS1SlicePtr(input []interface{}) *[]types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypeS1, 0)
	for _, elem := range input {
		output = append(output, expandTypeS1(elem.([]interface{})))
	}
	return &output
}
func expandTypeS1(input []interface{}) types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return types.TypeS1{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS1PtrSlice(input []interface{}) []*types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.TypeS1, 0)
	for _, elem := range input {
		output = append(output, expandTypeS1Ptr(elem.([]interface{})))
	}
	return output
}
func expandTypeS1Ptr(input []interface{}) *types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS1PtrSlicePtr(input []interface{}) *[]*types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.TypeS1, 0)
	for _, elem := range input {
		output = append(output, expandTypeS1Ptr(elem.([]interface{})))
	}
	return &output
}
func expandTypeS1Ptr(input []interface{}) *types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS1{I: b["i"].(int)}
	return output
}

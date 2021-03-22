package main

import types "types"

func expandTypePrimaryAliasPtrSliceCollection(input []interface{}) types.TypePrimaryAliasPtrSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasPtrSliceCollection{
		BoolPtrSlice:   expandBoolPtrSlice(b["bool_ptr_slice"].([]interface{})),
		FloatPtrSlice:  expandFloatPtrSlice(b["float_ptr_slice"].([]interface{})),
		IntPtrSlice:    expandIntPtrSlice(b["int_ptr_slice"].([]interface{})),
		StringPtrSlice: expandStrPtrSlice(b["string_ptr_slice"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlice(input []interface{}) []*types.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Bool, 0)
	for _, elem := range input {
		e := elem.(types.Bool)
		output = append(output, &e)
	}
	return output
}
func expandIntPtrSlice(input []interface{}) []*types.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Int, 0)
	for _, elem := range input {
		e := elem.(types.Int)
		output = append(output, &e)
	}
	return output
}
func expandStrPtrSlice(input []interface{}) []*types.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Str, 0)
	for _, elem := range input {
		e := elem.(types.Str)
		output = append(output, &e)
	}
	return output
}
func expandFloatPtrSlice(input []interface{}) []*types.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Float, 0)
	for _, elem := range input {
		e := elem.(types.Float)
		output = append(output, &e)
	}
	return output
}

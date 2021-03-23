package output

import types "types"

func expandTypePrimaryAliasPtrSlicePtrCollection(input []interface{}) types.TypePrimaryAliasPtrSlicePtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrSlicePtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasPtrSlicePtrCollection{
		BoolPtrSlicePtr:   expandBoolPtrSlicePtr(b["bool_ptr_slice_ptr"].([]interface{})),
		FloatPtrSlicePtr:  expandFloatPtrSlicePtr(b["float_ptr_slice_ptr"].([]interface{})),
		IntPtrSlicePtr:    expandIntPtrSlicePtr(b["int_ptr_slice_ptr"].([]interface{})),
		StringPtrSlicePtr: expandStrPtrSlicePtr(b["string_ptr_slice_ptr"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlicePtr(input []interface{}) *[]*types.Bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Bool, 0)
	for _, elem := range input {
		e := elem.(types.Bool)
		output = append(output, &e)
	}
	return &output
}
func expandIntPtrSlicePtr(input []interface{}) *[]*types.Int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Int, 0)
	for _, elem := range input {
		e := elem.(types.Int)
		output = append(output, &e)
	}
	return &output
}
func expandStrPtrSlicePtr(input []interface{}) *[]*types.Str {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Str, 0)
	for _, elem := range input {
		e := elem.(types.Str)
		output = append(output, &e)
	}
	return &output
}
func expandFloatPtrSlicePtr(input []interface{}) *[]*types.Float {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.Float, 0)
	for _, elem := range input {
		e := elem.(types.Float)
		output = append(output, &e)
	}
	return &output
}

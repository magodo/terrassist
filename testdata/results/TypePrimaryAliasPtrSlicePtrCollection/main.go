package types

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
func flattenTypePrimaryAliasPtrSlicePtrCollection(input types.TypePrimaryAliasPtrSlicePtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_slice_ptr":   flattenBoolPtrSlicePtr(input.BoolPtrSlicePtr),
		"float_ptr_slice_ptr":  flattenFloatPtrSlicePtr(input.FloatPtrSlicePtr),
		"int_ptr_slice_ptr":    flattenIntPtrSlicePtr(input.IntPtrSlicePtr),
		"string_ptr_slice_ptr": flattenStrPtrSlicePtr(input.StringPtrSlicePtr),
	}}
}
func flattenBoolPtrSlicePtr(input *[]*types.Bool) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e bool
		if elem != nil {
			e = bool(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenIntPtrSlicePtr(input *[]*types.Int) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e int
		if elem != nil {
			e = int(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenStrPtrSlicePtr(input *[]*types.Str) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e string
		if elem != nil {
			e = string(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenFloatPtrSlicePtr(input *[]*types.Float) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e float64
		if elem != nil {
			e = float64(*elem)
		}
		output = append(output, e)
	}
	return output
}

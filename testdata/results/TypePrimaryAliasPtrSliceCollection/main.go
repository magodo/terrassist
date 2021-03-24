package output

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
func flattenTypePrimaryAliasPtrSliceCollection(input types.TypePrimaryAliasPtrSliceCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_slice":   flattenBoolPtrSlice(input.BoolPtrSlice),
		"float_ptr_slice":  flattenFloatPtrSlice(input.FloatPtrSlice),
		"int_ptr_slice":    flattenIntPtrSlice(input.IntPtrSlice),
		"string_ptr_slice": flattenStrPtrSlice(input.StringPtrSlice),
	}}
}
func flattenBoolPtrSlice(input []*types.Bool) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e bool
		if elem != nil {
			e = bool(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenIntPtrSlice(input []*types.Int) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e int
		if elem != nil {
			e = int(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenStrPtrSlice(input []*types.Str) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e string
		if elem != nil {
			e = string(*elem)
		}
		output = append(output, e)
	}
	return output
}
func flattenFloatPtrSlice(input []*types.Float) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e float64
		if elem != nil {
			e = float64(*elem)
		}
		output = append(output, e)
	}
	return output
}

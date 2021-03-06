package types

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
		e := types.Bool(elem.(bool))
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
		e := types.Int(elem.(int))
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
		e := types.Str(elem.(string))
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
		e := types.Float(elem.(float64))
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
		var e types.Bool
		if elem != nil {
			e = *elem
		}
		output = append(output, bool(e))
	}
	return output
}
func flattenIntPtrSlice(input []*types.Int) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e types.Int
		if elem != nil {
			e = *elem
		}
		output = append(output, int(e))
	}
	return output
}
func flattenStrPtrSlice(input []*types.Str) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e types.Str
		if elem != nil {
			e = *elem
		}
		output = append(output, string(e))
	}
	return output
}
func flattenFloatPtrSlice(input []*types.Float) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var e types.Float
		if elem != nil {
			e = *elem
		}
		output = append(output, float64(e))
	}
	return output
}

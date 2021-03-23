package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrSliceCollection(input []interface{}) types.TypePrimaryPtrSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrSliceCollection{
		BoolPtrSlice:  expandBoolPtrSlice(b["bool_ptr_slice"].([]interface{})),
		FloatPtrSlice: expandFloat64PtrSlice(b["float_ptr_slice"].([]interface{})),
		IntPtrSlice:   expandIntPtrSlice(b["int_ptr_slice"].([]interface{})),
		StrPtrSlice:   expandStringPtrSlice(b["str_ptr_slice"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlice(input []interface{}) []*bool {
	if len(input) == 0 {
		return nil
	}
	output := make([]*bool, 0)
	for _, elem := range input {
		output = append(output, utils.Bool(elem.(bool)))
	}
	return output
}
func expandIntPtrSlice(input []interface{}) []*int {
	if len(input) == 0 {
		return nil
	}
	output := make([]*int, 0)
	for _, elem := range input {
		output = append(output, utils.Int(elem.(int)))
	}
	return output
}
func expandStringPtrSlice(input []interface{}) []*string {
	if len(input) == 0 {
		return nil
	}
	output := make([]*string, 0)
	for _, elem := range input {
		output = append(output, utils.String(elem.(string)))
	}
	return output
}
func expandFloat64PtrSlice(input []interface{}) []*float64 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*float64, 0)
	for _, elem := range input {
		output = append(output, utils.Float64(elem.(float64)))
	}
	return output
}
func flattenTypePrimaryPtrSliceCollection(input types.TypePrimaryPtrSliceCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_slice":  flattenBoolPtrSlice(input.BoolPtrSlice),
		"float_ptr_slice": flattenFloat64PtrSlice(input.FloatPtrSlice),
		"int_ptr_slice":   flattenIntPtrSlice(input.IntPtrSlice),
		"str_ptr_slice":   flattenStringPtrSlice(input.StrPtrSlice),
	}}
}
func flattenBoolPtrSlice(input []*bool) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		e := false
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenIntPtrSlice(input []*int) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		e := 0
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenStringPtrSlice(input []*string) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		e := ""
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenFloat64PtrSlice(input []*float64) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		e := 0.0
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}

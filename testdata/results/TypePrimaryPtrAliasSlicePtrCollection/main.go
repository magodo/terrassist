package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasSlicePtrCollection(input []interface{}) types.TypePrimaryPtrAliasSlicePtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasSlicePtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasSlicePtrCollection{
		BoolPtrSlicePtr: expandBoolPtrSlicePtr(b["bool_ptr_slice_ptr"].([]interface{})),
		FlatPtrSlicePtr: expandFloatPtrSlicePtr(b["flat_ptr_slice_ptr"].([]interface{})),
		IntPtrSlicePtr:  expandIntPtrSlicePtr(b["int_ptr_slice_ptr"].([]interface{})),
		StrPtrSlicePtr:  expandStrPtrSlicePtr(b["str_ptr_slice_ptr"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlicePtr(input []interface{}) *[]types.BoolPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.BoolPtr, 0)
	for _, elem := range input {
		output = append(output, types.BoolPtr(utils.Bool(elem.(bool))))
	}
	return &output
}
func expandIntPtrSlicePtr(input []interface{}) *[]types.IntPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.IntPtr, 0)
	for _, elem := range input {
		output = append(output, types.IntPtr(utils.Int(elem.(int))))
	}
	return &output
}
func expandStrPtrSlicePtr(input []interface{}) *[]types.StrPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.StrPtr, 0)
	for _, elem := range input {
		output = append(output, types.StrPtr(utils.String(elem.(string))))
	}
	return &output
}
func expandFloatPtrSlicePtr(input []interface{}) *[]types.FloatPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.FloatPtr, 0)
	for _, elem := range input {
		output = append(output, types.FloatPtr(utils.Float64(elem.(float64))))
	}
	return &output
}
func flattenTypePrimaryPtrAliasSlicePtrCollection(input types.TypePrimaryPtrAliasSlicePtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_slice_ptr": flattenBoolPtrSlicePtr(input.BoolPtrSlicePtr),
		"flat_ptr_slice_ptr": flattenFloatPtrSlicePtr(input.FlatPtrSlicePtr),
		"int_ptr_slice_ptr":  flattenIntPtrSlicePtr(input.IntPtrSlicePtr),
		"str_ptr_slice_ptr":  flattenStrPtrSlicePtr(input.StrPtrSlicePtr),
	}}
}
func flattenBoolPtrSlicePtr(input *[]types.BoolPtr) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e bool
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenIntPtrSlicePtr(input *[]types.IntPtr) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e int
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenStrPtrSlicePtr(input *[]types.StrPtr) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e string
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}
func flattenFloatPtrSlicePtr(input *[]types.FloatPtr) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		var e float64
		if elem != nil {
			e = *elem
		}
		output = append(output, e)
	}
	return output
}

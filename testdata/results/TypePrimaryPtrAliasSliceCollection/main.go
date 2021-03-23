package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasSliceCollection(input []interface{}) types.TypePrimaryPtrAliasSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasSliceCollection{
		BoolPtrSlice:  expandBoolPtrSlice(b["bool_ptr_slice"].([]interface{})),
		FloatPtrSlice: expandFloatPtrSlice(b["float_ptr_slice"].([]interface{})),
		IntPtrSlice:   expandIntPtrSlice(b["int_ptr_slice"].([]interface{})),
		StrPtrSlice:   expandStrPtrSlice(b["str_ptr_slice"].([]interface{})),
	}
	return output
}
func expandBoolPtrSlice(input []interface{}) []types.BoolPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.BoolPtr, 0)
	for _, elem := range input {
		output = append(output, types.BoolPtr(utils.Bool(elem.(bool))))
	}
	return output
}
func expandIntPtrSlice(input []interface{}) []types.IntPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.IntPtr, 0)
	for _, elem := range input {
		output = append(output, types.IntPtr(utils.Int(elem.(int))))
	}
	return output
}
func expandStrPtrSlice(input []interface{}) []types.StrPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.StrPtr, 0)
	for _, elem := range input {
		output = append(output, types.StrPtr(utils.String(elem.(string))))
	}
	return output
}
func expandFloatPtrSlice(input []interface{}) []types.FloatPtr {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.FloatPtr, 0)
	for _, elem := range input {
		output = append(output, types.FloatPtr(utils.Float64(elem.(float64))))
	}
	return output
}

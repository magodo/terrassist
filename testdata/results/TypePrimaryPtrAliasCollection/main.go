package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasCollection(input []interface{}) types.TypePrimaryPtrAliasCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasCollection{
		BoolPtr:  types.BoolPtr(utils.Bool(b["bool_ptr"].(bool))),
		FloatPtr: types.FloatPtr(utils.Float64(b["float_ptr"].(float64))),
		IntPtr:   types.IntPtr(utils.Int(b["int_ptr"].(int))),
		StrPtr:   types.StrPtr(utils.String(b["str_ptr"].(string))),
	}
	return output
}
func flattenTypePrimaryPtrAliasCollection(input types.TypePrimaryPtrAliasCollection) []interface{} {
	boolPtr := false
	if input.BoolPtr != nil {
		boolPtr = *input.BoolPtr
	}
	intPtr := 0
	if input.IntPtr != nil {
		intPtr = *input.IntPtr
	}
	strPtr := ""
	if input.StrPtr != nil {
		strPtr = *input.StrPtr
	}
	floatPtr := 0.0
	if input.FloatPtr != nil {
		floatPtr = *input.FloatPtr
	}
	return []interface{}{map[string]interface{}{
		"bool_ptr":  boolPtr,
		"float_ptr": floatPtr,
		"int_ptr":   intPtr,
		"str_ptr":   strPtr,
	}}
}

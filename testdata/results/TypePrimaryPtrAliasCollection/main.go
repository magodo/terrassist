package main

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
		BoolPtr:  BoolPtr(utils.Bool(b["bool_ptr"].(bool))),
		FloatPtr: FloatPtr(utils.Float64(b["float_ptr"].(float64))),
		IntPtr:   IntPtr(utils.Int(b["int_ptr"].(int))),
		StrPtr:   StrPtr(utils.String(b["str_ptr"].(string))),
	}
	return output
}

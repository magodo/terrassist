package types

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrCollection(input []interface{}) types.TypePrimaryPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrCollection{
		PtrBool:  utils.Bool(b["ptr_bool"].(bool)),
		PtrFloat: utils.Float64(b["ptr_float"].(float64)),
		PtrInt:   utils.Int(b["ptr_int"].(int)),
		PtrStr:   utils.String(b["ptr_str"].(string)),
	}
	return output
}
func flattenTypePrimaryPtrCollection(input types.TypePrimaryPtrCollection) []interface{} {
	var ptrBool bool
	if input.PtrBool != nil {
		ptrBool = *input.PtrBool
	}
	var ptrInt int
	if input.PtrInt != nil {
		ptrInt = *input.PtrInt
	}
	var ptrStr string
	if input.PtrStr != nil {
		ptrStr = *input.PtrStr
	}
	var ptrFloat float64
	if input.PtrFloat != nil {
		ptrFloat = *input.PtrFloat
	}
	return []interface{}{map[string]interface{}{
		"ptr_bool":  ptrBool,
		"ptr_float": ptrFloat,
		"ptr_int":   ptrInt,
		"ptr_str":   ptrStr,
	}}
}

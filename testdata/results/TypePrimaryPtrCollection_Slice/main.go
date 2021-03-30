package types

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrCollectionSlice(input []interface{}) []types.TypePrimaryPtrCollection {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypePrimaryPtrCollection, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, types.TypePrimaryPtrCollection{
			PtrBool:  utils.Bool(elem["ptr_bool"].(bool)),
			PtrFloat: utils.Float64(elem["ptr_float"].(float64)),
			PtrInt:   utils.Int(elem["ptr_int"].(int)),
			PtrStr:   utils.String(elem["ptr_str"].(string)),
		})
	}
	return output
}
func flattenTypePrimaryPtrCollectionSlice(input []types.TypePrimaryPtrCollection) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		var ptrBool bool
		if elem.PtrBool != nil {
			ptrBool = *elem.PtrBool
		}
		var ptrInt int
		if elem.PtrInt != nil {
			ptrInt = *elem.PtrInt
		}
		var ptrStr string
		if elem.PtrStr != nil {
			ptrStr = *elem.PtrStr
		}
		var ptrFloat float64
		if elem.PtrFloat != nil {
			ptrFloat = *elem.PtrFloat
		}
		output = append(output, map[string]interface{}{
			"ptr_bool":  ptrBool,
			"ptr_float": ptrFloat,
			"ptr_int":   ptrInt,
			"ptr_str":   ptrStr,
		})
	}
	return output
}

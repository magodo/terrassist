package types

import types "types"

func expandTypePrimaryAliasPtrCollection(input []interface{}) types.TypePrimaryAliasPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	ptrBool := types.Bool(b["ptr_bool"].(bool))
	ptrInt := types.Int(b["ptr_int"].(int))
	ptrStr := types.Str(b["ptr_str"].(string))
	ptrFloat := types.Float(b["ptr_float"].(float64))
	output := types.TypePrimaryAliasPtrCollection{
		PtrBool:  &ptrBool,
		PtrFloat: &ptrFloat,
		PtrInt:   &ptrInt,
		PtrStr:   &ptrStr,
	}
	return output
}
func flattenTypePrimaryAliasPtrCollection(input types.TypePrimaryAliasPtrCollection) []interface{} {
	var ptrBool bool
	if input.PtrBool != nil {
		ptrBool = bool(*input.PtrBool)
	}
	var ptrInt int
	if input.PtrInt != nil {
		ptrInt = int(*input.PtrInt)
	}
	var ptrStr string
	if input.PtrStr != nil {
		ptrStr = string(*input.PtrStr)
	}
	var ptrFloat float64
	if input.PtrFloat != nil {
		ptrFloat = float64(*input.PtrFloat)
	}
	return []interface{}{map[string]interface{}{
		"ptr_bool":  ptrBool,
		"ptr_float": ptrFloat,
		"ptr_int":   ptrInt,
		"ptr_str":   ptrStr,
	}}
}

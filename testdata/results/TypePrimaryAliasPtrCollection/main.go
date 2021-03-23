package output

import types "types"

func expandTypePrimaryAliasPtrCollection(input []interface{}) types.TypePrimaryAliasPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	ptrBool := b["ptr_bool"].(types.Bool)
	ptrInt := b["ptr_int"].(types.Int)
	ptrStr := b["ptr_str"].(types.Str)
	ptrFloat := b["ptr_float"].(types.Float)
	output := types.TypePrimaryAliasPtrCollection{
		PtrBool:  &ptrBool,
		PtrFloat: &ptrFloat,
		PtrInt:   &ptrInt,
		PtrStr:   &ptrStr,
	}
	return output
}
func flattenTypePrimaryAliasPtrCollection(input types.TypePrimaryAliasPtrCollection) []interface{} {
	ptrBool := false
	if input.PtrBool != nil {
		ptrBool = bool(*input.PtrBool)
	}
	ptrInt := 0
	if input.PtrInt != nil {
		ptrInt = int(*input.PtrInt)
	}
	ptrStr := ""
	if input.PtrStr != nil {
		ptrStr = string(*input.PtrStr)
	}
	ptrFloat := 0.0
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

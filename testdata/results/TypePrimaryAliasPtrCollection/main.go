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
	var ptrBool types.Bool
	if input.PtrBool != nil {
		ptrBool = *input.PtrBool
	}
	var ptrInt types.Int
	if input.PtrInt != nil {
		ptrInt = *input.PtrInt
	}
	var ptrStr types.Str
	if input.PtrStr != nil {
		ptrStr = *input.PtrStr
	}
	var ptrFloat types.Float
	if input.PtrFloat != nil {
		ptrFloat = *input.PtrFloat
	}
	return []interface{}{map[string]interface{}{
		"ptr_bool":  bool(ptrBool),
		"ptr_float": float64(ptrFloat),
		"ptr_int":   int(ptrInt),
		"ptr_str":   string(ptrStr),
	}}
}

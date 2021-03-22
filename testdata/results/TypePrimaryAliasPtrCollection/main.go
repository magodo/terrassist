package main

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

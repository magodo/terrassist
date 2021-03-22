package main

import types "types"

func expandTypePrimaryCollection(input []interface{}) types.TypePrimaryCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryCollection{
		Bool:  b["bool"].(bool),
		Float: b["float"].(float64),
		Int:   b["int"].(int),
		Str:   b["str"].(string),
	}
	return output
}

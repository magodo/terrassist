package types

import types "types"

func expandTypeNamedStructCollection(input []interface{}) types.TypeNamedStructCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructCollection{
		S1: expandTypeS1(b["s_1"].([]interface{})),
		S2: expandTypeS2(b["s_2"].([]interface{})),
	}
	return output
}
func expandTypeS1(input []interface{}) types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return types.TypeS1{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS2(input []interface{}) types.TypeS2 {
	if len(input) == 0 || input[0] == nil {
		return types.TypeS2{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeS2{S: b["s"].(string)}
	return output
}
func flattenTypeNamedStructCollection(input types.TypeNamedStructCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"s_1": flattenTypeS1(input.S1),
		"s_2": flattenTypeS2(input.S2),
	}}
}
func flattenTypeS1(input types.TypeS1) []interface{} {
	return []interface{}{map[string]interface{}{"i": input.I}}
}
func flattenTypeS2(input types.TypeS2) []interface{} {
	return []interface{}{map[string]interface{}{"s": input.S}}
}

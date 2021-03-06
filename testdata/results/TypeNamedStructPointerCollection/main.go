package types

import types "types"

func expandTypeNamedStructPointerCollection(input []interface{}) types.TypeNamedStructPointerCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructPointerCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructPointerCollection{
		S1: expandTypeS1Ptr(b["s_1"].([]interface{})),
		S2: expandTypeS2Ptr(b["s_2"].([]interface{})),
	}
	return output
}
func expandTypeS1Ptr(input []interface{}) *types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS2Ptr(input []interface{}) *types.TypeS2 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS2{S: b["s"].(string)}
	return output
}
func flattenTypeNamedStructPointerCollection(input types.TypeNamedStructPointerCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"s_1": flattenTypeS1Ptr(input.S1),
		"s_2": flattenTypeS2Ptr(input.S2),
	}}
}
func flattenTypeS1Ptr(input *types.TypeS1) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{"i": input.I}}
}
func flattenTypeS2Ptr(input *types.TypeS2) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{"s": input.S}}
}

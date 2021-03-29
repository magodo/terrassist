package types

import types "types"

func expandTypeNamedInterfaceMap(input map[string]interface{}) types.TypeNamedInterfaceMap {
	output := make(types.TypeNamedInterfaceMap)
	for k, v := range input {
		output[k] = expandTypeNamedInterface(v.([]interface{}))
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
func expandTypeS2Ptr(input []interface{}) *types.TypeS2 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS2{S: b["s"].(string)}
	return output
}
func expandTypeNamedInterface(input []interface{}) types.TypeNamedInterface {
	// TODO
	return nil
}
func flattenTypeNamedInterfaceMap(input types.TypeNamedInterfaceMap) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = flattenTypeNamedInterface(v)
	}
	return output
}
func flattenTypeS1(input types.TypeS1) []interface{} {
	return []interface{}{map[string]interface{}{"i": input.I}}
}
func flattenTypeS2Ptr(input *types.TypeS2) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{"s": input.S}}
}
func flattenTypeNamedInterface(input types.TypeNamedInterface) []interface{} {
	// TODO
	return nil
}

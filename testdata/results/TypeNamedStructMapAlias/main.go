package types

import types "types"

func expandTypeNamedStructMapAlias(input map[string]interface{}) types.TypeNamedStructMapAlias {
	output := make(types.TypeNamedStructMapAlias)
	for k, v := range input {
		output[k] = expandTypeS1(v.([]interface{}))
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
func flattenTypeNamedStructMapAlias(input types.TypeNamedStructMapAlias) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = flattenTypeS1(v)
	}
	return output
}
func flattenTypeS1(input types.TypeS1) []interface{} {
	return []interface{}{map[string]interface{}{"i": input.I}}
}

package output

import types "types"

func expandTypeNamedStructWithJSONIgnore(input []interface{}) types.TypeNamedStructWithJSONIgnore {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructWithJSONIgnore{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructWithJSONIgnore{
		I: b["i"].(int),
		J: b["j"].(string),
	}
	return output
}
func flattenTypeNamedStructWithJSONIgnore(input types.TypeNamedStructWithJSONIgnore) []interface{} {
	return []interface{}{map[string]interface{}{
		"i": input.I,
		"j": input.J,
	}}
}

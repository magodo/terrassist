package output

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

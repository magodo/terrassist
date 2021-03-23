package output

import types "types"

func expandTypeNamedStructSliceAlias(input []interface{}) types.TypeNamedStructSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeNamedStructSliceAlias, 0)
	for _, elem := range input {
		output = append(output, expandTypeS1(elem.([]interface{})))
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

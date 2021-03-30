package types

import types "types"

func expandTypeS1Slice(input []interface{}) []types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypeS1, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, types.TypeS1{I: elem["i"].(int)})
	}
	return output
}
func flattenTypeS1Slice(input []types.TypeS1) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}

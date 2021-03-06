package types

import types "types"

func expandTypeNamedStructSliceAlias(input []interface{}) types.TypeNamedStructSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeNamedStructSliceAlias, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, types.TypeS1{I: elem["i"].(int)})
	}
	return output
}
func flattenTypeNamedStructSliceAliasSlice(input types.TypeNamedStructSliceAlias) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}

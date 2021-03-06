package types

import types "types"

func expandTypeNamedBoolSliceAlias(input []interface{}) types.TypeNamedBoolSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeNamedBoolSliceAlias, 0)
	for _, elem := range input {
		output = append(output, types.Bool(elem.(bool)))
	}
	return output
}
func flattenTypeNamedBoolSliceAliasSlice(input types.TypeNamedBoolSliceAlias) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, bool(elem))
	}
	return output
}

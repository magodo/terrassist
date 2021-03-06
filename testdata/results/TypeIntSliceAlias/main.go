package types

import types "types"

func expandTypeIntSliceAlias(input []interface{}) types.TypeIntSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeIntSliceAlias, 0)
	for _, elem := range input {
		output = append(output, elem.(int))
	}
	return output
}
func flattenTypeIntSliceAliasSlice(input types.TypeIntSliceAlias) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, elem)
	}
	return output
}

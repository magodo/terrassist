package types

import types "types"

func expandTypeIntMapAlias(input map[string]interface{}) types.TypeIntMapAlias {
	output := make(types.TypeIntMapAlias)
	for k, v := range input {
		output[k] = v.(int)
	}
	return output
}
func flattenTypeIntMapAlias(input types.TypeIntMapAlias) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}
	return output
}

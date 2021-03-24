package types

import types "types"

func expandTypeNamedBoolMapAlias(input map[string]interface{}) types.TypeNamedBoolMapAlias {
	output := make(types.TypeNamedBoolMapAlias)
	for k, v := range input {
		output[k] = v.(types.Bool)
	}
	return output
}
func flattenTypeNamedBoolMapAlias(input types.TypeNamedBoolMapAlias) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = bool(v)
	}
	return output
}

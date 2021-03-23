package output

import types "types"

func expandTypeNamedBoolMapAlias(input map[string]interface{}) types.TypeNamedBoolMapAlias {
	output := make(types.TypeNamedBoolMapAlias)
	for k, v := range input {
		output[k] = v.(types.Bool)
	}
	return output
}

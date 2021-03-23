package main

import types "types"

func expandTypeIntMapAlias(input map[string]interface{}) types.TypeIntMapAlias {
	output := make(types.TypeIntMapAlias)
	for k, v := range input {
		output[k] = v.(int)
	}
	return output
}

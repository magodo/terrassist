package output

import (
	types "types"
	utils "types/utils"
)

func expandTypeIntPtrMapAlias(input map[string]interface{}) types.TypeIntPtrMapAlias {
	output := make(types.TypeIntPtrMapAlias)
	for k, v := range input {
		output[k] = utils.Int(v.(int))
	}
	return output
}

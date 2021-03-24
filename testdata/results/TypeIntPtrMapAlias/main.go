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
func flattenTypeIntPtrMapAlias(input types.TypeIntPtrMapAlias) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		e := 0
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}

package main

import (
	types "types"
	utils "types/utils"
)

func expandTypeIntPtrSliceAlias(input []interface{}) types.TypeIntPtrSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeIntPtrSliceAlias, 0)
	for _, elem := range input {
		output = append(output, utils.Int(elem.(int)))
	}
	return output
}

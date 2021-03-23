package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrCollection(input []interface{}) types.TypePrimaryPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrCollection{
		PtrBool:  utils.Bool(b["ptr_bool"].(bool)),
		PtrFloat: utils.Float64(b["ptr_float"].(float64)),
		PtrInt:   utils.Int(b["ptr_int"].(int)),
		PtrStr:   utils.String(b["ptr_str"].(string)),
	}
	return output
}

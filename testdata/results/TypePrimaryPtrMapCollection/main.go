package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrMapCollection(input []interface{}) types.TypePrimaryPtrMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrMapCollection{
		BoolPtrMap:  expandBoolPtrMap(b["bool_ptr_map"].(map[string]interface{})),
		FloatPtrMap: expandFloat64PtrMap(b["float_ptr_map"].(map[string]interface{})),
		IntPtrMap:   expandIntPtrMap(b["int_ptr_map"].(map[string]interface{})),
		StrPtrMap:   expandStringPtrMap(b["str_ptr_map"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMap(input map[string]interface{}) map[string]*bool {
	output := make(map[string]*bool)
	for k, v := range input {
		output[k] = utils.Bool(v.(bool))
	}
	return output
}
func expandIntPtrMap(input map[string]interface{}) map[string]*int {
	output := make(map[string]*int)
	for k, v := range input {
		output[k] = utils.Int(v.(int))
	}
	return output
}
func expandStringPtrMap(input map[string]interface{}) map[string]*string {
	output := make(map[string]*string)
	for k, v := range input {
		output[k] = utils.String(v.(string))
	}
	return output
}
func expandFloat64PtrMap(input map[string]interface{}) map[string]*float64 {
	output := make(map[string]*float64)
	for k, v := range input {
		output[k] = utils.Float64(v.(float64))
	}
	return output
}

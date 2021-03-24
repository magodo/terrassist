package types

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasMapPtrCollection(input []interface{}) types.TypePrimaryPtrAliasMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasMapPtrCollection{
		BoolPtrMapPtr: expandBoolPtrMapPtr(b["bool_ptr_map_ptr"].(map[string]interface{})),
		FlatPtrMapPtr: expandFloatPtrMapPtr(b["flat_ptr_map_ptr"].(map[string]interface{})),
		IntPtrMapPtr:  expandIntPtrMapPtr(b["int_ptr_map_ptr"].(map[string]interface{})),
		StrPtrMapPtr:  expandStrPtrMapPtr(b["str_ptr_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMapPtr(input map[string]interface{}) *map[string]types.BoolPtr {
	output := make(map[string]types.BoolPtr)
	for k, v := range input {
		output[k] = types.BoolPtr(utils.Bool(v.(bool)))
	}
	return &output
}
func expandIntPtrMapPtr(input map[string]interface{}) *map[string]types.IntPtr {
	output := make(map[string]types.IntPtr)
	for k, v := range input {
		output[k] = types.IntPtr(utils.Int(v.(int)))
	}
	return &output
}
func expandStrPtrMapPtr(input map[string]interface{}) *map[string]types.StrPtr {
	output := make(map[string]types.StrPtr)
	for k, v := range input {
		output[k] = types.StrPtr(utils.String(v.(string)))
	}
	return &output
}
func expandFloatPtrMapPtr(input map[string]interface{}) *map[string]types.FloatPtr {
	output := make(map[string]types.FloatPtr)
	for k, v := range input {
		output[k] = types.FloatPtr(utils.Float64(v.(float64)))
	}
	return &output
}
func flattenTypePrimaryPtrAliasMapPtrCollection(input types.TypePrimaryPtrAliasMapPtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_map_ptr": flattenBoolPtrMapPtr(input.BoolPtrMapPtr),
		"flat_ptr_map_ptr": flattenFloatPtrMapPtr(input.FlatPtrMapPtr),
		"int_ptr_map_ptr":  flattenIntPtrMapPtr(input.IntPtrMapPtr),
		"str_ptr_map_ptr":  flattenStrPtrMapPtr(input.StrPtrMapPtr),
	}}
}
func flattenBoolPtrMapPtr(input *map[string]types.BoolPtr) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e bool
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}
func flattenIntPtrMapPtr(input *map[string]types.IntPtr) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e int
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}
func flattenStrPtrMapPtr(input *map[string]types.StrPtr) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e string
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}
func flattenFloatPtrMapPtr(input *map[string]types.FloatPtr) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e float64
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}

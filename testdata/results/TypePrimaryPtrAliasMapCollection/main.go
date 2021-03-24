package output

import (
	types "types"
	utils "types/utils"
)

func expandTypePrimaryPtrAliasMapCollection(input []interface{}) types.TypePrimaryPtrAliasMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryPtrAliasMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryPtrAliasMapCollection{
		BoolPtrMap:  expandBoolPtrMap(b["bool_ptr_map"].(map[string]interface{})),
		FloatPtrMap: expandFloatPtrMap(b["float_ptr_map"].(map[string]interface{})),
		IntPtrMap:   expandIntPtrMap(b["int_ptr_map"].(map[string]interface{})),
		StrPtrMap:   expandStrPtrMap(b["str_ptr_map"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMap(input map[string]interface{}) map[string]types.BoolPtr {
	output := make(map[string]types.BoolPtr)
	for k, v := range input {
		output[k] = types.BoolPtr(utils.Bool(v.(bool)))
	}
	return output
}
func expandIntPtrMap(input map[string]interface{}) map[string]types.IntPtr {
	output := make(map[string]types.IntPtr)
	for k, v := range input {
		output[k] = types.IntPtr(utils.Int(v.(int)))
	}
	return output
}
func expandStrPtrMap(input map[string]interface{}) map[string]types.StrPtr {
	output := make(map[string]types.StrPtr)
	for k, v := range input {
		output[k] = types.StrPtr(utils.String(v.(string)))
	}
	return output
}
func expandFloatPtrMap(input map[string]interface{}) map[string]types.FloatPtr {
	output := make(map[string]types.FloatPtr)
	for k, v := range input {
		output[k] = types.FloatPtr(utils.Float64(v.(float64)))
	}
	return output
}
func flattenTypePrimaryPtrAliasMapCollection(input types.TypePrimaryPtrAliasMapCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_map":  flattenBoolPtrMap(input.BoolPtrMap),
		"float_ptr_map": flattenFloatPtrMap(input.FloatPtrMap),
		"int_ptr_map":   flattenIntPtrMap(input.IntPtrMap),
		"str_ptr_map":   flattenStrPtrMap(input.StrPtrMap),
	}}
}
func flattenBoolPtrMap(input map[string]types.BoolPtr) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		e := false
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}
func flattenIntPtrMap(input map[string]types.IntPtr) map[string]interface{} {
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
func flattenStrPtrMap(input map[string]types.StrPtr) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		e := ""
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}
func flattenFloatPtrMap(input map[string]types.FloatPtr) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		e := 0.0
		if v != nil {
			e = *v
		}
		output[k] = e
	}
	return output
}

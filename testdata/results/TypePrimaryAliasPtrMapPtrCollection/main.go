package types

import types "types"

func expandTypePrimaryAliasPtrMapPtrCollection(input []interface{}) types.TypePrimaryAliasPtrMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasPtrMapPtrCollection{
		BoolPtrMapPtr:   expandBoolPtrMapPtr(b["bool_ptr_map_ptr"].(map[string]interface{})),
		FloatPtrMapPtr:  expandFloatPtrMapPtr(b["float_ptr_map_ptr"].(map[string]interface{})),
		IntPtrMapPtr:    expandIntPtrMapPtr(b["int_ptr_map_ptr"].(map[string]interface{})),
		StringPtrMapPtr: expandStrPtrMapPtr(b["string_ptr_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMapPtr(input map[string]interface{}) *map[string]*types.Bool {
	output := make(map[string]*types.Bool)
	for k, v := range input {
		e := types.Bool(v.(bool))
		output[k] = &e
	}
	return &output
}
func expandIntPtrMapPtr(input map[string]interface{}) *map[string]*types.Int {
	output := make(map[string]*types.Int)
	for k, v := range input {
		e := types.Int(v.(int))
		output[k] = &e
	}
	return &output
}
func expandStrPtrMapPtr(input map[string]interface{}) *map[string]*types.Str {
	output := make(map[string]*types.Str)
	for k, v := range input {
		e := types.Str(v.(string))
		output[k] = &e
	}
	return &output
}
func expandFloatPtrMapPtr(input map[string]interface{}) *map[string]*types.Float {
	output := make(map[string]*types.Float)
	for k, v := range input {
		e := types.Float(v.(float64))
		output[k] = &e
	}
	return &output
}
func flattenTypePrimaryAliasPtrMapPtrCollection(input types.TypePrimaryAliasPtrMapPtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_map_ptr":   flattenBoolPtrMapPtr(input.BoolPtrMapPtr),
		"float_ptr_map_ptr":  flattenFloatPtrMapPtr(input.FloatPtrMapPtr),
		"int_ptr_map_ptr":    flattenIntPtrMapPtr(input.IntPtrMapPtr),
		"string_ptr_map_ptr": flattenStrPtrMapPtr(input.StringPtrMapPtr),
	}}
}
func flattenBoolPtrMapPtr(input *map[string]*types.Bool) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e types.Bool
		if v != nil {
			e = *v
		}
		output[k] = bool(e)
	}
	return output
}
func flattenIntPtrMapPtr(input *map[string]*types.Int) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e types.Int
		if v != nil {
			e = *v
		}
		output[k] = int(e)
	}
	return output
}
func flattenStrPtrMapPtr(input *map[string]*types.Str) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e types.Str
		if v != nil {
			e = *v
		}
		output[k] = string(e)
	}
	return output
}
func flattenFloatPtrMapPtr(input *map[string]*types.Float) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		var e types.Float
		if v != nil {
			e = *v
		}
		output[k] = float64(e)
	}
	return output
}

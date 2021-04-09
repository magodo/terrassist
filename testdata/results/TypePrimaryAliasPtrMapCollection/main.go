package types

import types "types"

func expandTypePrimaryAliasPtrMapCollection(input []interface{}) types.TypePrimaryAliasPtrMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasPtrMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasPtrMapCollection{
		BoolPtrMap:   expandBoolPtrMap(b["bool_ptr_map"].(map[string]interface{})),
		FloatPtrMap:  expandFloatPtrMap(b["float_ptr_map"].(map[string]interface{})),
		IntPtrMap:    expandIntPtrMap(b["int_ptr_map"].(map[string]interface{})),
		StringPtrMap: expandStrPtrMap(b["string_ptr_map"].(map[string]interface{})),
	}
	return output
}
func expandBoolPtrMap(input map[string]interface{}) map[string]*types.Bool {
	output := make(map[string]*types.Bool)
	for k, v := range input {
		e := types.Bool(v.(bool))
		output[k] = &e
	}
	return output
}
func expandIntPtrMap(input map[string]interface{}) map[string]*types.Int {
	output := make(map[string]*types.Int)
	for k, v := range input {
		e := types.Int(v.(int))
		output[k] = &e
	}
	return output
}
func expandStrPtrMap(input map[string]interface{}) map[string]*types.Str {
	output := make(map[string]*types.Str)
	for k, v := range input {
		e := types.Str(v.(string))
		output[k] = &e
	}
	return output
}
func expandFloatPtrMap(input map[string]interface{}) map[string]*types.Float {
	output := make(map[string]*types.Float)
	for k, v := range input {
		e := types.Float(v.(float64))
		output[k] = &e
	}
	return output
}
func flattenTypePrimaryAliasPtrMapCollection(input types.TypePrimaryAliasPtrMapCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_ptr_map":   flattenBoolPtrMap(input.BoolPtrMap),
		"float_ptr_map":  flattenFloatPtrMap(input.FloatPtrMap),
		"int_ptr_map":    flattenIntPtrMap(input.IntPtrMap),
		"string_ptr_map": flattenStrPtrMap(input.StringPtrMap),
	}}
}
func flattenBoolPtrMap(input map[string]*types.Bool) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e types.Bool
		if v != nil {
			e = *v
		}
		output[k] = bool(e)
	}
	return output
}
func flattenIntPtrMap(input map[string]*types.Int) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e types.Int
		if v != nil {
			e = *v
		}
		output[k] = int(e)
	}
	return output
}
func flattenStrPtrMap(input map[string]*types.Str) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e types.Str
		if v != nil {
			e = *v
		}
		output[k] = string(e)
	}
	return output
}
func flattenFloatPtrMap(input map[string]*types.Float) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e types.Float
		if v != nil {
			e = *v
		}
		output[k] = float64(e)
	}
	return output
}

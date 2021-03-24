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
		e := v.(types.Bool)
		output[k] = &e
	}
	return output
}
func expandIntPtrMap(input map[string]interface{}) map[string]*types.Int {
	output := make(map[string]*types.Int)
	for k, v := range input {
		e := v.(types.Int)
		output[k] = &e
	}
	return output
}
func expandStrPtrMap(input map[string]interface{}) map[string]*types.Str {
	output := make(map[string]*types.Str)
	for k, v := range input {
		e := v.(types.Str)
		output[k] = &e
	}
	return output
}
func expandFloatPtrMap(input map[string]interface{}) map[string]*types.Float {
	output := make(map[string]*types.Float)
	for k, v := range input {
		e := v.(types.Float)
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
		var e bool
		if v != nil {
			e = bool(*v)
		}
		output[k] = e
	}
	return output
}
func flattenIntPtrMap(input map[string]*types.Int) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e int
		if v != nil {
			e = int(*v)
		}
		output[k] = e
	}
	return output
}
func flattenStrPtrMap(input map[string]*types.Str) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e string
		if v != nil {
			e = string(*v)
		}
		output[k] = e
	}
	return output
}
func flattenFloatPtrMap(input map[string]*types.Float) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		var e float64
		if v != nil {
			e = float64(*v)
		}
		output[k] = e
	}
	return output
}

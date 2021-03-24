package types

import types "types"

func expandTypePrimaryAliasMapPtrCollection(input []interface{}) types.TypePrimaryAliasMapPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasMapPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasMapPtrCollection{
		BoolMapPtr:   expandBoolMapPtr(b["bool_map_ptr"].(map[string]interface{})),
		FloatMapPtr:  expandFloatMapPtr(b["float_map_ptr"].(map[string]interface{})),
		IntMapPtr:    expandIntMapPtr(b["int_map_ptr"].(map[string]interface{})),
		StringMapPtr: expandStrMapPtr(b["string_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandBoolMapPtr(input map[string]interface{}) *map[string]types.Bool {
	output := make(map[string]types.Bool)
	for k, v := range input {
		output[k] = v.(types.Bool)
	}
	return &output
}
func expandIntMapPtr(input map[string]interface{}) *map[string]types.Int {
	output := make(map[string]types.Int)
	for k, v := range input {
		output[k] = v.(types.Int)
	}
	return &output
}
func expandStrMapPtr(input map[string]interface{}) *map[string]types.Str {
	output := make(map[string]types.Str)
	for k, v := range input {
		output[k] = v.(types.Str)
	}
	return &output
}
func expandFloatMapPtr(input map[string]interface{}) *map[string]types.Float {
	output := make(map[string]types.Float)
	for k, v := range input {
		output[k] = v.(types.Float)
	}
	return &output
}
func flattenTypePrimaryAliasMapPtrCollection(input types.TypePrimaryAliasMapPtrCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_map_ptr":   flattenBoolMapPtr(input.BoolMapPtr),
		"float_map_ptr":  flattenFloatMapPtr(input.FloatMapPtr),
		"int_map_ptr":    flattenIntMapPtr(input.IntMapPtr),
		"string_map_ptr": flattenStrMapPtr(input.StringMapPtr),
	}}
}
func flattenBoolMapPtr(input *map[string]types.Bool) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = bool(v)
	}
	return output
}
func flattenIntMapPtr(input *map[string]types.Int) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = int(v)
	}
	return output
}
func flattenStrMapPtr(input *map[string]types.Str) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = string(v)
	}
	return output
}
func flattenFloatMapPtr(input *map[string]types.Float) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = float64(v)
	}
	return output
}

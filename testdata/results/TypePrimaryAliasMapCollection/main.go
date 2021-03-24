package output

import types "types"

func expandTypePrimaryAliasMapCollection(input []interface{}) types.TypePrimaryAliasMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasMapCollection{
		BoolMap:  expandBoolMap(b["bool_map"].(map[string]interface{})),
		FloatMap: expandFloatMap(b["float_map"].(map[string]interface{})),
		IntMap:   expandIntMap(b["int_map"].(map[string]interface{})),
		StrMap:   expandStrMap(b["str_map"].(map[string]interface{})),
	}
	return output
}
func expandBoolMap(input map[string]interface{}) map[string]types.Bool {
	output := make(map[string]types.Bool)
	for k, v := range input {
		output[k] = v.(types.Bool)
	}
	return output
}
func expandIntMap(input map[string]interface{}) map[string]types.Int {
	output := make(map[string]types.Int)
	for k, v := range input {
		output[k] = v.(types.Int)
	}
	return output
}
func expandStrMap(input map[string]interface{}) map[string]types.Str {
	output := make(map[string]types.Str)
	for k, v := range input {
		output[k] = v.(types.Str)
	}
	return output
}
func expandFloatMap(input map[string]interface{}) map[string]types.Float {
	output := make(map[string]types.Float)
	for k, v := range input {
		output[k] = v.(types.Float)
	}
	return output
}
func flattenTypePrimaryAliasMapCollection(input types.TypePrimaryAliasMapCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool_map":  flattenBoolMap(input.BoolMap),
		"float_map": flattenFloatMap(input.FloatMap),
		"int_map":   flattenIntMap(input.IntMap),
		"str_map":   flattenStrMap(input.StrMap),
	}}
}
func flattenBoolMap(input map[string]types.Bool) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = bool(v)
	}
	return output
}
func flattenIntMap(input map[string]types.Int) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = int(v)
	}
	return output
}
func flattenStrMap(input map[string]types.Str) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = string(v)
	}
	return output
}
func flattenFloatMap(input map[string]types.Float) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = float64(v)
	}
	return output
}

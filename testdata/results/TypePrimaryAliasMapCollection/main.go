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

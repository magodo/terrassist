package output

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

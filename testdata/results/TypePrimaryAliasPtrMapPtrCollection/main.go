package output

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
		e := v.(types.Bool)
		output[k] = &e
	}
	return &output
}
func expandIntPtrMapPtr(input map[string]interface{}) *map[string]*types.Int {
	output := make(map[string]*types.Int)
	for k, v := range input {
		e := v.(types.Int)
		output[k] = &e
	}
	return &output
}
func expandStrPtrMapPtr(input map[string]interface{}) *map[string]*types.Str {
	output := make(map[string]*types.Str)
	for k, v := range input {
		e := v.(types.Str)
		output[k] = &e
	}
	return &output
}
func expandFloatPtrMapPtr(input map[string]interface{}) *map[string]*types.Float {
	output := make(map[string]*types.Float)
	for k, v := range input {
		e := v.(types.Float)
		output[k] = &e
	}
	return &output
}

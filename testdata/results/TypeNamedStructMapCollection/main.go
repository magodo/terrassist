package types

import types "types"

func expandTypeNamedStructMapCollection(input []interface{}) types.TypeNamedStructMapCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructMapCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructMapCollection{
		S1Map:       expandTypeS1Map(b["s_1_map"].(map[string]interface{})),
		S1MapPtr:    expandTypeS1MapPtr(b["s_1_map_ptr"].(map[string]interface{})),
		S1PtrMap:    expandTypeS1PtrMap(b["s_1_ptr_map"].(map[string]interface{})),
		S1PtrMapPtr: expandTypeS1PtrMapPtr(b["s_1_ptr_map_ptr"].(map[string]interface{})),
	}
	return output
}
func expandTypeS1Map(input map[string]interface{}) map[string]types.TypeS1 {
	output := make(map[string]types.TypeS1)
	for k, v := range input {
		output[k] = expandTypeS1(v.([]interface{}))
	}
	return output
}
func expandTypeS1(input []interface{}) types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return types.TypeS1{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS1MapPtr(input map[string]interface{}) *map[string]types.TypeS1 {
	output := make(map[string]types.TypeS1)
	for k, v := range input {
		output[k] = expandTypeS1(v.([]interface{}))
	}
	return &output
}
func expandTypeS1PtrMap(input map[string]interface{}) map[string]*types.TypeS1 {
	output := make(map[string]*types.TypeS1)
	for k, v := range input {
		output[k] = expandTypeS1Ptr(v.([]interface{}))
	}
	return output
}
func expandTypeS1Ptr(input []interface{}) *types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS1{I: b["i"].(int)}
	return output
}
func expandTypeS1PtrMapPtr(input map[string]interface{}) *map[string]*types.TypeS1 {
	output := make(map[string]*types.TypeS1)
	for k, v := range input {
		output[k] = expandTypeS1Ptr(v.([]interface{}))
	}
	return &output
}
func flattenTypeNamedStructMapCollection(input types.TypeNamedStructMapCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"s_1_map":         flattenTypeS1Map(input.S1Map),
		"s_1_map_ptr":     flattenTypeS1MapPtr(input.S1MapPtr),
		"s_1_ptr_map":     flattenTypeS1PtrMap(input.S1PtrMap),
		"s_1_ptr_map_ptr": flattenTypeS1PtrMapPtr(input.S1PtrMapPtr),
	}}
}
func flattenTypeS1Map(input map[string]types.TypeS1) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = flattenTypeS1(v)
	}
	return output
}
func flattenTypeS1(input types.TypeS1) []interface{} {
	return []interface{}{map[string]interface{}{"i": input.I}}
}
func flattenTypeS1MapPtr(input *map[string]types.TypeS1) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = flattenTypeS1(v)
	}
	return output
}
func flattenTypeS1PtrMap(input map[string]*types.TypeS1) map[string]interface{} {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = flattenTypeS1Ptr(v)
	}
	return output
}
func flattenTypeS1Ptr(input *types.TypeS1) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{"i": input.I}}
}
func flattenTypeS1PtrMapPtr(input *map[string]*types.TypeS1) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	output := make(map[string]interface{})
	for k, v := range *input {
		output[k] = flattenTypeS1Ptr(v)
	}
	return output
}

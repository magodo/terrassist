package types

import types "types"

func expandTypeNamedStructSliceCollection(input []interface{}) types.TypeNamedStructSliceCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructSliceCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructSliceCollection{
		S1PtrSlice:    expandTypeS1PtrSlice(b["s_1_ptr_slice"].([]interface{})),
		S1PtrSlicePtr: expandTypeS1PtrSlicePtr(b["s_1_ptr_slice_ptr"].([]interface{})),
		S1Slice:       expandTypeS1Slice(b["s_1_slice"].([]interface{})),
		S1SlicePtr:    expandTypeS1SlicePtr(b["s_1_slice_ptr"].([]interface{})),
	}
	return output
}
func expandTypeS1Slice(input []interface{}) []types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypeS1, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, types.TypeS1{I: elem["i"].(int)})
	}
	return output
}
func expandTypeS1SlicePtr(input []interface{}) *[]types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]types.TypeS1, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, types.TypeS1{I: elem["i"].(int)})
	}
	return &output
}
func expandTypeS1PtrSlice(input []interface{}) []*types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.TypeS1, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, &types.TypeS1{I: elem["i"].(int)})
	}
	return output
}
func expandTypeS1PtrSlicePtr(input []interface{}) *[]*types.TypeS1 {
	if len(input) == 0 {
		return nil
	}
	output := make([]*types.TypeS1, 0)
	for _, elem := range input {
		elem := elem.(map[string]interface{})
		output = append(output, &types.TypeS1{I: elem["i"].(int)})
	}
	return &output
}
func flattenTypeNamedStructSliceCollection(input types.TypeNamedStructSliceCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"s_1_ptr_slice":     flattenTypeS1PtrSlice(input.S1PtrSlice),
		"s_1_ptr_slice_ptr": flattenTypeS1PtrSlicePtr(input.S1PtrSlicePtr),
		"s_1_slice":         flattenTypeS1Slice(input.S1Slice),
		"s_1_slice_ptr":     flattenTypeS1SlicePtr(input.S1SlicePtr),
	}}
}
func flattenTypeS1Slice(input []types.TypeS1) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}
func flattenTypeS1SlicePtr(input *[]types.TypeS1) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}
func flattenTypeS1PtrSlice(input []*types.TypeS1) []interface{} {
	output := make([]interface{}, 0)
	for _, elem := range input {
		if input == nil {
			continue
		}
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}
func flattenTypeS1PtrSlicePtr(input *[]*types.TypeS1) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	output := make([]interface{}, 0)
	for _, elem := range *input {
		if input == nil {
			continue
		}
		output = append(output, map[string]interface{}{"i": elem.I})
	}
	return output
}

package types

import types "types"

func expandTypeCyclicRefStructPtr(input []interface{}) *types.TypeCyclicRefStruct {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeCyclicRefStruct{Self: expandTypeCyclicRefStructPtr(b["self"].([]interface{}))}
	return output
}
func flattenTypeCyclicRefStructPtr(input *types.TypeCyclicRefStruct) []interface{} {
	if input == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{"self": flattenTypeCyclicRefStructPtr(input.Self)}}
}

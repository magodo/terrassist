package output

import types "types"

func expandTypeS1Ptr(input []interface{}) *types.TypeS1 {
	if len(input) == 0 || input[0] == nil {
		return nil
	}
	b := input[0].(map[string]interface{})
	output := &types.TypeS1{I: b["i"].(int)}
	return output
}

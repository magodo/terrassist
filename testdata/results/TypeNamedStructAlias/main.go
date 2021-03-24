package types

import types "types"

func expandTypeNamedStructAlias(input []interface{}) types.TypeNamedStructAlias {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructAlias{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructAlias{I: b["i"].(int)}
	return output
}
func flattenTypeNamedStructAlias(input types.TypeNamedStructAlias) []interface{} {
	return []interface{}{map[string]interface{}{"i": input.I}}
}

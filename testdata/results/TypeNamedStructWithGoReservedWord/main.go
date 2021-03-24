package types

import (
	types "types"
	utils "types/utils"
)

func expandTypeNamedStructWithGoReservedWord(input []interface{}) types.TypeNamedStructWithGoReservedWord {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNamedStructWithGoReservedWord{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNamedStructWithGoReservedWord{
		Go: utils.String(b["go"].(string)),
		If: utils.String(b["if"].(string)),
	}
	return output
}
func flattenTypeNamedStructWithGoReservedWord(input types.TypeNamedStructWithGoReservedWord) []interface{} {
	var g string
	if input.Go != nil {
		g = *input.Go
	}
	var iF string
	if input.If != nil {
		iF = *input.If
	}
	return []interface{}{map[string]interface{}{
		"go": g,
		"if": iF,
	}}
}

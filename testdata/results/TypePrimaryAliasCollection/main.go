package output

import types "types"

func expandTypePrimaryAliasCollection(input []interface{}) types.TypePrimaryAliasCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasCollection{
		Bool:  b["bool"].(types.Bool),
		Float: b["float"].(types.Float),
		Int:   b["int"].(types.Int),
		Str:   b["str"].(types.Str),
	}
	return output
}

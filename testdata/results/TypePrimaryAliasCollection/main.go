package types

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
func flattenTypePrimaryAliasCollection(input types.TypePrimaryAliasCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"bool":  bool(input.Bool),
		"float": float64(input.Float),
		"int":   int(input.Int),
		"str":   string(input.Str),
	}}
}

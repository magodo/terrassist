package types

import types "types"

func expandTypePrimaryAliasCollection(input []interface{}) types.TypePrimaryAliasCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypePrimaryAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypePrimaryAliasCollection{
		Bool:  types.Bool(b["bool"].(bool)),
		Float: types.Float(b["float"].(float64)),
		Int:   types.Int(b["int"].(int)),
		Str:   types.Str(b["str"].(string)),
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

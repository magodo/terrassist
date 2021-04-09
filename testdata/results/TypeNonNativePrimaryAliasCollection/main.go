package types

import types "types"

func expandTypeNonNativePrimaryAliasCollection(input []interface{}) types.TypeNonNativePrimaryAliasCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNonNativePrimaryAliasCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNonNativePrimaryAliasCollection{
		Float32: types.Float32(b["float_32"].(float32)),
		Int16:   types.Int16(b["int_16"].(int16)),
		Int32:   types.Int32(b["int_32"].(int32)),
		Int64:   types.Int64(b["int_64"].(int64)),
		Int8:    types.Int8(b["int_8"].(int8)),
		Uint16:  types.Uint16(b["uint_16"].(uint16)),
		Uint32:  types.Uint32(b["uint_32"].(uint32)),
		Uint64:  types.Uint64(b["uint_64"].(uint64)),
		Uint8:   types.Uint8(b["uint_8"].(uint8)),
	}
	return output
}
func flattenTypeNonNativePrimaryAliasCollection(input types.TypeNonNativePrimaryAliasCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"float_32": float64(input.Float32),
		"int_16":   int(input.Int16),
		"int_32":   int(input.Int32),
		"int_64":   int(input.Int64),
		"int_8":    int(input.Int8),
		"uint_16":  int(input.Uint16),
		"uint_32":  int(input.Uint32),
		"uint_64":  int(input.Uint64),
		"uint_8":   int(input.Uint8),
	}}
}

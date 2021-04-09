package types

import types "types"

func expandTypeNonNativePrimaryCollection(input []interface{}) types.TypeNonNativePrimaryCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNonNativePrimaryCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNonNativePrimaryCollection{
		Float32: b["float_32"].(float32),
		Int16:   b["int_16"].(int16),
		Int32:   b["int_32"].(int32),
		Int64:   b["int_64"].(int64),
		Int8:    b["int_8"].(int8),
		Uint:    b["uint"].(uint),
		Uint16:  b["uint_16"].(uint16),
		Uint32:  b["uint_32"].(uint32),
		Uint64:  b["uint_64"].(uint64),
		Uint8:   b["uint_8"].(uint8),
	}
	return output
}
func flattenTypeNonNativePrimaryCollection(input types.TypeNonNativePrimaryCollection) []interface{} {
	return []interface{}{map[string]interface{}{
		"float_32": float64(input.Float32),
		"int_16":   int(input.Int16),
		"int_32":   int(input.Int32),
		"int_64":   int(input.Int64),
		"int_8":    int(input.Int8),
		"uint":     int(input.Uint),
		"uint_16":  int(input.Uint16),
		"uint_32":  int(input.Uint32),
		"uint_64":  int(input.Uint64),
		"uint_8":   int(input.Uint8),
	}}
}

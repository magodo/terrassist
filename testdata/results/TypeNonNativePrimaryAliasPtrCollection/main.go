package types

import types "types"

func expandTypeNonNativePrimaryAliasPtrCollection(input []interface{}) types.TypeNonNativePrimaryAliasPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNonNativePrimaryAliasPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	ptrInt8 := types.Int8(b["ptr_int_8"].(int))
	ptrInt16 := types.Int16(b["ptr_int_16"].(int))
	ptrInt32 := types.Int32(b["ptr_int_32"].(int))
	ptrInt64 := types.Int64(b["ptr_int_64"].(int))
	ptrUint8 := types.Uint8(b["ptr_uint_8"].(int))
	ptrUint16 := types.Uint16(b["ptr_uint_16"].(int))
	ptrUint32 := types.Uint32(b["ptr_uint_32"].(int))
	ptrUint64 := types.Uint64(b["ptr_uint_64"].(int))
	ptrFloat32 := types.Float32(b["ptr_float_32"].(float64))
	output := types.TypeNonNativePrimaryAliasPtrCollection{
		PtrFloat32: &ptrFloat32,
		PtrInt16:   &ptrInt16,
		PtrInt32:   &ptrInt32,
		PtrInt64:   &ptrInt64,
		PtrInt8:    &ptrInt8,
		PtrUint16:  &ptrUint16,
		PtrUint32:  &ptrUint32,
		PtrUint64:  &ptrUint64,
		PtrUint8:   &ptrUint8,
	}
	return output
}
func flattenTypeNonNativePrimaryAliasPtrCollection(input types.TypeNonNativePrimaryAliasPtrCollection) []interface{} {
	var ptrInt8 types.Int8
	if input.PtrInt8 != nil {
		ptrInt8 = *input.PtrInt8
	}
	var ptrInt16 types.Int16
	if input.PtrInt16 != nil {
		ptrInt16 = *input.PtrInt16
	}
	var ptrInt32 types.Int32
	if input.PtrInt32 != nil {
		ptrInt32 = *input.PtrInt32
	}
	var ptrInt64 types.Int64
	if input.PtrInt64 != nil {
		ptrInt64 = *input.PtrInt64
	}
	var ptrUint8 types.Uint8
	if input.PtrUint8 != nil {
		ptrUint8 = *input.PtrUint8
	}
	var ptrUint16 types.Uint16
	if input.PtrUint16 != nil {
		ptrUint16 = *input.PtrUint16
	}
	var ptrUint32 types.Uint32
	if input.PtrUint32 != nil {
		ptrUint32 = *input.PtrUint32
	}
	var ptrUint64 types.Uint64
	if input.PtrUint64 != nil {
		ptrUint64 = *input.PtrUint64
	}
	var ptrFloat32 types.Float32
	if input.PtrFloat32 != nil {
		ptrFloat32 = *input.PtrFloat32
	}
	return []interface{}{map[string]interface{}{
		"ptr_float_32": float64(ptrFloat32),
		"ptr_int_16":   int(ptrInt16),
		"ptr_int_32":   int(ptrInt32),
		"ptr_int_64":   int(ptrInt64),
		"ptr_int_8":    int(ptrInt8),
		"ptr_uint_16":  int(ptrUint16),
		"ptr_uint_32":  int(ptrUint32),
		"ptr_uint_64":  int(ptrUint64),
		"ptr_uint_8":   int(ptrUint8),
	}}
}

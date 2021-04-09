package types

import (
	types "types"
	utils "types/utils"
)

func expandTypeNonNativePrimaryPtrCollection(input []interface{}) types.TypeNonNativePrimaryPtrCollection {
	if len(input) == 0 || input[0] == nil {
		return types.TypeNonNativePrimaryPtrCollection{}
	}
	b := input[0].(map[string]interface{})
	output := types.TypeNonNativePrimaryPtrCollection{
		PtrFloat32: utils.Float32(b["ptr_float_32"].(float32)),
		PtrInt16:   utils.Int16(b["ptr_int_16"].(int16)),
		PtrInt32:   utils.Int32(b["ptr_int_32"].(int32)),
		PtrInt64:   utils.Int64(b["ptr_int_64"].(int64)),
		PtrInt8:    utils.Int8(b["ptr_int_8"].(int8)),
		PtrUint:    utils.Uint(b["ptr_uint"].(uint)),
		PtrUint16:  utils.Uint16(b["ptr_uint_16"].(uint16)),
		PtrUint32:  utils.Uint32(b["ptr_uint_32"].(uint32)),
		PtrUint64:  utils.Uint64(b["ptr_uint_64"].(uint64)),
		PtrUint8:   utils.Uint8(b["ptr_uint_8"].(uint8)),
	}
	return output
}
func flattenTypeNonNativePrimaryPtrCollection(input types.TypeNonNativePrimaryPtrCollection) []interface{} {
	var ptrInt8 int8
	if input.PtrInt8 != nil {
		ptrInt8 = *input.PtrInt8
	}
	var ptrInt16 int16
	if input.PtrInt16 != nil {
		ptrInt16 = *input.PtrInt16
	}
	var ptrInt32 int32
	if input.PtrInt32 != nil {
		ptrInt32 = *input.PtrInt32
	}
	var ptrInt64 int64
	if input.PtrInt64 != nil {
		ptrInt64 = *input.PtrInt64
	}
	var ptrUint uint
	if input.PtrUint != nil {
		ptrUint = *input.PtrUint
	}
	var ptrUint8 uint8
	if input.PtrUint8 != nil {
		ptrUint8 = *input.PtrUint8
	}
	var ptrUint16 uint16
	if input.PtrUint16 != nil {
		ptrUint16 = *input.PtrUint16
	}
	var ptrUint32 uint32
	if input.PtrUint32 != nil {
		ptrUint32 = *input.PtrUint32
	}
	var ptrUint64 uint64
	if input.PtrUint64 != nil {
		ptrUint64 = *input.PtrUint64
	}
	var ptrFloat32 float32
	if input.PtrFloat32 != nil {
		ptrFloat32 = *input.PtrFloat32
	}
	return []interface{}{map[string]interface{}{
		"ptr_float_32": float64(ptrFloat32),
		"ptr_int_16":   int(ptrInt16),
		"ptr_int_32":   int(ptrInt32),
		"ptr_int_64":   int(ptrInt64),
		"ptr_int_8":    int(ptrInt8),
		"ptr_uint":     int(ptrUint),
		"ptr_uint_16":  int(ptrUint16),
		"ptr_uint_32":  int(ptrUint32),
		"ptr_uint_64":  int(ptrUint64),
		"ptr_uint_8":   int(ptrUint8),
	}}
}

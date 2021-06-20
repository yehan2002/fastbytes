//go:build !no_unsafe
// +build !no_unsafe

package bytes

import (
	"reflect"

	"github.com/yehan2002/bytes/internal/unsafe"
)

var rotateBigEndian = unsafe.IsLittleEndian

func fromI8(src []int8, dst []byte) int                          { return unsafe.FromI8(src, dst) }
func fromI16(src []int16, dst []byte, rot bool) int              { return unsafe.FromI16(src, dst, rot) }
func fromI32(src []int32, dst []byte, rot bool) int              { return unsafe.FromI32(src, dst, rot) }
func fromI64(src []int64, dst []byte, rot bool) int              { return unsafe.FromI64(src, dst, rot) }
func fromU16(src []uint16, dst []byte, rot bool) int             { return unsafe.FromU16(src, dst, rot) }
func fromU32(src []uint32, dst []byte, rot bool) int             { return unsafe.FromU32(src, dst, rot) }
func fromU64(src []uint64, dst []byte, rot bool) int             { return unsafe.FromU64(src, dst, rot) }
func fromF32(sr []float32, dst []byte, rot bool) int             { return unsafe.FromF32(sr, dst, rot) }
func fromF64(sr []float64, dst []byte, rot bool) int             { return unsafe.FromF64(sr, dst, rot) }
func toI8(src []uint8, dst []int8) int                           { return unsafe.ToI8(src, dst) }
func toI16(src []byte, dst []int16, rot bool) int                { return unsafe.ToI16(src, dst, rot) }
func toI32(src []byte, dst []int32, rot bool) int                { return unsafe.ToI32(src, dst, rot) }
func toI64(src []byte, dst []int64, rot bool) int                { return unsafe.ToI64(src, dst, rot) }
func toU16(src []byte, dst []uint16, rot bool) int               { return unsafe.ToU16(src, dst, rot) }
func toU32(src []byte, dst []uint32, rot bool) int               { return unsafe.ToU32(src, dst, rot) }
func toU64(src []byte, dst []uint64, rot bool) int               { return unsafe.ToU64(src, dst, rot) }
func toF32(src []byte, dst []float32, rot bool) int              { return unsafe.ToF32(src, dst, rot) }
func toF64(src []byte, dst []float64, rot bool) int              { return unsafe.ToF64(src, dst, rot) }
func fromSlice(s interface{}, d []byte, rot bool) (int, error)   { return unsafe.FromSlice(s, d, rot) }
func toSlice(s []byte, d interface{}, rot bool) (int, error)     { return unsafe.ToSlice(s, d, rot) }
func fromValue(s reflect.Value, d []byte, rot bool) (int, error) { return unsafe.FromValue(s, d, rot) }
func toValue(s []byte, d reflect.Value, rot bool) (int, error)   { return unsafe.ToValue(s, d, rot) }

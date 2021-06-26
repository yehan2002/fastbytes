//go:build !go1.17
// +build !go1.17

package unsafe

import (
	"reflect"
	"runtime"
	"unsafe"
)

// sliceOf creates a new byte slice from the given pointer.
// The capacity of the returned slice is the same as the length.
// This function assumes that the caller keeps the ptr reachable.
func sliceOf(data unsafe.Pointer, length int) (v []byte) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(data)
	slice.Len, slice.Cap = length, length
	return
}

// u8Tou16 converts the given byte slice to a uint16 slice
// The returned  slice has a length of `len(v)/2`
// This function panics if slice is shorter than 2
func u8Tou16(d []byte) (v []uint16) {
	_ = d[:2]
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/uint16Bytes, len(d)/uint16Bytes
	runtime.KeepAlive(d)
	return
}

// u8Tou32 converts the given byte slice to a uint32 slice
// The returned  slice has a length of `len(v)/4`
// This function panics if slice is shorter than 4
func u8Tou32(d []byte) (v []uint32) {
	_ = d[:4]
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/uint32Bytes, len(d)/uint32Bytes
	runtime.KeepAlive(d)
	return
}

// u8Tou64 converts the given byte slice to a uint64 slice
// The returned  slice has a length of `len(v)/8`
// This function panics if slice is shorter than 8
func u8Tou64(d []byte) (v []uint64) {
	_ = d[:8]
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/uint64Bytes, len(d)/uint64Bytes
	runtime.KeepAlive(d)
	return
}

// u16Tou64 converts the given uint16 slice to a uint64 slice
// The returned  slice has a length of `len(v)/4`
// This function panics if slice is shorter than 4
func u16Tou64(d []uint16) (v []uint64) {
	_ = d[:4]
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/uint64Uint16s, len(d)/uint64Uint16s
	runtime.KeepAlive(d)
	return
}

// f64Tou64 converts the given float64 slice to a uint64 slice
func f64Tou64(v []float64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

// i64Tou64 converts the given int64 slice to a uint64 slice
func i64Tou64(v []int64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

// f32Tou32 converts the given float32 slice to a uint32 slice
func f32Tou32(v []float32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

// i64Tou32 converts the given int32 slice to a uint32 slice
func i32Tou32(v []int32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

// i16Tou16 converts the given int16 slice to a uint16 slice
func i16Tou16(v []int16) []uint16 { return *(*[]uint16)(unsafe.Pointer(&v)) }

// i8Tou8 converts the given int8 slice to a uint8 slice
func i8Tou8(v []int8) []byte { return *(*[]byte)(unsafe.Pointer(&v)) }

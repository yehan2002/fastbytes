//go:build !go1.17
// +build !go1.17

package unsafe

import (
	"reflect"
	"runtime"
	"unsafe"
)

//sliceOf creates a new byte slice from the given pointer.
//The capacity of the returned slice is the same as the length.
//This function assumes that the caller keeps the ptr reachable.
func sliceOf(data unsafe.Pointer, length int) (v []byte) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(data)
	slice.Len, slice.Cap = length, length
	return
}

func u8Tou16(d []byte) (v []uint16) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/2, len(d)/2
	runtime.KeepAlive(d)
	return
}

func u8Tou32(d []byte) (v []uint32) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/4, len(d)/4
	runtime.KeepAlive(d)
	return
}

func u8Tou64(d []byte) (v []uint64) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/8, len(d)/8
	runtime.KeepAlive(d)
	return
}

func u16Tou64(d []uint16) (v []uint64) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	slice.Data = uintptr(unsafe.Pointer(&d[0]))
	slice.Len, slice.Cap = len(d)/4, len(d)/4
	runtime.KeepAlive(d)
	return
}

func f64Tou64(v []float64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }
func i64Tou64(v []int64) []uint64   { return *(*[]uint64)(unsafe.Pointer(&v)) }
func f32Tou32(v []float32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }
func i32Tou32(v []int32) []uint32   { return *(*[]uint32)(unsafe.Pointer(&v)) }
func i16Tou16(v []int16) []uint16   { return *(*[]uint16)(unsafe.Pointer(&v)) }
func i8Tou8(v []int8) []byte        { return *(*[]byte)(unsafe.Pointer(&v)) }

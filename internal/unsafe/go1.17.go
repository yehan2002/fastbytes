//go:build go1.17
// +build go1.17

package unsafe

import "unsafe"

//sliceOf creates a new byte slice from the given pointer.
//The capacity of the returned slice is the same as the length.
//This function assumes that the caller keeps the ptr reachable.
func sliceOf(ptr unsafe.Pointer, length int) []byte { return unsafe.Slice((*byte)(ptr), length) }

func u8Tou16(d []byte) []uint16 { return unsafe.Slice((*uint16)(unsafe.Pointer(&d[0])), len(d)/2) }

func u8Tou32(d []byte) []uint32 { return unsafe.Slice((*uint32)(unsafe.Pointer(&d[0])), len(d)/4) }

func u8Tou64(d []byte) []uint64 { return unsafe.Slice((*uint64)(unsafe.Pointer(&d[0])), len(d)/8) }

func u16Tou64(d []uint16) []uint64 { return unsafe.Slice((*uint64)(unsafe.Pointer(&d[0])), len(d)/4) }

func i8Tou8(v []int8) []byte { return *(*[]byte)(unsafe.Pointer(&v)) }

func i16Tou16(v []int16) []uint16 { return *(*[]uint16)(unsafe.Pointer(&v)) }

func i32Tou32(v []int32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

func f32Tou32(v []float32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

func i64Tou64(v []int64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

func f64Tou64(v []float64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

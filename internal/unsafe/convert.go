//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import "unsafe"

// u8Tou16 converts the given byte slice to a uint16 slice
// The returned  slice has a length of `len(v)/2`.
func u8Tou16(d []byte) []uint16 {
	if len(d) == 0 {
		return nil
	}
	return unsafe.Slice((*uint16)(unsafe.Pointer(&d[0])), len(d)/2)
}

// u8Tou32 converts the given byte slice to a uint32 slice
// The returned  slice has a length of `len(v)/4`
func u8Tou32(d []byte) []uint32 {
	if len(d) == 0 {
		return nil
	}
	return unsafe.Slice((*uint32)(unsafe.Pointer(&d[0])), len(d)/4)
}

// u8Tou64 converts the given byte slice to a uint64 slice
// The returned  slice has a length of `len(v)/8`
func u8Tou64(d []byte) []uint64 {
	if len(d) == 0 {
		return nil
	}
	return unsafe.Slice((*uint64)(unsafe.Pointer(&d[0])), len(d)/8)
}

// u16Tou64 converts the given uint16 slice to a uint64 slice
// The returned  slice has a length of `len(v)/4`
func u16Tou64(d []uint16) []uint64 {
	if len(d) == 0 {
		return nil
	}
	return unsafe.Slice((*uint64)(unsafe.Pointer(&d[0])), len(d)/4)
}

// i8Tou8 converts the given int8 slice to a uint8 slice
func i8Tou8(v []int8) []byte { return *(*[]byte)(unsafe.Pointer(&v)) }

// i16Tou16 converts the given int16 slice to a uint16 slice
func i16Tou16(v []int16) []uint16 { return *(*[]uint16)(unsafe.Pointer(&v)) }

// i64Tou32 converts the given int32 slice to a uint32 slice
func i32Tou32(v []int32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

// f32Tou32 converts the given float32 slice to a uint32 slice
func f32Tou32(v []float32) []uint32 { return *(*[]uint32)(unsafe.Pointer(&v)) }

// i64Tou64 converts the given int64 slice to a uint64 slice
func i64Tou64(v []int64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

// f64Tou64 converts the given float64 slice to a uint64 slice
func f64Tou64(v []float64) []uint64 { return *(*[]uint64)(unsafe.Pointer(&v)) }

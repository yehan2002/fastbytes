//go:build !no_unsafe
// +build !no_unsafe

package unsafe

//Bytes a bytes provider that uses unsafe
type Bytes struct{}

// FromI8 copy bytes from []int8
func (Bytes) FromI8(s []int8, d []byte) int {
	return copy(d, i8Tou8(s))
}

// FromU16 copy bytes from []uint16
func (Bytes) FromU16(s []uint16, d []byte, r bool) int {
	return copy16(s, u8Tou16(d), r)
}

// FromI16 copy bytes from []int16
func (Bytes) FromI16(s []int16, d []byte, r bool) int {
	return copy16(i16Tou16(s), u8Tou16(d), r)
}

// FromU32 copy bytes from []uint32
func (Bytes) FromU32(s []uint32, d []byte, r bool) int {
	return copy32(s, u8Tou32(d), r)
}

// FromI32 copy bytes from []int32
func (Bytes) FromI32(s []int32, d []byte, r bool) int {
	return copy32(i32Tou32(s), u8Tou32(d), r)
}

// FromF32 copy bytes from []float32
func (Bytes) FromF32(s []float32, d []byte, r bool) int {
	return copy32(f32Tou32(s), u8Tou32(d), r)
}

// FromU64 copy bytes from []uint64
func (Bytes) FromU64(s []uint64, d []byte, r bool) int {
	return copy64(s, u8Tou64(d), r)
}

// FromI64 copy bytes from []int64
func (Bytes) FromI64(s []int64, d []byte, r bool) int {
	return copy64(i64Tou64(s), u8Tou64(d), r)
}

// FromF64 copy bytes from []float64
func (Bytes) FromF64(s []float64, d []byte, r bool) int {
	return copy64(f64Tou64(s), u8Tou64(d), r)
}

// ToI8 copy bytes to []int8
func (Bytes) ToI8(s []byte, d []int8) int {
	return copy(i8Tou8(d), s)
}

// ToU16 copy bytes to []uint16
func (Bytes) ToU16(s []byte, d []uint16, r bool) int {
	return copy16(u8Tou16(s), d, r)
}

// ToI16 copy bytes to []int16
func (Bytes) ToI16(s []byte, d []int16, r bool) int {
	return copy16(u8Tou16(s), i16Tou16(d), r)
}

// ToU32 copy bytes to []uint32
func (Bytes) ToU32(s []byte, d []uint32, r bool) int {
	return copy32(u8Tou32(s), d, r)
}

// ToI32 copy bytes to []int32
func (Bytes) ToI32(s []byte, d []int32, r bool) int {
	return copy32(u8Tou32(s), i32Tou32(d), r)
}

// ToF32 copy bytes to []float32
func (Bytes) ToF32(s []byte, d []float32, r bool) int {
	return copy32(u8Tou32(s), f32Tou32(d), r)
}

// ToU64 copy bytes to []uint64
func (Bytes) ToU64(s []byte, d []uint64, r bool) int {
	return copy64(u8Tou64(s), d, r)
}

// ToI64 copy bytes to []int64
func (Bytes) ToI64(s []byte, d []int64, r bool) int {
	return copy64(u8Tou64(s), i64Tou64(d), r)
}

// ToF64 copy bytes to []float64
func (Bytes) ToF64(s []byte, d []float64, r bool) int {
	return copy64(u8Tou64(s), f64Tou64(d), r)
}

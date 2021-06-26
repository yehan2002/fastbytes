package safe

import (
	"encoding/binary"
	"math"

	"github.com/yehan2002/bytes/internal"
)

//Bytes byte provider that does not use unsafe
type Bytes struct{}

// ToI8 copy bytes to []int8
func (Bytes) ToI8(src []byte, dst []int8) int {
	if len(dst) >= len(src) {
		for i := 0; i < len(src); i++ {
			dst[i] = int8(src[i])
		}
	}
	return len(dst) * internal.Uint8Bytes
}

// ToI16 copy bytes to []int16
func (Bytes) ToI16(src []byte, dst []int16, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint16Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int16(binary.BigEndian.Uint16(src[i*internal.Uint16Bytes:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int16(binary.LittleEndian.Uint16(src[i*internal.Uint16Bytes:]))
		}
	}
	return len(dst) * internal.Uint16Bytes
}

// ToU16 copy bytes to []uint16
func (Bytes) ToU16(src []byte, dst []uint16, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint16Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint16(src[i*internal.Uint16Bytes:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint16(src[i*internal.Uint16Bytes:])
		}
	}
	return len(dst) * internal.Uint16Bytes
}

// ToI32 copy bytes to []int32
func (Bytes) ToI32(src []byte, dst []int32, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int32(binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int32(binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:]))
		}
	}
	return len(dst) * internal.Uint32Bytes
}

// ToU32 copy bytes to []uint32
func (Bytes) ToU32(src []byte, dst []uint32, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:])
		}
	}
	return len(dst) * internal.Uint32Bytes
}

// ToF32 copy bytes to []float32
func (Bytes) ToF32(src []byte, dst []float32, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float32frombits(binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float32frombits(binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:]))
		}
	}
	return len(dst) * internal.Uint32Bytes
}

// ToI64 copy bytes to []int64
func (Bytes) ToI64(src []byte, dst []int64, bigEndian bool) int {
	_ = src[:len(dst)*8]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int64(binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int64(binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:]))
		}
	}
	return len(dst) * internal.Uint64Bytes
}

// ToU64 copy bytes to []uint64
func (Bytes) ToU64(src []byte, dst []uint64, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint64Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:])
		}
	}
	return len(dst) * internal.Uint64Bytes
}

// ToF64 copy bytes to []float64
func (Bytes) ToF64(src []byte, dst []float64, bigEndian bool) int {
	_ = src[:len(dst)*internal.Uint64Bytes]
	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float64frombits(binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float64frombits(binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:]))
		}
	}
	return len(dst) * internal.Uint64Bytes
}

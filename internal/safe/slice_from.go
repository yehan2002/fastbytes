package safe

import (
	"encoding/binary"
	"math"

	"github.com/yehan2002/fastbytes/internal"
)

// FromI8 copy bytes from []int8
func (Bytes) FromI8(src []int8, dst []byte) int {
	if len(dst) >= len(src) {
		for i := 0; i < len(src); i++ {
			dst[i] = uint8(src[i])
		}
	}
	return len(src)
}

// FromI16 copy bytes from []int16
func (Bytes) FromI16(src []int16, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint16Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src[i]))
		}
	}

	return len(src) * internal.Uint16Bytes
}

// FromU16 copy bytes from []uint16
func (Bytes) FromU16(src []uint16, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint16Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint16(dst[i*internal.Uint16Bytes:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint16(dst[i*internal.Uint16Bytes:], src[i])
		}
	}
	return len(src) * internal.Uint16Bytes
}

// FromI32 copy bytes from []int32
func (Bytes) FromI32(src []int32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src[i]))
		}
	}
	return len(src) * internal.Uint32Bytes
}

// FromU32 copy bytes from []uint32
func (Bytes) FromU32(src []uint32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], src[i])
		}
	}
	return len(src) * internal.Uint32Bytes
}

// FromF32 copy bytes from []float32
func (Bytes) FromF32(src []float32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint32Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], math.Float32bits(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], math.Float32bits(src[i]))
		}
	}
	return len(src) * internal.Uint32Bytes
}

// FromI64 copy bytes from []int64
func (Bytes) FromI64(src []int64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint64Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], uint64(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], uint64(src[i]))
		}
	}
	return len(src) * internal.Uint64Bytes
}

// FromU64 copy bytes from []uint64
func (Bytes) FromU64(src []uint64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint64Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], src[i])
		}
	}
	return len(src) * internal.Uint64Bytes
}

// FromF64 copy bytes from []float64
func (Bytes) FromF64(src []float64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*internal.Uint64Bytes]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], math.Float64bits(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], math.Float64bits(src[i]))
		}
	}
	return len(src) * internal.Uint64Bytes
}

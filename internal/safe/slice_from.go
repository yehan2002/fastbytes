package safe

import (
	"encoding/binary"
	"math"
)

//FromI8 int8
func FromI8(src []int8, dst []byte) int {
	if len(dst) >= len(src) {
		for i := 0; i < len(src); i++ {
			dst[i] = uint8(src[i])
		}
	}
	return len(src)
}

//FromI16 int16
func FromI16(src []int16, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*2]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint16(dst[i*2:], uint16(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint16(dst[i*2:], uint16(src[i]))
		}
	}

	return len(src) * 2
}

//FromU16 uint16
func FromU16(src []uint16, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*2]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint16(dst[i*2:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint16(dst[i*2:], src[i])
		}
	}
	return len(src) * 2
}

//FromI32 int32
func FromI32(src []int32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*4]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*4:], uint32(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*4:], uint32(src[i]))
		}
	}
	return len(src) * 4
}

//FromU32 uint32
func FromU32(src []uint32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*4]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*4:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*4:], src[i])
		}
	}
	return len(src) * 4
}

//FromF32 float32
func FromF32(src []float32, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*4]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint32(dst[i*4:], math.Float32bits(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint32(dst[i*4:], math.Float32bits(src[i]))
		}
	}
	return len(src) * 4
}

//FromI64 int64
func FromI64(src []int64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*8]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*8:], uint64(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*8:], uint64(src[i]))
		}
	}
	return len(src) * 8
}

//FromU64 uint64
func FromU64(src []uint64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*8]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*8:], src[i])
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*8:], src[i])
		}
	}
	return len(src) * 8
}

//FromF64 float64
func FromF64(src []float64, dst []byte, bigEndian bool) int {
	_ = dst[:len(src)*8]
	if bigEndian {
		for i := 0; i < len(src); i++ {
			binary.BigEndian.PutUint64(dst[i*8:], math.Float64bits(src[i]))
		}
	} else {
		for i := 0; i < len(src); i++ {
			binary.LittleEndian.PutUint64(dst[i*8:], math.Float64bits(src[i]))
		}
	}
	return len(src) * 8
}

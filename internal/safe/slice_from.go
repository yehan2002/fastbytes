package safe

import (
	"encoding/binary"
	"math"
)

// FromI8 copy bytes from []int8
func (Bytes) FromI8(src []int8, dst []byte) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst) {
		src = src[:len(dst)]
	}

	for i := 0; i < len(src); i++ {
		dst[i] = uint8(src[i])
	}
	return len(src)
}

// FromI16 copy bytes from []int16
func (Bytes) FromI16(src []int16, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*2 > len(dst) {
		src = src[:len(dst)/2]
	}

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

// FromU16 copy bytes from []uint16
func (Bytes) FromU16(src []uint16, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*2 > len(dst) {
		src = src[:len(dst)/2]
	}

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

// FromI32 copy bytes from []int32
func (Bytes) FromI32(src []int32, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*4 > len(dst) {
		src = src[:len(dst)/4]
	}

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

// FromU32 copy bytes from []uint32
func (Bytes) FromU32(src []uint32, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*4 > len(dst) {
		src = src[:len(dst)/4]
	}

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

// FromF32 copy bytes from []float32
func (Bytes) FromF32(src []float32, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*4 > len(dst) {
		src = src[:len(dst)/4]
	}

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

// FromI64 copy bytes from []int64
func (Bytes) FromI64(src []int64, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*8 > len(dst) {
		src = src[:len(dst)/8]
	}

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

// FromU64 copy bytes from []uint64
func (Bytes) FromU64(src []uint64, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*8 > len(dst) {
		src = src[:len(dst)/8]
	}

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

// FromF64 copy bytes from []float64
func (Bytes) FromF64(src []float64, dst []byte, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src)*8 > len(dst) {
		src = src[:len(dst)/8]
	}

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

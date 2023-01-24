package safe

import (
	"encoding/binary"
	"math"
)

// Bytes byte provider that does not use unsafe
type Bytes struct{}

// ToI8 copy bytes to []int8
func (Bytes) ToI8(src []byte, dst []int8) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst) {
		src = src[:len(dst)]
	}

	for i := 0; i < len(src); i++ {
		dst[i] = int8(src[i])
	}
	return len(dst)
}

// ToI16 copy bytes to []int16
func (Bytes) ToI16(src []byte, dst []int16, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*2 {
		src = src[:len(dst)*2]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int16(binary.BigEndian.Uint16(src[i*2:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int16(binary.LittleEndian.Uint16(src[i*2:]))
		}
	}

	return len(dst) * 2
}

// ToU16 copy bytes to []uint16
func (Bytes) ToU16(src []byte, dst []uint16, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*2 {
		src = src[:len(dst)*2]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint16(src[i*2:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint16(src[i*2:])
		}
	}
	return len(dst) * 2
}

// ToI32 copy bytes to []int32
func (Bytes) ToI32(src []byte, dst []int32, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*4 {
		src = src[:len(dst)*4]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int32(binary.BigEndian.Uint32(src[i*4:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int32(binary.LittleEndian.Uint32(src[i*4:]))
		}
	}
	return len(dst) * 4
}

// ToU32 copy bytes to []uint32
func (Bytes) ToU32(src []byte, dst []uint32, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*4 {
		src = src[:len(dst)*4]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint32(src[i*4:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint32(src[i*4:])
		}
	}
	return len(dst) * 4
}

// ToF32 copy bytes to []float32
func (Bytes) ToF32(src []byte, dst []float32, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*4 {
		src = src[:len(dst)*4]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float32frombits(binary.BigEndian.Uint32(src[i*4:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float32frombits(binary.LittleEndian.Uint32(src[i*4:]))
		}
	}
	return len(dst) * 4
}

// ToI64 copy bytes to []int64
func (Bytes) ToI64(src []byte, dst []int64, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*8 {
		src = src[:len(dst)*8]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = int64(binary.BigEndian.Uint64(src[i*8:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = int64(binary.LittleEndian.Uint64(src[i*8:]))
		}
	}
	return len(dst) * 8
}

// ToU64 copy bytes to []uint64
func (Bytes) ToU64(src []byte, dst []uint64, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*8 {
		src = src[:len(dst)*8]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.BigEndian.Uint64(src[i*8:])
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = binary.LittleEndian.Uint64(src[i*8:])
		}
	}
	return len(dst) * 8
}

// ToF64 copy bytes to []float64
func (Bytes) ToF64(src []byte, dst []float64, bigEndian bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst)*8 {
		src = src[:len(dst)*8]
	}

	if bigEndian {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float64frombits(binary.BigEndian.Uint64(src[i*8:]))
		}
	} else {
		for i := 0; i < len(dst); i++ {
			dst[i] = math.Float64frombits(binary.LittleEndian.Uint64(src[i*8:]))
		}
	}
	return len(dst) * 8
}

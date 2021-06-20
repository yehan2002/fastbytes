package safe

import (
	"encoding/binary"
	"math"
)

//ToI8 int8
func ToI8(src []byte, dst []int8) int {
	if len(dst) >= len(src) {
		for i := 0; i < len(src); i++ {
			dst[i] = int8(src[i])
		}
	}
	return len(dst)
}

//ToI16 int16
func ToI16(src []byte, dst []int16, bigEndian bool) int {
	_ = src[:len(dst)*2]
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

//ToU16 uint16
func ToU16(src []byte, dst []uint16, bigEndian bool) int {
	_ = src[:len(dst)*2]
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

//ToI32 int32
func ToI32(src []byte, dst []int32, bigEndian bool) int {
	_ = src[:len(dst)*4]
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

//ToU32 uint32
func ToU32(src []byte, dst []uint32, bigEndian bool) int {
	_ = src[:len(dst)*4]
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

//ToF32 float32
func ToF32(src []byte, dst []float32, bigEndian bool) int {
	_ = src[:len(dst)*4]
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

//ToI64 int64
func ToI64(src []byte, dst []int64, bigEndian bool) int {
	_ = src[:len(dst)*8]
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

//ToU64 uint64
func ToU64(src []byte, dst []uint64, bigEndian bool) int {
	_ = src[:len(dst)*8]
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

//ToF64 float64
func ToF64(src []byte, dst []float64, bigEndian bool) int {
	_ = src[:len(dst)*8]
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

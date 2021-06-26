package unsafe

import (
	"math/bits"
)

const (
	// rot16Msb this has all MSB bits set for 4 uint16 represented as a single uint64.
	rot16Msb uint64 = 0xff00ff00ff00ff00
	// rot16Lsb this has all LSB bits set for 4 uint16 represented as a single uint64.
	rot16Lsb uint64 = 0x00ff00ff00ff00ff
)

// copy16 copy uint16 from src to dst.
// If rotate is set the bytes of the uint16 are rotated
func copy16(src, dst []uint16, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		rotate16(dst)
	}
	return n * uint16Bytes
}

// copy16 copy uint32 from src to dst.
// If rotate is set the bytes of the uint32 are rotated
func copy32(src, dst []uint32, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		for j := 0; j < len(dst); j++ {
			dst[j] = bits.ReverseBytes32(dst[j])
		}
	}
	return n * uint32Bytes
}

// copy16 copy uint64 from src to dst.
// If rotate is set the bytes of the uint64 are rotated
func copy64(src, dst []uint64, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		for j := 0; j < len(dst); j++ {
			dst[j] = bits.ReverseBytes64(dst[j])
		}
	}
	return n * uint64Bytes
}

// copySlice copy bytes from src to dst.
// This converts the given src and dst to a slice with the element size of `size`.
// The size must be one 1,2,4 or 8. Any other values will case this function to panic.
// If rotate is set the bytes of the dst are rotated according to the given size.
// This function assumes that the given slices are at least one byte long.
func copySlice(src, dst []byte, size int, rotate bool) int {
	switch size {
	case uint8Bytes:
		return copy(dst, src)
	case uint16Bytes:
		return FromU16(u8Tou16(src), dst, rotate)
	case uint32Bytes:
		return FromU32(u8Tou32(src), dst, rotate)
	case uint64Bytes:
		return FromU64(u8Tou64(src), dst, rotate)
	}
	panic("invalid byte size provided")
}

// rotate16SmallSize slices smaller than this will not be optimized
const rotate16SmallSize = 12

// rotate16 rotate the bytes in the given slice.
// This function is used since individually rotating uint16 is slow.
func rotate16(dst []uint16) {
	var l int
	if len(dst) >= rotate16SmallSize {
		d := u16Tou64(dst)
		for j := 0; j < len(d); j++ {
			var u = d[j]
			d[j] = (u&rot16Msb)>>uint8Bits | (u&rot16Lsb)<<uint8Bits
		}
		l = len(d) * uint64Uint16s
	}

	for ; l < len(dst); l++ {
		dst[l] = bits.ReverseBytes16(dst[l])
	}
}

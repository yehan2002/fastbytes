//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"math/bits"
	"unsafe"

	"github.com/yehan2002/fastbytes/v2/internal/unsafe/asm"
)

const (
	// rot16Msb this has all MSB bits set for 4 uint16 represented as a single uint64.
	rot16Msb uint64 = 0xff00ff00ff00ff00
	// rot16Lsb this has all LSB bits set for 4 uint16 represented as a single uint64.
	rot16Lsb uint64 = 0x00ff00ff00ff00ff
)

// copy16 copy uint16 from src to dst.
// If rotate is set the bytes of the uint16 are rotated.
func copy16(src, dst []uint16, rotate bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst) {
		src = src[:len(dst)]
	}

	var n int
	if !rotate || !asm.CanASM || checkOverlap16(src, dst) {
		n = copy(dst, src)
		if rotate {
			rotate16(dst)
		}
	} else {
		n = int(asm.Copy16(src, dst))
	}
	return n * 2
}

// copy16 copy uint32 from src to dst.
// If rotate is set the bytes of the uint32 are rotated.
func copy32(src, dst []uint32, rotate bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst) {
		src = src[:len(dst)]
	}

	var n int
	if !rotate || !asm.CanASM || checkOverlap32(src, dst) {
		n = copy(dst, src)
		if rotate {
			rotate32(dst)
		}
	} else {
		n = int(asm.Copy32(src, dst))
	}
	return n * 4
}

// copy16 copy uint64 from src to dst.
// If rotate is set the bytes of the uint64 are rotated.
func copy64(src, dst []uint64, rotate bool) int {
	if len(src) == 0 || len(dst) == 0 {
		return 0
	} else if len(src) > len(dst) {
		src = src[:len(dst)]
	}

	var n int
	if !rotate || !asm.CanASM || checkOverlap64(src, dst) {
		n = copy(dst, src)
		if rotate {
			rotate64(dst)
		}
	} else {
		n = int(asm.Copy64(src, dst))
	}
	return n * 8
}

// copySlice copy bytes from src to dst.
// This converts the given src and dst to a slice with the element size of `size`.
// The size must be one 1,2,4 or 8. Any other values will case this function to panic.
// If rotate is set the bytes of the dst are rotated according to the given size.
// This function assumes that the given slices are at least one byte long.
func copySlice(s, d []byte, size int, rotate bool) int {
	switch size {
	case 1:
		return copy(d, s)
	case 2:
		return copy16(u8Tou16(s), u8Tou16(d), rotate)
	case 4:
		return copy32(u8Tou32(s), u8Tou32(d), rotate)
	case 8:
		return copy64(u8Tou64(s), u8Tou64(d), rotate)
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
			d[j] = (u&rot16Msb)>>8 | (u&rot16Lsb)<<8
		}
		l = len(d) * 4
	}

	for ; l < len(dst); l++ {
		dst[l] = bits.ReverseBytes16(dst[l])
	}
}

// rotate32 rotate the bytes in the given slice.
func rotate32(dst []uint32) {
	for j := 0; j < len(dst); j++ {
		dst[j] = bits.ReverseBytes32(dst[j])
	}
}

// rotate64 rotate the bytes in the given slice.
func rotate64(dst []uint64) {
	for j := 0; j < len(dst); j++ {
		dst[j] = bits.ReverseBytes64(dst[j])
	}
}

// checkOverlap16 checks if slices overlap.
// This panics if `len(s1) < 1` or `len(s2) < 1`.
// See comment on `sliceOverlap`
func checkOverlap16(s1, s2 []uint16) (overlap bool) {
	p1, p2 := uintptr(unsafe.Pointer(&s1[0])), uintptr(unsafe.Pointer(&s2[0]))
	return sliceOverlap(p1, p2, len(s1), len(s2), 2)
}

// checkOverlap32 checks if slices overlap.
// This panics if `len(s1) < 1` or `len(s2) < 1`.
// See comment on `sliceOverlap`
func checkOverlap32(s1, s2 []uint32) (overlap bool) {
	p1, p2 := uintptr(unsafe.Pointer(&s1[0])), uintptr(unsafe.Pointer(&s2[0]))
	return sliceOverlap(p1, p2, len(s1), len(s2), 4)
}

// checkOverlap64 checks if slices overlap.
// This panics if `len(s1) < 1` or `len(s2) < 1`.
// See comment on `sliceOverlap`
func checkOverlap64(s1, s2 []uint64) (overlap bool) {
	p1, p2 := uintptr(unsafe.Pointer(&s1[0])), uintptr(unsafe.Pointer(&s2[0]))
	return sliceOverlap(p1, p2, len(s1), len(s2), 8)
}

// sliceOverlap checks if the given slices overlap.
// If the given slices overlap it is not safe to call `asm.Copy*`.
// This function assumes that both slices are kept reachable by the caller.
func sliceOverlap(ptr1 uintptr, ptr2 uintptr, length1 int, length2 int, size uintptr) bool {
	if ptr1 < ptr2 {
		return ptr1+(uintptr(length1)*8) >= ptr2
	}
	return ptr2 == ptr1 || ptr2+(uintptr(length2)*size) >= ptr1
}

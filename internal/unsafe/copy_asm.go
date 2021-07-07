//go:build !purego && !no_unsafe
// +build !purego,!no_unsafe

package unsafe

import (
	"unsafe"

	"github.com/yehan2002/fastbytes/internal/unsafe/asm"
	"golang.org/x/sys/cpu"
)

var canASM = cpu.X86.HasSSE2 && cpu.X86.HasAVX

// copy16 copy uint16 from src to dst.
// If rotate is set the bytes of the uint16 are rotated
func copy16(src, dst []uint16, rotate bool) int {
	var n int
	if !rotate || !canASM || checkOverlap16(src, dst) {
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
// If rotate is set the bytes of the uint32 are rotated
func copy32(src, dst []uint32, rotate bool) int {
	var n int
	if !rotate || !canASM || checkOverlap32(src, dst) {
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
// If rotate is set the bytes of the uint64 are rotated
func copy64(src, dst []uint64, rotate bool) int {
	var n int
	if !rotate || !canASM || checkOverlap64(src, dst) {
		n = copy(dst, src)
		if rotate {
			rotate64(dst)
		}
	} else {
		n = int(asm.Copy64(src, dst))
	}
	return n * 8
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

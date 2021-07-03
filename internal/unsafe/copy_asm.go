//go:build !purego && !no_unsafe
// +build !purego,!no_unsafe

package unsafe

import (
	"math/bits"

	"github.com/yehan2002/fastbytes/internal/unsafe/asm"
	"golang.org/x/sys/cpu"
)

var canASM = cpu.X86.HasSSE2 && cpu.X86.HasAVX

// copy16 copy uint16 from src to dst.
// If rotate is set the bytes of the uint16 are rotated
func copy16(src, dst []uint16, rotate bool) int {
	var n int
	if !rotate || !canASM {
		n = copy(dst, src)
		if rotate {
			rotate16(dst)
		}
	} else {
		n = int(asm.Copy16(src, dst))
		for ; n < len(dst); n++ {
			dst[n] = bits.ReverseBytes16(src[n])
		}
	}
	return n * 2
}

// copy16 copy uint32 from src to dst.
// If rotate is set the bytes of the uint32 are rotated
func copy32(src, dst []uint32, rotate bool) int {
	var n int
	if !rotate || !canASM {
		n = copy(dst, src)
		if rotate {
			rotate32(dst)
		}
	} else {
		n = int(asm.Copy32(src, dst))
		for ; n < len(dst); n++ {
			dst[n] = bits.ReverseBytes32(src[n])
		}
	}
	return n * 4
}

// copy16 copy uint64 from src to dst.
// If rotate is set the bytes of the uint64 are rotated
func copy64(src, dst []uint64, rotate bool) int {
	var n int
	if !rotate || !canASM {
		n = copy(dst, src)
		if rotate {
			rotate64(dst)
		}
	} else {
		n = int(asm.Copy64(src, dst))
		for ; n < len(dst); n++ {
			dst[n] = bits.ReverseBytes64(src[n])
		}
	}
	return n * 8
}

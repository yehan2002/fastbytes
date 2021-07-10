//go:build (!amd64 || purego) && !no_unsafe
// +build !amd64 purego
// +build !no_unsafe

package unsafe

// copy16 copy uint16 from src to dst.
// If rotate is set the bytes of the uint16 are rotated
func copy16(src, dst []uint16, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		rotate16(dst)
	}
	return n * 2
}

// copy16 copy uint32 from src to dst.
// If rotate is set the bytes of the uint32 are rotated
func copy32(src, dst []uint32, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		rotate32(dst)
	}
	return n * 4
}

// copy16 copy uint64 from src to dst.
// If rotate is set the bytes of the uint64 are rotated
func copy64(src, dst []uint64, rotate bool) int {
	n := copy(dst, src)
	if rotate {
		rotate64(dst)
	}
	return n * 8
}

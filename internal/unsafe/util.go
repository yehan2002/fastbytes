//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"errors"
	"unsafe"
)

var errAddress = errors.New("cannot address")

// IsLittleEndian this checks if the current system is little endian.
var IsLittleEndian = func() bool {
	test := u8Tou16([]byte{0xEF, 0xBE})

	return checkEndianess(test[0])
}()

func init() {
	var v int64
	var z any = &v
	if (*int64)(ifaceAddr(z)) != &v {
		panic("fastbytes: internal error: ifaceAddr is incompatible with this go version. Use `no_unsafe` build tag to fix this.")
	}
}

func checkEndianess(w uint16) bool {
	if w == 0xBEEF || w == 0xEFBE {
		return w == 0xBEEF //nolint
	}

	// This should be unreachable
	// This is kept here in case golang changes the internal representation of slices
	panic("fastbytes: internal error: unable to get host byte order. Use `no_unsafe` build tag to fix this.")
}

// ifaceAddr gets a pointer to the value contained inside the given interface.
// This function depends on the internal representation of interfaces in golang and may break in future versions.
func ifaceAddr(i interface{}) unsafe.Pointer {
	return (*[2]unsafe.Pointer)(unsafe.Pointer(&i))[1]
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
	if ptr2 == ptr1 {
		return false // special case: allow copying in place
	}

	if ptr1 < ptr2 {
		return ptr1+(uintptr(length1)*size) >= ptr2
	}
	return ptr2+(uintptr(length2)*size) >= ptr1
}

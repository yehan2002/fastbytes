package unsafe

import (
	"reflect"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/bytes/internal/safe"
)

// FromSlice copies bytes from the given interface.
// The provided interface must be a type that can be safely copied.
// The given slice must be large enough to fit all the bytes in `s`
func FromSlice(s interface{}, dst []byte, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = ifaceBytes(s, true); err == nil && len(src) != 0 {
		if internal.CanFitCopyFrom(len(src)/size, len(dst), size) {
			return copySlice(src, dst, size, rotate), nil
		}
		return -1, internal.ErrShort
	}
	return
}

// FromValue copies bytes from the given value.
// The provided value must be a type that can be safely converted to bytes.
// The given slice must be large enough to fit all bytes in `s`
func FromValue(s reflect.Value, dst []byte, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = valueBytes(s); err == nil && len(src) != 0 {
		if internal.CanFitCopyFrom(len(src)/size, len(dst), size) {
			return copySlice(src, dst, size, rotate), nil
		}
		return 0, internal.ErrShort
	}
	if err == errAddress {
		return safe.FromValue(s, dst, rotate && IsLittleEndian) //nolint: wrapcheck
	}
	return
}

// ToSlice copies bytes from `s` into the given slice.
// The given interface must be a type  that can be safely written to.
// `d` must be large enough to fit all the bytes in `s`.
func ToSlice(src []byte, d interface{}, rotate bool) (n int, err error) {
	var dst []byte
	var size int
	if dst, size, err = ifaceBytes(d, false); err == nil && len(dst) != 0 {
		if internal.CanFitCopyTo(len(src), len(dst)/size, size) {
			return copySlice(src, dst, size, rotate), nil
		}
		return -1, internal.ErrShort
	}
	return
}

// ToValue copies bytes from `src` into the given value
// The given interface must be a type that can be safely written to.
// `d` must be large enough to fit all the bytes in `src`
func ToValue(src []byte, d reflect.Value, rotate bool) (n int, err error) {
	var dst []byte
	var size int
	if dst, size, err = valueBytes(d); err == nil && len(dst) != 0 {
		if internal.CanFitCopyTo(len(src), len(dst)/size, size) {
			return copySlice(src, dst, size, rotate), nil
		}
		return -1, internal.ErrShort
	}
	if err == errAddress {
		return 0, internal.ErrUnaddressable
	}
	return
}

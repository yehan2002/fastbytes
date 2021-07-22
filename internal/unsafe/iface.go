//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"

	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/fastbytes/v2/internal/safe"
)

var safeBytes = safe.Bytes{}

// FromSlice copies bytes from the given interface.
// The provided interface must be a type that can be safely copied.
// The given slice must be large enough to fit all the bytes in `s`
func (Bytes) FromSlice(s interface{}, dst []byte, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = ifaceBytes(s, true); err == nil && len(src) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	return
}

// FromValue copies bytes from the given value.
// The provided value must be a type that can be safely converted to bytes.
// The given slice must be large enough to fit all bytes in `s`
func (Bytes) FromValue(s reflect.Value, dst []byte, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = valueBytes(s); err == nil && len(src) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	if err == errAddress {
		return safeBytes.FromValue(s, dst, rotate && IsLittleEndian) //nolint: wrapcheck
	}
	return
}

// ToSlice copies bytes from `s` into the given slice.
// The given interface must be a type  that can be safely written to.
// `d` must be large enough to fit all the bytes in `s`.
func (Bytes) ToSlice(src []byte, d interface{}, rotate bool) (n int, err error) {
	var dst []byte
	var size int
	if dst, size, err = ifaceBytes(d, false); err == nil && len(dst) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	return
}

// ToValue copies bytes from `src` into the given value
// The given interface must be a type that can be safely written to.
// `d` must be large enough to fit all the bytes in `src`
func (Bytes) ToValue(src []byte, d reflect.Value, rotate bool) (n int, err error) {
	var dst []byte
	var size int
	if dst, size, err = valueBytes(d); err == nil && len(dst) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	if err == errAddress {
		return 0, internal.ErrUnaddressable
	}
	return
}

//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"unsafe"

	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/fastbytes/v2/internal/safe"
)

var safeBytes = safe.Bytes{}

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

// ToValue copies bytes from `src` into the given value
// The given [reflect.Value] must be a type that can be safely written to.
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

// FromValueOffset copies bytes from the given value.
// The provided value must be a type that can be safely converted to bytes.
// The given slice must be large enough to fit all bytes in `s`
func (Bytes) FromValueOffset(s reflect.Value, dst []byte, start, end int, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = valueBytes(s); err == nil && len(src) != 0 {
		if err = internal.CheckOffsets(start, end, len(dst)/size); err != nil {
			return 0, err
		}

		start *= size
		end *= size
		return copySlice(src, dst[start:end], size, rotate), nil
	}

	if err == errAddress {
		return safeBytes.FromValue(s, dst, rotate && IsLittleEndian) //nolint: wrapcheck
	}
	return
}

// ToValueOffset copies bytes from `src` into the given value
// The given [reflect.Value] must be a type that can be safely written to.
// `d` must be large enough to fit all the bytes in `src`
func (Bytes) ToValueOffset(src []byte, d reflect.Value, start, end int, rotate bool) (n int, err error) {
	var dst []byte
	var size int

	if dst, size, err = valueBytes(d); err == nil && len(dst) != 0 {
		if err = internal.CheckOffsets(start, end, len(dst)/size); err != nil {
			return 0, err
		}

		start *= size
		end *= size

		return copySlice(src, dst[start:end], size, rotate), nil
	}

	if err == errAddress {
		return 0, internal.ErrUnaddressable
	}
	return
}

// valueBytes converts the given reflect.Value to a byte slice.
// The slice returned by this function has a length and capacity of `v.Len()*element size`.
func valueBytes(v reflect.Value) (data []byte, elementSize int, err error) {
	if empty, err := isValidValue(v); empty || err != nil {
		return nil, 0, err
	}

	var dataPtr unsafe.Pointer

	switch v.Kind() {
	case reflect.Slice:
		dataPtr = v.UnsafePointer()
	case reflect.Ptr:
		v = v.Elem()
		fallthrough
	case reflect.Array:
		// This will not panic because isValidValue has already checked if the array is addressable.
		dataPtr = v.Addr().UnsafePointer()
	default:
		return nil, 0, internal.ErrUnsupported
	}

	size := int(v.Type().Elem().Size())
	length := v.Len() * size

	return unsafe.Slice((*byte)(dataPtr), length), size, nil
}

// isValidValue checks if the given value is a addressable slice, array or a pointer to a slice
// with a element type that can be safety converted to bytes.
func isValidValue(v reflect.Value) (empty bool, err error) {
	typ := v.Type()
	if safe := internal.IsSafeSlice(typ); !v.IsValid() || !safe {
		return false, internal.ErrUnsupported
	} else if !(v.CanAddr() || v.Kind() == reflect.Ptr || (v.Len() > 0 && v.Index(0).CanAddr())) {
		if v.Len() == 0 {
			return true, nil
		}
		return false, errAddress
	}
	return
}

//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"unsafe"

	"github.com/yehan2002/fastbytes/v2/internal"
)

// FromSlice copies bytes from the given interface.
// The provided interface must be a type that can be safely copied.
// The given slice must be large enough to fit all the bytes in `s`
func (Bytes) FromSlice(s any, dst []byte, rotate bool) (n int, err error) {
	var src []byte
	var size int
	if src, size, err = ifaceBytes(s, true); err == nil && len(src) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	return
}

// ToSlice copies bytes from `s` into the given slice.
// The given interface must be a type  that can be safely written to.
// `d` must be large enough to fit all the bytes in `s`.
func (Bytes) ToSlice(src []byte, d any, rotate bool) (n int, err error) {
	var dst []byte
	var size int
	if dst, size, err = ifaceBytes(d, false); err == nil && len(dst) != 0 {
		return copySlice(src, dst, size, rotate), nil
	}
	return
}

// ifaceBytes converts the given interface into a byte slice
func ifaceBytes(i any, arrayOk bool) (v []byte, size int, err error) {
	typ := reflect.TypeOf(i)

	if i == nil {
		return nil, 0, nil
	}

	if !internal.IsSafeSlice(typ) {
		return nil, 0, internal.ErrUnsupported
	}

	var ptr unsafe.Pointer
	var len int

	switch typ.Kind() { //nolint
	case reflect.Slice:
		ptr, size, len = sliceInfo(i, typ)
	case reflect.Array:
		if !arrayOk {
			return nil, 0, internal.ErrUnaddressable
		}
		ptr, size, len = arrayInfo(i, typ)
	case reflect.Ptr:
		ptr, size, len = arrayInfo(i, typ.Elem())
	}

	if len == 0 {
		return nil, 0, nil
	}

	return unsafe.Slice((*byte)(ptr), len), size, nil
}

// sliceInfo gets the pointer to first element of the slice and the number of bytes till the end of the slice.
// This function assumes that the given value is not nil and that it is a slice
// The caller must keep the value reachable.
func sliceInfo(i any, _ reflect.Type) (data unsafe.Pointer, size int, len int) {
	slice := (*reflect.SliceHeader)(ifaceAddr(i))
	data = unsafe.Pointer(slice.Data)
	size = int(reflect.TypeOf(i).Elem().Size())
	len = size * slice.Len
	return
}

// arrayInfo gets the pointer to first element of the array and the number of bytes till the end of the array.
// This function assumes that the given value is not nil and that it is an array.
// The caller must keep the value reachable.
func arrayInfo(i any, typ reflect.Type) (data unsafe.Pointer, size int, length int) {
	data = ifaceAddr(i)
	size = int(typ.Elem().Size())
	length = typ.Len() * size
	return
}

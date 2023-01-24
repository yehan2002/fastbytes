//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/yehan2002/fastbytes/v2/internal"
)

var errAddress = errors.New("cannot address")

// IsLittleEndian this checks if the current system is little endian.
var IsLittleEndian = func() bool {
	test := u8Tou16([]byte{0xEF, 0xBE})

	return checkEndianess(test[0])
}()

func checkEndianess(w uint16) bool {
	if w == 0xBEEF || w == 0xEFBE {
		return w == 0xBEEF //nolint
	}

	// This should be unreachable
	// This is kept here in case golang changes the internal representation of slices
	panic("fastbytes: internal error: unable to get host byte order. Use `no_unsafe` build tag to fix this.")
}

func init() {
	var v int64
	var z any = &v
	if (*int64)(ifaceAddr(z)) != &v {
		panic("fastbytes: internal error: ifaceAddr is incompatible with this go version. Use `no_unsafe` build tag to fix this.")
	}
}

// ifaceAddr gets a pointer to the value contained inside the given interface.
// This function depends on the internal representation of interfaces in golang and may break in future versions.
func ifaceAddr(i interface{}) unsafe.Pointer {
	return (*[2]unsafe.Pointer)(unsafe.Pointer(&i))[1]
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

// ifaceBytes converts the given interface into a byte slice
func ifaceBytes(i interface{}, arrayOk bool) (v []byte, size int, err error) {
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
func sliceInfo(i interface{}, _ reflect.Type) (data unsafe.Pointer, size int, len int) {
	slice := (*reflect.SliceHeader)(ifaceAddr(i))
	data = unsafe.Pointer(slice.Data)
	size = int(reflect.TypeOf(i).Elem().Size())
	len = size * slice.Len
	return
}

// arrayInfo gets the pointer to first element of the array and the number of bytes till the end of the array.
// This function assumes that the given value is not nil and that it is an array.
// The caller must keep the value reachable.
func arrayInfo(i interface{}, typ reflect.Type) (data unsafe.Pointer, size int, length int) {
	data = ifaceAddr(i)
	size = int(typ.Elem().Size())
	length = typ.Len() * size
	return
}

package unsafe

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/yehan2002/bytes/internal"
)

var errAddress error = errors.New("cannot address")

//IsLittleEndian this checks if the current system is little endian
var IsLittleEndian = func() bool {
	test := u8Tou16([]byte{0xEF, 0xBE})

	return checkEndianess(test[0])
}()

func checkEndianess(w uint16) bool {
	if w == 0xBEEF || w == 0xEFBE {
		return w == 0xBEEF
	}

	//This should be unreachable
	//This is kept here in case golang changes the internal representation of slices
	panic("internal error while attempting to get hosts byte order. Use `no_unsafe` build tag to fix this.")
}

//ifaceAddrGC gets a pointer to the value contained inside the given interface.
//This function depends on the internal representation of interfaces in golang and may break in future versions.
func ifaceAddr(i interface{}) unsafe.Pointer { return (*[2]unsafe.Pointer)(unsafe.Pointer(&i))[1] }

//valueBytes converts the given reflect.Value to a byte slice.
//The slice returned by this function has a length and capacity of `v.Len()*element size`.
func valueBytes(v reflect.Value) ([]byte, int, error) {
	typ := v.Type()

	if safe := internal.IsSafeSlice(typ); !v.IsValid() || !safe {
		return nil, 0, internal.ErrUnsupported
	} else if !(v.CanAddr() || v.Kind() == reflect.Ptr || (v.Len() > 0 && v.Index(0).CanAddr())) {
		if v.Kind() == reflect.Slice {
			if v.IsNil() || v.Len() == 0 {
				return nil, 0, nil
			}
		} else {
			if v.Len() == 0 {
				return nil, 0, nil
			}
		}
		return nil, 0, errAddress
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var dataPtr unsafe.Pointer
	if v.Kind() == reflect.Slice {
		dataPtr = unsafe.Pointer(v.Pointer())
	} else if v.Kind() == reflect.Array {
		dataPtr = unsafe.Pointer(v.UnsafeAddr())
	}

	size := int(v.Type().Elem().Size())
	length := v.Len() * size

	return sliceOf(dataPtr, length), size, nil
}

//ifaceBytes converts the given interface into a byte slice
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

	if k := typ.Kind(); k == reflect.Slice {
		ptr, size, len = sliceInfo(i, typ)
	} else if k == reflect.Ptr {
		ptr, size, len = arrayInfo(i, typ.Elem())
	} else {
		if !arrayOk {
			return nil, 0, internal.ErrUnaddressable
		}
		ptr, size, len = arrayInfo(i, typ)
	}

	if len == 0 {
		return nil, 0, nil
	}

	return sliceOf(ptr, len), size, nil
}

//sliceInfo gets the pointer to first element of the slice and the number of bytes till the end of the slice.
//This function assumes that the given value is not nil and that it is a slice
//The caller must keep the value reachable.
func sliceInfo(i interface{}, typ reflect.Type) (data unsafe.Pointer, size int, len int) {
	slice := (*reflect.SliceHeader)(ifaceAddr(i))
	data = unsafe.Pointer(slice.Data)
	size = int(reflect.TypeOf(i).Elem().Size())
	len = size * slice.Len
	return
}

//arrayInfo gets the pointer to first element of the array and the number of bytes till the end of the array.
//This function assumes that the given value is not nil and that it is an array.
//The caller must keep the value reachable.
func arrayInfo(i interface{}, typ reflect.Type) (data unsafe.Pointer, size int, len int) {
	data = ifaceAddr(i)
	size = int(typ.Elem().Size())
	len = typ.Len() * size
	return
}

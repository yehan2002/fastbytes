package internal

import (
	"errors"
	"fmt"
	"reflect"
)

//Exported from bytes
var (
	ErrShort         error = errors.New("short")
	ErrUnsupported   error = errors.New("unsupported")
	ErrUnaddressable error = errors.New("unaddressable")
)

//IsSafeSlice returns if the given type is a slice or array, with a element type that can be safely converted to bytes.
//All signed and unsigned intergers except uint and int, and floats are considered to be safe types.
//uint and int are considered to be unsafe since their size is platform dependent.
func IsSafeSlice(t reflect.Type) bool {
	if k := t.Kind(); k != reflect.Array && k != reflect.Slice {
		if k != reflect.Ptr || t.Elem().Kind() != reflect.Array {
			return false
		}
		t = t.Elem()
	}

	k := t.Elem().Kind()
	return k == reflect.Int8 || k == reflect.Uint8 ||
		k == reflect.Int16 || k == reflect.Uint16 ||
		k == reflect.Int32 || k == reflect.Uint32 || k == reflect.Float32 ||
		k == reflect.Int64 || k == reflect.Uint64 || k == reflect.Float64
}

//CanFitCopyFrom a
func CanFitCopyFrom(src, dst, size int) bool {
	return dst >= size && src*size <= dst
}

//CanFitCopyTo b
func CanFitCopyTo(src, dst, size int) bool { return src >= size && src <= dst*size }

//ShouldCopyFrom 1
func ShouldCopyFrom(src, dst, size int) (bool, error) {
	fmt.Println(src, dst, size)
	if src == 0 {
		return false, nil
	}
	if dst >= size && src*size <= dst {
		return true, nil
	}
	return false, ErrShort
}

//ShouldCopyTo 2
func ShouldCopyTo(src, dst, size int) (bool, error) {
	fmt.Println(src, dst, size)
	if src == 0 {
		return false, nil
	}
	if src >= size && src <= dst*size {
		return true, nil
	}
	return false, ErrShort
}

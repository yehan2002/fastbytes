package internal

import (
	"errors"
	"reflect"
)

// Error to be used by sub packages.
// These will be overridden by the init function in fastbytes
var (
	ErrUnsupported   = errors.New("unsupported")
	ErrUnaddressable = errors.New("unaddressable")
)

// IsSafeSlice returns if the given type is a slice or array, with a element type that can be safely converted to bytes.
// Slices with the element type of uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32 and float64
// can be safely converted to bytes. uint, int, uintptr cannot safely be converted to bytes since their
// size is platform dependant.
func IsSafeSlice(t reflect.Type) bool {
	if k := t.Kind(); k != reflect.Array && k != reflect.Slice {
		if k != reflect.Ptr || t.Elem().Kind() != reflect.Array {
			return false
		}
		t = t.Elem()
	}

	switch t.Elem().Kind() {
	case reflect.Int8, reflect.Uint8,
		reflect.Int16, reflect.Uint16,
		reflect.Int32, reflect.Uint32, reflect.Float32,
		reflect.Int64, reflect.Uint64, reflect.Float64:
		return true
	default:
		return false
	}
}

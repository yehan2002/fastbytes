package internal

import (
	"errors"
	"reflect"
)

// Error to be used by sub packages.
// These will be overridden by bytes
var (
	ErrUnsupported   = errors.New("unsupported")
	ErrUnaddressable = errors.New("unaddressable")
)

// IsSafeSlice returns if the given type is a slice or array, with a element type that can be safely converted to bytes.
// All signed and unsigned intergers except uint and int, and floats are considered to be safe types.
// uint and int are considered to be unsafe since their size is platform dependent.
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

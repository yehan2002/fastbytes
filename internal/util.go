package internal

import (
	"errors"
	"reflect"
)

const (
	// Uint8Bytes the number of bytes in an uint8.
	Uint8Bytes = 1 << iota
	// Uint16Bytes the number of bytes in an uint16.
	Uint16Bytes
	// Uint32Bytes the number of bytes in an uint32.
	Uint32Bytes
	// Uint64Bytes the number of bytes in an uint64.
	Uint64Bytes
)

// Error to be used by sub packages.
// These will be overridden by bytes
var (
	ErrShort         = errors.New("short")
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

// CanFitCopyFrom returns if dst is large enough to fit src if src has a element size of `size`
func CanFitCopyFrom(src, dst, size int) bool {
	return dst >= size && src*size <= dsthttps://github.com/yehan2002/fastbytes/blob/main/internal/util.go
}

// CanFitCopyTo returns if dst is large enough to fit src if dst has a element size of `size`
func CanFitCopyTo(src, dst, size int) bool { return src >= size && src <= dst*size }

// CanCopyFrom returns if src can be copied to dst
// This returns true if src can fit in dst and len(src)>0
func CanCopyFrom(src, dst, size int) (bool, error) {
	if src == 0 {
		return false, nil
	}
	if CanFitCopyFrom(src, dst, size) {
		return true, nil
	}
	return false, ErrShort
}

// CanCopyTo returns if src can be copied to dst
// This returns true if src can fit in dst and len(src)>0
func CanCopyTo(src, dst, size int) (bool, error) {
	if src == 0 {
		return false, nil
	}
	if CanFitCopyTo(src, dst, size) {
		return true, nil
	}
	return false, ErrShort
}

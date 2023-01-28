package internal

import (
	"fmt"
	"reflect"

	"github.com/yehan2002/errors"
)

// Error to be used by sub packages.
// These will be overridden by the init function in fastbytes
var (
	ErrUnsupported   = errors.New("unsupported")
	ErrUnaddressable = errors.New("unaddressable")
	ErrOffset        = errors.New("offset")
)

// IsSafeSliceValue like [IsSafeSlice] but uses a [reflect.Value] instead.
// This function also checks if the given value is the zero value of [reflect.Value].
func IsSafeSliceValue(v reflect.Value) bool {
	return v.IsValid() && IsSafeSlice(v.Type())
}

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

// CheckOffsets checks if the given offsets are valid for a slice with the given length
func CheckOffsets(start, end, length int) error {
	if start < 0 || end < 0 {
		return errors.CauseStr(ErrOffset, fmt.Sprintf("offset is negative: start %d end %d", start, end))
	} else if start > end {
		return errors.CauseStr(ErrOffset, fmt.Sprintf("start > end: start %d end %d", start, end))
	} else if start > length || end > length {
		return errors.CauseStr(ErrOffset, fmt.Sprintf("offset is outside slice bounds: length %d start %d end %d", length, start, end))
	}

	return nil
}

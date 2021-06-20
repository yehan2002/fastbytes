package safe

import (
	"fmt"
	"reflect"

	"github.com/yehan2002/bytes/internal"
)

//FromSlice copy byte from `i` into `dst`
func FromSlice(i interface{}, dst []byte, bigEndian bool) (n int, err error) {
	var ok bool
	switch src := i.(type) {
	case []uint8:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 1); ok {
			n = copy(dst[:len(src)], src)
		}
	case []int8:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 1); ok {
			n = FromI8(src, dst)
		}
	case []uint16:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 2); ok {
			n = FromU16(src, dst, bigEndian)
		}
	case []int16:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 2); ok {
			n = FromI16(src, dst, bigEndian)
		}
	case []uint32:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 4); ok {
			n = FromU32(src, dst, bigEndian)
		}
	case []float32:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 4); ok {
			n = FromF32(src, dst, bigEndian)
		}
	case []int32:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 4); ok {
			n = FromI32(src, dst, bigEndian)
		}
	case []int64:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 8); ok {
			n = FromI64(src, dst, bigEndian)
		}
	case []uint64:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 8); ok {
			n = FromU64(src, dst, bigEndian)
		}
	case []float64:
		if ok, err = internal.ShouldCopyFrom(len(src), len(dst), 8); ok {
			n = FromF64(src, dst, bigEndian)
		}
	default:
		return FromValue(reflect.ValueOf(i), dst, bigEndian)
	}
	return
}

//ToSlice copy byte from `src` into `i`
func ToSlice(src []byte, i interface{}, bigEndian bool) (n int, err error) {
	var ok bool
	fmt.Println(src, i, reflect.TypeOf(i))
	switch dst := i.(type) {
	case []uint8:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 1); ok {
			n = copy(dst, src[:len(dst)])
		}
	case []int8:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 1); ok {
			n = ToI8(src, dst)
		}
	case []uint16:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 2); ok {
			n = ToU16(src, dst, bigEndian)
		}
	case []int16:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 2); ok {
			n = ToI16(src, dst, bigEndian)
		}
	case []uint32:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 4); ok {
			n = ToU32(src, dst, bigEndian)
		}
	case []int32:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 4); ok {
			n = ToI32(src, dst, bigEndian)
		}
	case []float32:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 4); ok {
			n = ToF32(src, dst, bigEndian)
		}
	case []uint64:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 8); ok {
			n = ToU64(src, dst, bigEndian)
		}
	case []int64:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 8); ok {
			n = ToI64(src, dst, bigEndian)
		}
	case []float64:
		if ok, err = internal.ShouldCopyTo(len(src), len(dst), 8); ok {
			n = ToF64(src, dst, bigEndian)
		}
	default:
		return toValue(src, reflect.ValueOf(i), bigEndian)
	}
	return
}

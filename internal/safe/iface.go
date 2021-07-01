package safe

import (
	"reflect"

	"github.com/yehan2002/fastbytes/internal"
)

// FromSlice copy byte from `i` into `dst`
func (b Bytes) FromSlice(i interface{}, dst []byte, bigEndian bool) (n int, err error) { //nolint: cyclop
	var ok bool
	switch src := i.(type) {
	case []uint8:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint8Bytes); ok {
			n = copy(dst[:len(src)], src)
		}
	case []int8:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint8Bytes); ok {
			n = b.FromI8(src, dst)
		}
	case []uint16:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint16Bytes); ok {
			n = b.FromU16(src, dst, bigEndian)
		}
	case []int16:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint16Bytes); ok {
			n = b.FromI16(src, dst, bigEndian)
		}
	case []uint32:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.FromU32(src, dst, bigEndian)
		}
	case []float32:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.FromF32(src, dst, bigEndian)
		}
	case []int32:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.FromI32(src, dst, bigEndian)
		}
	case []int64:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.FromI64(src, dst, bigEndian)
		}
	case []uint64:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.FromU64(src, dst, bigEndian)
		}
	case []float64:
		if ok, err = internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.FromF64(src, dst, bigEndian)
		}
	default:
		return fromValue(reflect.ValueOf(i), dst, bigEndian)
	}
	return
}

// ToSlice copy byte from `src` into `i`
func (b Bytes) ToSlice(src []byte, i interface{}, bigEndian bool) (n int, err error) { //nolint: cyclop
	var ok bool
	switch dst := i.(type) {
	case []uint8:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint8Bytes); ok {
			n = copy(dst, src[:len(dst)])
		}
	case []int8:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint8Bytes); ok {
			n = b.ToI8(src, dst)
		}
	case []uint16:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint16Bytes); ok {
			n = b.ToU16(src, dst, bigEndian)
		}
	case []int16:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint16Bytes); ok {
			n = b.ToI16(src, dst, bigEndian)
		}
	case []uint32:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.ToU32(src, dst, bigEndian)
		}
	case []int32:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.ToI32(src, dst, bigEndian)
		}
	case []float32:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); ok {
			n = b.ToF32(src, dst, bigEndian)
		}
	case []uint64:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.ToU64(src, dst, bigEndian)
		}
	case []int64:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.ToI64(src, dst, bigEndian)
		}
	case []float64:
		if ok, err = internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); ok {
			n = b.ToF64(src, dst, bigEndian)
		}
	default:
		return toValue(src, reflect.ValueOf(i), bigEndian)
	}
	return
}

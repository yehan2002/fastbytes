package safe

import (
	"reflect"
)

// FromSlice copy byte from `i` into `dst`
func (b Bytes) FromSlice(i any, dst []byte, bigEndian bool) (n int, err error) { //nolint: cyclop
	switch src := i.(type) {
	case []uint8:
		n = copy(dst[:len(src)], src)
	case []int8:
		n = b.FromI8(src, dst)
	case []uint16:
		n = b.FromU16(src, dst, bigEndian)
	case []int16:
		n = b.FromI16(src, dst, bigEndian)
	case []uint32:
		n = b.FromU32(src, dst, bigEndian)
	case []float32:
		n = b.FromF32(src, dst, bigEndian)
	case []int32:
		n = b.FromI32(src, dst, bigEndian)
	case []int64:
		n = b.FromI64(src, dst, bigEndian)
	case []uint64:
		n = b.FromU64(src, dst, bigEndian)
	case []float64:
		n = b.FromF64(src, dst, bigEndian)
	default:
		return fromValue(reflect.ValueOf(i), dst, bigEndian)
	}
	return
}

// ToSlice copy byte from `src` into `i`
func (b Bytes) ToSlice(src []byte, i any, bigEndian bool) (n int, err error) { //nolint: cyclop
	switch dst := i.(type) {
	case []uint8:
		n = copy(dst, src[:len(dst)])
	case []int8:
		n = b.ToI8(src, dst)
	case []uint16:
		n = b.ToU16(src, dst, bigEndian)
	case []int16:
		n = b.ToI16(src, dst, bigEndian)
	case []uint32:
		n = b.ToU32(src, dst, bigEndian)
	case []int32:
		n = b.ToI32(src, dst, bigEndian)
	case []float32:
		n = b.ToF32(src, dst, bigEndian)
	case []uint64:
		n = b.ToU64(src, dst, bigEndian)
	case []int64:
		n = b.ToI64(src, dst, bigEndian)
	case []float64:
		n = b.ToF64(src, dst, bigEndian)
	default:
		return toValue(src, reflect.ValueOf(i), bigEndian)
	}
	return
}

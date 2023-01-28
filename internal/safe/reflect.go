package safe

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/yehan2002/fastbytes/v2/internal"
)

// FromValue copy from value
func (b Bytes) FromValue(src reflect.Value, dst []byte, bigEndian bool) (n int, err error) {
	if !internal.IsSafeSlice(src.Type()) {
		return 0, internal.ErrUnsupported
	}

	if src.CanInterface() && src.Kind() == reflect.Slice && src.Type().Elem().PkgPath() == "" {
		return b.FromSlice(src.Interface(), dst, bigEndian)
	}

	return fromValue(src, dst, bigEndian)
}

// ToValue copy from value
func (b Bytes) ToValue(src []byte, dst reflect.Value, bigEndian bool) (n int, err error) {
	if !internal.IsSafeSlice(dst.Type()) {
		return 0, internal.ErrUnsupported
	}

	if dst.CanInterface() && dst.Kind() == reflect.Slice && dst.Type().Elem().PkgPath() == "" {
		return b.ToSlice(src, dst.Interface(), bigEndian)
	}

	return toValue(src, dst, bigEndian)
}

// FromValueOffset copy from value
func (b Bytes) FromValueOffset(src reflect.Value, dst []byte, start, end int, bigEndian bool) (n int, err error) {
	if !internal.IsSafeSlice(src.Type()) {
		return 0, internal.ErrUnsupported
	}

	if internal.CheckOffsets(start, end, src.Len()); err != nil {
		return 0, err
	}

	return b.FromValue(src.Slice(start, end), dst, bigEndian)
}

// ToValueOffset copy from value
func (b Bytes) ToValueOffset(src []byte, dst reflect.Value, start, end int, bigEndian bool) (n int, err error) {
	if !internal.IsSafeSlice(dst.Type()) {
		return 0, internal.ErrUnsupported
	}

	if internal.CheckOffsets(start, end, dst.Len()); err != nil {
		return 0, err
	}

	return b.ToValue(src, dst.Slice(start, end), bigEndian)
}

func fromValue(src reflect.Value, dst []byte, bigEndian bool) (n int, err error) { //nolint: cyclop, funlen
	if src.Kind() == reflect.Ptr { // ptr to array
		src = src.Elem()
	}

	size := int(src.Type().Elem().Size())
	if src.Len() == 0 || len(dst) == 0 {
		return 0, nil
	} else if src.Len()*size > len(dst) {
		src = src.Slice(0, len(dst)/size)
	}

	switch src.Type().Elem().Kind() {
	case reflect.Uint8:
		for i := 0; i < src.Len(); i++ {
			dst[i] = uint8(src.Index(i).Uint())
		}
		n = src.Len()
	case reflect.Int8:
		for i := 0; i < src.Len(); i++ {
			dst[i] = uint8(src.Index(i).Int())
		}
		n = src.Len()
	case reflect.Uint16:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint16(dst[i*2:], uint16(src.Index(i).Uint()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint16(dst[i*2:], uint16(src.Index(i).Uint()))
			}
		}
		n = src.Len() * 2
	case reflect.Int16:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint16(dst[i*2:], uint16(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint16(dst[i*2:], uint16(src.Index(i).Int()))
			}
		}
		n = src.Len() * 2
	case reflect.Uint32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*4:], uint32(src.Index(i).Uint()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*4:], uint32(src.Index(i).Uint()))
			}
		}
		n = src.Len() * 4
	case reflect.Int32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*4:], uint32(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*4:], uint32(src.Index(i).Int()))
			}
		}
		n = src.Len() * 4
	case reflect.Float32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*4:], math.Float32bits(float32(src.Index(i).Float())))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*4:], math.Float32bits(float32(src.Index(i).Float())))
			}
		}
		n = src.Len() * 4
	case reflect.Uint64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*8:], src.Index(i).Uint())
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*8:], src.Index(i).Uint())
			}
		}
		n = src.Len() * 8
	case reflect.Int64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*8:], uint64(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*8:], uint64(src.Index(i).Int()))
			}
		}
		n = src.Len() * 8
	case reflect.Float64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*8:], math.Float64bits(src.Index(i).Float()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*8:], math.Float64bits(src.Index(i).Float()))
			}
		}
		n = src.Len() * 8
	default:
		return 0, internal.ErrUnsupported // unreachable
	}
	return
}

func toValue(src []byte, dst reflect.Value, bigEndian bool) (n int, err error) { //nolint: cyclop, funlen
	if dst.Kind() == reflect.Ptr { // ptr to array
		dst = dst.Elem()
	}
	if !dst.CanAddr() {
		return 0, internal.ErrUnaddressable
	}

	size := int(dst.Type().Elem().Size())
	if len(src) == 0 || dst.Len() == 0 {
		return 0, nil
	} else if len(src) > dst.Len()*size {
		src = src[:dst.Len()*size]
	}

	switch dst.Type().Elem().Kind() {
	case reflect.Uint8:
		for i := 0; i < dst.Len(); i++ {
			dst.Index(i).SetUint(uint64(src[i]))
		}
		n = dst.Len()
	case reflect.Int8:
		for i := 0; i < dst.Len(); i++ {
			dst.Index(i).SetInt(int64(src[i]))
		}
		n = dst.Len()
	case reflect.Uint16:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.BigEndian.Uint16(src[i*2:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.LittleEndian.Uint16(src[i*2:])))
			}
		}
		n = dst.Len() * 2
	case reflect.Int16:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint16(src[i*2:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint16(src[i*2:])))
			}
		}
		n = dst.Len() * 2
	case reflect.Uint32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.BigEndian.Uint32(src[i*4:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.LittleEndian.Uint32(src[i*4:])))
			}
		}
		n = dst.Len() * 4
	case reflect.Int32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint32(src[i*4:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint32(src[i*4:])))
			}
		}
		n = dst.Len() * 4
	case reflect.Float32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(float64(math.Float32frombits(binary.BigEndian.Uint32(src[i*4:]))))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(src[i*4:]))))
			}
		}
		n = dst.Len() * 4
	case reflect.Uint64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(binary.BigEndian.Uint64(src[i*8:]))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(binary.LittleEndian.Uint64(src[i*8:]))
			}
		}
		n = dst.Len() * 8
	case reflect.Int64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint64(src[i*8:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint64(src[i*8:])))
			}
		}
		n = dst.Len() * 8
	case reflect.Float64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(math.Float64frombits(binary.BigEndian.Uint64(src[i*8:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(src[i*8:])))
			}
		}
		n = dst.Len() * 8
	default:
		return 0, internal.ErrUnsupported // unreachable
	}
	return
}

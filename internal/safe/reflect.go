package safe

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/yehan2002/bytes/internal"
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

func fromValue(src reflect.Value, dst []byte, bigEndian bool) (n int, err error) { //nolint: cyclop, funlen
	if src.Kind() == reflect.Ptr { // ptr to array
		src = src.Elem()
	}
	if ok, err := internal.CanCopyFrom(src.Len(), len(dst), int(src.Type().Elem().Size())); !ok {
		return 0, err
	}

	switch src.Type().Elem().Kind() {
	case reflect.Uint8:
		_ = dst[:src.Len()]
		for i := 0; i < src.Len(); i++ {
			dst[i] = uint8(src.Index(i).Uint())
		}
		n = src.Len()
	case reflect.Int8:
		_ = dst[:src.Len()]
		for i := 0; i < src.Len(); i++ {
			dst[i] = uint8(src.Index(i).Int())
		}
		n = src.Len()
	case reflect.Uint16:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src.Index(i).Uint()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src.Index(i).Uint()))
			}
		}
		n = src.Len() * internal.Uint16Bytes
	case reflect.Int16:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint16(dst[i*internal.Uint16Bytes:], uint16(src.Index(i).Int()))
			}
		}
		n = src.Len() * internal.Uint16Bytes
	case reflect.Uint32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src.Index(i).Uint()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src.Index(i).Uint()))
			}
		}
		n = src.Len() * internal.Uint32Bytes
	case reflect.Int32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], uint32(src.Index(i).Int()))
			}
		}
		n = src.Len() * internal.Uint32Bytes
	case reflect.Float32:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint32(dst[i*internal.Uint32Bytes:], math.Float32bits(float32(src.Index(i).Float())))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint32(dst[i*internal.Uint32Bytes:], math.Float32bits(float32(src.Index(i).Float())))
			}
		}
		n = src.Len() * internal.Uint32Bytes
	case reflect.Uint64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], src.Index(i).Uint())
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], src.Index(i).Uint())
			}
		}
		n = src.Len() * internal.Uint64Bytes
	case reflect.Int64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], uint64(src.Index(i).Int()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], uint64(src.Index(i).Int()))
			}
		}
		n = src.Len() * internal.Uint64Bytes
	case reflect.Float64:
		if bigEndian {
			for i := 0; i < src.Len(); i++ {
				binary.BigEndian.PutUint64(dst[i*internal.Uint64Bytes:], math.Float64bits(src.Index(i).Float()))
			}
		} else {
			for i := 0; i < src.Len(); i++ {
				binary.LittleEndian.PutUint64(dst[i*internal.Uint64Bytes:], math.Float64bits(src.Index(i).Float()))
			}
		}
		n = src.Len() * internal.Uint64Bytes
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

	if ok, err := internal.CanCopyFrom(dst.Len(), len(src), int(dst.Type().Elem().Size())); !ok {
		return 0, err
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
				dst.Index(i).SetUint(uint64(binary.BigEndian.Uint16(src[i*internal.Uint16Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.LittleEndian.Uint16(src[i*internal.Uint16Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint16Bytes
	case reflect.Int16:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint16(src[i*internal.Uint16Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint16(src[i*internal.Uint16Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint16Bytes
	case reflect.Uint32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(uint64(binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint32Bytes
	case reflect.Int32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint32Bytes
	case reflect.Float32:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(float64(math.Float32frombits(binary.BigEndian.Uint32(src[i*internal.Uint32Bytes:]))))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(src[i*internal.Uint32Bytes:]))))
			}
		}
		n = dst.Len() * internal.Uint32Bytes
	case reflect.Uint64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:]))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetUint(binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:]))
			}
		}
		n = dst.Len() * internal.Uint64Bytes
	case reflect.Int64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetInt(int64(binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint64Bytes
	case reflect.Float64:
		if bigEndian {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(math.Float64frombits(binary.BigEndian.Uint64(src[i*internal.Uint64Bytes:])))
			}
		} else {
			for i := 0; i < dst.Len(); i++ {
				dst.Index(i).SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(src[i*internal.Uint64Bytes:])))
			}
		}
		n = dst.Len() * internal.Uint64Bytes
	default:
		return 0, internal.ErrUnsupported // unreachable
	}
	return
}

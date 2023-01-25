package fastbytes

import "reflect"

const (
	// ErrUnadressable typo
	//
	// Deprecated: Use [ErrUnaddressable] instead.
	ErrUnadressable = ErrUnaddressable
)

// ByteOrder a byteorder.
//
// It is recommended to use [Endianess] instead.
type ByteOrder interface {
	// FromI8 converts and copies bytes from `src` into `dst`.
	// The number of bytes copied is min(len(src), len(dst))
	FromI8(src []int8, dst []byte) (n int)
	// FromI16 converts and copies []int16 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*2, len(dst))
	FromI16(src []int16, dst []byte) (n int)
	// FromU16 converts and copies []uint16 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*2, len(dst))
	FromU16(src []uint16, dst []byte) (n int)
	// FromI32 converts and copies []int32 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*4, len(dst))
	FromI32(src []int32, dst []byte) (n int)
	// FromU32 converts and copies []uint32 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*4, len(dst))
	FromU32(src []uint32, dst []byte) (n int)
	// FromF32 converts and copies []float32 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*4, len(dst))
	FromF32(src []float32, dst []byte) (n int)
	// FromI64 converts and copies []int64 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*8, len(dst))
	FromI64(src []int64, dst []byte) (n int)
	// FromU64 converts and copies []int64 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*8, len(dst))
	FromU64(src []uint64, dst []byte) (n int)
	// FromF64 converts and copies []float64 from `src` into `dst`.
	// The number of bytes copied is min(len(src)*8, len(dst))
	FromF64(src []float64, dst []byte) (n int)
	// ToI8 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst))
	ToI8(src []byte, dst []int8) (n int)
	// ToI16 converts and copies bytes form `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*2)
	ToI16(src []byte, dst []int16) (n int)
	// ToU16 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*2)
	ToU16(src []byte, dst []uint16) (n int)
	// ToI32 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*4)
	ToI32(src []byte, dst []int32) (n int)
	// ToU32 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*4)
	ToU32(src []byte, dst []uint32) (n int)
	// ToF32 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*4)
	ToF32(src []byte, dst []float32) (n int)
	// ToI64 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*8)
	ToI64(src []byte, dst []int64) (n int)
	// ToU64 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*8)
	ToU64(src []byte, dst []uint64) (n int)
	// ToU64 converts and copies bytes from `src` into `dst`
	// The number of bytes copied is min(len(src), len(dst)*8)
	ToF64(src []byte, dst []float64) (n int)

	// To copies bytes from `s` into the given slice.
	// The given interface must be a type  that can be safely written to.
	// The number of bytes copied is min(len(src), len(dst)* element size of dst)
	To(src []byte, dst interface{}) (n int, err error)
	// From copies bytes from the given interface.
	// The provided interface must be a type that can be safely copied.
	// The number of bytes copied is min(len(src)* element size of dst, len(dst))
	From(src interface{}, dst []byte) (n int, err error)
	// ToValue copies bytes from `src` into the given value
	// The given interface must be a type that can be safely written to.
	// The number of bytes copied is min(len(src), len(dst)* element size of dst)
	ToValue(src []byte, dst reflect.Value) (n int, err error)
	// FromValue copies bytes from the given value.
	// The provided value must be a type that can be safely converted to bytes.
	// The number of bytes copied is min(len(src)* element size of dst, len(dst))
	FromValue(src reflect.Value, dst []byte) (n int, err error)

	fastbytes()
}

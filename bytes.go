// A package for reading and writing values from byte slices quickly

package bytes

import (
	"reflect"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/errors"
)

const (
	// ErrShort the given source is too short to fill destination
	// or the given destination is too short to contain all bytes.
	ErrShort = errors.Error("bytes: source/destination is too short")
	// ErrUnsupported the given type is not supported.
	// All signed and unsigned intergers except uint and int, and floats are supported
	// uint and int are unsupported since their size is platform dependent.
	ErrUnsupported = errors.Error("bytes: unsupported target/source type")
	// ErrUnadressable the give reflect.Value cannot be addressed
	ErrUnadressable = errors.Error("bytes: un-addressable value")
)

var (
	// BigEndian copies bytes to and from big endian byte slices
	BigEndian = bytes{rotate: rotateBigEndian}
	// LittleEndian copies bytes to and from little endian byte slices
	LittleEndian = bytes{rotate: !rotateBigEndian}

	_ ByteOrder = BigEndian
	_ ByteOrder = LittleEndian
)

// ByteOrder the byteorder
type ByteOrder interface {
	// FromI8 converts and copies bytes from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromI8(src []int8, dst []byte) (n int, err error)
	// FromI16 converts and copies []int16 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromI16(src []int16, dst []byte) (n int, err error)
	// FromU16 converts and copies []uint16 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromU16(src []uint16, dst []byte) (n int, err error)
	// FromI32 converts and copies []int32 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromI32(src []int32, dst []byte) (n int, err error)
	// FromU32 converts and copies []uint32 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromU32(src []uint32, dst []byte) (n int, err error)
	// FromF32 converts and copies []float32 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromF32(src []float32, dst []byte) (n int, err error)
	// FromI64 converts and copies []int64 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromI64(src []int64, dst []byte) (n int, err error)
	// FromU64 converts and copies []int64 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromU64(src []uint64, dst []byte) (n int, err error)
	// FromF64 converts and copies []float64 from `src` into `dst`.
	// This returns an error if src is longer than dst
	FromF64(src []float64, dst []byte) (n int, err error)
	// ToI8 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than dst
	ToI8(src []byte, dst []int8) (n int, err error)
	// ToI16 converts ans copies bytes form `src` into `dst`
	// This returns an error if src is longer than `len(dst)*2`
	ToI16(src []byte, dst []int16) (n int, err error)
	// ToU16 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*2`
	ToU16(src []byte, dst []uint16) (n int, err error)
	// ToI32 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*4`
	ToI32(src []byte, dst []int32) (n int, err error)
	// ToU32 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*4`
	ToU32(src []byte, dst []uint32) (n int, err error)
	// ToF32 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*4`
	ToF32(src []byte, dst []float32) (n int, err error)
	// ToI64 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*8`
	ToI64(src []byte, dst []int64) (n int, err error)
	// ToU64 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*8`
	ToU64(src []byte, dst []uint64) (n int, err error)
	// ToU64 converts ans copies bytes from `src` into `dst`
	// This returns an error if src is longer than `len(dst)*8`
	ToF64(src []byte, dst []float64) (n int, err error)

	// To copies bytes from `s` into the given slice.
	// The given interface must be a type  that can be safely written to.
	// `d` must be large enough to fit all the bytes in `s`.
	To(src []byte, dst interface{}) (n int, err error)
	// From copies bytes from the given interface.
	// The provided interface must be a type that can be safely copied.
	// The given slice must be large enough to fit all the bytes in `s`
	From(src interface{}, dst []byte) (n int, err error)
	// ToValue copies bytes from `src` into the given value
	// The given interface must be a type that can be safely written to.
	// `d` must be large enough to fit all the bytes in `src`
	ToValue(src []byte, dst reflect.Value) (n int, err error)
	// FromValue copies bytes from the given value.
	// The provided value must be a type that can be safely converted to bytes.
	// The given slice must be large enough to fit all bytes in `s`
	FromValue(src reflect.Value, dst []byte) (n int, err error)
}

var _ internal.Provider = provider{}

// export errors to `internal`
func init() {
	internal.ErrShort, internal.ErrUnsupported, internal.ErrUnaddressable = ErrShort, ErrUnsupported, ErrUnadressable
}

//A package for reading and writing values from byte slices quickly

package bytes

import (
	"reflect"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/errors"
)

const (
	//ErrShort the given source is too short to fill destination
	//or the given destination is too short to contain all bytes.
	ErrShort = errors.Error("bytes: source/destination is too short")
	//ErrUnsupported the given type is not supported.
	//All signed and unsigned intergers except uint and int, and floats are supported
	//uint and int are unsupported since their size is platform dependent.
	ErrUnsupported = errors.Error("bytes: unsupported target/source type")
	//ErrUnadressable the give reflect.Value cannot be addressed
	ErrUnadressable = errors.Error("bytes: un-addressable value")
)

var (
	//BigEndian bytes
	BigEndian = bytes{rotate: rotateBigEndian}
	//LittleEndian bytes
	LittleEndian = bytes{rotate: !rotateBigEndian}

	_ ByteOrder = BigEndian
	_ ByteOrder = LittleEndian
)

//ByteOrder the byteorder
type ByteOrder interface {
	//FromI8 converts and copies bytes from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromI8(src []int8, dst []byte) (n int, err error)
	//FromI16 converts and copies []int16 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromI16(src []int16, dst []byte) (n int, err error)
	//FromU16 converts and copies []uint16 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromU16(src []uint16, dst []byte) (n int, err error)
	//FromI32 converts and copies []int32 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromI32(src []int32, dst []byte) (n int, err error)
	//FromU32 converts and copies []uint32 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromU32(src []uint32, dst []byte) (n int, err error)
	//FromF32 converts and copies []float32 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromF32(src []float32, dst []byte) (n int, err error)
	//FromI64 converts and copies []int64 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromI64(src []int64, dst []byte) (n int, err error)
	//FromU64 converts and copies []int64 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromU64(src []uint64, dst []byte) (n int, err error)
	//FromF64 converts and copies []float64 from `src` into `dst`.
	//This returns an error if src is longer than dst
	FromF64(src []float64, dst []byte) (n int, err error)
	ToI8(src []byte, dst []int8) (n int, err error)
	ToI16(src []byte, dst []int16) (n int, err error)
	ToU16(src []byte, dst []uint16) (n int, err error)
	ToI32(src []byte, dst []int32) (n int, err error)
	ToU32(src []byte, dst []uint32) (n int, err error)
	ToF32(src []byte, dst []float32) (n int, err error)
	ToI64(src []byte, dst []int64) (n int, err error)
	ToU64(src []byte, dst []uint64) (n int, err error)
	ToF64(src []byte, dst []float64) (n int, err error)
	To(src []byte, dst interface{}) (n int, err error)
	From(src interface{}, dst []byte) (n int, err error)
	ToValue(src []byte, dst reflect.Value) (n int, err error)
	FromValue(src reflect.Value, dst []byte) (n int, err error)
}

type bytes struct{ rotate bool }

//FromI8 converts and copies bytes from `src` into `dst`.
//This returns an error if src is longer than dst
func (bytes) FromI8(src []int8, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 1); !Copy {
		return 0, err
	}
	return fromI8(src, dst), nil
}

//FromI16 converts and copies []int16 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromI16(src []int16, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 2); !Copy {
		return 0, err
	}
	return fromI16(src, dst, b.rotate), nil
}

//FromU16 converts and copies []uint16 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromU16(src []uint16, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 2); !Copy {
		return 0, err
	}
	return fromU16(src, dst, b.rotate), nil
}

//FromI32 converts and copies []int32 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromI32(src []int32, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return fromI32(src, dst, b.rotate), nil
}

//FromU32 converts and copies []uint32 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromU32(src []uint32, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return fromU32(src, dst, b.rotate), nil
}

//FromF32 converts and copies []float32 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromF32(src []float32, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return fromF32(src, dst, b.rotate), nil
}

//FromI64 converts and copies []int64 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromI64(src []int64, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return fromI64(src, dst, b.rotate), nil
}

//FromU64 converts and copies []int64 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromU64(src []uint64, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return fromU64(src, dst, b.rotate), nil
}

//FromF64 converts and copies []float64 from `src` into `dst`.
//This returns an error if src is longer than dst
func (b bytes) FromF64(src []float64, dst []byte) (n int, err error) {
	if Copy, err := internal.ShouldCopyFrom(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return fromF64(src, dst, b.rotate), nil
}

func (bytes) ToI8(src []byte, dst []int8) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 1); !Copy {
		return 0, err
	}
	return toI8(src, dst), nil
}

func (b bytes) ToI16(src []byte, dst []int16) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 2); !Copy {
		return 0, err
	}
	return toI16(src, dst, b.rotate), nil
}

func (b bytes) ToU16(src []byte, dst []uint16) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 2); !Copy {
		return 0, err
	}
	return toU16(src, dst, b.rotate), nil
}

func (b bytes) ToI32(src []byte, dst []int32) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return toI32(src, dst, b.rotate), nil
}

func (b bytes) ToU32(src []byte, dst []uint32) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return toU32(src, dst, b.rotate), nil
}

func (b bytes) ToF32(src []byte, dst []float32) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 4); !Copy {
		return 0, err
	}
	return toF32(src, dst, b.rotate), nil
}

func (b bytes) ToI64(src []byte, dst []int64) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return toI64(src, dst, b.rotate), nil
}

func (b bytes) ToU64(src []byte, dst []uint64) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return toU64(src, dst, b.rotate), nil
}

func (b bytes) ToF64(src []byte, dst []float64) (n int, err error) {
	if Copy, err := internal.ShouldCopyTo(len(src), len(dst), 8); !Copy {
		return 0, err
	}
	return toF64(src, dst, b.rotate), nil
}

func (b bytes) To(src []byte, dst interface{}) (int, error) {
	return toSlice(src, dst, b.rotate)
}

func (b bytes) ToValue(src []byte, dst reflect.Value) (int, error) {
	return toValue(src, dst, b.rotate)
}

func (b bytes) From(src interface{}, dst []byte) (int, error) {
	return fromSlice(src, dst, b.rotate)
}

func (b bytes) FromValue(src reflect.Value, dst []byte) (int, error) {
	return fromValue(src, dst, b.rotate)
}

//export errors to `internal`
func init() {
	internal.ErrShort, internal.ErrUnsupported, internal.ErrUnaddressable = ErrShort, ErrUnsupported, ErrUnadressable
}

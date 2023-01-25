// A package for reading and writing values from byte slices quickly

package fastbytes

import (
	"reflect"

	"github.com/yehan2002/errors"
	"github.com/yehan2002/fastbytes/v2/internal"
)

const (
	// ErrUnsupported the given type is not supported.
	// All signed and unsigned integers except uint and int, and floats are supported
	// uint and int are unsupported since their size is platform dependent.
	ErrUnsupported = errors.Const("fastbytes: unsupported target/source type")
	// ErrUnaddressable the given [reflect.Value] cannot be addressed
	ErrUnaddressable = errors.Const("fastbytes: un-addressable value")
)

var (
	// BigEndian copies bytes to and from big endian byte slices
	BigEndian = Endianess{rotate: rotateBigEndian}
	// LittleEndian copies bytes to and from little endian byte slices
	LittleEndian = Endianess{rotate: !rotateBigEndian}
)

// Endianess provides methods to read and write integer/float slices.
type Endianess struct {
	p      provider
	rotate bool
}

var _ internal.Provider = provider{}
var _ ByteOrder = (*Endianess)(nil)

// FromI8 converts and copies bytes from `src` into `dst`.
// The number of bytes copied is min(len(src), len(dst))
func (b *Endianess) FromI8(src []int8, dst []byte) (n int) { return b.p.FromI8(src, dst) }

// FromI16 converts and copies []int16 from `src` into `dst`.
// The number of bytes copied is min(len(src)*2, len(dst))
func (b *Endianess) FromI16(src []int16, dst []byte) (n int) { return b.p.FromI16(src, dst, b.rotate) }

// FromU16 converts and copies []uint16 from `src` into `dst`.
// The number of bytes copied is min(len(src)*2, len(dst))
func (b *Endianess) FromU16(src []uint16, dst []byte) (n int) { return b.p.FromU16(src, dst, b.rotate) }

// FromI32 converts and copies []int32 from `src` into `dst`.
// The number of bytes copied is min(len(src)*4, len(dst))
func (b *Endianess) FromI32(src []int32, dst []byte) (n int) { return b.p.FromI32(src, dst, b.rotate) }

// FromU32 converts and copies []uint32 from `src` into `dst`.
// The number of bytes copied is min(len(src)*4, len(dst))
func (b *Endianess) FromU32(src []uint32, dst []byte) (n int) { return b.p.FromU32(src, dst, b.rotate) }

// FromF32 converts and copies []float32 from `src` into `dst`.
// The number of bytes copied is min(len(src)*4, len(dst))
func (b *Endianess) FromF32(src []float32, dst []byte) (n int) {
	return b.p.FromF32(src, dst, b.rotate)
}

// FromI64 converts and copies []int64 from `src` into `dst`.
// The number of bytes copied is min(len(src)*8, len(dst))
func (b *Endianess) FromI64(src []int64, dst []byte) (n int) { return b.p.FromI64(src, dst, b.rotate) }

// FromU64 converts and copies []int64 from `src` into `dst`.
// The number of bytes copied is min(len(src)*8, len(dst))
func (b *Endianess) FromU64(src []uint64, dst []byte) (n int) { return b.p.FromU64(src, dst, b.rotate) }

// FromF64 converts and copies []float64 from `src` into `dst`.
// The number of bytes copied is min(len(src)*8, len(dst))
func (b *Endianess) FromF64(src []float64, dst []byte) (n int) {
	return b.p.FromF64(src, dst, b.rotate)
}

// ToI8 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst))
func (b *Endianess) ToI8(src []byte, dst []int8) (n int) { return b.p.ToI8(src, dst) }

// ToI16 converts and copies bytes form `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*2)
func (b *Endianess) ToI16(src []byte, dst []int16) (n int) { return b.p.ToI16(src, dst, b.rotate) }

// ToU16 converts and copies bytes form `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*2)
func (b *Endianess) ToU16(src []byte, dst []uint16) (n int) { return b.p.ToU16(src, dst, b.rotate) }

// ToI32 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*4)
func (b *Endianess) ToI32(src []byte, dst []int32) (n int) { return b.p.ToI32(src, dst, b.rotate) }

// ToU32 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*4)
func (b *Endianess) ToU32(src []byte, dst []uint32) (n int) { return b.p.ToU32(src, dst, b.rotate) }

// ToF32 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*4)
func (b *Endianess) ToF32(src []byte, dst []float32) (n int) { return b.p.ToF32(src, dst, b.rotate) }

// ToI64 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*8)
func (b *Endianess) ToI64(src []byte, dst []int64) (n int) { return b.p.ToI64(src, dst, b.rotate) }

// ToU64 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*8)
func (b *Endianess) ToU64(src []byte, dst []uint64) (n int) { return b.p.ToU64(src, dst, b.rotate) }

// ToF64 converts and copies bytes from `src` into `dst`
// The number of bytes copied is min(len(src), len(dst)*8)
func (b *Endianess) ToF64(src []byte, dst []float64) (n int) { return b.p.ToF64(src, dst, b.rotate) }

// To copies bytes from `s` into the given slice.
// The given interface must be a type  that can be safely written to.
// The number of bytes copied is min(len(src), len(dst)* element size of dst)
func (b *Endianess) To(src []byte, dst interface{}) (int, error) {
	return b.p.ToSlice(src, dst, b.rotate)
}

// ToValue copies bytes from `src` into the given value
// The given interface must be a type that can be safely written to.
// The number of bytes copied is min(len(src), len(dst)* element size of dst)
func (b *Endianess) ToValue(src []byte, dst reflect.Value) (int, error) {
	return b.p.ToValue(src, dst, b.rotate)
}

// From copies bytes from the given interface.
// The provided interface must be a type that can be safely copied.
// The number of bytes copied is min(len(src)* element size of dst, len(dst))
func (b *Endianess) From(src interface{}, dst []byte) (int, error) {
	return b.p.FromSlice(src, dst, b.rotate)
}

// FromValue copies bytes from the given value.
// The provided value must be a type that can be safely converted to bytes.
// The number of bytes copied is min(len(src)* element size of dst, len(dst))
func (b *Endianess) FromValue(src reflect.Value, dst []byte) (int, error) {
	return b.p.FromValue(src, dst, b.rotate)
}

// IsBigEndian returns if the endianess of this struct is big endian.
// If this returns false, the endianess is little endian.
// This returns the endianess this struct read/writes not the system endianess.
func (b *Endianess) IsBigEndian() bool { return b.rotate == rotateBigEndian }

func (b *Endianess) fastbytes() {}

// export errors to `internal`
func init() {
	internal.ErrUnsupported, internal.ErrUnaddressable = ErrUnsupported, ErrUnaddressable
}

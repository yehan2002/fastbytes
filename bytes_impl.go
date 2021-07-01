package fastbytes

import (
	"reflect"

	"github.com/yehan2002/fastbytes/internal"
)

type bytes struct {
	p      provider
	rotate bool
}

// FromI8 converts and copies bytes from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromI8(src []int8, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint8Bytes); !ok {
		return 0, err
	}
	return b.p.FromI8(src, dst), nil
}

// FromI16 converts and copies []int16 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromI16(src []int16, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint16Bytes); !ok {
		return 0, err
	}
	return b.p.FromI16(src, dst, b.rotate), nil
}

// FromU16 converts and copies []uint16 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromU16(src []uint16, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint16Bytes); !ok {
		return 0, err
	}
	return b.p.FromU16(src, dst, b.rotate), nil
}

// FromI32 converts and copies []int32 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromI32(src []int32, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.FromI32(src, dst, b.rotate), nil
}

// FromU32 converts and copies []uint32 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromU32(src []uint32, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.FromU32(src, dst, b.rotate), nil
}

// FromF32 converts and copies []float32 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromF32(src []float32, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.FromF32(src, dst, b.rotate), nil
}

// FromI64 converts and copies []int64 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromI64(src []int64, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.FromI64(src, dst, b.rotate), nil
}

// FromU64 converts and copies []int64 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromU64(src []uint64, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.FromU64(src, dst, b.rotate), nil
}

// FromF64 converts and copies []float64 from `src` into `dst`.
// This returns an error if src is longer than dst
func (b bytes) FromF64(src []float64, dst []byte) (n int, err error) {
	if ok, err := internal.CanCopyFrom(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.FromF64(src, dst, b.rotate), nil
}

// ToI8 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than dst
func (b bytes) ToI8(src []byte, dst []int8) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint8Bytes); !ok {
		return 0, err
	}
	return b.p.ToI8(src, dst), nil
}

// ToI16 converts ans copies bytes form `src` into `dst`
// This returns an error if src is longer than `len(dst)*2`
func (b bytes) ToI16(src []byte, dst []int16) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint16Bytes); !ok {
		return 0, err
	}
	return b.p.ToI16(src, dst, b.rotate), nil
}

// ToU16 converts ans copies bytes form `src` into `dst`
// This returns an error if src is longer than `len(dst)*2`
func (b bytes) ToU16(src []byte, dst []uint16) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint16Bytes); !ok {
		return 0, err
	}
	return b.p.ToU16(src, dst, b.rotate), nil
}

// ToI32 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*4`
func (b bytes) ToI32(src []byte, dst []int32) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.ToI32(src, dst, b.rotate), nil
}

// ToU32 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*4`
func (b bytes) ToU32(src []byte, dst []uint32) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.ToU32(src, dst, b.rotate), nil
}

// ToF32 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*4`
func (b bytes) ToF32(src []byte, dst []float32) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint32Bytes); !ok {
		return 0, err
	}
	return b.p.ToF32(src, dst, b.rotate), nil
}

// ToI64 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*8`
func (b bytes) ToI64(src []byte, dst []int64) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.ToI64(src, dst, b.rotate), nil
}

// ToU64 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*8`
func (b bytes) ToU64(src []byte, dst []uint64) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.ToU64(src, dst, b.rotate), nil
}

// ToF64 converts ans copies bytes from `src` into `dst`
// This returns an error if src is longer than `len(dst)*8`
func (b bytes) ToF64(src []byte, dst []float64) (n int, err error) {
	if ok, err := internal.CanCopyTo(len(src), len(dst), internal.Uint64Bytes); !ok {
		return 0, err
	}
	return b.p.ToF64(src, dst, b.rotate), nil
}

// To copies bytes from `s` into the given slice.
// The given interface must be a type  that can be safely written to.
// `d` must be large enough to fit all the bytes in `s`.
func (b bytes) To(src []byte, dst interface{}) (int, error) {
	return b.p.ToSlice(src, dst, b.rotate)
}

// ToValue copies bytes from `src` into the given value
// The given interface must be a type that can be safely written to.
// `d` must be large enough to fit all the bytes in `src`
func (b bytes) ToValue(src []byte, dst reflect.Value) (int, error) {
	return b.p.ToValue(src, dst, b.rotate)
}

// From copies bytes from the given interface.
// The provided interface must be a type that can be safely copied.
// The given slice must be large enough to fit all the bytes in `s`
func (b bytes) From(src interface{}, dst []byte) (int, error) {
	return b.p.FromSlice(src, dst, b.rotate)
}

// FromValue copies bytes from the given value.
// The provided value must be a type that can be safely converted to bytes.
// The given slice must be large enough to fit all bytes in `s`
func (b bytes) FromValue(src reflect.Value, dst []byte) (int, error) {
	return b.p.FromValue(src, dst, b.rotate)
}

//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/internal"
	"github.com/yehan2002/is"
)

func TestUnsafe(t *testing.T) { is.Suite(t, &testUnsafe{}) }

type testUnsafe struct{}

func (*testUnsafe) TestIfacePtr(is is.IS) {
	var i = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	is.True(*(*[10]int)(ifaceAddr(i)) == i, "IfacePtr returned an invalid pointer")
}

//nolint
func (*testUnsafe) TestIfaceBytes(is is.IS) {
	u32 := [8]uint32{0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210, 0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210}
	t := reflect.TypeOf(u32)

	ptr, size, l := arrayInfo(u32, t)
	is.True(*(*[len(u32)]uint32)(ptr) == u32, "arrayInfo returned invalid pointer")
	is.True(size == 4, "arrayInfo returned incorrect type")
	is.True(l == len(u32)*4, "arrayInfo returned invalid length")

	v1, sizeOf, err := ifaceBytes(u32, true)
	is.Nil(err, "ifaceBytes returned unexpected error")
	is.True(sizeOf == size, "ifaceBytes returned incorrect size")
	is.True(l == len(v1), "arrayInfo returned invalid length")

	v2, sizeOf, err := ifaceBytes(u32[:], true)
	is.Nil(err, "ifaceBytes returned unexpected error")
	is.True(sizeOf == size, "ifaceBytes returned incorrect size")
	is.True(l == len(v2), "arrayInfo returned invalid length")

	v3, sizeOf, err := ifaceBytes(&u32, true)
	is.Nil(err, "ifaceBytes returned unexpected error")
	is.True(sizeOf == size, "ifaceBytes returned incorrect size")
	is.True(l == len(v1), "arrayInfo returned invalid length")

	r1 := u8Tou32(v1)
	r2 := u8Tou32(v2)
	r3 := u8Tou32(v3)
	for i := range u32 {
		v := u32[i]
		is.True(v == r1[i], "invalid pointer returned")
		is.True(v == r2[i], "invalid pointer returned")
		is.True(v == r3[i], "invalid pointer returned")
	}

	v, l, err := ifaceBytes(nil, true)
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for nil iface")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([0]uint32{}, true)
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for zero length array")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([]uint32{}, true)
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for zero length slice")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([]int{1}, true)
	is.True(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == internal.ErrUnsupported, "ifaceBytes returned unexpected error")
}

//nolint
func (*testUnsafe) TestValueBytes(is is.IS) {
	u32 := [8]uint32{0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210, 0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210}
	u32V := reflect.ValueOf(&u32)
	u32SV := reflect.ValueOf(u32[:])

	t := reflect.TypeOf(u32)

	ptr, size, l := arrayInfo(u32, t)
	is.True(*(*[len(u32)]uint32)(ptr) == u32, "arrayInfo returned invalid pointer")
	is.True(size == 4, "arrayInfo returned incorrect type")
	is.True(l == len(u32)*4, "arrayInfo returned invalid length")

	v1, sizeOf, err := valueBytes(u32V)
	is.Nil(err, "ifaceBytes returned unexpected error")
	is.True(sizeOf == size, "ifaceBytes returned incorrect size", sizeOf)
	is.True(l == len(v1), "arrayInfo returned invalid length")

	v2, sizeOf, err := valueBytes(u32SV)
	is.Nil(err, "ifaceBytes returned unexpected error")
	is.True(sizeOf == size, "ifaceBytes returned incorrect size")
	is.True(l == len(v2), "arrayInfo returned invalid length")

	r1 := u8Tou32(v1)
	r2 := u8Tou32(v2)
	for i := range u32 {
		w := u32[i]
		is.True(w == r1[i] && w == r2[i], "invalid pointer returned")
	}

	v, l, err := valueBytes(reflect.ValueOf([]uint32(nil)))
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for nil iface")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([0]uint32{}))
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for zero length array")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([]uint32{}))
	is.True(v == nil, "ifaceBytes returned non-nill byte slice for zero length slice")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([]int{1}))
	is.True(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == internal.ErrUnsupported, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf(u32))
	is.True(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is.True(l == 0, "ifaceBytes returned incorrect length")
	is.True(err == errAddress, "ifaceBytes returned unexpected error", err)

}

//nolint
func (*testUnsafe) TestPanic(is is.IS) {
	//This should never happen
	is.MustPanicCall(func() { checkEndianess(0) })
}

//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/is/v2"
)

func TestUnsafe(t *testing.T) {
	t.Parallel()
	is.Suite(t, &testUnsafe{})
}

type testUnsafe struct{}

func (*testUnsafe) TestIfacePtr(is is.Is) {
	var i = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	is(*(*[10]int)(ifaceAddr(i)) == i, "IfacePtr returned an invalid pointer")
}

//nolint
func (*testUnsafe) TestIfaceBytes(is is.Is) {
	u32 := [8]uint32{0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210, 0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210}
	t := reflect.TypeOf(u32)

	ptr, size, l := arrayInfo(u32, t)
	is(*(*[len(u32)]uint32)(ptr) == u32, "arrayInfo returned invalid pointer")
	is(size == 4, "arrayInfo returned incorrect type")
	is(l == len(u32)*4, "arrayInfo returned invalid length")

	v1, sizeOf, err := ifaceBytes(u32, true)
	is(err == nil, "ifaceBytes returned unexpected error")
	is(sizeOf == size, "ifaceBytes returned incorrect size")
	is(l == len(v1), "arrayInfo returned invalid length")

	v2, sizeOf, err := ifaceBytes(u32[:], true)
	is(err == nil, "ifaceBytes returned unexpected error")
	is(sizeOf == size, "ifaceBytes returned incorrect size")
	is(l == len(v2), "arrayInfo returned invalid length")

	v3, sizeOf, err := ifaceBytes(&u32, true)
	is(err == nil, "ifaceBytes returned unexpected error")
	is(sizeOf == size, "ifaceBytes returned incorrect size")
	is(l == len(v1), "arrayInfo returned invalid length")

	r1 := u8Tou32(v1)
	r2 := u8Tou32(v2)
	r3 := u8Tou32(v3)
	for i := range u32 {
		v := u32[i]
		is(v == r1[i], "invalid pointer returned")
		is(v == r2[i], "invalid pointer returned")
		is(v == r3[i], "invalid pointer returned")
	}

	v, l, err := ifaceBytes(nil, true)
	is(v == nil, "ifaceBytes returned non-nill byte slice for nil iface")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([0]uint32{}, true)
	is(v == nil, "ifaceBytes returned non-nill byte slice for zero length array")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([]uint32{}, true)
	is(v == nil, "ifaceBytes returned non-nill byte slice for zero length slice")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = ifaceBytes([]int{1}, true)
	is(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == internal.ErrUnsupported, "ifaceBytes returned unexpected error")
}

//nolint
func (*testUnsafe) TestValueBytes(is is.Is) {
	u32 := [8]uint32{0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210, 0x01234567, 0x89ABCDEF, 0xFEDBCA98, 0x76543210}
	u32V := reflect.ValueOf(&u32)
	u32SV := reflect.ValueOf(u32[:])

	t := reflect.TypeOf(u32)

	ptr, size, l := arrayInfo(u32, t)
	is(*(*[len(u32)]uint32)(ptr) == u32, "arrayInfo returned invalid pointer")
	is(size == 4, "arrayInfo returned incorrect type")
	is(l == len(u32)*4, "arrayInfo returned invalid length")

	v1, sizeOf, err := valueBytes(u32V)
	is(err == nil, "ifaceBytes returned unexpected error")
	is(sizeOf == size, "ifaceBytes returned incorrect size", sizeOf)
	is(l == len(v1), "arrayInfo returned invalid length")

	v2, sizeOf, err := valueBytes(u32SV)
	is(err == nil, "ifaceBytes returned unexpected error")
	is(sizeOf == size, "ifaceBytes returned incorrect size")
	is(l == len(v2), "arrayInfo returned invalid length")

	r1 := u8Tou32(v1)
	r2 := u8Tou32(v2)
	for i := range u32 {
		w := u32[i]
		is(w == r1[i] && w == r2[i], "invalid pointer returned")
	}

	v, l, err := valueBytes(reflect.ValueOf([]uint32(nil)))
	is(v == nil, "ifaceBytes returned non-nill byte slice for nil iface")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([0]uint32{}))
	is(v == nil, "ifaceBytes returned non-nill byte slice for zero length array")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([]uint32{}))
	is(v == nil, "ifaceBytes returned non-nill byte slice for zero length slice")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == nil, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf([]int{1}))
	is(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == internal.ErrUnsupported, "ifaceBytes returned unexpected error")

	v, l, err = valueBytes(reflect.ValueOf(u32))
	is(v == nil, "ifaceBytes returned non-nil byte slice for unsafe type")
	is(l == 0, "ifaceBytes returned incorrect length")
	is(err == errAddress, "ifaceBytes returned unexpected error", err)

}

func (*testUnsafe) TestPanic(is is.Is) {
	is.Panic(func() { checkEndianess(0) }, "This should never happen")
}

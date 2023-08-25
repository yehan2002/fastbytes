//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/fastbytes/v2/internal/testutil"
	"github.com/yehan2002/is/v2"
)

var bytes = Bytes{}

func TestSlice(t *testing.T) { testutil.Test(t, bytes, IsLittleEndian) }

func TestSliceErrors(t *testing.T) {
	t.Parallel()
	is.Suite(t, &testSlice{})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestError(is is.Is) {
	v := [1]byte{}
	v2 := [1]uint64{}
	n, err := bytes.FromSlice([2]uint32{1}, v[:], t.rotateBigEndian)
	is(err == nil, "dst is addressable")
	is(n == 0, "no bytes should be copied")
	n, err = bytes.FromValue(reflect.ValueOf([]uint16{1, 2}), v[:], t.rotateBigEndian)
	is(err == nil, "dst is addressable")
	is(n == 0, "no bytes should be copied")
	n, err = bytes.ToSlice([]byte{1, 3, 4}, v2[:], t.rotateBigEndian)
	is(err == nil, "dst is addressable")
	is(n == 0, "no bytes should be copied")
	n, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v2[:]), t.rotateBigEndian)
	is(err == nil, "dst is addressable")
	is(n == 0, "no bytes should be copied")

	_, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v), t.rotateBigEndian)
	is(err == internal.ErrUnaddressable, "dst must not be addressable")
}

func (t *testSlice) TestZero(is is.Is) {
	v := [0]byte{}
	n, err := bytes.FromSlice([]uint16{}, v[:], t.rotateBigEndian)
	is(err == nil, "unexpected error")
	is(n == 0, "Incorrect bytes copied")

	n, err = bytes.FromValue(reflect.ValueOf([]uint16{}), v[:], t.rotateBigEndian)
	is(err == nil, "unexpected error")
	is(n == 0, "Incorrect bytes copied")

	n, err = bytes.ToSlice([]byte{}, v[:], t.rotateBigEndian)
	is(err == nil, "unexpected error")
	is(n == 0, "Incorrect bytes copied")

	n, err = bytes.ToValue([]byte{}, reflect.ValueOf(v[:]), t.rotateBigEndian)
	is(err == nil, "unexpected error")
	is(n == 0, "Incorrect bytes copied")
}

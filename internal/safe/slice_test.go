package safe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/fastbytes/v2/internal/testutil"
	"github.com/yehan2002/is/v2"
)

var bytes = Bytes{}

func TestSlice(t *testing.T) { testutil.Test(t, Bytes{}, true) }

func TestSliceErrors(t *testing.T) {
	is.Suite(t, &testSlice{})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestError(is is.Is) {
	v := [1]byte{}
	_, err := bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v), t.rotateBigEndian)
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

package safe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal/testdata"
	"github.com/yehan2002/is/v2"
)

var bytes = Bytes{}

func TestSlice(t *testing.T) { testdata.Test(t, Bytes{}, true) }

func TestSliceErrors(t *testing.T) {
	is.Suite(t, &testSlice{})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestError(is is.Is) {
	/* 	v := [1]byte{}
	   	_, err := bytes.FromSlice([2]uint32{1}, v[:], t.rotateBigEndian)
	   	is(err == internal.ErrShort, "dst must be too short")
	   	_, err = bytes.FromValue(reflect.ValueOf([]uint16{1, 2}), v[:], t.rotateBigEndian)
	   	is(err == internal.ErrShort, "dst must be too short")
	   	_, err = bytes.ToSlice([]byte{1, 3, 4}, v[:], t.rotateBigEndian)
	   	is(err == internal.ErrShort, "dst must be too short")
	   	_, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v[:]), t.rotateBigEndian)
	   	is(err == internal.ErrShort, "dst must be too short")

	   	_, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v), t.rotateBigEndian)
	   	is(err == internal.ErrUnaddressable, "dst must not be addressable") */
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

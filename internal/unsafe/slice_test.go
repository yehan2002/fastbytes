//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/internal"
	"github.com/yehan2002/fastbytes/internal/testdata"
	"github.com/yehan2002/is"
)

var bytes = Bytes{}

func TestSlice(t *testing.T)      { testdata.Test(t, bytes, IsLittleEndian) }
func BenchmarkSlice(t *testing.B) { testdata.Benchmark(t, bytes, IsLittleEndian) }

func TestSliceErrors(t *testing.T) {
	is.Suite(t, &testSlice{})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestError(is is.IS) {
	v := [1]byte{}
	_, err := bytes.FromSlice([2]uint32{1}, v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = bytes.FromValue(reflect.ValueOf([]uint16{1, 2}), v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = bytes.ToSlice([]byte{1, 3, 4}, v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v[:]), t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")

	_, err = bytes.ToValue([]byte{1, 3, 4}, reflect.ValueOf(v), t.rotateBigEndian)
	is.True(err == internal.ErrUnaddressable, "dst must not be addressable")
}

func (t *testSlice) TestZero(is is.IS) {
	v := [0]byte{}
	n, err := bytes.FromSlice([]uint16{}, v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = bytes.FromValue(reflect.ValueOf([]uint16{}), v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = bytes.ToSlice([]byte{}, v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = bytes.ToValue([]byte{}, reflect.ValueOf(v[:]), t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")
}

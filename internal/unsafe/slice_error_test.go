package unsafe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/bytes/internal/data"
	"github.com/yehan2002/is"
)

func TestSlice(t *testing.T) {
	is.Suite(t, &testSlice{rotateBigEndian: rotateBigEndian})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestError(is is.IS) {
	v := [len(data.Bytes) / 2]byte{}
	_, err := FromSlice(data.Uint16BigEndian, v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = FromValue(reflect.ValueOf(data.Uint16BigEndian[:]), v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = ToSlice(data.Bytes[:], v[:], t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")
	_, err = ToValue(data.Bytes[:], reflect.ValueOf(v[:]), t.rotateBigEndian)
	is.True(err == internal.ErrShort, "dst must be too short")

	_, err = ToValue(data.Bytes[:], reflect.ValueOf(v), t.rotateBigEndian)
	is.True(err == internal.ErrUnaddressable, "dst must not be addressable")
}

func (t *testSlice) TestZero(is is.IS) {
	v := [0]byte{}
	n, err := FromSlice(data.Uint16BigEndian[:0], v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = FromValue(reflect.ValueOf(data.Uint16BigEndian[:0]), v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = ToSlice(data.Bytes[:0], v[:], t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")

	n, err = ToValue(data.Bytes[:0], reflect.ValueOf(v[:]), t.rotateBigEndian)
	is.True(err == nil, "unexpected error")
	is.True(n == 0, "Incorrect bytes copied")
}

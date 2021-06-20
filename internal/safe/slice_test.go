package safe

import (
	"reflect"
	"testing"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/bytes/internal/data"
	"github.com/yehan2002/is"
)

func TestSlice(t *testing.T) {
	is.Suite(t, &testSlice{rotateBigEndian: true})
}

type testSlice struct{ rotateBigEndian bool }

func (t *testSlice) TestFrom8(is is.IS) {
	var dst [len(data.Bytes)]byte
	n := FromI8(data.Int8[:], dst[:])
	is.True(n == len(data.Int8), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")
}

func (t *testSlice) TestTo8(is is.IS) {
	var dst [len(data.Int8)]int8
	n := ToI8(data.Bytes[:], dst[:])
	is.True(n == len(data.Int8), "Not all bytes copied")
	is.True(dst == data.Int8, "Incorrect copy")
}

func (t *testSlice) TestFrom16(is is.IS) {
	var dst [len(data.Bytes)]byte
	n := FromI16(data.Int16BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromI16(data.Int16LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU16(data.Uint16BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU16(data.Uint16LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")
}

func (t *testSlice) TestTo16(is is.IS) {
	var dst [len(data.Int16BigEndian)]int16
	var dst2 [len(data.Int16BigEndian)]uint16

	n := ToI16(data.Bytes[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int16BigEndian, "Incorrect copy")

	n = ToI16(data.Bytes[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int16LittleEndian, "Incorrect copy")

	n = ToU16(data.Bytes[:], dst2[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint16BigEndian, "Incorrect copy")

	n = ToU16(data.Bytes[:], dst2[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint16LittleEndian, "Incorrect copy")
}

func (t *testSlice) TestFrom32(is is.IS) {
	var dst [len(data.Bytes)]byte
	n := FromI32(data.Int32BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromI32(data.Int32LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU32(data.Uint32BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU32(data.Uint32LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromF32(data.Float32BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromF32(data.Float32LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")
}

func (t *testSlice) TestTo32(is is.IS) {
	var dst [len(data.Int32BigEndian)]int32
	var dst2 [len(data.Int32BigEndian)]uint32
	var dst3 [len(data.Int32BigEndian)]float32

	n := ToI32(data.Bytes[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int32BigEndian, "Incorrect copy")

	n = ToI32(data.Bytes[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int32LittleEndian, "Incorrect copy")

	n = ToU32(data.Bytes[:], dst2[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint32BigEndian, "Incorrect copy")

	n = ToU32(data.Bytes[:], dst2[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint32LittleEndian, "Incorrect copy")

	n = ToF32(data.Bytes[:], dst3[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst3 == data.Float32BigEndian, "Incorrect copy")

	n = ToF32(data.Bytes[:], dst3[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst3 == data.Float32LittleEndian, "Incorrect copy")
}

func (t *testSlice) TestFrom64(is is.IS) {
	var dst [len(data.Bytes)]byte
	n := FromI64(data.Int64BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromI64(data.Int64LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU64(data.Uint64BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromU64(data.Uint64LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromF64(data.Float64BigEndian[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")

	n = FromF64(data.Float64LittleEndian[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Bytes, "Incorrect copy")
}

func (t *testSlice) TestTo64(is is.IS) {
	var dst [len(data.Int64BigEndian)]int64
	var dst2 [len(data.Int64BigEndian)]uint64
	var dst3 [len(data.Int64BigEndian)]float64

	n := ToI64(data.Bytes[:], dst[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int64BigEndian, "Incorrect copy")

	n = ToI64(data.Bytes[:], dst[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst == data.Int64LittleEndian, "Incorrect copy")

	n = ToU64(data.Bytes[:], dst2[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint64BigEndian, "Incorrect copy")

	n = ToU64(data.Bytes[:], dst2[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst2 == data.Uint64LittleEndian, "Incorrect copy")

	n = ToF64(data.Bytes[:], dst3[:], t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst3 == data.Float64BigEndian, "Incorrect copy")

	n = ToF64(data.Bytes[:], dst3[:], !t.rotateBigEndian)
	is.True(n == len(data.Bytes), "Not all bytes copied")
	is.True(dst3 == data.Float64LittleEndian, "Incorrect copy")
}

func (t *testSlice) TestFrom(is is.IS) {
	var dst [len(data.Bytes)]byte

	for i, d := range data.AllBigEndian {
		n, err := FromSlice(d, dst[:], t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(dst == data.Bytes, "Incorrect copy", dst, data.Bytes, i)

		n, err = FromValue(reflect.ValueOf(d), dst[:], t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(dst == data.Bytes, "Incorrect copy")
	}

	for _, d := range data.AllLittleEndian {
		n, err := FromValue(reflect.ValueOf(d), dst[:], !t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(dst == data.Bytes, "Incorrect copy")

		n, err = FromSlice(d, dst[:], !t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(dst == data.Bytes, "Incorrect copy")
	}

}

func (t *testSlice) NewOf(orig reflect.Value) reflect.Value {
	typ := orig.Type()
	var v reflect.Value
	switch typ.Kind() {
	case reflect.Slice:
		v = reflect.MakeSlice(typ, orig.Len(), orig.Len())
	case reflect.Array:
		v = reflect.New(typ).Elem()
	case reflect.Ptr:
		v = reflect.New(typ.Elem())
	}
	return v
}

func (t *testSlice) Compare(v1, v2 reflect.Value) bool {
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
		v2 = v2.Elem()
	}
	for i := 0; i < v1.Len(); i++ {
		if v1.Index(i).Interface() != v2.Index(i).Interface() {
			return false
		}
	}
	return true
}

func (t *testSlice) TestSlice(is is.IS) {
	for _, d := range data.AllBigEndian {
		orig := reflect.ValueOf(d)
		iface := t.NewOf(orig).Interface()
		n, err := ToSlice(data.Bytes[:], iface, t.rotateBigEndian)
		if orig.Kind() == reflect.Array {
			is.True(err == internal.ErrUnaddressable, "value should be unadressable")
		} else {
			is.Nil(err, "unexpected error")
			is.True(n == len(data.Bytes), "Not All bytes copied")
			is.True(t.Compare(reflect.ValueOf(iface), orig), "incorrect copy")
		}

		v := t.NewOf(orig)
		n, err = ToValue(data.Bytes[:], v, t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(t.Compare(v, orig), "incorrect copy")
	}

	for _, d := range data.AllLittleEndian {
		orig := reflect.ValueOf(d)
		iface := t.NewOf(orig).Interface()
		n, err := ToSlice(data.Bytes[:], iface, !t.rotateBigEndian)
		if orig.Kind() == reflect.Array {
			is.True(err == internal.ErrUnaddressable, "value should be unadressable")
		} else {
			is.Nil(err, "unexpected error")
			is.True(n == len(data.Bytes), "Not All bytes copied")
			is.True(t.Compare(reflect.ValueOf(iface), orig), "incorrect copy")
		}

		v := t.NewOf(orig)
		n, err = ToValue(data.Bytes[:], v, !t.rotateBigEndian)
		is.Nil(err, "unexpected error")
		is.True(n == len(data.Bytes), "Not All bytes copied")
		is.True(t.Compare(v, orig), "incorrect copy")
	}
}

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

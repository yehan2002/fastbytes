package testutil

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/yehan2002/fastbytes/v2/internal"
	"github.com/yehan2002/is/v2"
)

type testRunner struct {
	t               *testing.T
	provider        internal.Provider
	rotateBigEndian bool
}

// Test this tests the given provider.
func Test(t *testing.T, provider internal.Provider, rotateBigEndian bool) {
	is.Suite(t, &testRunner{t: t, provider: provider, rotateBigEndian: rotateBigEndian})
}

func (r *testRunner) TestFrom8(i is.Is) {
	i.RunP("FromI8", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI8(int8Slice[:], dst[:])
		r.checkCopyFrom(is, "From8", dst, n)
	})
}

func (r *testRunner) TestFrom16(i is.Is) {
	i.RunP("FromI16BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI16(Int16BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From16", dst, n)
	})
	i.RunP("FromI16LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI16(int16LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From16", dst, n)
	})
	i.RunP("FromU16BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU16(uint16BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From16", dst, n)
	})
	i.RunP("FromU16LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU16(uint16LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From16", dst, n)
	})
}

func (r *testRunner) TestFrom32(i is.Is) {
	i.RunP("FromI32BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI32(int32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
	i.RunP("FromI32LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI32(int32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
	i.RunP("FromU32BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU32(uint32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
	i.RunP("FromU32LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU32(uint32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
	i.RunP("FromF32BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromF32(float32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
	i.RunP("FromF32LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromF32(float32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From32", dst, n)
	})
}

func (r *testRunner) TestFrom64(i is.Is) {
	i.RunP("FromI64BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI64(int64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
	i.RunP("FromI64LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromI64(int64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
	i.RunP("FromU64BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU64(uint64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
	i.RunP("FromU64LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromU64(uint64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
	i.RunP("FromF64BigEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromF64(float64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
	i.RunP("FromF64LittleEndian", func(is is.Is) {
		var dst [len(bytes)]byte
		n := r.provider.FromF64(float64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(is, "From64", dst, n)
	})
}

func (r *testRunner) TestTo8(i is.Is) {
	i.RunP("TestTo8", func(is is.Is) {
		var dst [len(bytes)]int8
		n := r.provider.ToI8(bytes[:], dst[:])
		r.checkCopy(is, "ToI8", dst, n, int8Slice)
	})

}

func (r *testRunner) TestTo16(i is.Is) {
	i.RunP("ToI16BigEndian", func(is is.Is) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI16", dst, n, Int16BigEndian)
	})
	i.RunP("ToI16LittleEndian", func(is is.Is) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI16", dst, n, int16LittleEndian)
	})
	i.RunP("ToU16BigEndian", func(is is.Is) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI16", dst, n, uint16BigEndian)
	})
	i.RunP("ToU16LittleEndian", func(is is.Is) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI16", dst, n, uint16LittleEndian)
	})
}

func (r *testRunner) TestTo32(i is.Is) {
	i.RunP("ToI32BigEndian", func(is is.Is) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, int32BigEndian)
	})
	i.RunP("ToI32LittleEndian", func(is is.Is) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, int32LittleEndian)
	})
	i.RunP("ToU32BigEndian", func(is is.Is) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, uint32BigEndian)
	})
	i.RunP("ToU32LittleEndian", func(is is.Is) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, uint32LittleEndian)
	})
	i.RunP("ToF32BigEndian", func(is is.Is) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, float32BigEndian)
	})
	i.RunP("ToF32LittleEndian", func(is is.Is) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI32", dst, n, float32LittleEndian)
	})
}

func (r *testRunner) TestTo64(i is.Is) {
	i.RunP("ToI64BigEndian", func(is is.Is) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, int64BigEndian)
	})
	i.RunP("ToI64LittleEndian", func(is is.Is) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, int64LittleEndian)
	})
	i.RunP("ToU64BigEndian", func(is is.Is) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, uint64BigEndian)
	})
	i.RunP("ToU64LittleEndian", func(is is.Is) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, uint64LittleEndian)
	})
	i.RunP("ToF64BigEndian", func(is is.Is) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, float64BigEndian)
	})
	i.RunP("ToF64LittleEndian", func(is is.Is) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(is, "ToI64", dst, n, float64LittleEndian)
	})
}

func (r *testRunner) TestFromSlice(i is.Is) {
	fromSlice := func(name string, src func(bigEndian bool) any) {
		i.RunP(name, func(is is.Is) {
			var dst [len(bytes)]byte
			n, err := r.provider.FromSlice(src(true), dst[:], r.rotateBigEndian)
			r.checkSlice(is, name+" (BigEndian)", err, dst, n, bytes)

			dst = [len(bytes)]byte{}
			n, err = r.provider.FromSlice(src(false), dst[:], !r.rotateBigEndian)
			r.checkSlice(is, name+" (LittleEndian)", err, dst, n, bytes)
		})
	}

	for _, d := range TestData {
		name := d.Name()
		fromSlice(name, d.Array)
		fromSlice(name, d.ArrayPtr)
		fromSlice(name, d.Slice)
	}
}

func (r *testRunner) TestFromValue(i is.Is) {
	fromValue := func(name string, src func(bigEndian bool) reflect.Value) {
		i.RunP(name, func(is is.Is) {
			var dst [len(bytes)]byte
			n, err := r.provider.FromValue(src(true), dst[:], r.rotateBigEndian)
			r.checkSlice(is, name+" (BigEndian)", err, dst, n, bytes)

			dst = [len(bytes)]byte{}
			n, err = r.provider.FromValue(src(false), dst[:], !r.rotateBigEndian)
			r.checkSlice(is, name+" (LittleEndian)", err, dst, n, bytes)
		})
	}

	for _, d := range TestData {
		name := d.Name()
		fromValue(name, d.ValueArray)
		fromValue(name, d.ValueArrayUnaddressable)
		fromValue(name, d.ValueArrayPtr)
		fromValue(name, d.ValueSlice)
	}
}

func (r *testRunner) TestToSlice(i is.Is) {
	toSlice := func(name string, create func() any, expected func(bigEndian bool) any) {
		i.RunP(name, func(is is.Is) {
			dst := create()
			n, err := r.provider.ToSlice(bytes[:], dst, r.rotateBigEndian)
			r.checkSlice(is, name+" (BigEndian)", err, dst, n, expected(true))

			dst = create()
			n, err = r.provider.ToSlice(bytes[:], dst, !r.rotateBigEndian)
			r.checkSlice(is, name+" (LittleEndian)", err, dst, n, expected(false))
		})
	}

	for _, d := range TestData {
		name := d.Name()
		toSlice(name, d.NewArrayPtr, d.ArrayPtr)
		toSlice(name, d.NewSlice, d.Slice)

		_, err := r.provider.ToSlice(bytes[:], d.NewArray(), true)
		if !errors.Is(err, internal.ErrUnaddressable) {
			r.t.Error("ToSlice allowed unsafe copy")
			r.t.FailNow()
		}
	}
}

func (r *testRunner) TestToValue(i is.Is) {
	toValue := func(name string, create func() reflect.Value, expected func(bigEndian bool) reflect.Value) {
		i.RunP(name, func(is is.Is) {
			dst := create()
			n, err := r.provider.ToValue(bytes[:], dst, r.rotateBigEndian)
			r.checkSlice(is, name+" (BigEndian)", err, dst.Interface(), n, expected(true).Interface())

			dst = create()
			n, err = r.provider.ToValue(bytes[:], dst, !r.rotateBigEndian)
			r.checkSlice(is, name+" (LittleEndian)", err, dst.Interface(), n, expected(false).Interface())
		})
	}
	for _, d := range TestData {
		name := d.Name()
		toValue(name, d.NewArrayValue, d.ValueArray)
		toValue(name, d.NewArrayPtrValue, d.ValueArrayPtr)
		toValue(name, d.NewSliceValue, d.ValueSlice)
	}
}

func (r *testRunner) checkSlice(is is.Is, name string, err error, v any, length int, expectedValue any) {
	is.T().Helper()
	is(err == nil, "%s unexpected error %s %s %s", name, err, v, expectedValue)
	r.checkCopy(is, name, v, length, expectedValue)
}

func (r *testRunner) checkCopy(is is.Is, name string, v any, length int, expectedValue any) {
	is.T().Helper()
	is(length == len(bytes), "%s copied an incorrect number of bytes. copied %d expected %d", name, length, len(bytes))

	diff := cmp.Diff(v, expectedValue, cmpopts.EquateNaNs())
	if diff != "" {
		is.Fail(name + " copied incorrectly\nValues are not equal:\n" + diff)
	}
}

func (r *testRunner) checkCopyFrom(is is.Is, name string, v any, length int) {
	is.T().Helper()
	r.checkCopy(is, name, v, length, bytes)
}

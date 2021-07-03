package testdata

import (
	"errors"
	"reflect"
	"testing"

	"github.com/yehan2002/fastbytes/internal"
	"github.com/yehan2002/is"
)

type testRunner struct {
	t               *testing.T
	provider        internal.Provider
	rotateBigEndian bool
}

//Test this tests the given provider.
func Test(t *testing.T, provider internal.Provider, rotateBigEndian bool) {
	r := testRunner{t: t, provider: provider, rotateBigEndian: rotateBigEndian}
	r.testFrom8()
	r.testFrom16()
	r.testFrom32()
	r.testFrom64()
	r.testTo8()
	r.testTo16()
	r.testTo32()
	r.testTo64()
	r.testFromSlice()
	r.testToSlice()
	r.testFromValue()
	r.testToValue()
}

func (r *testRunner) run(name string, test func(t *testing.T)) {
	r.t.Run(name, func(t *testing.T) { t.Parallel(); test(t) })
}

func (r *testRunner) testFrom8() {
	r.run("FromI8", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI8(int8Slice[:], dst[:])
		r.checkCopyFrom(t, "From8", dst, n)
	})
}

func (r *testRunner) testFrom16() {
	r.run("FromI16BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI16(Int16BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From16", dst, n)
	})
	r.run("FromI16LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI16(int16LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From16", dst, n)
	})
	r.run("FromU16BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU16(uint16BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From16", dst, n)
	})
	r.run("FromU16LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU16(uint16LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From16", dst, n)
	})
}

func (r *testRunner) testFrom32() {
	r.run("FromI32BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI32(int32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
	r.run("FromI32LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI32(int32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
	r.run("FromU32BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU32(uint32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
	r.run("FromU32LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU32(uint32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
	r.run("FromF32BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromF32(float32BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
	r.run("FromF32LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromF32(float32LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From32", dst, n)
	})
}

func (r *testRunner) testFrom64() {
	r.run("FromI64BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI64(int64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
	r.run("FromI64LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI64(int64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
	r.run("FromU64BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU64(uint64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
	r.run("FromU64LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromU64(uint64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
	r.run("FromF64BigEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromF64(float64BigEndian[:], dst[:], r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
	r.run("FromF64LittleEndian", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromF64(float64LittleEndian[:], dst[:], !r.rotateBigEndian)
		r.checkCopyFrom(t, "From64", dst, n)
	})
}

func (r *testRunner) testTo8() {
	r.run("TestTo8", func(t *testing.T) {
		var dst [len(bytes)]int8
		n := r.provider.ToI8(bytes[:], dst[:])
		r.checkCopy(t, "ToI8", dst, n, int8Slice)
	})

}

func (r *testRunner) testTo16() {
	r.run("ToI16BigEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, Int16BigEndian)
	})
	r.run("ToI16LittleEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, int16LittleEndian)
	})
	r.run("ToU16BigEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, uint16BigEndian)
	})
	r.run("ToU16LittleEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, uint16LittleEndian)
	})
}

func (r *testRunner) testTo32() {
	r.run("ToI32BigEndian", func(t *testing.T) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, int32BigEndian)
	})
	r.run("ToI32LittleEndian", func(t *testing.T) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, int32LittleEndian)
	})
	r.run("ToU32BigEndian", func(t *testing.T) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, uint32BigEndian)
	})
	r.run("ToU32LittleEndian", func(t *testing.T) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, uint32LittleEndian)
	})
	r.run("ToF32BigEndian", func(t *testing.T) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, float32BigEndian)
	})
	r.run("ToF32LittleEndian", func(t *testing.T) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, float32LittleEndian)
	})
}

func (r *testRunner) testTo64() {
	r.run("ToI64BigEndian", func(t *testing.T) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, int64BigEndian)
	})
	r.run("ToI64LittleEndian", func(t *testing.T) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, int64LittleEndian)
	})
	r.run("ToU64BigEndian", func(t *testing.T) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, uint64BigEndian)
	})
	r.run("ToU64LittleEndian", func(t *testing.T) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, uint64LittleEndian)
	})
	r.run("ToF64BigEndian", func(t *testing.T) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, float64BigEndian)
	})
	r.run("ToF64LittleEndian", func(t *testing.T) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, float64LittleEndian)
	})
}

func (r *testRunner) testFromSlice() {
	fromSlice := func(name string, src func(bigEndian bool) interface{}) {
		r.t.Run(name, func(t *testing.T) {
			var dst [len(bytes)]byte
			n, err := r.provider.FromSlice(src(true), dst[:], r.rotateBigEndian)
			r.checkSlice(t, name+" (BigEndian)", err, dst, n, bytes)

			dst = [len(bytes)]byte{}
			n, err = r.provider.FromSlice(src(false), dst[:], !r.rotateBigEndian)
			r.checkSlice(t, name+" (LittleEndian)", err, dst, n, bytes)
		})
	}

	for _, d := range TestData {
		name := "FromSlice with element type " + d.Name()
		fromSlice(name, d.Array)
		fromSlice(name, d.ArrayPtr)
		fromSlice(name, d.Slice)
	}
}

func (r *testRunner) testFromValue() {
	fromValue := func(name string, src func(bigEndian bool) reflect.Value) {
		r.t.Run(name, func(t *testing.T) {
			var dst [len(bytes)]byte
			n, err := r.provider.FromValue(src(true), dst[:], r.rotateBigEndian)
			r.checkSlice(t, name+" (BigEndian)", err, dst, n, bytes)

			dst = [len(bytes)]byte{}
			n, err = r.provider.FromValue(src(false), dst[:], !r.rotateBigEndian)
			r.checkSlice(t, name+" (LittleEndian)", err, dst, n, bytes)
		})
	}

	for _, d := range TestData {
		name := "FromValue with element type " + d.Name()
		fromValue(name, d.ValueArray)
		fromValue(name, d.ValueArrayUnaddressable)
		fromValue(name, d.ValueArrayPtr)
		fromValue(name, d.ValueSlice)
	}
}

func (r *testRunner) testToSlice() {
	toSlice := func(name string, create func() interface{}, expected func(bigEndian bool) interface{}) {
		r.t.Run(name, func(t *testing.T) {
			dst := create()
			n, err := r.provider.ToSlice(bytes[:], dst, r.rotateBigEndian)
			r.checkSlice(t, name+" (BigEndian)", err, dst, n, expected(true))

			dst = create()
			n, err = r.provider.ToSlice(bytes[:], dst, !r.rotateBigEndian)
			r.checkSlice(t, name+" (LittleEndian)", err, dst, n, expected(false))
		})
	}

	for _, d := range TestData {
		name := "ToSlice with element type " + d.Name()
		toSlice(name, d.NewArrayPtr, d.ArrayPtr)
		toSlice(name, d.NewSlice, d.Slice)

		_, err := r.provider.ToSlice(bytes[:], d.NewArray(), true)
		if !errors.Is(err, internal.ErrUnaddressable) {
			r.t.Error("ToSlice allowed unsafe copy")
			r.t.FailNow()
		}
	}
}

func (r *testRunner) testToValue() {
	toValue := func(name string, create func() reflect.Value, expected func(bigEndian bool) reflect.Value) {
		r.t.Run(name, func(t *testing.T) {
			dst := create()
			n, err := r.provider.ToValue(bytes[:], dst, r.rotateBigEndian)
			r.checkSlice(t, name+" (BigEndian)", err, dst.Interface(), n, expected(true).Interface())

			dst = create()
			n, err = r.provider.ToValue(bytes[:], dst, !r.rotateBigEndian)
			r.checkSlice(t, name+" (LittleEndian)", err, dst.Interface(), n, expected(false).Interface())
		})
	}

	for _, d := range TestData {
		name := "ToValue with element type " + d.Name()
		toValue(name, d.NewArrayValue, d.ValueArray)
		toValue(name, d.NewArrayPtrValue, d.ValueArrayPtr)
		toValue(name, d.NewSliceValue, d.ValueSlice)
	}
}

func (r *testRunner) checkSlice(t *testing.T, name string, err error, v interface{}, length int, expectedValue interface{}) {
	if err != nil {
		t.Error(name, " unexpected error", err, v, expectedValue)
		panic(err)
	}
	r.checkCopy(t, name, v, length, expectedValue)
}

func (r *testRunner) checkCopy(t *testing.T, name string, v interface{}, length int, expectedValue interface{}) {
	if length != len(bytes) {
		t.Error("copied", length, "expected", len(bytes))
		panic(name + " copied an incorrect number of bytes")
	}
	if !reflect.DeepEqual(v, expectedValue) {
		is := is.New(t)
		is.Equal(v, expectedValue)
		/* t.Errorf("Got: %#v", v)
		t.Errorf("Expected: %#v", expectedValue)

		panic(name + " copied incorrectly") */
	}
}

func (r *testRunner) checkCopyFrom(t *testing.T, name string, v interface{}, length int) {
	r.checkCopy(t, name, v, length, bytes)
}

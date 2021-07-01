package testdata

import (
	"errors"
	"reflect"
	"testing"

	"github.com/yehan2002/bytes/internal"
)

type testRunner struct {
	t               *testing.T
	provider        internal.Provider
	rotateBigEndian bool
}

//Test this tests the given provider.
func Test(t *testing.T, provider internal.Provider, rotateBigEndian bool) {
	r := testRunner{t: t, provider: provider, rotateBigEndian: rotateBigEndian}
	r.testFrom()
	r.testTo()
	r.testFromSlice()
	r.testToSlice()
}

func (r *testRunner) run(name string, test func(t *testing.T)) {
	r.t.Run(name, func(t *testing.T) { t.Parallel(); test(t) })
}

func (r *testRunner) testFrom() {
	r.run("FromI8", func(t *testing.T) {
		var dst [len(bytes)]byte
		n := r.provider.FromI8(int8Slice[:], dst[:])
		r.checkCopyFrom(t, "From8", dst, n)
	})
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

func (r *testRunner) testTo() {
	r.run("TestTo8", func(t *testing.T) {
		var dst [len(bytes)]int8
		n := r.provider.ToI8(bytes[:], dst[:])
		r.checkCopy(t, "ToI8", dst, n, int8Slice, len(bytes))
	})
	r.run("ToI16BigEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, Int16BigEndian, len(bytes))
	})
	r.run("ToI16LittleEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]int16
		n := r.provider.ToI16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, int16LittleEndian, len(bytes))
	})
	r.run("ToU16BigEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, uint16BigEndian, len(bytes))
	})
	r.run("ToU16LittleEndian", func(t *testing.T) {
		var dst [len(Int16BigEndian)]uint16
		n := r.provider.ToU16(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI16", dst, n, uint16LittleEndian, len(bytes))
	})
	r.run("ToI32BigEndian", func(t *testing.T) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, int32BigEndian, len(bytes))
	})
	r.run("ToI32LittleEndian", func(t *testing.T) {
		var dst [len(int32BigEndian)]int32
		n := r.provider.ToI32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, int32LittleEndian, len(bytes))
	})
	r.run("ToU32BigEndian", func(t *testing.T) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, uint32BigEndian, len(bytes))
	})
	r.run("ToU32LittleEndian", func(t *testing.T) {
		var dst [len(uint32BigEndian)]uint32
		n := r.provider.ToU32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, uint32LittleEndian, len(bytes))
	})
	r.run("ToF32BigEndian", func(t *testing.T) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, float32BigEndian, len(bytes))
	})
	r.run("ToF32LittleEndian", func(t *testing.T) {
		var dst [len(float32BigEndian)]float32
		n := r.provider.ToF32(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI32", dst, n, float32LittleEndian, len(bytes))
	})
	r.run("ToI64BigEndian", func(t *testing.T) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, int64BigEndian, len(bytes))
	})
	r.run("ToI64LittleEndian", func(t *testing.T) {
		var dst [len(int64BigEndian)]int64
		n := r.provider.ToI64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, int64LittleEndian, len(bytes))
	})
	r.run("ToU64BigEndian", func(t *testing.T) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, uint64BigEndian, len(bytes))
	})
	r.run("ToU64LittleEndian", func(t *testing.T) {
		var dst [len(uint64BigEndian)]uint64
		n := r.provider.ToU64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, uint64LittleEndian, len(bytes))
	})
	r.run("ToF64BigEndian", func(t *testing.T) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, float64BigEndian, len(bytes))
	})
	r.run("ToF64LittleEndian", func(t *testing.T) {
		var dst [len(float64BigEndian)]float64
		n := r.provider.ToF64(bytes[:], dst[:], !r.rotateBigEndian)
		r.checkCopy(t, "ToI64", dst, n, float64LittleEndian, len(bytes))
	})
}

func (r *testRunner) fromSlice(name string, src func(bigEndian bool) interface{}) {
	r.t.Run(name, func(t *testing.T) {
		var dst [len(bytes)]byte
		n, err := r.provider.FromSlice(src(true), dst[:], r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, bytes, len(bytes))

		dst = [len(bytes)]byte{}
		n, err = r.provider.FromSlice(src(false), dst[:], !r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, bytes, len(bytes))
	})
}

func (r *testRunner) fromValue(name string, src func(bigEndian bool) reflect.Value) {
	r.t.Run(name, func(t *testing.T) {
		var dst [len(bytes)]byte
		n, err := r.provider.FromValue(src(true), dst[:], r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, bytes, len(bytes))

		dst = [len(bytes)]byte{}
		n, err = r.provider.FromValue(src(false), dst[:], !r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, bytes, len(bytes))
	})
}

func (r *testRunner) testFromSlice() {

	for _, d := range TestData {
		name := "FromSlice with element type " + d.Name()
		r.fromSlice(name, d.Array)
		r.fromSlice(name, d.ArrayPtr)
		r.fromSlice(name, d.Slice)

		name = "FromValue with element type " + d.Name()
		r.fromValue(name, d.ValueArray)
		r.fromValue(name, d.ValueArrayUnaddressable)
		r.fromValue(name, d.ValueArrayPtr)
		r.fromValue(name, d.ValueSlice)
	}
}

func (r *testRunner) toSlice(name string, create func() interface{}, expected func(bigEndian bool) interface{}) {
	r.t.Run(name, func(t *testing.T) {
		dst := create()
		n, err := r.provider.ToSlice(bytes[:], dst, r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, expected(true), len(bytes))

		dst = create()
		n, err = r.provider.ToSlice(bytes[:], dst, !r.rotateBigEndian)
		r.checkSlice(t, name, err, dst, n, expected(false), len(bytes))
	})
}

func (r *testRunner) toValue(name string, create func() reflect.Value, expected func(bigEndian bool) reflect.Value) {
	r.t.Run(name, func(t *testing.T) {
		dst := create()
		n, err := r.provider.ToValue(bytes[:], dst, r.rotateBigEndian)
		r.checkSlice(t, name, err, dst.Interface(), n, expected(true).Interface(), len(bytes))

		dst = create()
		n, err = r.provider.ToValue(bytes[:], dst, !r.rotateBigEndian)
		r.checkSlice(t, name, err, dst.Interface(), n, expected(false).Interface(), len(bytes))
	})
}

func (r *testRunner) testToSlice() {

	for _, d := range TestData {
		name := "ToSlice with element type " + d.Name()
		r.toSlice(name, d.NewArrayPtr, d.ArrayPtr)
		r.toSlice(name, d.NewSlice, d.Slice)

		name = "ToValue with element type " + d.Name()
		r.toValue(name, d.NewArrayValue, d.ValueArray)
		r.toValue(name, d.NewArrayPtrValue, d.ValueArrayPtr)
		r.toValue(name, d.NewSliceValue, d.ValueSlice)
	}

	for _, d := range TestData {
		dst := d.NewArray()
		_, err := r.provider.ToSlice(bytes[:], dst, true)
		if !errors.Is(err, internal.ErrUnaddressable) {
			r.t.Error("ToSlice allowed unsafe copy")
			r.t.FailNow()
		}
	}
}

func (r *testRunner) checkSlice(t *testing.T, name string, err error, v interface{}, length int, expectedValue interface{}, expectedLength int) {
	if err != nil {
		t.Error(name, " unexpected error", err, v, expectedValue)
		panic(err)
	}
	r.checkCopy(t, name, v, length, expectedValue, expectedLength)
}

func (r *testRunner) checkCopy(t *testing.T, name string, v interface{}, length int, expectedValue interface{}, expectedLength int) {
	if length != expectedLength {
		t.Error("copied", length, "expected", expectedLength)
		panic(name + " copied an incorrect number of bytes")
	}
	if !reflect.DeepEqual(v, expectedValue) {
		t.Error(v, expectedValue)
		panic(name + " copied incorrectly")
	}
}

func (r *testRunner) checkCopyFrom(t *testing.T, name string, v interface{}, length int) {
	r.checkCopy(t, name, v, length, bytes, len(bytes))
}

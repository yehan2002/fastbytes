package unsafe

import (
	"errors"
	"reflect"
	"testing"

	"github.com/yehan2002/bytes/internal"
	"github.com/yehan2002/bytes/internal/data"
)

func TestFrom8(t *testing.T) {
	t.Parallel()
	var dst [len(data.Bytes)]byte
	n := FromI8(data.Int8[:], dst[:])
	checkCopyFrom(t, "From8", dst, n)
}

func TestFrom16(t *testing.T) {
	t.Parallel()
	t.Run("FromI16BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI16(data.Int16BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From16", dst, n)
	})
	t.Run("FromI16LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI16(data.Int16LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From16", dst, n)
	})
	t.Run("FromU16BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU16(data.Uint16BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From16", dst, n)
	})
	t.Run("FromU16LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU16(data.Uint16LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From16", dst, n)
	})
}

func TestFrom32(t *testing.T) {
	t.Parallel()
	t.Run("FromI32BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI32(data.Int32BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
	t.Run("FromI32LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI32(data.Int32LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
	t.Run("FromU32BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU32(data.Uint32BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
	t.Run("FromU32LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU32(data.Uint32LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
	t.Run("FromF32BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromF32(data.Float32BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
	t.Run("FromF32LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromF32(data.Float32LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From32", dst, n)
	})
}

func TestFrom64(t *testing.T) {
	t.Parallel()
	t.Run("FromI64BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI64(data.Int64BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
	t.Run("FromI64LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromI64(data.Int64LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
	t.Run("FromU64BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU64(data.Uint64BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
	t.Run("FromU64LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromU64(data.Uint64LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
	t.Run("FromF64BigEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromF64(data.Float64BigEndian[:], dst[:], rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
	t.Run("FromF64LittleEndian", func(t *testing.T) {
		var dst [len(data.Bytes)]byte
		n := FromF64(data.Float64LittleEndian[:], dst[:], !rotateBigEndian)
		checkCopyFrom(t, "From64", dst, n)
	})
}

func TestTo8(t *testing.T) {
	t.Parallel()
	var dst [len(data.Bytes)]int8
	n := ToI8(data.Bytes[:], dst[:])
	checkCopy(t, "ToI8", dst, n, data.Int8, len(data.Bytes))
}

func TestTo16(t *testing.T) {
	t.Parallel()
	t.Run("ToI16BigEndian", func(t *testing.T) {
		var dst [len(data.Int16BigEndian)]int16
		n := ToI16(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI16", dst, n, data.Int16BigEndian, len(data.Bytes))
	})
	t.Run("ToI16LittleEndian", func(t *testing.T) {
		var dst [len(data.Int16BigEndian)]int16
		n := ToI16(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI16", dst, n, data.Int16LittleEndian, len(data.Bytes))
	})
	t.Run("ToU16BigEndian", func(t *testing.T) {
		var dst [len(data.Int16BigEndian)]uint16
		n := ToU16(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI16", dst, n, data.Uint16BigEndian, len(data.Bytes))
	})
	t.Run("ToU16LittleEndian", func(t *testing.T) {
		var dst [len(data.Int16BigEndian)]uint16
		n := ToU16(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI16", dst, n, data.Uint16LittleEndian, len(data.Bytes))
	})
}

func TestTo32(t *testing.T) {
	t.Parallel()
	t.Run("ToI32BigEndian", func(t *testing.T) {
		var dst [len(data.Int32BigEndian)]int32
		n := ToI32(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Int32BigEndian, len(data.Bytes))
	})
	t.Run("ToI32LittleEndian", func(t *testing.T) {
		var dst [len(data.Int32BigEndian)]int32
		n := ToI32(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Int32LittleEndian, len(data.Bytes))
	})
	t.Run("ToU32BigEndian", func(t *testing.T) {
		var dst [len(data.Uint32BigEndian)]uint32
		n := ToU32(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Uint32BigEndian, len(data.Bytes))
	})
	t.Run("ToU32LittleEndian", func(t *testing.T) {
		var dst [len(data.Uint32BigEndian)]uint32
		n := ToU32(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Uint32LittleEndian, len(data.Bytes))
	})
	t.Run("ToF32BigEndian", func(t *testing.T) {
		var dst [len(data.Float32BigEndian)]float32
		n := ToF32(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Float32BigEndian, len(data.Bytes))
	})
	t.Run("ToF32LittleEndian", func(t *testing.T) {
		var dst [len(data.Float32BigEndian)]float32
		n := ToF32(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI32", dst, n, data.Float32LittleEndian, len(data.Bytes))
	})
}

func TestTo64(t *testing.T) {
	t.Parallel()
	t.Run("ToI64BigEndian", func(t *testing.T) {
		var dst [len(data.Int64BigEndian)]int64
		n := ToI64(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Int64BigEndian, len(data.Bytes))
	})
	t.Run("ToI64LittleEndian", func(t *testing.T) {
		var dst [len(data.Int64BigEndian)]int64
		n := ToI64(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Int64LittleEndian, len(data.Bytes))
	})
	t.Run("ToU64BigEndian", func(t *testing.T) {
		var dst [len(data.Uint64BigEndian)]uint64
		n := ToU64(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Uint64BigEndian, len(data.Bytes))
	})
	t.Run("ToU64LittleEndian", func(t *testing.T) {
		var dst [len(data.Uint64BigEndian)]uint64
		n := ToU64(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Uint64LittleEndian, len(data.Bytes))
	})
	t.Run("ToF64BigEndian", func(t *testing.T) {
		var dst [len(data.Float64BigEndian)]float64
		n := ToF64(data.Bytes[:], dst[:], rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Float64BigEndian, len(data.Bytes))
	})
	t.Run("ToF64LittleEndian", func(t *testing.T) {
		var dst [len(data.Float64BigEndian)]float64
		n := ToF64(data.Bytes[:], dst[:], !rotateBigEndian)
		checkCopy(t, "ToI64", dst, n, data.Float64LittleEndian, len(data.Bytes))
	})
}

func TestFromSlice(t *testing.T) {
	t.Parallel()

	fromSlice := func(name string, src func(bigEndian bool) interface{}) {
		var dst [len(data.Bytes)]byte
		n, err := FromSlice(src(true), dst[:], rotateBigEndian)
		checkSlice(t, name, err, dst, n, data.Bytes, len(data.Bytes))

		dst = [len(data.Bytes)]byte{}
		n, err = FromSlice(src(false), dst[:], !rotateBigEndian)
		checkSlice(t, name, err, dst, n, data.Bytes, len(data.Bytes))
	}

	fromValue := func(name string, src func(bigEndian bool) reflect.Value) {
		var dst [len(data.Bytes)]byte
		n, err := FromValue(src(true), dst[:], rotateBigEndian)
		checkSlice(t, name, err, dst, n, data.Bytes, len(data.Bytes))

		dst = [len(data.Bytes)]byte{}
		n, err = FromValue(src(false), dst[:], !rotateBigEndian)
		checkSlice(t, name, err, dst, n, data.Bytes, len(data.Bytes))
	}

	for _, d := range data.TestData {
		name := "FromSlice with element type " + d.Name()
		fromSlice(name, d.Array)
		fromSlice(name, d.ArrayPtr)
		fromSlice(name, d.Slice)

		name = "FromValue with element type " + d.Name()
		fromValue(name, d.ValueArray)
		fromValue(name, d.ValueArrayUnaddressable)
		fromValue(name, d.ValueArrayPtr)
		fromValue(name, d.ValueSlice)
	}
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	toSlice := func(name string, create func() interface{}, expected func(bigEndian bool) interface{}) {
		dst := create()
		n, err := ToSlice(data.Bytes[:], dst, rotateBigEndian)
		checkSlice(t, name, err, dst, n, expected(true), len(data.Bytes))

		dst = create()
		n, err = ToSlice(data.Bytes[:], dst, !rotateBigEndian)
		checkSlice(t, name, err, dst, n, expected(false), len(data.Bytes))
	}

	toValue := func(name string, create func() reflect.Value, expected func(bigEndian bool) reflect.Value) {
		dst := create()
		n, err := ToValue(data.Bytes[:], dst, rotateBigEndian)
		checkSlice(t, name, err, dst.Interface(), n, expected(true).Interface(), len(data.Bytes))

		dst = create()
		n, err = ToValue(data.Bytes[:], dst, !rotateBigEndian)
		checkSlice(t, name, err, dst.Interface(), n, expected(false).Interface(), len(data.Bytes))
	}

	for _, d := range data.TestData {
		name := "ToSlice with element type " + d.Name()
		toSlice(name, d.NewArrayPtr, d.ArrayPtr)
		toSlice(name, d.NewSlice, d.Slice)

		name = "ToValue with element type " + d.Name()
		toValue(name, d.NewArrayValue, d.ValueArray)
		toValue(name, d.NewArrayPtrValue, d.ValueArrayPtr)
		toValue(name, d.NewSliceValue, d.ValueSlice)
	}

	for _, d := range data.TestData {
		dst := d.NewArray()
		_, err := ToSlice(data.Bytes[:], dst, true)
		if !errors.Is(err, internal.ErrUnaddressable) {
			t.Error("ToSlice allowed unsafe copy")
			t.FailNow()
		}
	}
}

func checkSlice(t *testing.T, name string, err error, v interface{}, length int, expectedValue interface{}, expectedLength int) {
	if err != nil {
		t.Error(name, " unexpected error", err, v, expectedValue)
		panic(err)
	}
	checkCopy(t, name, v, length, expectedValue, expectedLength)
}

func checkCopy(t *testing.T, name string, v interface{}, length int, expectedValue interface{}, expectedLength int) {
	if length != expectedLength {
		t.Error("copied", length, "expected", expectedLength)
		panic(name + " copied an incorrect number of bytes")
	}
	if !reflect.DeepEqual(v, expectedValue) {
		t.Error(v, expectedValue)
		panic(name + " copied incorrectly")
	}
}

func checkCopyFrom(t *testing.T, name string, v interface{}, length int) {
	checkCopy(t, name, v, length, data.Bytes, len(data.Bytes))
}

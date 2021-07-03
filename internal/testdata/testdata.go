package testdata

import (
	"encoding/binary"
	"math"
	"reflect"
)

const size = 256 * 16

var bytes = func() (v [size]byte) {
	for h := 0; h < 16; h++ {
		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				v[uint16(i<<4|j)|uint16(h<<8)] = byte(i<<4|j) ^ byte(h)
			}
		}
	}
	return
}()

// testdata
var (
	int8Slice [size]int8

	uint16BigEndian [size >> 1]uint16
	Int16BigEndian  [size >> 1]int16

	uint16LittleEndian [size >> 1]uint16
	int16LittleEndian  [size >> 1]int16

	uint32BigEndian  [size >> 2]uint32
	int32BigEndian   [size >> 2]int32
	float32BigEndian [size >> 2]float32

	uint32LittleEndian  [size >> 2]uint32
	int32LittleEndian   [size >> 2]int32
	float32LittleEndian [size >> 2]float32

	uint64BigEndian  [size >> 3]uint64
	int64BigEndian   [size >> 3]int64
	float64BigEndian [size >> 3]float64

	uint64LittleEndian  [size >> 3]uint64
	int64LittleEndian   [size >> 3]int64
	float64LittleEndian [size >> 3]float64
)

func init() {
	for i, b := range bytes {
		int8Slice[i] = int8(b)
	}

	appendData(bytes, bytes)
	appendData(int8Slice, int8Slice)

	for i := 0; i < len(bytes); i += 2 {
		j := i >> 1

		value := binary.BigEndian.Uint16(bytes[i:])
		uint16BigEndian[j] = value
		Int16BigEndian[j] = int16(value)

		value = binary.LittleEndian.Uint16(bytes[i:])
		uint16LittleEndian[j] = value
		int16LittleEndian[j] = int16(value)
	}

	appendData(uint16LittleEndian, uint16BigEndian)
	appendData(int16LittleEndian, Int16BigEndian)

	for i := 0; i < len(bytes); i += 4 {
		j := i >> 2 //nolint

		value := binary.LittleEndian.Uint32(bytes[i:])
		uint32LittleEndian[j] = value
		int32LittleEndian[j] = int32(value)
		float32LittleEndian[j] = math.Float32frombits(value)

		value = binary.BigEndian.Uint32(bytes[i:])
		uint32BigEndian[j] = value
		int32BigEndian[j] = int32(value)
		float32BigEndian[j] = math.Float32frombits(value)
	}

	appendData(uint32LittleEndian, uint32BigEndian)
	appendData(int32LittleEndian, int32BigEndian)
	appendData(float32LittleEndian, float32BigEndian)

	for i := 0; i < len(bytes); i += 8 {
		j := i >> 3 //nolint

		value := binary.LittleEndian.Uint64(bytes[i:])
		uint64LittleEndian[j] = value
		int64LittleEndian[j] = int64(value)
		float64LittleEndian[j] = math.Float64frombits(value)

		value = binary.BigEndian.Uint64(bytes[i:])
		uint64BigEndian[j] = value
		int64BigEndian[j] = int64(value)
		float64BigEndian[j] = math.Float64frombits(value)
	}

	appendData(uint64LittleEndian, uint64BigEndian)
	appendData(int64LittleEndian, int64BigEndian)
	appendData(float64LittleEndian, float64BigEndian)
}

func appendData(little, big interface{}) {
	typ := reflect.TypeOf(little)
	length := typ.Len()

	littleVPtr, bigVPtr := reflect.New(typ), reflect.New(typ)
	littleV, bigV := littleVPtr.Elem(), bigVPtr.Elem()

	reflect.Copy(littleV, reflect.ValueOf(little))
	reflect.Copy(bigV, reflect.ValueOf(big))

	littleSliceV, bigSliceV := littleV.Slice(0, length), bigV.Slice(0, length)

	d := &Data{
		len:  length,
		name: typ.Elem().Kind().String(),
		typ:  typ,

		v:          [2]interface{}{little, big},
		vValue:     [2]reflect.Value{littleV, bigV},
		vPtr:       [2]interface{}{littleVPtr.Interface(), bigVPtr.Interface()},
		vPtrValue:  [2]reflect.Value{littleVPtr, bigVPtr},
		slice:      [2]interface{}{littleSliceV.Interface(), bigSliceV.Interface()},
		sliceValue: [2]reflect.Value{littleSliceV, bigSliceV},
	}

	TestData = append(TestData, d)
}

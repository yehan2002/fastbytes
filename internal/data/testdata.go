package data

import (
	"encoding/binary"
	"math"
	"reflect"
)

const size = len(Bytes)

// Bytes bytes used for tests
var Bytes = [...]uint8{
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
}

// testdata
var (
	Int8 [size]int8

	Uint16BigEndian [size >> 1]uint16
	Int16BigEndian  [size >> 1]int16

	Uint16LittleEndian [size >> 1]uint16
	Int16LittleEndian  [size >> 1]int16

	Uint32BigEndian  [size >> 2]uint32
	Int32BigEndian   [size >> 2]int32
	Float32BigEndian [size >> 2]float32

	Uint32LittleEndian  [size >> 2]uint32
	Int32LittleEndian   [size >> 2]int32
	Float32LittleEndian [size >> 2]float32

	Uint64BigEndian  [size >> 3]uint64
	Int64BigEndian   [size >> 3]int64
	Float64BigEndian [size >> 3]float64

	Uint64LittleEndian  [size >> 3]uint64
	Int64LittleEndian   [size >> 3]int64
	Float64LittleEndian [size >> 3]float64
)

// AllBigEndian all bigEndian value
var AllBigEndian []interface{}

// AllLittleEndian all bigEndian value
var AllLittleEndian []interface{}

func init() {
	for i, b := range Bytes {
		Int8[i] = int8(b)
	}

	for i := 0; i < len(Bytes); i += 2 {
		j := i >> 1

		value := binary.BigEndian.Uint16(Bytes[i:])
		Uint16BigEndian[j] = value
		Int16BigEndian[j] = int16(value)

		value = binary.LittleEndian.Uint16(Bytes[i:])
		Uint16LittleEndian[j] = value
		Int16LittleEndian[j] = int16(value)
	}

	appendData(Uint16LittleEndian, Uint16BigEndian)
	appendData(Int16LittleEndian, Int16BigEndian)

	for i := 0; i < len(Bytes); i += 4 {
		j := i >> 2 //nolint

		value := binary.LittleEndian.Uint32(Bytes[i:])
		Uint32LittleEndian[j] = value
		Int32LittleEndian[j] = int32(value)
		Float32LittleEndian[j] = math.Float32frombits(value)

		value = binary.BigEndian.Uint32(Bytes[i:])
		Uint32BigEndian[j] = value
		Int32BigEndian[j] = int32(value)
		Float32BigEndian[j] = math.Float32frombits(value)
	}

	for i := 0; i < len(Bytes); i += 8 {
		j := i >> 3 //nolint

		value := binary.LittleEndian.Uint64(Bytes[i:])
		Uint64LittleEndian[j] = value
		Int64LittleEndian[j] = int64(value)
		Float64LittleEndian[j] = math.Float64frombits(value)

		value = binary.BigEndian.Uint64(Bytes[i:])
		Uint64BigEndian[j] = value
		Int64BigEndian[j] = int64(value)
		Float64BigEndian[j] = math.Float64frombits(value)
	}

	AllBigEndian = []interface{}{
		Bytes, Int8, Uint16BigEndian, Int16BigEndian, Uint32BigEndian, Int32BigEndian, Float32BigEndian, Uint64BigEndian, Int64BigEndian, Float64BigEndian,
		&Bytes, &Int8, &Uint16BigEndian, &Int16BigEndian, &Uint32BigEndian, &Int32BigEndian, &Float32BigEndian, &Uint64BigEndian, &Int64BigEndian, &Float64BigEndian,
		Bytes[:], Int8[:], Uint16BigEndian[:], Int16BigEndian[:], Uint32BigEndian[:], Int32BigEndian[:], Float32BigEndian[:], Uint64BigEndian[:], Int64BigEndian[:], Float64BigEndian[:],
	}

	AllLittleEndian = []interface{}{
		Bytes, Int8, Uint16LittleEndian, Int16LittleEndian, Uint32LittleEndian, Int32LittleEndian, Float32LittleEndian, Uint64LittleEndian, Int64LittleEndian, Float64LittleEndian,
		&Bytes, &Int8, &Uint16LittleEndian, &Int16LittleEndian, &Uint32LittleEndian, &Int32LittleEndian, &Float32LittleEndian, &Uint64LittleEndian, &Int64LittleEndian, &Float64LittleEndian,
		Bytes[:], Int8[:], Uint16LittleEndian[:], Int16LittleEndian[:], Uint32LittleEndian[:], Int32LittleEndian[:], Float32LittleEndian[:], Uint64LittleEndian[:], Int64LittleEndian[:], Float64LittleEndian[:],
	}
}

// Data test data
type Data struct {
	len  int
	name string
	typ  reflect.Type

	v     [2]interface{}
	vPtr  [2]interface{}
	slice [2]interface{}

	vValue     [2]reflect.Value
	vPtrValue  [2]reflect.Value
	sliceValue [2]reflect.Value
}

func bToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (d *Data) Array(bigEndian bool) interface{}    { return d.v[bToI(bigEndian)] }
func (d *Data) ArrayPtr(bigEndian bool) interface{} { return d.vPtr[bToI(bigEndian)] }
func (d *Data) ValueArrayUnaddressable(bigEndian bool) reflect.Value {
	return reflect.ValueOf(d.v[bToI(bigEndian)])
}
func (d *Data) Slice(bigEndian bool) interface{} { return d.slice[bToI(bigEndian)] }

func (d *Data) ValueArray(bigEndian bool) reflect.Value    { return d.vValue[bToI(bigEndian)] }
func (d *Data) ValueArrayPtr(bigEndian bool) reflect.Value { return d.vPtrValue[bToI(bigEndian)] }
func (d *Data) ValueSlice(bigEndian bool) reflect.Value    { return d.sliceValue[bToI(bigEndian)] }

func (d *Data) NewArrayValue() reflect.Value    { return reflect.New(d.typ).Elem() }
func (d *Data) NewArrayPtrValue() reflect.Value { return reflect.New(d.typ) }
func (d *Data) NewSliceValue() reflect.Value    { return reflect.New(d.typ).Elem().Slice(0, d.len) }

func (d *Data) NewArray() interface{}    { return d.NewArrayValue().Interface() }
func (d *Data) NewArrayPtr() interface{} { return d.NewArrayPtrValue().Interface() }
func (d *Data) NewSlice() interface{}    { return d.NewSliceValue().Interface() }

func (d *Data) Bytes() int   { return d.len * int(d.typ.Elem().Size()) }
func (d *Data) Name() string { return d.name }

// TestData test data
var TestData []*Data

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

package testdata

import "reflect"

// TestData test data
var TestData []*Data

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

//Array get an array with the given endianess
func (d *Data) Array(bigEndian bool) interface{} {
	return d.v[bToI(bigEndian)]
}

//ArrayPtr get a pointer to an array with the given endianess
func (d *Data) ArrayPtr(bigEndian bool) interface{} {
	return d.vPtr[bToI(bigEndian)]
}

//Slice get a slice with the given endianess
func (d *Data) Slice(bigEndian bool) interface{} {
	return d.slice[bToI(bigEndian)]
}

//ValueArray get an array with the given endianness
func (d *Data) ValueArray(bigEndian bool) reflect.Value {
	return d.vValue[bToI(bigEndian)]
}

//ValueArrayPtr get a pointer to an array with the given endianness
func (d *Data) ValueArrayPtr(bigEndian bool) reflect.Value {
	return d.vPtrValue[bToI(bigEndian)]
}

//ValueArrayUnaddressable get an unaddressable array with the given endianness
func (d *Data) ValueArrayUnaddressable(bigEndian bool) reflect.Value {
	return reflect.ValueOf(d.v[bToI(bigEndian)])
}

//ValueSlice get a slice with the given endianess
func (d *Data) ValueSlice(bigEndian bool) reflect.Value {
	return d.sliceValue[bToI(bigEndian)]
}

//Name gets the name of this type
func (d *Data) Name() string { return d.name }

//Bytes gets the length of data in bytes
func (d *Data) Bytes() int {
	return d.len * int(d.typ.Elem().Size())
}

//NewArrayValue returns a new array with the same element type and size as the array returned by `Array()`
func (d *Data) NewArrayValue() reflect.Value {
	return reflect.New(d.typ).Elem()
}

//NewArrayPtrValue returns a new pointer to an array with the same element type and size as the array returned by `ArrayPtr()`
func (d *Data) NewArrayPtrValue() reflect.Value {
	return reflect.New(d.typ)
}

//NewSliceValue returns a new slice with the same element type and size as the slice returned by `Slice()`
func (d *Data) NewSliceValue() reflect.Value {
	return reflect.New(d.typ).Elem().Slice(0, d.len)
}

//NewArray returns a new array with the same element type and size as the array returned by `Array()`
func (d *Data) NewArray() interface{} {
	return d.NewArrayValue().Interface()
}

//NewArrayPtr returns a new pointer to an array with the same element type and size as the array returned by `ArrayPtr()`
func (d *Data) NewArrayPtr() interface{} {
	return d.NewArrayPtrValue().Interface()
}

//NewSlice returns a new slice with the same element type and size as the slice returned by `Slice()`
func (d *Data) NewSlice() interface{} {
	return d.NewSliceValue().Interface()
}

func bToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

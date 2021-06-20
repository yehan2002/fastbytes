package data

import (
	"encoding/binary"
	"math"
)

const size = 64

//Bytes uint8
var Bytes = [size]uint8{
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
	0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
}

//testdata
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

//AllBigEndian all bigEndian value
var AllBigEndian []interface{}

//AllLittleEndian all bigEndian value
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

	for i := 0; i < len(Bytes); i += 4 {
		j := i >> 2

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
		j := i >> 3

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

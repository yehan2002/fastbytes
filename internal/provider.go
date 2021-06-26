package internal

import "reflect"

//Provider a provider
type Provider interface {
	FromI8([]int8, []byte) int
	FromI16([]int16, []byte, bool) int
	FromI32([]int32, []byte, bool) int
	FromI64([]int64, []byte, bool) int
	FromU16([]uint16, []byte, bool) int
	FromU32([]uint32, []byte, bool) int
	FromU64([]uint64, []byte, bool) int
	FromF32([]float32, []byte, bool) int
	FromF64([]float64, []byte, bool) int
	ToI8([]uint8, []int8) int
	ToI16([]byte, []int16, bool) int
	ToI32([]byte, []int32, bool) int
	ToI64([]byte, []int64, bool) int
	ToU16([]byte, []uint16, bool) int
	ToU32([]byte, []uint32, bool) int
	ToU64([]byte, []uint64, bool) int
	ToF32([]byte, []float32, bool) int
	ToF64([]byte, []float64, bool) int
	FromSlice(interface{}, []byte, bool) (int, error)
	ToSlice([]byte, interface{}, bool) (int, error)
	FromValue(reflect.Value, []byte, bool) (int, error)
	ToValue([]byte, reflect.Value, bool) (int, error)
}

//go:build !amd64

package asm

// Copy64 copy and rotate uint64 from `src` to `dst`
func Copy64(src []uint64, dst []uint64) uint64 { panic("not available") }

// Copy32 copy and rotate uint32 from `src` to `dst`
func Copy32(src []uint32, dst []uint32) uint64 { panic("not available") }

// Copy32 copy and rotate uint32 from `src` to `dst`
func Copy16(src []uint16, dst []uint16) uint64 { panic("not available") }

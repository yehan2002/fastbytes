//go:build !purego && !no_unsafe
// +build !purego,!no_unsafe

package unsafe

func init() { canASMPtr = &canASM }

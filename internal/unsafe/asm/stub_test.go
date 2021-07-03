package asm

import "testing"

func TestSwapEndian(t *testing.T) {
	a := []uint16{0x1112, 0x2122, 0x3132, 0x4142, 0x5152, 0x6162, 0x7172, 0x8182}
	b := make([]uint16, 8)

	Copy16(a, b)
	t.Errorf("%#v %#v", a, b)
}

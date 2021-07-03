package main

import (
	"fmt"

	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

//go:generate go run . -out ./../copy_amd64.s -stubs ./../stub_amd64.go -pkg asm

func main() {
	copySlice(16, 0x0E0F_0C0D_0A0B_0809, 0x0607_0405_0203_0001)
	copySlice(32, 0x0C0D0E0F_08090A0B, 0x040506070_0010203)
	copySlice(64, 0x08090A0B0C0D0E0F, 0x0001020304050607)
	build.Generate()
}

//https://stackoverflow.com/questions/56407741/reverse-byte-order-in-xmm-or-ymm-register

func copySlice(size int, shuffle1, shuffle2 uint64) {
	build.TEXT(fmt.Sprintf("Copy%d", size), build.NOSPLIT, fmt.Sprintf("func(src []uint%d, dst []uint%d) uint64", size, size))
	build.Pragma("noescape")
	build.Doc(fmt.Sprintf("Copy%d copy and rotate uint64 from `src` to `dst`", size))
	src := build.Load(build.Param("src").Base(), build.GP64())
	dst := build.Load(build.Param("dst").Base(), build.GP64())
	n := build.Load(build.Param("src").Len(), build.GP64())

	mask := build.XMM()
	xmm := build.XMM()
	tmp := build.GP64()

	/* build.MOVQ(operand.U64(uint64(0x0001020304050607)), t)
	build.PINSRQ(operand.Imm(uint64(0)), t, mask)
	build.MOVQ(operand.U64(uint64(0x08090A0B0C0D0E0F)), t)
	build.PINSRQ(operand.Imm(uint64(1)), t, mask) */

	build.Comment("Setup byte mask for byte shuffling")
	build.MOVQ(operand.U64(shuffle1), tmp)
	build.MOVQ(tmp, xmm)
	build.MOVQ(operand.U64(shuffle2), tmp)
	build.MOVQ(tmp, mask)
	build.MOVLHPS(xmm, mask)

	build.Label("loop")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(uint64(128/size)-1))
	build.JLE(operand.LabelRef("done"))

	build.Comment("Read 16 bytes from src, rotate them and copy them to dst.")
	build.MOVOU(operand.Mem{Base: src}, xmm)
	build.VPSHUFB(mask, xmm, xmm)
	build.MOVOU(xmm, operand.Mem{Base: dst})

	build.Comment("Advance pointers, decrement byte count")
	build.ADDQ(operand.I32(-(128 / size)), n)
	build.ADDQ(operand.Imm(16), src)
	build.ADDQ(operand.Imm(16), dst)
	build.JMP(operand.LabelRef("loop"))

	build.Label("done")
	length := build.Load(build.Param("src").Len(), tmp)
	build.SUBQ(n, length)
	build.Store(length, build.ReturnIndex(0))
	build.RET()
}

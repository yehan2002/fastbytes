package main

import (
	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

//https://stackoverflow.com/questions/56407741/reverse-byte-order-in-xmm-or-ymm-register

func copy64() {
	build.TEXT("Copy64", build.NOSPLIT, "func(src []uint64, dst []uint64) uint64")
	build.Pragma("noescape")
	build.Doc("Copy64 copy and rotate uint64 from `src` to `dst`")
	src := build.Load(build.Param("src").Base(), build.GP64())
	dst := build.Load(build.Param("dst").Base(), build.GP64())
	n := build.Load(build.Param("src").Len(), build.GP64())

	mask := build.XMM()
	xmm := build.XMM()
	t := build.GP64()

	/* build.MOVQ(operand.U64(uint64(0x0001020304050607)), t)
	build.PINSRQ(operand.Imm(uint64(0)), t, mask)
	build.MOVQ(operand.U64(uint64(0x08090A0B0C0D0E0F)), t)
	build.PINSRQ(operand.Imm(uint64(1)), t, mask) */

	build.Comment("Setup byte mask for byte shuffling")
	build.MOVQ(operand.U64(uint64(0x08090A0B0C0D0E0F)), t)
	build.MOVQ(t, xmm)
	build.MOVQ(operand.U64(uint64(0x0001020304050607)), t)
	build.MOVQ(t, mask)
	build.MOVLHPS(xmm, mask)

	build.Label("loop")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(1))
	build.JLE(operand.LabelRef("done"))

	build.Comment("Read 16 bytes from src, rotate them and copy them to dst.")
	build.MOVOU(operand.Mem{Base: src}, xmm)
	build.VPSHUFB(mask, xmm, xmm)
	build.MOVOU(xmm, operand.Mem{Base: dst})

	build.Comment("Advance pointers, decrement byte count")
	build.ADDQ(operand.I32(-2), n)
	build.ADDQ(operand.Imm(16), src)
	build.ADDQ(operand.Imm(16), dst)
	build.JMP(operand.LabelRef("loop"))

	build.Label("done")
	length := build.Load(build.Param("src").Len(), t)
	build.SUBQ(n, length)
	build.Store(length, build.ReturnIndex(0))
	build.RET()
}

func copy32() {
	build.TEXT("Copy32", build.NOSPLIT, "func(src []uint32, dst []uint32) uint64")
	build.Pragma("noescape")
	build.Doc("Copy32 copy and rotate uint32 from `src` to `dst`")
	src := build.Load(build.Param("src").Base(), build.GP64())
	dst := build.Load(build.Param("dst").Base(), build.GP64())
	n := build.Load(build.Param("src").Len(), build.GP64())

	mask := build.XMM()
	xmm := build.XMM()
	t := build.GP64()

	/* build.MOVQ(operand.U64(uint64(0x0001020304050607)), t)
	build.PINSRQ(operand.Imm(uint64(0)), t, mask)
	build.MOVQ(operand.U64(uint64(0x08090A0B0C0D0E0F)), t)
	build.PINSRQ(operand.Imm(uint64(1)), t, mask) */

	build.Comment("Setup byte mask for byte shuffling")
	build.MOVQ(operand.U64(uint64(0x0C0D0E0F08090A0B)), t)
	build.MOVQ(t, xmm)
	build.MOVQ(operand.U64(uint64(0x0405060700010203)), t)
	build.MOVQ(t, mask)
	build.MOVLHPS(xmm, mask)

	build.Label("loop")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(3))
	build.JLE(operand.LabelRef("done"))

	build.Comment("Read 16 bytes from src, rotate them and copy them to dst.")
	build.MOVOU(operand.Mem{Base: src}, xmm)
	build.VPSHUFB(mask, xmm, xmm)
	build.MOVOU(xmm, operand.Mem{Base: dst})

	build.Comment("Advance pointers, decrement byte count")
	build.ADDQ(operand.I32(-4), n)
	build.ADDQ(operand.Imm(16), src)
	build.ADDQ(operand.Imm(16), dst)
	build.JMP(operand.LabelRef("loop"))

	build.Label("done")
	length := build.Load(build.Param("src").Len(), t)
	build.SUBQ(n, length)
	build.Store(length, build.ReturnIndex(0))
	build.RET()
}

func copy16() {
	build.TEXT("Copy16", build.NOSPLIT, "func(src []uint16, dst []uint16) uint64")
	build.Pragma("noescape")
	build.Doc("Copy32 copy and rotate uint32 from `src` to `dst`")
	src := build.Load(build.Param("src").Base(), build.GP64())
	dst := build.Load(build.Param("dst").Base(), build.GP64())
	n := build.Load(build.Param("src").Len(), build.GP64())

	mask := build.XMM()
	xmm := build.XMM()
	t := build.GP64()

	/* build.MOVQ(operand.U64(uint64(0x0001020304050607)), t)
	build.PINSRQ(operand.Imm(uint64(0)), t, mask)
	build.MOVQ(operand.U64(uint64(0x08090A0B0C0D0E0F)), t)
	build.PINSRQ(operand.Imm(uint64(1)), t, mask) */

	build.Comment("Setup byte mask for byte shuffling")
	build.MOVQ(operand.U64(uint64(0x0E0F0C0D0A0B0809)), t)
	build.MOVQ(t, xmm)
	build.MOVQ(operand.U64(uint64(0x0607040502030001)), t)
	build.MOVQ(t, mask)
	build.MOVLHPS(xmm, mask)

	build.Label("loop")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(7))
	build.JLE(operand.LabelRef("done"))

	build.Comment("Read 16 bytes from src, rotate them and copy them to dst.")
	build.MOVOU(operand.Mem{Base: src}, xmm)
	build.VPSHUFB(mask, xmm, xmm)
	build.MOVOU(xmm, operand.Mem{Base: dst})

	build.Comment("Advance pointers, decrement byte count")
	build.ADDQ(operand.I32(-8), n)
	build.ADDQ(operand.Imm(16), src)
	build.ADDQ(operand.Imm(16), dst)
	build.JMP(operand.LabelRef("loop"))

	build.Label("done")
	length := build.Load(build.Param("src").Len(), t)
	build.SUBQ(n, length)
	build.Store(length, build.ReturnIndex(0))
	build.RET()
}

package main

import (
	"fmt"

	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
	"github.com/mmcloughlin/avo/reg"
)

//go:generate go run . -out ./../copy_amd64.s -stubs ./../stub_amd64.go -pkg asm

func main() {
	copySlice(16, 0x0E0F_0C0D_0A0B_0809, 0x0607_0405_0203_0001)
	copySlice(32, 0x0C0D0E0F_08090A0B, 0x040506070_0010203)
	copySlice(64, 0x08090A0B0C0D0E0F, 0x0001020304050607)
	build.ConstraintExpr("amd64,!purego")
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

	build.Comment("Setup byte mask for byte shuffling")
	build.MOVQ(operand.U64(shuffle1), tmp)
	build.MOVQ(tmp, xmm)
	build.MOVQ(operand.U64(shuffle2), tmp)
	build.MOVQ(tmp, mask)
	build.MOVLHPS(xmm, mask)

	build.Label("loop1024")
	build.CMPQ(n, operand.Imm(uint64(128/size)*8-1))
	build.JLE(operand.LabelRef("loop128"))
	var registers = []reg.Virtual{xmm}
	for i := 0; i < 7; i++ {
		registers = append(registers, build.XMM())
	}

	for i, reg := range registers {
		build.MOVOU(operand.Mem{Base: src}.Offset(i*16), reg)
	}

	for _, reg := range registers {
		build.VPSHUFB(mask, reg, reg)
	}

	for i, reg := range registers {
		build.MOVOU(reg, operand.Mem{Base: dst}.Offset(i*16))
	}

	build.ADDQ(operand.I32(-(128/size)*8), n)
	build.ADDQ(operand.Imm(128), src)
	build.ADDQ(operand.Imm(128), dst)
	build.JMP(operand.LabelRef("loop1024"))

	build.Label("loop128")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(uint64(128/size)-1))
	build.JLE(operand.LabelRef("tail"))

	build.Comment("Read 16 bytes from src, rotate them and copy them to dst.")
	build.MOVOU(operand.Mem{Base: src}, xmm)
	build.VPSHUFB(mask, xmm, xmm)
	build.MOVOU(xmm, operand.Mem{Base: dst})

	build.Comment("Advance pointers, decrement byte count")
	build.ADDQ(operand.I32(-(128 / size)), n)
	build.ADDQ(operand.Imm(16), src)
	build.ADDQ(operand.Imm(16), dst)
	build.JMP(operand.LabelRef("loop128"))

	build.Label("tail")
	build.CMPQ(n, operand.Imm(0))
	build.JE(operand.LabelRef("done"))
	var bytes uint64
	switch size {
	case 64:
		build.MOVQ(operand.Mem{Base: src}, tmp)
		build.BSWAPQ(tmp)
		build.MOVQ(tmp, operand.Mem{Base: dst})
		bytes = 8
	case 32:
		tmp := build.GP32()
		build.MOVL(operand.Mem{Base: src}, tmp)
		build.BSWAPL(tmp)
		build.MOVL(tmp, operand.Mem{Base: dst})
		bytes = 4
	case 16:
		tmp := build.GP16()
		build.MOVW(operand.Mem{Base: src}, tmp)
		build.ROLW(operand.U8(8), tmp)
		build.MOVW(tmp, operand.Mem{Base: dst})
		bytes = 2
	}
	build.ADDQ(operand.Imm(bytes), src)
	build.ADDQ(operand.Imm(bytes), dst)
	build.DECQ(n)

	build.JMP(operand.LabelRef("tail"))

	build.Label("done")
	length := build.Load(build.Param("src").Len(), tmp)
	build.SUBQ(n, length)

	build.Store(length, build.ReturnIndex(0))
	build.RET()
}

package main

import "github.com/mmcloughlin/avo/build"

//go:generate go run . -out ./../copy_amd64.s -stubs ./../stub_amd64.go -pkg asm

func main() {
	copy64()
	copy32()
	copy16()
	build.Generate()
}

/* func main2() {
	build.TEXT("Rot64", build.NOSPLIT, "func(src []uint64, dst []uint64)")
	build.Doc("Rot64 rotate the bytes in the given slice")
	src := build.Load(build.Param("src").Base(), build.GP64())
	dst := build.Load(build.Param("dst").Base(), build.GP64())
	n := build.Load(build.Param("src").Len(), build.GP64())

	build.Label("loop")
	build.Comment("Loop until zero bytes remain.")
	build.CMPQ(n, operand.Imm(0))
	build.JE(operand.LabelRef("done"))

	build.Comment("Load, rotate and store the uint64")
	tmp := build.GP64()
	build.MOVQ(operand.Mem{Base: src}, tmp)
	build.BSWAPQ(tmp)
	build.MOVQ(tmp, operand.Mem{Base: dst})

	build.Comment("Advance pointer, decrement byte count.")
	build.ADDQ(operand.Imm(8), src)
	build.ADDQ(operand.Imm(8), dst)
	build.DECQ(n)
	build.JMP(operand.LabelRef("loop"))

	build.Label("done")
	build.RET()
	build.Generate()
} */

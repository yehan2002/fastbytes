// +build amd64,!purego

package asm

import "golang.org/x/sys/cpu"

//CanASM this reports if assembly functions provided by this package can be used
var CanASM = cpu.X86.HasSSE2 && cpu.X86.HasAVX

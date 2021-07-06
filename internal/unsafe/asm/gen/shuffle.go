package main

import "encoding/binary"

// shuffleMask the mask use for byte shuffling (the `VPSHUFB` instruction).
// Each byte in the shuffle mask coresponds to a byte in a uint128.
// The following pseudo-code shows the operation of the VPSHUFB instruction.
// `v` and `ret` are 128bit registers.
//
//  func VPSHUFB(mask shuffleMask, v [16]byte) (ret [16]byte) {
//  	for i := range mask {
//	    	if m := mask[i]; m > 0x0F {
//		    	ret[i] = 0
//		    } else {
//			    ret[i] = v[m]
//		    }
//	     }
//   }
//
type shuffleMask [16]byte

func (s shuffleMask) LSB() (lsb uint64) { return binary.LittleEndian.Uint64(s[:8]) }
func (s shuffleMask) MSB() (msb uint64) { return binary.LittleEndian.Uint64(s[8:]) }

var shuffleUint16 = shuffleMask{1, 0, 3, 2, 5, 4, 7, 6, 9, 8, 11, 10, 13, 12, 15, 14}
var shuffleUint32 = shuffleMask{3, 2, 1, 0, 7, 6, 5, 4, 11, 10, 9, 8, 15, 14, 13, 12}
var shuffleUint64 = shuffleMask{7, 6, 5, 4, 3, 2, 1, 0, 15, 14, 13, 12, 11, 10, 9, 8}

//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"testing"

	"github.com/yehan2002/fastbytes/internal/unsafe/asm"
	"github.com/yehan2002/is"
)

func TestCopy(t *testing.T) {
	// This test must not be called parellely since this mutates `canASM`
	is.Suite(t, &copyTest{})
	if asm.CanASM {
		asm.CanASM = false
		is.Suite(t, &copyTest{})
		asm.CanASM = true
	}
}

type copyTest struct{}

// TestCopy16Large test copy16 when len(dst) < rotate16SmallSize
func (c *copyTest) TestCopy16Small(is is.IS) {
	var (
		test    = [...]uint16{0xDEAD, 0xBEEF, 0xAD0B, 0xFECA}
		testRot = [...]uint16{0xADDE, 0xEFBE, 0xBAD, 0xCAFE}
	)
	var dst [len(test)]uint16

	if len(test) >= rotate16SmallSize {
		is.Fail("copy16 test is incorrect. The value of rotate16SmallSize has changed")
	}

	// test copying without rotation
	n := copy16(test[:], dst[:], false)
	c.checkCopy(is, "copy16", dst, n, test, len(test)*2)

	n = copy16(test[:], dst[:], true)
	c.checkCopy(is, "copy16", dst, n, testRot, len(test)*2)
}

// TestCopy16Large test copy16 when len(dst)>= rotate16SmallSlice
func (c *copyTest) TestCopy16Large(is is.IS) {
	var (
		test    = [...]uint16{0x0123, 0x4567, 0x89AB, 0xCDEF, 0xFEDC, 0xBA98, 0x7654, 0x3210, 0x1234, 0x5678, 0x9ABC, 0xDEFE, 0xDCBA, 0x9876}
		testRot = [...]uint16{0x2301, 0x6745, 0xAB89, 0xEFCD, 0xDCFE, 0x98BA, 0x5476, 0x1032, 0x3412, 0x7856, 0xBC9A, 0xFEDE, 0xBADC, 0x7698}
	)
	var dst [len(test)]uint16

	if len(test) < rotate16SmallSize {
		is.Fail("copy16 test is incorrect. The value of rotate16SmallSize has changed")
	}

	// test copying without rotation
	n := copy16(test[:], dst[:], false)
	c.checkCopy(is, "copy16", dst, n, test, len(test)*2)

	n = copy16(test[:], dst[:], true)
	c.checkCopy(is, "copy16", dst, n, testRot, len(test)*2)
}

func (c *copyTest) TestCopy32(is is.IS) {
	var (
		test    = [...]uint32{0x01234567, 0x89ABCDEF, 0xEDCBA987, 0x65432101, 0x23456789, 0xABCDEFED, 0xCBA98765}
		testRot = [...]uint32{0x67452301, 0xEFCDAB89, 0x87A9CBED, 0x01214365, 0x89674523, 0xEDEFCDAB, 0x6587A9CB}
	)
	var dst [len(test)]uint32

	// test copying without rotation
	n := copy32(test[:], dst[:], false)
	c.checkCopy(is, "copy32", dst, n, test, len(test)*4)

	n = copy32(test[:], dst[:], true)
	c.checkCopy(is, "copy32", dst, n, testRot, len(test)*4)
}

func (c *copyTest) TestCopy64(is is.IS) {
	var (
		test    = [...]uint64{0x0123456789ABCDEF, 0xEDCBA98765432101, 0x23456789ABCDEFED, 0xCBA9876543210123}
		testRot = [...]uint64{0xEFCDAB8967452301, 0x0121436587A9CBED, 0xEDEFCDAB89674523, 0x230121436587A9CB}
	)
	var dst [len(test)]uint64

	// test copying without rotation
	n := copy64(test[:], dst[:], false)
	c.checkCopy(is, "copy64", dst, n, test, len(test)*8)

	n = copy64(test[:], dst[:], true)
	c.checkCopy(is, "copy64", dst, n, testRot, len(test)*8)
}

func (c *copyTest) TestCopySlice(is is.IS) {
	var (
		test      = [...]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
		testRot16 = [...]byte{0x23, 0x01, 0x67, 0x45, 0xAB, 0x89, 0xEF, 0xCD, 0xDC, 0xFE, 0x98, 0xBA, 0x54, 0x76, 0x10, 0x32}
		testRot32 = [...]byte{0x67, 0x45, 0x23, 0x01, 0xEF, 0xCD, 0xAB, 0x89, 0x98, 0xBA, 0xDC, 0xFE, 0x10, 0x32, 0x54, 0x76}
		testRot64 = [...]byte{0xEF, 0xCD, 0xAB, 0x89, 0x67, 0x45, 0x23, 0x01, 0x10, 0x32, 0x54, 0x76, 0x98, 0xBA, 0xDC, 0xFE}
	)
	var dst [len(test)]byte

	n := copySlice(test[:], dst[:], 1, true)
	c.checkCopy(is, "copySlice with size=1", dst, n, test, len(test))

	n = copySlice(test[:], dst[:], 2, true)
	c.checkCopy(is, "copySlice with size=2", dst, n, testRot16, len(test))

	n = copySlice(test[:], dst[:], 4, true)
	c.checkCopy(is, "copySlice with size=4", dst, n, testRot32, len(test))

	n = copySlice(test[:], dst[:], 8, true)
	c.checkCopy(is, "copySlice with size=8", dst, n, testRot64, len(test))

	is.MustPanicCall(func() {
		copySlice(nil, nil, 3, true) // This should never happen
	})
}

func (c *copyTest) checkCopy(is is.IS, name string, v interface{}, length int, expectedValue interface{}, expectedLength int) {
	is.True(length == expectedLength, name, " copied an incorrect number of bytes")
	is.True(v == expectedValue, name, " copied incorrectly", v, expectedValue)
}

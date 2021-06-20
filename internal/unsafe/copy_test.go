package unsafe

import (
	"testing"

	"github.com/yehan2002/is"
)

func TestCopy(t *testing.T) {
	is.Suite(t, &copyTest{})
}

type copyTest struct{}

//TestCopy16Large test copy16 when len(dst) < rotate16SmallSize
func (c *copyTest) TestCopy16Small(is is.IS) {
	var (
		test    = [...]uint16{0xDEAD, 0xBEEF, 0xFECA}
		testRot = [...]uint16{0xADDE, 0xEFBE, 0xCAFE}
	)
	var dst [len(test)]uint16

	if len(test) >= rotate16SmallSize {
		is.Fail("copy16 test is incorrect. The value of rotate16SmallSize has changed")
	}

	//test copying without rotation
	n := copy16(test[:], dst[:], false)
	is.True(n == len(test)*2, "copy16 copied an incorrect number of bytes")
	is.True(test == dst, "copy16 copied incorrectly")

	n = copy16(test[:], dst[:], true)
	is.True(n == len(test)*2, "copy16 copied an incorrect number of bytes")
	is.True(testRot == dst, "copy16 copied incorrectly")

}

//TestCopy16Large test copy16 when len(dst)>= rotate16SmallSlice
func (c *copyTest) TestCopy16Large(is is.IS) {
	var (
		test    = [...]uint16{0x0123, 0x4567, 0x89AB, 0xCDEF, 0xFEDC, 0xBA98, 0x7654, 0x3210, 0x1234, 0x5678, 0x9ABC, 0xDEFE, 0xDCBA, 0x9876}
		testRot = [...]uint16{0x2301, 0x6745, 0xAB89, 0xEFCD, 0xDCFE, 0x98BA, 0x5476, 0x1032, 0x3412, 0x7856, 0xBC9A, 0xFEDE, 0xBADC, 0x7698}
	)
	var dst [len(test)]uint16

	if len(test) < rotate16SmallSize {
		is.Fail("copy16 test is incorrect. The value of rotate16SmallSize has changed")
	}

	//test copying without rotation
	n := copy16(test[:], dst[:], false)
	is.True(n == len(test)*2, "copy16 copied an incorrect number of bytes")
	is.True(test == dst, "copy16 copied incorrectly")

	n = copy16(test[:], dst[:], true)
	is.True(n == len(test)*2, "copy16 copied an incorrect number of bytes")
	is.True(dst == testRot, "copy16 copied incorrectly")
}

func (c *copyTest) TestCopy32(is is.IS) {
	var (
		test    = [...]uint32{0x01234567, 0x89ABCDEF, 0xEDCBA987, 0x65432101, 0x23456789, 0xABCDEFED, 0xCBA98765}
		testRot = [...]uint32{0x67452301, 0xEFCDAB89, 0x87A9CBED, 0x01214365, 0x89674523, 0xEDEFCDAB, 0x6587A9CB}
	)
	var dst [len(test)]uint32

	//test copying without rotation
	n := copy32(test[:], dst[:], false)
	is.True(n == len(test)*4, "copy32 copied an incorrect number of bytes")
	is.True(test == dst, "copy32 copied incorrectly")

	n = copy32(test[:], dst[:], true)
	is.True(n == len(test)*4, "copy32 copied an incorrect number of bytes")
	is.True(dst == testRot, "copy32 copied incorrectly")
}

func (c *copyTest) TestCopy64(is is.IS) {
	var (
		test    = [...]uint64{0x0123456789ABCDEF, 0xEDCBA98765432101, 0x23456789ABCDEFED, 0xCBA9876543210123}
		testRot = [...]uint64{0xEFCDAB8967452301, 0x0121436587A9CBED, 0xEDEFCDAB89674523, 0x230121436587A9CB}
	)
	var dst [len(test)]uint64

	//test copying without rotation
	n := copy64(test[:], dst[:], false)
	is.True(n == len(test)*8, "copy64 copied an incorrect number of bytes")
	is.True(test == dst, "copy64 copied incorrectly")

	n = copy64(test[:], dst[:], true)
	is.True(n == len(test)*8, "copy64 copied an incorrect number of bytes")
	is.True(dst == testRot, "copy64 copied incorrectly")
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
	is.True(n == len(test), "copySlice copied an incorrect number of bytes")
	is.True(test == dst, "copySlice with size 1 copied incorrectly")

	n = copySlice(test[:], dst[:], 2, true)
	is.True(n == len(test), "copySlice copied an incorrect number of bytes")
	is.True(testRot16 == dst, "copySlice with size 2 copied incorrectly")

	n = copySlice(test[:], dst[:], 4, true)
	is.True(n == len(test), "copySlice copied an incorrect number of bytes")
	is.True(testRot32 == dst, "copySlice with size 4 copied incorrectly")

	n = copySlice(test[:], dst[:], 8, true)
	is.True(n == len(test), "copySlice copied an incorrect number of bytes")
	is.True(testRot64 == dst, "copySlice with size 8 copied incorrectly")

	is.MustPanicCall(func() {
		copySlice(nil, nil, 3, true) // This should never happen
	})
}

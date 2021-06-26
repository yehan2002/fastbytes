//go:build !no_unsafe
// +build !no_unsafe

package bytes

import (
	"github.com/yehan2002/bytes/internal/unsafe"
)

var rotateBigEndian = unsafe.IsLittleEndian

type provider = unsafe.Bytes

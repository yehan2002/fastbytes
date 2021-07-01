//go:build !no_unsafe
// +build !no_unsafe

package fastbytes

import (
	"github.com/yehan2002/fastbytes/internal/unsafe"
)

var rotateBigEndian = unsafe.IsLittleEndian

type provider = unsafe.Bytes

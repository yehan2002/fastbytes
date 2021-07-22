//go:build no_unsafe
// +build no_unsafe

package fastbytes

import (
	"github.com/yehan2002/fastbytes/v2/internal/safe"
)

type provider = safe.Bytes

var rotateBigEndian = true

//go:build no_unsafe
// +build no_unsafe

package fastbytes

import (
	"reflect"

	"github.com/yehan2002/fastbytes/internal/safe"
)

var rotateBigEndian = true
var provider = safe.Bytes

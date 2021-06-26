//go:build no_unsafe
// +build no_unsafe

package bytes

import (
	"reflect"

	"github.com/yehan2002/bytes/internal/safe"
)

var rotateBigEndian = true
var provider = safe.BytesSafe

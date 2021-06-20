//go:build no_unsafe
// +build no_unsafe

package unsafe

func init() { panic("bytes/internal/unsafe: Imported when the tag `no_unsafe` was set") }

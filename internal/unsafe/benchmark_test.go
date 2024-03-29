//go:build !no_unsafe
// +build !no_unsafe

package unsafe

import (
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal/testutil"
)

var benchmark = testutil.Benchmark(bytes, IsLittleEndian)

func BenchmarkFrom8(b *testing.B) { benchmark.BenchmarkFrom8(b) }

func BenchmarkFrom16(b *testing.B) { benchmark.BenchmarkFrom16(b) }

func BenchmarkFrom32(b *testing.B) { benchmark.BenchmarkFrom32(b) }

func BenchmarkFrom64(b *testing.B) { benchmark.BenchmarkFrom64(b) }

func BenchmarkTo8(b *testing.B) { benchmark.BenchmarkTo8(b) }

func BenchmarkTo16(b *testing.B) { benchmark.BenchmarkTo16(b) }

func BenchmarkTo32(b *testing.B) { benchmark.BenchmarkTo32(b) }

func BenchmarkTo64(b *testing.B) { benchmark.BenchmarkTo64(b) }

package testdata

import (
	"testing"

	"github.com/yehan2002/fastbytes/v2/internal"
)

//Bench benchmarks the given provider.
type Bench struct {
	provider        internal.Provider
	rotateBigEndian bool
}

//Benchmark benchmarks the given provider
func Benchmark(pr internal.Provider, rot bool) *Bench {
	return &Bench{provider: pr, rotateBigEndian: rot}
}

//BenchmarkFrom8 benchmarks `From8`
func (r *Bench) BenchmarkFrom8(b *testing.B) {
	b.Run("FromI8", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI8(int8Slice[:], dst[:])
		}
	})
}

//BenchmarkFrom16 benchmarks `From16`
func (r *Bench) BenchmarkFrom16(b *testing.B) {
	b.Run("FromI16BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI16(Int16BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromI16LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI16(int16LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("FromU16BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU16(uint16BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromU16LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU16(uint16LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
}

//BenchmarkFrom32 benchmarks `From32`
func (r *Bench) BenchmarkFrom32(b *testing.B) {
	b.Run("FromI32BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI32(int32BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromI32LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI32(int32LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("FromU32BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU32(uint32BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromU32LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU32(uint32LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("FromF32BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromF32(float32BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromF32LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromF32(float32LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
}

//BenchmarkFrom64 benchmarks `From64`
func (r *Bench) BenchmarkFrom64(b *testing.B) {
	b.Run("FromI64BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI64(int64BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromI64LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromI64(int64LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("FromU64BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU64(uint64BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromU64LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromU64(uint64LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("FromF64BigEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromF64(float64BigEndian[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("FromF64LittleEndian", func(b *testing.B) {
		var dst [len(bytes)]byte
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.FromF64(float64LittleEndian[:], dst[:], !r.rotateBigEndian)
		}
	})
}

//BenchmarkTo8 benchmarks `From8`
func (r *Bench) BenchmarkTo8(b *testing.B) {
	b.Run("TestTo8", func(b *testing.B) {
		var dst [len(bytes)]int8
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI8(bytes[:], dst[:])
		}
	})
}

//BenchmarkTo16 benchmarks `From16`
func (r *Bench) BenchmarkTo16(b *testing.B) {
	b.Run("ToI16BigEndian", func(b *testing.B) {
		var dst [len(Int16BigEndian)]int16
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI16(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToI16LittleEndian", func(b *testing.B) {
		var dst [len(Int16BigEndian)]int16
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI16(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("ToU16BigEndian", func(b *testing.B) {
		var dst [len(Int16BigEndian)]uint16
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU16(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToU16LittleEndian", func(b *testing.B) {
		var dst [len(Int16BigEndian)]uint16
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU16(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
}

//BenchmarkTo32 benchmarks `From32`
func (r *Bench) BenchmarkTo32(b *testing.B) {
	b.Run("ToI32BigEndian", func(b *testing.B) {
		var dst [len(int32BigEndian)]int32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI32(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToI32LittleEndian", func(b *testing.B) {
		var dst [len(int32BigEndian)]int32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI32(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("ToU32BigEndian", func(b *testing.B) {
		var dst [len(uint32BigEndian)]uint32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU32(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToU32LittleEndian", func(b *testing.B) {
		var dst [len(uint32BigEndian)]uint32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU32(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("ToF32BigEndian", func(b *testing.B) {
		var dst [len(float32BigEndian)]float32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToF32(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToF32LittleEndian", func(b *testing.B) {
		var dst [len(float32BigEndian)]float32
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToF32(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
}

//BenchmarkTo64 benchmarks `From64`
func (r *Bench) BenchmarkTo64(b *testing.B) {
	b.Run("ToI64BigEndian", func(b *testing.B) {
		var dst [len(int64BigEndian)]int64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI64(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToI64LittleEndian", func(b *testing.B) {
		var dst [len(int64BigEndian)]int64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToI64(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("ToU64BigEndian", func(b *testing.B) {
		var dst [len(uint64BigEndian)]uint64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU64(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToU64LittleEndian", func(b *testing.B) {
		var dst [len(uint64BigEndian)]uint64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToU64(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
	b.Run("ToF64BigEndian", func(b *testing.B) {
		var dst [len(float64BigEndian)]float64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToF64(bytes[:], dst[:], r.rotateBigEndian)
		}
	})
	b.Run("ToF64LittleEndian", func(b *testing.B) {
		var dst [len(float64BigEndian)]float64
		b.ReportAllocs()
		b.SetBytes(int64(len(bytes)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r.provider.ToF64(bytes[:], dst[:], !r.rotateBigEndian)
		}
	})
}

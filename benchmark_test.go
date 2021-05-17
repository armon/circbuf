package circbuf

import (
	"testing"
)

// k is the number of repetitions per operation.
const k = 1000

// benchmarkWrite benchmarks Write, by writing writeSize to a Buffer that is
// bufSize bytes long.
func benchmarkWrite(b *testing.B, bufSize, writeSize int64) {
	data := make([]byte, writeSize)
	buf, err := NewBuffer(bufSize)
	if err != nil {
		b.Fatalf("creating buffer of size %v: %v", bufSize, err)
	}
	b.SetBytes(k * writeSize)

	for i := 0; i < b.N; i++ {
		for j := 0; j < k; j++ {
			_, _ = buf.Write(data)
		}
	}
}

func Benchmark_Write_1024_500(b *testing.B) {
	benchmarkWrite(b, 1024, 500)
}

func Benchmark_Write_1025_500(b *testing.B) {
	benchmarkWrite(b, 1025, 500)
}

func Benchmark_Write_1024_5000(b *testing.B) {
	benchmarkWrite(b, 1024, 5000)
}

func Benchmark_Write_1025_5000(b *testing.B) {
	benchmarkWrite(b, 1025, 5000)
}

func Benchmark_Write_65536_5000(b *testing.B) {
	benchmarkWrite(b, 65536, 5000)
}

func Benchmark_Write_65537_5000(b *testing.B) {
	benchmarkWrite(b, 65537, 5000)
}

func Benchmark_Write_1024_5(b *testing.B) {
	benchmarkWrite(b, 1024, 5)
}

func Benchmark_Write_1025_5(b *testing.B) {
	benchmarkWrite(b, 1025, 5)
}

// benchmarkWriteByte benchmarks WriteByte, on a Buffer that is bufSize bytes
// long.
func benchmarkWriteByte(b *testing.B, bufSize int64) {
	buf, err := NewBuffer(bufSize)
	if err != nil {
		b.Fatalf("creating buffer of size %v: %v", bufSize, err)
	}
	b.SetBytes(k)

	for i := 0; i < b.N; i++ {
		for j := 0; j < k; j++ {
			_ = buf.WriteByte(0xba)
		}
	}
}

func Benchmark_WriteByte_1024(b *testing.B) {
	benchmarkWriteByte(b, 1024)
}

func Benchmark_WriteByte_1025(b *testing.B) {
	benchmarkWriteByte(b, 1025)
}

func Benchmark_WriteByte_65536(b *testing.B) {
	benchmarkWriteByte(b, 65536)
}

func Benchmark_WriteByte_65537(b *testing.B) {
	benchmarkWriteByte(b, 65537)
}

// benchmarkGet benchmarks Get with a buffer that is bufSize bytes long and was
// written to proportionally to fillRatio (so, for example, fillRatio == 0.5
// means write to half of the buffer, by fillRatio == 1 means write to every
// byte of the buffer, and fillRatio == 2 means write twice as much data as it
// fits in the buffer).
func benchmarkGet(b *testing.B, bufSize int64, fillRatio float64) {
	buf, err := NewBuffer(bufSize)
	if err != nil {
		b.Fatalf("creating buffer of size %v: %v", bufSize, err)
	}

	writeSize := int64(float64(bufSize) * fillRatio)
	data := make([]byte, writeSize)
	_, err = buf.Write(data)
	if err != nil {
		b.Fatalf("writing data to buffer: %v", err)
	}

	readLimit := bufSize
	if bufSize > writeSize {
		readLimit = writeSize
	}

	b.SetBytes(k)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < k; j++ {
			for h := int64(0); h < readLimit; h++ {
				_, _ = buf.Get(h)
			}
		}
	}
}

func Benchmark_Get_HalfFull_1024(b *testing.B) {
	benchmarkGet(b, 1024, 0.5)
}

func Benchmark_Get_HalfFull_1025(b *testing.B) {
	benchmarkGet(b, 1025, 0.5)
}

func Benchmark_Get_Full_1024(b *testing.B) {
	benchmarkGet(b, 1024, 1.0)
}

func Benchmark_Get_Full_1025(b *testing.B) {
	benchmarkGet(b, 1025, 1.0)
}

func Benchmark_Get_TwiceFull_1024(b *testing.B) {
	benchmarkGet(b, 1024, 2.0)
}

func Benchmark_Get_TwiceFull_1025(b *testing.B) {
	benchmarkGet(b, 1025, 2.0)
}

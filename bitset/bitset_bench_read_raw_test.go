package bitset

import (
	"testing"
	"time"

	wb "github.com/willf/bitset"
)

func BenchmarkBigIntRawRead(b *testing.B) {
	bs := NewBigInt()
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs.Int.Bit(i) == 1
		}
	}
}

func BenchmarkBoolRawRead(b *testing.B) {
	bs := make(BitsetBool, M)
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs[j]
		}
	}
}

func BenchmarkBoolDynRawRead(b *testing.B) {
	bs := make(BitsetBoolDyn, M)
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs[j]
		}
	}
}

func BenchmarkBoolDyn0RawRead(b *testing.B) {
	var bs BitsetBoolDyn
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs.GetBit(j)
		}
	}
}

func BenchmarkMapRawRead(b *testing.B) {
	bs := make(BitsetMap, M)
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs[j]
		}
	}
}

func BenchmarkUint8RawRead(b *testing.B) {
	bs := NewUint8(M)
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs.GetBit(j)
		}
	}
}

func BenchmarkUint64RawRead(b *testing.B) {
	bs := NewUint64(M)
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs.GetBit(j)
		}
	}
}

func BenchmarkWillfRawRead(b *testing.B) {
	bs := BitsetWillf{BitSet: *wb.New(uint(M))}
	for i := 0; i < b.N; i++ {
		for j := 3; j < M; j += 6 {
			Sinkb = bs.BitSet.Test(uint(i))
		}
	}
}

// touch prevents the compiler from treating the result as constant
func touch(x int) int {
	if time.Now().Year() > 3000 {
		return 0
	}
	return x
}

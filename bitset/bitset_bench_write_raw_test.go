package bitset

import (
	"testing"

	wb "github.com/willf/bitset"
)

func BenchmarkBigIntRawWrite(b *testing.B) {
	bs := NewBigInt()
	M = touch(M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs.Int.SetBit(bs.Int, j, 1)
		}
		for j := 1; j < M; j += 7 {
			bs.Int.SetBit(bs.Int, j, 0)
		}
	}
}

func BenchmarkBoolRawWrite(b *testing.B) {
	bs := make(BitsetBool, M)
	M = touch(M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs[j] = true
		}
		for j := 1; j < M; j += 7 {
			bs[j] = false
		}
	}
}

func BenchmarkBoolDynRawWrite(b *testing.B) {
	bs := make(BitsetBoolDyn, M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs[j] = true
		}
		for j := 1; j < M; j += 7 {
			bs[j] = false
		}
	}
}

func BenchmarkBoolDyn0RawWrite(b *testing.B) {
	var bs BitsetBoolDyn
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, true)
		}
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, false)
		}
	}
}

func BenchmarkMapRawWrite(b *testing.B) {
	bs := make(BitsetMap, M)
	M = touch(M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs[j] = true
		}
		for j := 1; j < M; j += 7 {
			bs[j] = false
		}
	}
}

func BenchmarkUint8RawWrite(b *testing.B) {
	bs := NewUint8(M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, true)
		}
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, false)
		}
	}
}

func BenchmarkUint64RawWrite(b *testing.B) {
	bs := NewUint64(M)
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, true)
		}
		for j := 1; j < M; j += 7 {
			bs.SetBit(j, false)
		}
	}
}
func BenchmarkWillfRawWrite(b *testing.B) {
	bs := BitsetWillf{BitSet: *wb.New(uint(M))}
	for i := 0; i < b.N; i++ {
		for j := 1; j < M; j += 7 {
			bs.BitSet.SetTo(uint(j), true)
		}
		for j := 1; j < M; j += 7 {
			bs.BitSet.SetTo(uint(j), false)
		}
	}
}

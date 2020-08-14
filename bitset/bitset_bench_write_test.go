package bitset

import (
	"testing"

	wb "github.com/willf/bitset"
)

func BenchmarkBigIntIfcWrite(b *testing.B) {
	bs := NewBigInt()
	benchmarkWrite(bs, b, M)
}

func BenchmarkBoolIfcWrite(b *testing.B) {
	bs := make(BitsetBool, M)
	benchmarkWrite(bs, b, M)
}

func BenchmarkBoolDynIfcWrite(b *testing.B) {
	bs := make(BitsetBoolDyn, M)
	benchmarkWrite(&bs, b, M)
}

func BenchmarkBoolDyn0IfcWrite(b *testing.B) {
	var bs BitsetBoolDyn
	benchmarkWrite(&bs, b, M)
}

func BenchmarkMapIfcWrite(b *testing.B) {
	bs := make(BitsetMap, M)
	benchmarkWrite(bs, b, M)
}

func BenchmarkUint8IfcWrite(b *testing.B) {
	bs := NewUint8(M)
	benchmarkWrite(bs, b, M)
}

func BenchmarkUint64IfcWrite(b *testing.B) {
	bs := NewUint64(M)
	benchmarkWrite(bs, b, M)
}
func BenchmarkWillfWrite(b *testing.B) {
	bs := BitsetWillf{BitSet: *wb.New(uint(M))}
	benchmarkWrite(bs, b, M)
}

func benchmarkWrite(bs Bitset, b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 1; j < n; j += 7 {
			bs.SetBit(j, true)
		}
		for j := 1; j < n; j += 7 {
			bs.SetBit(j, false)
		}
	}
}

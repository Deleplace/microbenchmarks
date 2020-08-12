package bitset

import (
	"testing"

	wb "github.com/willf/bitset"
)

var M = 100000

func BenchmarkBoolWrite(b *testing.B) {
	bs := make(BitsetBool, M)
	benchmarkWrite(bs, b, M)
}

func BenchmarkBoolDynWrite(b *testing.B) {
	bs := make(BitsetBoolDyn, M)
	benchmarkWrite(&bs, b, M)
}

func BenchmarkBoolDyn0Write(b *testing.B) {
	var bs BitsetBoolDyn
	benchmarkWrite(&bs, b, M)
}

func BenchmarkUint8Write(b *testing.B) {
	bs := NewUint8(M)
	benchmarkWrite(bs, b, M)
}

func BenchmarkUint64Write(b *testing.B) {
	bs := NewUint64(M)
	benchmarkWrite(bs, b, M)
}
func BenchmarkWillfWrite(b *testing.B) {
	bs := BitsetWillf{BitSet: *wb.New(uint(M))}
	benchmarkWrite(bs, b, M)
}

func BenchmarkBoolRead(b *testing.B) {
	bs := make(BitsetBool, M)
	benchmarkRead(bs, b, M)
}

func BenchmarkBoolDynRead(b *testing.B) {
	bs := make(BitsetBoolDyn, M)
	benchmarkRead(&bs, b, M)
}

func BenchmarkBoolDyn0Read(b *testing.B) {
	var bs BitsetBoolDyn
	benchmarkRead(&bs, b, M)
}

func BenchmarkUint8Read(b *testing.B) {
	bs := NewUint8(M)
	benchmarkRead(bs, b, M)
}

func BenchmarkUint64Read(b *testing.B) {
	bs := NewUint64(M)
	benchmarkRead(bs, b, M)
}

func BenchmarkWillfRead(b *testing.B) {
	bs := BitsetWillf{BitSet: *wb.New(uint(M))}
	benchmarkRead(bs, b, M)
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

var Sinkb = false

func benchmarkRead(bs Bitset, b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := 3; j < n; j += 6 {
			Sinkb = bs.GetBit(j)
		}
	}
}

func TestEquivalence(t *testing.T) {
	sets := []Bitset{
		make(BitsetBool, M),
		NewUint8(M),
		NewUint64(M),
		&BitsetBoolDyn{},
		BitsetWillf{},
	}
	for _, bs := range sets {
		for j := 2; j < M; j += 13 {
			bs.SetBit(j, true)
		}
	}
	testEqual(t, sets...)
}

func testEqual(t testing.TB, sets ...Bitset) {
	if len(sets) <= 1 {
		return
	}
	ref := sets[0]
	for i, bs := range sets[1:] {
		if bs.Len() != ref.Len() {
			return
		}
		for j := 0; j < bs.Len(); j++ {
			if bs.GetBit(j) != ref.GetBit(j) {
				t.Errorf("bitset %d differs at index %d", 1+i, j)
				return
			}
		}
	}
}

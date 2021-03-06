package bitset

import (
	"testing"
)

// M is the size of the bitsets created for tests and benchmarks.
var M = 100000

func TestEquivalence(t *testing.T) {
	sets := []Bitset{
		NewBigInt(),
		make(BitsetBool, M),
		make(BitsetMap, M),
		NewUint8(M),
		NewUint64(M),
		&BitsetBoolDyn{},
		&BitsetWillf{},
	}
	for _, bs := range sets {
		for j := 2; j < M; j += 13 {
			bs.SetBit(j, true)
		}
		for j := 1; j < M; j += 5 {
			bs.SetBit(j, false)
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
		for j, n := 0, max(bs.Len(), ref.Len()); j < n; j++ {
			if bs.GetBit(j) != ref.GetBit(j) {
				t.Errorf("bitset %d differs at index %d", 1+i, j)
				return
			}
		}
	}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

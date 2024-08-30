package sliceseq

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// This benchmark tests a potential performance penalty when using slice
// iterators, compared to a traditional range over the slice keys/values.
//
// Iterators were introduced in Go 1.23
//
// Results on a MacBook M1 show a 50% penalty when ranging over the values,
// and a 300% penalty when ranging over the indices and values.

// Size of list. Tune manually to produce interesting values.
const M = 1_000

// Exported global var, to make sure the compiler doesn't wipe away the computations
var Sink int

type V [64]byte

// BenchmarkIterateValues iterates over slice values, using the
// traditional range over slice style
func BenchmarkIterateValues(b *testing.B) {
	r.Seed([32]byte{42})
	s := randomizedSlice()
	b.ResetTimer()

	needle := V{1, 2, 3}
	for range b.N {
		for _, v := range s {
			if v == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateValuesSeq iterates over slice values, using the
// new range style over an iter.Seq iterator returned by
// slices.Values, introduced in Go 1.23
func BenchmarkIterateValuesSeq(b *testing.B) {
	r.Seed([32]byte{42})
	s := randomizedSlice()
	b.ResetTimer()

	needle := V{1, 2, 3}
	for range b.N {
		for v := range slices.Values(s) {
			if v == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateKeysValues iterates over slice indices and values, using the
// traditional range over slice style
func BenchmarkIterateKeysValues(b *testing.B) {
	r.Seed([32]byte{42})
	s := randomizedSlice()
	b.ResetTimer()

	for range b.N {
		for i, v := range s {
			if v[i%64] == 42 {
				Sink++
			}
		}
	}
}

// BenchmarkIterateKeysValuesSeq iterates over slice indices and values, using the
// new range style over an iter.Seq2 iterator returned by slices.All, introduced
// in Go 1.23
func BenchmarkIterateKeysValuesSeq(b *testing.B) {
	r.Seed([32]byte{42})
	s := randomizedSlice()
	b.ResetTimer()

	for range b.N {
		for i, v := range slices.All(s) {
			if v[i%64] == 42 {
				Sink++
			}
		}
	}
}

var r = rand.NewChaCha8([32]byte{42})

func randomValue() (v [64]byte) {
	r.Read(v[:])
	return v
}

func randomizedSlice() []V {
	s := make([]V, M)
	for i := range M {
		s[i] = randomValue()
	}
	return s
}

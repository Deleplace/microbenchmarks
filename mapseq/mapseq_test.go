package dedupe

import (
	"maps"
	"math/rand/v2"
	"testing"
)

// This benchmark tests a potential performance penalty when using map
// iterators, compared to a traditional range over the map keys/values.
//
// Iterators were introduced in Go 1.23
//
// Results on a MacBook M1 show a 18% penalty, which is an acceptable cost
// for the flexibility of using iterators.
// Ranging over the iterator is slightly slower, but same order of magnitude.

// Size of list. Tune manually to produce interesting values.
const M = 10_000

// Exported global var, to make sure the compiler doesn't wipe away the computations
var Sink int

type (
	K [64]byte
	V [64]byte
)

// BenchmarkIterateKeys iterates over map keys, using the
// traditional range over map style
func BenchmarkIterateKeys(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	needle := K{1, 2, 3}
	for range b.N {
		for k := range m {
			if k == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateKeysSeq iterates over map keys, using the
// new range style over an iter.Seq iterator returned by
// maps.Keys, introduced in Go 1.23
func BenchmarkIterateKeysSeq(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	needle := K{1, 2, 3}
	for range b.N {
		for k := range maps.Keys(m) {
			if k == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateValues iterates over map values, using the
// traditional range over map style
func BenchmarkIterateValues(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	needle := V{1, 2, 3}
	for range b.N {
		for _, v := range m {
			if v == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateValuesSeq iterates over map values, using the
// new range style over an iter.Seq iterator returned by
// maps.Values, introduced in Go 1.23
func BenchmarkIterateValuesSeq(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	needle := V{1, 2, 3}
	for range b.N {
		for v := range maps.Values(m) {
			if v == needle {
				Sink++
				break
			}
		}
	}
}

// BenchmarkIterateKeysValues iterates over map keys and values, using the
// traditional range over map style
func BenchmarkIterateKeysValues(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	for range b.N {
		for k, v := range m {
			if k == K(v) {
				Sink++
			}
		}
	}
}

// BenchmarkIterateKeysValuesSeq iterates over map keys and values, using the
// new range style over an iter.Seq iterator returned by maps.All, introduced
// in Go 1.23
func BenchmarkIterateKeysValuesSeq(b *testing.B) {
	r.Seed([32]byte{42})
	m := randomizedMap()
	b.ResetTimer()

	for range b.N {
		for k, v := range maps.All(m) {
			if k == K(v) {
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

func randomizedMap() map[K]V {
	m := make(map[K]V, M)
	for range M {
		k, v := randomValue(), randomValue()
		m[k] = v
	}
	return m
}

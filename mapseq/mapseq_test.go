package dedupe

import (
	"maps"
	"math/rand/v2"
	"testing"
)

// Size of list. Tune manually to produce interesting values.
const (
	M   = 100_000
	max = 1e4
)

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

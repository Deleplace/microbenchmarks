package shift

import (
	"math/rand"
	"testing"
)

func BenchmarkShift(b *testing.B) {
	M := 1000
	r := make([]uint, M)
	for i := range r {
		r[i] = uint(rand.Intn(12345678))
	}
	a := make([]uint, M)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(a, r)
		for j := range a {
			a[j] >>= 1
		}
	}
}

func BenchmarkDiv(b *testing.B) {
	M := 1000
	r := make([]uint, M)
	for i := range r {
		r[i] = uint(rand.Intn(12345678))
	}
	a := make([]uint, M)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(a, r)
		for j := range a {
			a[j] /= 2
		}
	}
}

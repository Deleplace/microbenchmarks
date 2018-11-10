package foreach

import (
	"math/rand"
	"testing"
)

func BenchmarkRange1(b *testing.B) {
	M := 1000
	a := make([]uint, M)
	for i := range a {
		a[i] = uint(rand.Intn(12345678))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum uint
		for j := range a {
			sum += a[j]
		}
	}
}

func BenchmarkRange2(b *testing.B) {
	M := 1000
	a := make([]uint, M)
	for i := range a {
		a[i] = uint(rand.Intn(12345678))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum uint
		for _, v := range a {
			sum += v
		}
	}
}

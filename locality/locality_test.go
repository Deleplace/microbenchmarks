package locality

import (
	"math/rand"
	"testing"
)

const M = 100000

// How fast is it to access a[0], a[1], ..., a[M-1] ?
// (thru an indirection array)
func BenchmarkAccessForward(b *testing.B) {
	rand.Seed(42)

	a := make([]uint, M)
	index := make([]int, M)
	for i := range a {
		a[i] = uint(i)
		index[i] = i
	}

	b.ResetTimer()
	var sum uint
	for i := 0; i < b.N; i++ {
		for j := range index {
			sum += a[index[j]]
		}
	}
	if expected := uint(b.N*M*(M-1)) / 2; sum != expected {
		b.Errorf("Computed %d, expected %d", sum, expected)
	}
}

// How fast is it to access a[M-1], a[M-2], ..., a[0] ?
// (thru an indirection array)
func BenchmarkAccessBackward(b *testing.B) {
	rand.Seed(42)

	a := make([]uint, M)
	index := make([]int, M)
	for i := range a {
		a[i] = uint(i)
		index[i] = M - i - 1
	}

	b.ResetTimer()
	var sum uint
	for i := 0; i < b.N; i++ {
		for j := range index {
			sum += a[index[j]]
		}
	}
	if expected := uint(b.N*M*(M-1)) / 2; sum != expected {
		b.Errorf("Computed %d, expected %d", sum, expected)
	}
}

// How fast is it to access all elements of a in random order ?
// (thru an indirection array)
func BenchmarkAccessRandom(b *testing.B) {
	rand.Seed(42)

	M := 100000
	a := make([]uint, M)
	index := rand.Perm(M)
	for i := range a {
		a[i] = uint(i)
	}

	b.ResetTimer()
	var sum uint
	for i := 0; i < b.N; i++ {
		for j := range index {
			sum += a[index[j]]
		}
	}
	if expected := uint(b.N*M*(M-1)) / 2; sum != expected {
		b.Errorf("Computed %d, expected %d", sum, expected)
	}
}

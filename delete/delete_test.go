package bench

import (
	"math/rand"
	"slices"
	"testing"
)

func customDelete[S ~[]E, E any](s S, i, j int) S {
	// this is copied vebatim from the standard library
	// thus the performance is expected to be the same
	// as slices.Delete
	_ = s[i:j] // bounds check
	return append(s[:i], s[j:]...)
}

func customDeleteZero[S ~[]E, E any](s S, i, j int) S {
	_ = s[i:j]
	s2 := append(s[:i], s[j:]...)
	clear(s[len(s)-j+i:])
	return s2
}

func BenchmarkSlicesDeleteOne(b *testing.B) {
	// Standard library
	benchmarkDeleteOne(b, slices.Delete[[]int, int])
}

func BenchmarkCustomDeleteOne(b *testing.B) {
	benchmarkDeleteOne(b, customDelete[[]int, int])
}

func BenchmarkCustomDeleteOneZero(b *testing.B) {
	benchmarkDeleteOne(b, customDeleteZero[[]int, int])
}

func BenchmarkSlicesDeleteRange(b *testing.B) {
	// Standard library
	benchmarkDeleteRange(b, slices.Delete[[]int, int])
}

func BenchmarkCustomDeleteRange(b *testing.B) {
	benchmarkDeleteRange(b, customDelete[[]int, int])
}

func BenchmarkCustomDeleteRangeZero(b *testing.B) {
	benchmarkDeleteRange(b, customDeleteZero[[]int, int])
}

func benchmarkDeleteOne[T any](b *testing.B, deleter func([]T, int, int) []T) {
	var random [1_000]int
	r := rand.New(rand.NewSource(42))
	for i := range random {
		random[i] = r.Intn(1_000)
	}
	a := make([]T, 1_000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		k := 0
		for j := 0; j < 500; j++ {
			// Consume 2 numbers, use delete
			size := random[k]
			k++
			a = a[:size]
			// Delete 1 element
			i := random[k] % len(a)
			k++
			a = deleter(a, i, i+1)
			// Prevent optimizing out
			Sink = a
		}

	}
}

func benchmarkDeleteRange[T any](b *testing.B, deleter func([]T, int, int) []T) {
	var random [1_000]int
	r := rand.New(rand.NewSource(42))
	for i := range random {
		random[i] = r.Intn(1_000)
	}
	a := make([]T, 1_000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		k := 0
		for j := 0; j < 200; j++ {
			// Consume 5 numbers, use delete
			size := random[k]
			k++
			a = a[:size]
			{
				// Delete a range
				i := random[k] % len(a)
				j := i + (random[k+1] % (len(a) - i))
				k += 2
				a = deleter(a, i, j)
			}
			{
				// Delete a range
				i := random[k] % len(a)
				j := i + (random[k+1] % (len(a) - i))
				k += 2
				a = deleter(a, i, j)
			}
			// Prevent optimizing out
			Sink = a
		}

	}
}

var Sink any

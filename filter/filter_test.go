package filter

import (
	"math/rand"
	"testing"
)

const M = 10_000_000

var x []T

var expectedSize int

func init() {
	x = make([]T, M)
	for i := range x {
		x[i] = T(rand.Intn(9999))
	}

	for _, v := range x {
		if p(v) {
			expectedSize++
		}
	}
}

func BenchmarkFilterOnePass(b *testing.B) {
	var y []T
	for i := 0; i < b.N; i++ {
		y = FilterOnePass(x)
	}
	if len(y) != expectedSize {
		b.Fatalf("Expected %d, got %d", expectedSize, len(y))
	}
}

func BenchmarkFilterTwoPasses(b *testing.B) {
	var y []T
	for i := 0; i < b.N; i++ {
		y = FilterTwoPasses(x)
	}
	if len(y) != expectedSize {
		b.Fatalf("Expected %d, got %d", expectedSize, len(y))
	}
}

func BenchmarkFilterTwoPassesOptimized(b *testing.B) {
	var y []T
	for i := 0; i < b.N; i++ {
		y = FilterTwoPassesOptimized(x)
	}
	if len(y) != expectedSize {
		b.Fatalf("Expected %d, got %d", expectedSize, len(y))
	}
}

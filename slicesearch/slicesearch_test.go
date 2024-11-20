package slicesearch

import (
	"math/rand"
	"testing"
)

func BenchmarkSeq1_000(b *testing.B) {
	N := 1_000
	benchmarkSeq(b, N)
}

func BenchmarkSeq10_000(b *testing.B) {
	N := 10_000
	benchmarkSeq(b, N)
}

func BenchmarkSeq100_000(b *testing.B) {
	N := 100_000
	benchmarkSeq(b, N)
}

func BenchmarkSeq1_000_000(b *testing.B) {
	N := 1_000_000
	benchmarkSeq(b, N)
}

func benchmarkSeq(b *testing.B, size int) {
	needle := 4242
	negative := randomValues(size, 4000)
	positive := randomValues(size, 4000)
	positive[len(positive)-1] = needle

	for i := 0; i < b.N; i++ {
		found := searchSequential(negative, needle)
		if found {
			b.Errorf("Should not have found %d in negative haystack", needle)
		}
		found = searchSequential(positive, needle)
		if !found {
			b.Errorf("Should have found %d in positive haystack", needle)
		}
	}
}

func BenchmarkConc1_000(b *testing.B) {
	N := 1_000
	benchmarkConc(b, N)
}

func BenchmarkConc10_000(b *testing.B) {
	N := 10_000
	benchmarkConc(b, N)
}

func BenchmarkConc100_000(b *testing.B) {
	N := 100_000
	benchmarkConc(b, N)
}

func BenchmarkConc1_000_000(b *testing.B) {
	N := 1_000_000
	benchmarkConc(b, N)
}

func benchmarkConc(b *testing.B, size int) {
	needle := 4242
	negative := randomValues(size, 4000)
	positive := randomValues(size, 4000)
	positive[len(positive)-1] = needle

	for i := 0; i < b.N; i++ {
		found := searchConcurrent(negative, needle)
		if found {
			b.Errorf("Should not have found %d in negative haystack", needle)
		}
		found = searchConcurrent(positive, needle)
		if !found {
			b.Errorf("Should have found %d in positive haystack", needle)
		}
	}
}

func randomValues(size int, maxValue int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = rand.Intn(maxValue)
	}
	return a
}

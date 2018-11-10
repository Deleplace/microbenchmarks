package shift

import (
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	M := 1000
	r := make([]int, M)
	rand.Seed(123)
	for i := range r {
		r[i] = rand.Intn(600)
	}

	b.ResetTimer()
	distinct := 0
	for i := 0; i < b.N; i++ {
		seen := make([]bool, M)
		for _, v := range r {
			if !seen[v] {
				seen[v] = true
				distinct++
			}
		}
	}
	// log.Println("N =", b.N, "=> distinct =", distinct)
}

func BenchmarkMap(b *testing.B) {
	M := 1000
	r := make([]int, M)
	rand.Seed(123)
	for i := range r {
		r[i] = rand.Intn(600)
	}

	b.ResetTimer()
	distinct := 0
	for i := 0; i < b.N; i++ {
		seen := make(map[int]bool, M)
		for _, v := range r {
			seen[v] = true
		}
		distinct += len(seen)
	}
	// log.Println("N =", b.N, "=> distinct =", distinct)
}

func BenchmarkBigInt(b *testing.B) {
	M := 1000
	r := make([]int, M)
	rand.Seed(123)
	for i := range r {
		r[i] = rand.Intn(600)
	}

	b.ResetTimer()
	distinct := 0
	for i := 0; i < b.N; i++ {
		// a *big.Int serves as a very cromulent bitset
		var seen big.Int
		for _, v := range r {
			if seen.Bit(v) == 0 {
				(&seen).SetBit(&seen, v, 1)
				distinct++
			}
		}
	}
	// log.Println("N =", b.N, "=> distinct =", distinct)
}

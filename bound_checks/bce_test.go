package bench

import (
	"math/rand"
	"testing"
)

var data []uint64

func init() {
	data = make([]uint64, 1_000_000)
	r := rand.New(rand.NewSource(42))
	for i := range data {
		data[i] = r.Uint64()
	}
}

// max4A computes the max out of 4 numbers
func max4A(x []uint64) uint64 {
	max := x[0]
	for _, v := range x[1:4] {
		if v > max {
			max = v
		}
	}
	return max
}

// max4B computes the max out of 4 numbers
func max4B(x []uint64) uint64 {
	max := x[3]
	for _, v := range x[0:3] {
		if v > max {
			max = v
		}
	}
	return max
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := uint64(0)
		for j := 0; j < len(data); j += 4 {
			sum += max4A(data[j : j+4])
		}
		Sink = sum
	}
	// b.Log(Sink)
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := uint64(0)
		for j := 0; j < len(data); j += 4 {
			sum += max4B(data[j : j+4])
		}
		Sink = sum
	}
	// b.Log(Sink)
}

// Prevents optimizing things away
var Sink uint64

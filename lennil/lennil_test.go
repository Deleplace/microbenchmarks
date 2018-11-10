package lennil

import (
	"math/rand"
	"testing"
)

func BenchmarkLen(b *testing.B) {
	const M = 177
	a := make([][]int, M)
	for i := range a {
		if rand.Float32() < .8 {
			a[i] = []int{22, 33}
		}
	}

	b.ResetTimer()
	var count = 0
	for i := 0; i < b.N; i++ {
		for _, v := range a {
			if len(v) > 0 {
				count++
			}
		}
	}
	if count < 0 {
		b.Errorf("Dummy error")
	}
}

// Surprise! this overly verbose version seems... much much faster, when
// there are lots of nils!
func BenchmarkNilLen(b *testing.B) {
	const M = 177
	a := make([][]int, M)
	for i := range a {
		if rand.Float32() < .8 {
			a[i] = []int{22, 33}
		}
	}

	b.ResetTimer()
	var count = 0
	for i := 0; i < b.N; i++ {
		for _, v := range a {
			if v != nil && len(v) > 0 {
				count++
			}
		}
	}
	if count < 0 {
		b.Errorf("Dummy error")
	}
}

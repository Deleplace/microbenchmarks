package rangeto

import "testing"

func to(n int) []struct{} {
	return make([]struct{}, n)
}

func BenchmarkIterationClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum = 0
		for i := 0; i < 1; i++ {
			sum += i
		}
		for i := 0; i < 10; i++ {
			sum += i
		}
		for i := 0; i < 100; i++ {
			sum += i
		}
		// for i := 0; i < 1000; i++ {
		// 	sum += i
		// }
	}
}

func BenchmarkRangeTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum = 0
		for i := range to(1) {
			sum += i
		}
		for i := range to(10) {
			sum += i
		}
		for i := range to(100) {
			sum += i
		}
		// for i := range to(1000) {
		// 	sum += i
		// }
	}
}

var sum = 0

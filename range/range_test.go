package rangeto

import "testing"

var a = []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

func BenchmarkIterationClassic(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(a); i++ {
			sum += len(a[i])
		}
	}
	if expected := 35 * b.N; sum != expected {
		b.Errorf("Expected %d, got %d", expected, sum)
	}
}

func BenchmarkRangeIndex(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		for i := range a {
			sum += len(a[i])
		}
	}
	if expected := 35 * b.N; sum != expected {
		b.Errorf("Expected %d, got %d", expected, sum)
	}
}

func BenchmarkRangeValue(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		for _, s := range a {
			sum += len(s)
		}
	}
	if expected := 35 * b.N; sum != expected {
		b.Errorf("Expected %d, got %d", expected, sum)
	}
}

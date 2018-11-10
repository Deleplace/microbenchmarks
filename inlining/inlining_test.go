package inlining

import "testing"

func BenchmarkNoCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
	}
}

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func BenchmarkCheckInFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			if false {
				check(i)
			}
			sum += i
		}
	}
}

func BenchmarkFalseInCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			falseInCheck(i)
			sum += i
		}
	}
}

func BenchmarkVariadicCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			variadicCheck(i)
			sum += i
		}
	}
}

func BenchmarkVariadicCheckInFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			if false {
				variadicCheck(i)
			}
			sum += i
		}
	}
}

func BenchmarkFalseInVariadicCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			falseInVariadicCheck(i)
			sum += i
		}
	}
}

func BenchmarkFail(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || fail(i)
			sum += i
		}
	}
}

func BenchmarkVariadicFail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || variadicFail(i)
			sum += i
		}
	}
}

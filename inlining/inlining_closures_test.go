package inlining

import (
	"fmt"
	"testing"
)

func BenchmarkClosureNoCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
	}
}

func BenchmarkClosureCheck(b *testing.B) {
	check := func(i int) {
		if i < 0 {
			panic("Negative number")
		}
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func BenchmarkClosureCheckInFalse(b *testing.B) {
	check := func(i int) {
		if i < 0 {
			panic("Negative number")
		}
	}

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

func BenchmarkClosureFalseInCheck(b *testing.B) {
	check := func(i int) {
		if false {
			if i < 0 {
				panic("Negative number")
			}
		}
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func BenchmarkClosureVariadicCheck(b *testing.B) {
	check := func(i ...int) {
		if i[0] < 0 {
			panic("Negative number")
		}
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func BenchmarkClosureVariadicCheckInFalse(b *testing.B) {
	check := func(i ...int) {
		if i[0] < 0 {
			panic("Negative number")
		}
	}

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

func BenchmarkClosureFalseInVariadicCheck(b *testing.B) {
	check := func(i ...int) {
		if false {
			if i[0] < 0 {
				panic("Negative number")
			}
		}
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func BenchmarkClosureFail(b *testing.B) {
	fail := func(i int) bool {
		panic(fmt.Sprintf("Negative number: %d", i))
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || fail(i)
			sum += i
		}
	}
}

func BenchmarkClosureVariadicFail(b *testing.B) {
	fail := func(i int) bool {
		panic(fmt.Sprintf("Negative things: %v", i))
	}

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || fail(i)
			sum += i
		}
	}
}

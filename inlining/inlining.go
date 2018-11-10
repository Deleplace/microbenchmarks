package inlining

import (
	"fmt"
	"testing"
)

func check(i int) {
	if i < 0 {
		panic("Negative number")
	}
}

func falseInCheck(i int) {
	if false {
		if i < 0 {
			panic("Negative number")
		}
	}
}

func variadicCheck(i ...int) {
	if i[0] < 0 {
		panic("Negative number")
	}
}

func falseInVariadicCheck(i ...int) {
	if false {
		if i[0] < 0 {
			panic("Negative number")
		}
	}
}

func fail(i int) bool {
	panic(fmt.Sprintf("Negative number: %d", i))
}

func variadicFail(i int) bool {
	panic(fmt.Sprintf("Negative things: %v", i))
}

/////////////

func NoCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
	}
}

func Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			check(i)
			sum += i
		}
	}
}

func CheckInFalse(b *testing.B) {
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

func FalseInCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			falseInCheck(i)
			sum += i
		}
	}
}

func VariadicCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			variadicCheck(i)
			sum += i
		}
	}
}

func VariadicCheckInFalse(b *testing.B) {
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

func FalseInVariadicCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			falseInVariadicCheck(i)
			sum += i
		}
	}
}

func Fail(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || fail(i)
			sum += i
		}
	}
}

func VariadicFail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < 10000; i++ {
			_ = i >= 0 || variadicFail(i)
			sum += i
		}
	}
}

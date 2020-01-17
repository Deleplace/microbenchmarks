package newerror

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var global = errors.New("global")

func Global() error {
	return global
}

func BenchmarkGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = Global()
	}
}

func ErrorsNew() error {
	return errors.New("constant message")
}

func BenchmarkErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = ErrorsNew()
	}
}

func FmtErrorf() error {
	return fmt.Errorf("constant message")
}

func BenchmarkFmtErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = FmtErrorf()
	}
}

func ErrorsNewArg(arg int) error {
	return errors.New("param message: " + strconv.Itoa(arg))
}

func BenchmarkErrorsNewArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = ErrorsNewArg(i)
	}
}

func FmtErrorfArg(arg int) error {
	return fmt.Errorf("param message: %d", arg)
}

func BenchmarkFmtErrorfArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = FmtErrorfArg(i)
	}
}

// Sink prevents some optimizing out
var Sink error

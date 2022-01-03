package bench

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// Convert a slice of n bytes to the equivalent hex string
// of 2n digits.
//
// With n=11

var buffer = []byte("Hello world")

func TestResult(t *testing.T) {
	s1 := fmt.Sprintf("%x", buffer)
	s2 := hex.EncodeToString(buffer)
	if s1 != s1 {
		t.Errorf("%q != %q", s1, s2)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = fmt.Sprintf("%x", buffer)
	}
}

func BenchmarkEncodeToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = hex.EncodeToString(buffer)
	}
}

// Prevents optimizing computations away (sort of)
var Sink string

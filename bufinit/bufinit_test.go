package bench

import "testing"

var M = 20 * 1024 * 1024

var Sink []byte

func BenchmarkAlloc(b *testing.B) {
	for range b.N {
		buf := make([]byte, M)
		// buf is already initialized to zero
		Sink = buf
	}
}

func BenchmarkAllocSetZero(b *testing.B) {
	for range b.N {
		buf := make([]byte, M)
		// buf is already initialized to zero
		// Setting elements to zero again
		for i := range buf {
			buf[i] = 0
		}
		Sink = buf
	}
}

func BenchmarkAllocSetOne(b *testing.B) {
	for range b.N {
		buf := make([]byte, M)
		// buf is already initialized to zero
		// Setting elements to one
		for i := range buf {
			buf[i] = 1
		}
		Sink = buf
	}
}

func BenchmarkAllocSetFirst(b *testing.B) {
	for range b.N {
		buf := make([]byte, M)
		// buf is already initialized to zero
		// Setting first to zero, repeatedly
		for range buf {
			buf[0] = 0
		}
		Sink = buf
	}
}

package indirectcall

import (
	"bytes"
	"io"
	"testing"
)

// Number of int32 to be written by a call
const M = 100000

var a []int32
var buffer bytes.Buffer

func init() {
	a = make([]int32, M)
	for i := range a {
		a[i] = 12 + 3*int32(i)
	}
}

func BenchmarkWrite1(b *testing.B) {
	var w io.Writer = &buffer
	buffer.Grow(4 * len(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = write1(w, a)
		buffer.Reset()
	}
}

func BenchmarkWrite2(b *testing.B) {
	var w *bytes.Buffer = &buffer
	buffer.Grow(4 * len(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = write2(w, a)
		buffer.Reset()
	}
}

func BenchmarkWrite3(b *testing.B) {
	var w io.Writer = &buffer
	buffer.Grow(4 * len(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = write3(w, a)
		buffer.Reset()
	}
}

func BenchmarkWrite4(b *testing.B) {
	var w *bytes.Buffer = &buffer
	buffer.Grow(4 * len(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = write4(w, a)
		buffer.Reset()
	}
}

func BenchmarkWrite3or4(b *testing.B) {
	var w io.Writer = &buffer
	buffer.Grow(4 * len(a))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = write3or4(w, a)
		buffer.Reset()
	}
}

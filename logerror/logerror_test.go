package logerror

import (
	"fmt"
	"testing"
)

var data = []byte("trololololo")

func doWork() {
	for i, c := range data {
		switch c {
		case 'l':
			data[i] = 'L'
		case 'L':
			data[i] = 'l'
		}
	}
}

func TestWork(t *testing.T) {
	fmt.Println(string(data))
	doWork()
	fmt.Println(string(data))
	if expected, got := "troLoLoLoLo", string(data); got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
	doWork()
	fmt.Println(string(data))
	if expected, got := "trololololo", string(data); got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func BenchmarkWork1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWork()
	}
}

func BenchmarkWork2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWork()
		if a := len(data); a != 11 {
			b.Errorf("Wrong length %d", a)
		}
	}
}

func assert(condition bool, t testing.TB, pattern string, args ...interface{}) {
	if !condition {
		t.Errorf(pattern, args...)
	}
}

func BenchmarkWork3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWork()
		assert(len(data) == 11, b, "Wrong length %d", len(data))
	}
}

var assert2 = func(condition bool, t testing.TB, pattern string, args ...interface{}) {
	if !condition {
		t.Errorf(pattern, args...)
	}
}

func BenchmarkWork4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWork()
		assert2(len(data) == 11, b, "Wrong length %d", len(data))
	}
}

func BenchmarkWork5(b *testing.B) {
	assert2 = func(condition bool, t testing.TB, pattern string, args ...interface{}) {}
	for i := 0; i < b.N; i++ {
		doWork()
		assert2(len(data) == 11, b, "Wrong length %d", len(data))
	}
}

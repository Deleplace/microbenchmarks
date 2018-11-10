package threeindexslice

import (
	"testing"
)

func BenchmarkTwoIndex(b *testing.B) {
	s := []byte("  qwertyuiop   ")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z := trimA(s)
		_ = z
	}
}
func BenchmarkThreeIndex(b *testing.B) {
	s := []byte("  qwertyuiop   ")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z := trimB(s)
		_ = z
	}
}

func trimA(s []byte) []byte {
	const space = ' '
	n := len(s)
	low, high := 0, n
	for low < n && s[low] == space {
		low++
	}
	for high > low && s[high-1] == space {
		high--
	}
	return s[low:high]
}

func trimB(s []byte) []byte {
	const space = ' '
	n := len(s)
	low, high := 0, n
	for low < n && s[low] == space {
		low++
	}
	for high > low && s[high-1] == space {
		high--
	}
	return s[low:high:high]
}

func TestTwoIndex(t *testing.T) {
	s := []byte("   Hello world  ")
	z := trimA(s)
	z = append(z, 'Z', 'Z')

	if string(z) != "Hello worldZZ" {
		t.Errorf("Got %q", string(z))
	}
	if string(s) != "   Hello worldZZ" {
		t.Errorf("Got %q", string(s))
	}
}

func TestThreeIndex(t *testing.T) {
	s := []byte("   Hello world  ")
	z := trimB(s)
	z = append(z, 'Z', 'Z')

	if string(z) != "Hello worldZZ" {
		t.Errorf("Got %q", string(z))
	}
	if string(s) != "   Hello world  " {
		t.Errorf("Got %q", string(s))
	}
}

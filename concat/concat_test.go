package bench

import (
	"fmt"
	"strconv"
	"testing"
)

// Sink prevents optimizing computations away (somewhat)
var Sink string

func concat1(s string, i int) string {
	return s + strconv.Itoa(i)
}

func concat2(s string, i int) string {
	return fmt.Sprintf("%s%d", s, i)
}

func concat3(s string, i int) string {
	return fmt.Sprintf("%s%s", s, strconv.Itoa(i))
}

func TestConcat1(t *testing.T) {
	if got, expected := concat1("item-", 42), "item-42"; got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestConcat2(t *testing.T) {
	if got, expected := concat2("item-", 42), "item-42"; got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestConcat3(t *testing.T) {
	if got, expected := concat3("item-", 42), "item-42"; got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func BenchmarkConcat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = concat1("item-", 42)
	}
}

func BenchmarkConcat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = concat2("item-", 42)
	}
}

func BenchmarkConcat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = concat3("item-", 42)
	}
}

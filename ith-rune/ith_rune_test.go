package ithrune

import (
	"math/rand"
	"testing"
	"unicode/utf8"
)

func TestIthRune(t *testing.T) {
	for _, x := range []struct {
		s            string
		i            int
		expectedRune rune
		expectedOK   bool
	}{
		{"", -1, utf8.RuneError, false},
		{"", 0, utf8.RuneError, false},
		{"", 1, utf8.RuneError, false},
		{"a", -1, utf8.RuneError, false},
		{"a", 0, 'a', true},
		{"a", 1, utf8.RuneError, false},
		{"à", 0, 'à', true},
		{"à", 1, utf8.RuneError, false},
		{"Résumé", -1, utf8.RuneError, false},
		{"Résumé", 0, 'R', true},
		{"Résumé", 1, 'é', true},
		{"Résumé", 2, 's', true},
		{"Résumé", 3, 'u', true},
		{"Résumé", 4, 'm', true},
		{"Résumé", 5, 'é', true},
		{"Résumé", 6, utf8.RuneError, false},
	} {
		{
			buf := []byte(x.s)
			r, ok := IthRuneInBytes(buf, x.i)
			if r != x.expectedRune || ok != x.expectedOK {
				t.Errorf("IthRuneInBytes([]byte(%q), %d) -> expected %v, %v ; go %v, %v", x.s, x.i, x.expectedRune, x.expectedOK, r, ok)
			}
		}
		{
			r, ok := IthRuneInString(x.s, x.i)
			if r != x.expectedRune || ok != x.expectedOK {
				t.Errorf("IthRuneInBytes([]byte(%q), %d) -> expected %v, %v ; go %v, %v", x.s, x.i, x.expectedRune, x.expectedOK, r, ok)
			}
		}
	}
}

func Benchmark_10_1_runes_conv(b *testing.B) {
	benchRunesConv(b, 10, 1)
}

func Benchmark_10_1_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 10, 1)
}

func Benchmark_100_1_runes_conv(b *testing.B) {
	benchRunesConv(b, 100, 1)
}

func Benchmark_100_1_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 100, 1)
}

func Benchmark_1000_1_runes_conv(b *testing.B) {
	benchRunesConv(b, 100, 1)
}

func Benchmark_1000_1_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 100, 1)
}

func Benchmark_10_5_runes_conv(b *testing.B) {
	benchRunesConv(b, 10, 5)
}

func Benchmark_10_5_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 10, 5)
}

func Benchmark_100_5_runes_conv(b *testing.B) {
	benchRunesConv(b, 100, 5)
}

func Benchmark_100_5_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 100, 5)
}

func Benchmark_1000_5_runes_conv(b *testing.B) {
	benchRunesConv(b, 1000, 5)
}

func Benchmark_1000_5_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 1000, 5)
}

func Benchmark_100_50_runes_conv(b *testing.B) {
	benchRunesConv(b, 100, 50)
}

func Benchmark_100_50_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 100, 50)
}

func Benchmark_1000_500_runes_conv(b *testing.B) {
	benchRunesConv(b, 1000, 500)
}

func Benchmark_1000_500_decode_loop(b *testing.B) {
	benchDecodeLoop(b, 1000, 500)
}

func benchRunesConv(b *testing.B, stringLength int, k int) {
	rng = rand.New(rand.NewSource(42))
	s := randomString(stringLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runes := []rune(s)
		for j := 0; j < k; j++ {
			m := (j * 179) % stringLength
			r := runes[m]
			if r == utf8.RuneError {
				b.Errorf("Got RuneError in %q for index %d", s, m)
			}
		}
	}
}

func benchDecodeLoop(b *testing.B, stringLength int, k int) {
	rng = rand.New(rand.NewSource(42))
	s := randomString(stringLength)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < k; j++ {
			m := (j * 179) % stringLength
			r, ok := IthRuneInString(s, m)
			if r == utf8.RuneError || ok == false {
				b.Errorf("Got (%v, %v) in %q for index %d", r, ok, s, m)
			}
		}
	}
}

var alphabet = []rune("aàâäbcdeéèêëfghiîïjklmnoôöpqrstuûüvwxyz0123456789         ")

var rng = rand.New(rand.NewSource(42))

func randomString(n int) string {
	a := make([]rune, n)
	for i := range a {
		a[i] = alphabet[rng.Intn(len(alphabet))]
	}
	return string(a)
}

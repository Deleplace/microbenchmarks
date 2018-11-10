package dedupe

import (
	"math/rand"
	"sort"
	"testing"
)

// Size of list. Tune manually to produce interesting values.
const (
	M   = 10000
	max = 1e4
)

// Exported global var, to make sure the compiler doesn't wipe away the computations
var X []int

func BenchmarkDedupeSort(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		sort.Ints(a)
		d := []int{a[0]}
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				d = append(d, a[i])
			}
		}

		X = d
		if i == 0 {
			// fmt.Println("X has length", len(X))
			b.StopTimer()
			if h := hash(X); h != 5928143255714578365 {
				// b.Fatalf("hash is %v", h)
			}
			b.StartTimer()
		}
	}
}

func _BenchmarkDedupeSort2(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		sort.Ints(a)
		distinct := 1
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				distinct++
			}
		}

		d := make([]int, 1, distinct)
		d[0] = a[0]
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				d = append(d, a[i])
			}
		}
		X = d
		if i == 0 {
			b.StopTimer()
			if h := hash(X); h != 5928143255714578365 {
				// b.Fatalf("hash is %v", h)
			}
			b.StartTimer()
		}
	}
}

func _BenchmarkDedupeSort3(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		sort.Ints(a)
		distinct := 1
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				distinct++
			}
		}

		d := make([]int, distinct)
		d[0] = a[0]
		j := 1
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				d[j] = a[i]
				j++
			}
		}
		X = d
		if i == 0 {
			b.StopTimer()
			if h := hash(X); h != 5928143255714578365 {
				// b.Fatalf("hash is %v", h)
			}
			b.StartTimer()
		}
	}
}

func _BenchmarkDedupeSort4(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		sort.Ints(a)
		d := make([]int, len(a))
		d[0] = a[0]
		j := 1
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				d[j] = a[i]
				j++
			}
		}
		d = d[:j]
		X = d
		if i == 0 {
			b.StopTimer()
			if h := hash(X); h != 5928143255714578365 {
				// b.Fatalf("hash is %v", h)
			}
			b.StartTimer()
		}
	}
}
func BenchmarkDedupeSort5(b *testing.B) {
	rand.Seed(123)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomValues()
		b.StartTimer()

		// This one is destructive! Doesn't preserve the original values!

		sort.Ints(a)
		j := 1
		for i := 1; i < len(a); i++ {
			if a[i] != a[i-1] {
				a[j] = a[i]
				j++
			}
		}
		d := a[:j]
		X = d
		if i == 0 {
			b.StopTimer()
			if h := hash(X); h != 5928143255714578365 {
				// b.Fatalf("hash is %v", h)
			}
			b.StartTimer()
		}
	}
}

func BenchmarkDedupeHashMap(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make(map[int]bool)
		var d []int
		for _, v := range a {
			if !m[v] {
				d = append(d, v)
				m[v] = true
			}
		}

		X = d
	}
}

func _BenchmarkDedupeHashMap2(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make(map[int]bool, len(a))
		var d []int
		for _, v := range a {
			if !m[v] {
				d = append(d, v)
				m[v] = true
			}
		}

		X = d
	}
}

func _BenchmarkDedupeHashMap3(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make(map[int]bool, len(a))
		d := make([]int, 0, len(a))
		for _, v := range a {
			if !m[v] {
				d = append(d, v)
				m[v] = true
			}
		}

		X = d
	}
}

func _BenchmarkDedupeHashMap4(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make(map[int]bool, len(a))
		for _, v := range a {
			if !m[v] {
				m[v] = true
			}
		}
		d := make([]int, len(m))
		i := 0
		for v := range m {
			d[i] = v
			i++
		}

		X = d
	}
}

func BenchmarkDedupeHashMap5(b *testing.B) {
	rand.Seed(123)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomValues()
		b.StartTimer()

		// This one is destructive! Doesn't preserve the original values!

		m := make(map[int]bool, len(a))
		for _, v := range a {
			m[v] = true
		}
		i := 0
		for v := range m {
			a[i] = v
			i++
		}
		d := a[:i]

		X = d
	}
}

func BenchmarkDedupeBitset(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make([]byte, max)
		for _, v := range a {
			m[v] = 1
		}
		d := make([]int, len(a))
		j := 0
		for _, v := range a {
			if m[v] == 1 {
				d[j] = v
				j++
				m[v] = 2
			}
		}
		d = d[:j]

		X = d
	}
}

func BenchmarkDedupeBitsetBool(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make([]bool, max)
		for _, v := range a {
			m[v] = true
		}
		d := make([]int, len(a))
		j := 0
		for _, v := range a {
			if m[v] {
				d[j] = v
				j++
				m[v] = false
			}
		}
		d = d[:j]

		X = d
	}
}

func BenchmarkDedupeBitsetBoolAppend(b *testing.B) {
	rand.Seed(123)
	a := randomValues()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		shuffle(a)
		b.StartTimer()

		m := make([]bool, max)
		for _, v := range a {
			m[v] = true
		}
		var d []int
		for _, v := range a {
			if m[v] {
				d = append(d, v)
				m[v] = false
			}
		}

		X = d
	}
}

func BenchmarkDedupeBitset2(b *testing.B) {
	rand.Seed(123)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomValues()
		b.StartTimer()

		// This one is destructive! Doesn't preserve the original values!

		m := make([]byte, max)
		for _, v := range a {
			m[v] = 1
		}
		j := 0
		for _, v := range a {
			if m[v] == 1 {
				a[j] = v
				j++
				m[v] = 2
			}
		}
		d := a[:j]

		X = d
	}
}

func BenchmarkDedupeBitset2Bool(b *testing.B) {
	rand.Seed(123)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := randomValues()
		b.StartTimer()

		// This one is destructive! Doesn't preserve the original values!

		m := make([]bool, max)
		for _, v := range a {
			m[v] = true
		}
		j := 0
		for _, v := range a {
			if m[v] {
				a[j] = v
				j++
				m[v] = false
			}
		}
		d := a[:j]

		X = d
	}
}

func randomValues() []int {
	a := make([]int, M)
	for i := range a {
		a[i] = rand.Intn(max)
	}
	return a
}

func shuffle(a []int) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

func hash(a []int) int {
	x := 1
	for _, v := range a {
		x = 3*x + v
	}
	return x
}

package copypast

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkAlloc_1_1_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAlloc(b, 1, 1000, 1, 1000)
	}
}

func BenchmarkAlloc_1_1_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAlloc(b, 1, 5000, 1, 5000)
	}
}

func BenchmarkAlloc_4_4_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAlloc(b, 4, 5000, 4, 5000)
	}
}

func BenchmarkAppend_1_1_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAppend(b, 1, 1000, 1, 1000)
	}
}

func BenchmarkAppend_1_1_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAppend(b, 1, 5000, 1, 5000)
	}
}

func BenchmarkAppend_4_4_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchAppend(b, 4, 5000, 4, 5000)
	}
}

func BenchmarkRef_1_1_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRef(b, 1, 1000, 1, 1000)
	}
}

func BenchmarkRef_1_1_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRef(b, 1, 5000, 1, 5000)
	}
}

func BenchmarkRef_4_4_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchRef(b, 4, 5000, 4, 5000)
	}
}

func BenchmarkUnprotected_4_4_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		launchUnprotected(b, 4, 5000, 4, 5000)
	}
}

func launchAlloc(b *testing.B, R int, NR int, A int, maxsize int) {
	x := make([]int, 0, maxsize)
	var lock sync.Mutex

	var wg sync.WaitGroup
	wg.Add(R + A)
	for r := 0; r < R; r++ {
		go func() {
			for i := 0; i < NR; i++ {
				lock.Lock()
				xx := make([]int, len(x))
				copy(xx, x)
				lock.Unlock()

				if sum(xx) == -1 {
					b.Error("wut")
				}

				runtime.Gosched()
			}
			wg.Done()
		}()

	}
	for a := 0; a < A; a++ {
		a := a
		go func() {
			s := 0
			for s < maxsize {
				lock.Lock()
				x = append(x, a)
				s = len(x)
				lock.Unlock()
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func launchAppend(b *testing.B, R int, NR int, A int, maxsize int) {
	x := make([]int, 0, maxsize)
	var lock sync.Mutex

	var wg sync.WaitGroup
	wg.Add(R + A)
	for r := 0; r < R; r++ {
		go func() {
			var xx []int
			for i := 0; i < NR; i++ {
				lock.Lock()
				xx = xx[:0]
				xx = append(xx, x...)
				lock.Unlock()

				if sum(xx) == -1 {
					b.Error("wut")
				}

				runtime.Gosched()
			}
			wg.Done()
		}()

	}
	for a := 0; a < A; a++ {
		a := a
		go func() {
			s := 0
			for s < maxsize {
				lock.Lock()
				x = append(x, a)
				s = len(x)
				lock.Unlock()
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func launchRef(b *testing.B, R int, NR int, A int, maxsize int) {
	x := make([]int, 0, maxsize)
	var lock sync.Mutex

	var wg sync.WaitGroup
	wg.Add(R + A)
	for r := 0; r < R; r++ {
		go func() {
			for i := 0; i < NR; i++ {
				lock.Lock()
				xx := x
				lock.Unlock()

				if sum(xx) == -1 {
					b.Error("wut")
				}

				runtime.Gosched()
			}
			wg.Done()
		}()

	}
	for a := 0; a < A; a++ {
		a := a
		go func() {
			s := 0
			for s < maxsize {
				lock.Lock()
				x = append(x, a)
				s = len(x)
				lock.Unlock()
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func launchUnprotected(b *testing.B, R int, NR int, A int, maxsize int) {
	x := make([]int, 0, maxsize)

	var wg sync.WaitGroup
	wg.Add(R + A)
	for r := 0; r < R; r++ {
		go func() {
			for i := 0; i < NR; i++ {
				xx := x

				if sum(xx) == -1 {
					b.Error("wut")
				}

				runtime.Gosched()
			}
			wg.Done()
		}()

	}
	for a := 0; a < A; a++ {
		a := a
		go func() {
			s := 0
			for s < maxsize {
				x = append(x, a)
				s = len(x)
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
	}
	return s
}

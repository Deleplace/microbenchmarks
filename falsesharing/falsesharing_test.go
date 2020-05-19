package falsesharing

import (
	"sync"
	"testing"
)

const M = 20000000

// Sink makes sure results are not optimized away by the compiler
var Sink int

var SinkMu sync.Mutex

const expectedSum = (M * (M + 1)) / 2

// How fast is it to access a[0], a[1], ..., a[M-1] ?
func BenchmarkAccessForward(b *testing.B) {
	a := make([]int, M)
	for i := range a {
		a[i] = 1 + i
	}

	b.ResetTimer()
	var sum int
	for i := 0; i < b.N; i++ {
		for j := 0; j < M; j++ {
			sum += a[j]
		}
	}
	b.StopTimer()
	SinkMu.Lock()
	Sink = sum
	SinkMu.Unlock()
	if sum/b.N != expectedSum {
		b.Errorf("Wrong sum %d != %d", sum, expectedSum)
	}
}

// How fast is it to access a[M-1], a[M-2], ..., a[0] ?
func BenchmarkAccessBackward(b *testing.B) {
	a := make([]int, M)
	for i := range a {
		a[i] = 1 + i
	}

	b.ResetTimer()
	var sum int
	for i := 0; i < b.N; i++ {
		for j := M - 1; j >= 0; j-- {
			sum += a[j]
		}
	}
	b.StopTimer()
	SinkMu.Lock()
	Sink = sum
	SinkMu.Unlock()
	if sum/b.N != expectedSum {
		b.Errorf("Wrong sum %d != %d", sum, expectedSum)
	}
}

// How fast is it to access concurrently with 2 goroutines
// a[0], ..., a[M/2-1]
// a[M/2], ..., a[M-1]
func BenchmarkAccessConcurrentForwardForward(b *testing.B) {
	a := make([]int, M)
	for i := range a {
		a[i] = 1 + i
	}

	b.ResetTimer()
	var sum1, sum2 int
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		const middle = M / 2
		go func() {
			for j := 0; j < middle; j++ {
				sum1 += a[j]
			}
			wg.Done()
		}()
		go func() {
			for j := middle; j < M; j++ {
				sum2 += a[j]
			}
			wg.Done()
		}()
		wg.Wait()

	}
	b.StopTimer()
	sum := sum1 + sum2
	SinkMu.Lock()
	Sink = sum
	SinkMu.Unlock()
	if sum/b.N != expectedSum {
		b.Errorf("Wrong sum %d != %d", sum, expectedSum)
	}
}

// How fast is it to access concurrently with 2 goroutines
// a[0], ..., a[M/2-1]
// a[M-1], ..., a[M/2]
func BenchmarkAccessConcurrentForwardBackward(b *testing.B) {
	a := make([]int, M)
	for i := range a {
		a[i] = 1 + i
	}

	b.ResetTimer()
	var sum1, sum2 int
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		const middle = M / 2
		go func() {
			for j := 0; j < middle; j++ {
				sum1 += a[j]
			}
			wg.Done()
		}()
		go func() {
			for j := M - 1; j >= middle; j-- {
				sum2 += a[j]
			}
			wg.Done()
		}()
		wg.Wait()

	}
	b.StopTimer()
	sum := sum1 + sum2
	SinkMu.Lock()
	Sink = sum
	SinkMu.Unlock()
	if sum/b.N != expectedSum {
		b.Errorf("Wrong sum %d != %d", sum, expectedSum)
	}
}

// How fast is it to access concurrently with 2 goroutines
// a[0], a[2], ..., a[M/2-2]
// a[1], a[3], ..., a[M/2-1]
func BenchmarkAccessConcurrentInterleaved(b *testing.B) {
	a := make([]int, M)
	for i := range a {
		a[i] = 1 + i
	}

	b.ResetTimer()
	var sum1, sum2 int
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		const middle = M / 2
		go func() {
			for j := 0; j < M; j += 2 {
				sum1 += a[j]
			}
			wg.Done()
		}()
		go func() {
			for j := 1; j < M; j += 2 {
				sum2 += a[j]
			}
			wg.Done()
		}()
		wg.Wait()

	}
	b.StopTimer()
	sum := sum1 + sum2
	SinkMu.Lock()
	Sink = sum
	SinkMu.Unlock()
	if sum/b.N != expectedSum {
		b.Errorf("Wrong sum %d != %d", sum, expectedSum)
	}
}

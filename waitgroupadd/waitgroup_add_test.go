package waitgroupadd

import (
	"sync"
	"testing"
)

const M = 100000

func BenchmarkSequential_1000_1000(b *testing.B) {
	benchmarkSequential(b, 1000, 1000)
}
func BenchmarkAdd1_1000_1000(b *testing.B) {
	benchmarkAdd1(b, 1000, 1000)
}

func BenchmarkAddN_1000_1000(b *testing.B) {
	benchmarkAddN(b, 1000, 1000)
}

func BenchmarkSequential_1000_200(b *testing.B) {
	benchmarkSequential(b, 1000, 200)
}

func BenchmarkAdd1_1000_200(b *testing.B) {
	benchmarkAdd1(b, 1000, 200)
}

func BenchmarkAddN_1000_200(b *testing.B) {
	benchmarkAddN(b, 1000, 200)
}

func BenchmarkSequential_1000_20(b *testing.B) {
	benchmarkSequential(b, 1000, 20)
}
func BenchmarkAdd1_1000_20(b *testing.B) {
	benchmarkAdd1(b, 1000, 20)
}

func BenchmarkAddN_1000_20(b *testing.B) {
	benchmarkAddN(b, 1000, 20)
}

// T is the number of tasks
// M is the number of operations per task
func benchmarkSequential(b *testing.B, T, M int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < T; j++ {
			x := int64(123456789) + int64(j)
			for k := 0; k < M; k++ {
				if x%2 == 0 {
					x /= 2
				} else {
					x = 3*x + 1
				}
				if x == -1 {
					// Should never happen.
					break
				}
			}
		}
	}
}

// Call WaitGroup.Add(1) for each task
// T is the number of tasks
// M is the number of operations per task
func benchmarkAdd1(b *testing.B, T, M int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < T; j++ {
			wg.Add(1)
			x := int64(123456789) + int64(j)
			go func() {
				for k := 0; k < M; k++ {
					if x%2 == 0 {
						x /= 2
					} else {
						x = 3*x + 1
					}
					if x == -1 {
						// Should never happen.
						break
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// Call WaitGroup.Add(N) only once
// T is the number of tasks
// M is the number of operations per task
func benchmarkAddN(b *testing.B, T, M int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(T)
		for j := 0; j < T; j++ {
			x := int64(123456789) + int64(j)
			go func() {
				for k := 0; k < M; k++ {
					if x%2 == 0 {
						x /= 2
					} else {
						x = 3*x + 1
					}
					if x == -1 {
						// Should never happen.
						break
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

package slicesearch

import (
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
)

// What is faster when searching a value in a large slice:
// - sequentially?
// - concurrently?
//
// Of course this may depend on the size of the slice.
// Small slices are always better off without a concurrency overhead.

func searchSequential(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func searchConcurrent(haystack []int, needle int) bool {
	var wg sync.WaitGroup
	// 1 worker by CPU.
	// This is just a heuristic, the sweet spot might be different.
	w := runtime.NumCPU()
	m := len(haystack) / w
	var result atomic.Bool
	for chunk := range slices.Chunk(haystack, m) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range chunk {
				if v == needle {
					result.Store(true)
					return
				}
			}
		}()
	}
	wg.Wait()
	return result.Load()
}

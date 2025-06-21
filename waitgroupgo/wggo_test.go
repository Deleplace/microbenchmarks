package bench

import (
	"slices"
	"sync"
	"testing"
)

// Go 1.25 introduces the new method sync.WaitGroup.Go
//
// Does it come with a perf penalty?
//
// Typical execution on my laptop:
//
// % go1.25rc1 test -bench=. -benchtime=20s wggo_test.go
// goos: darwin
// goarch: arm64
// cpu: Apple M1 Pro
//
// BenchmarkSearchA-10         	 4597898	      5601 ns/op
// BenchmarkSearchB-10         	 3878198	      6158 ns/op
//
// BenchmarkSearchALarge-10    	   52866	    455116 ns/op
// BenchmarkSearchBLarge-10    	   41680	    573406 ns/op
//
// BenchmarkSearchCLarge-10    	 8665867	      2665 ns/op
// BenchmarkSearchDLarge-10    	 8367006	      2914 ns/op
//
// PASS

func searchA(hay []int, needle int) bool {
	var wg sync.WaitGroup

	var found bool
	proclaimFound := sync.OnceFunc(func() {
		// fmt.Println("FOUND IT")
		found = true
	})

	var lookup func(a []int)
	lookup = func(a []int) {
		if len(a) == 0 {
			return
		}
		if len(a) == 1 {
			if a[0] == needle {
				proclaimFound()
			}
			return
		}
		// Divide in half, recurse left and right
		m := len(a) / 2
		wg.Add(2)
		go func() {
			defer wg.Done()
			lookup(a[:m])
		}()
		go func() {
			defer wg.Done()
			lookup(a[m:])
		}()
	}

	lookup(hay)

	wg.Wait()
	return found
}

func searchB(hay []int, needle int) bool {
	var wg sync.WaitGroup

	var found bool
	proclaimFound := sync.OnceFunc(func() {
		// fmt.Println("FOUND IT")
		found = true
	})

	var lookup func(a []int)
	lookup = func(a []int) {
		if len(a) == 0 {
			return
		}
		if len(a) == 1 {
			if a[0] == needle {
				proclaimFound()
			}
			return
		}
		// Divide in half, recurse left and right
		m := len(a) / 2
		wg.Go(func() {
			lookup(a[:m])
		})
		wg.Go(func() {
			lookup(a[m:])
		})
	}

	lookup(hay)

	wg.Wait()

	return found
}

func TestX(t *testing.T) {
	//t.Error("TODO")
}

func BenchmarkSearchA(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchA(hay1, needle)
		found2 = found2 || searchA(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
		b.Fatal("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

func BenchmarkSearchB(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchB(hay1, needle)
		found2 = found2 || searchB(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

//
// Large: repeat the hay 100x
//

func BenchmarkSearchALarge(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	hay1 = slices.Repeat(hay1, 100)
	hay2 = slices.Repeat(hay2, 100)
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchA(hay1, needle)
		found2 = found2 || searchA(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
		b.Fatal("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

func BenchmarkSearchBLarge(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	hay1 = slices.Repeat(hay1, 100)
	hay2 = slices.Repeat(hay2, 100)
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchB(hay1, needle)
		found2 = found2 || searchB(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

//
// searchC == searchA but fallbacks to linear search for small input
// searchD == searchB but fallbacks to linear search for small input

func linearLookup(hay []int, needle int) bool {
	found := false
	for _, v := range hay {
		if v == needle {
			found = true
			// We intentionally do not return early
		}
	}
	return found
}

func searchC(hay []int, needle int) bool {
	var wg sync.WaitGroup

	var found bool
	proclaimFound := sync.OnceFunc(func() {
		// fmt.Println("FOUND IT")
		found = true
	})

	var lookup func(a []int)
	lookup = func(a []int) {
		if len(a) <= 500 {
			found := linearLookup(a, needle)
			if found {
				proclaimFound()
			}
			return
		}
		// Divide in half, recurse left and right
		m := len(a) / 2
		wg.Add(2)
		go func() {
			defer wg.Done()
			lookup(a[:m])
		}()
		go func() {
			defer wg.Done()
			lookup(a[m:])
		}()
	}

	lookup(hay)

	wg.Wait()
	return found
}

func searchD(hay []int, needle int) bool {
	var wg sync.WaitGroup

	var found bool
	proclaimFound := sync.OnceFunc(func() {
		// fmt.Println("FOUND IT")
		found = true
	})

	var lookup func(a []int)
	lookup = func(a []int) {
		if len(a) <= 500 {
			found := linearLookup(a, needle)
			if found {
				proclaimFound()
			}
			return
		}
		// Divide in half, recurse left and right
		m := len(a) / 2
		wg.Go(func() {
			lookup(a[:m])
		})
		wg.Go(func() {
			lookup(a[m:])
		})
	}

	lookup(hay)

	wg.Wait()

	return found
}

func BenchmarkSearchCLarge(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	hay1 = slices.Repeat(hay1, 100)
	hay2 = slices.Repeat(hay2, 100)
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchC(hay1, needle)
		found2 = found2 || searchC(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
		b.Fatal("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

func BenchmarkSearchDLarge(b *testing.B) {

	needle := 42
	hay1 := []int{-3, 0, 53, 0, 2_000_000, 42, -3, 42}
	hay2 := []int{-3, 0, 53, 0, 2_000_000, 0, -3, 76, 92, 267, 93, 76, 987, 7, 7, 76, -1}
	hay1 = slices.Repeat(hay1, 100)
	hay2 = slices.Repeat(hay2, 100)
	var found1, found2 bool

	for b.Loop() {
		found1 = found1 || searchD(hay1, needle)
		found2 = found2 || searchD(hay2, needle)
	}
	if !found1 {
		b.Error("found1: expected true")
	}
	if found2 {
		b.Error("found2: expected false")
	}
}

package filter

// In some cases (where the predicate is cheap), it's cheaper to do
// an extra pass to count the number of elements to be kept, in order
// to make a single allocation with the exact size needed for the result.

//go:noinline
func FilterOnePass(x []T) []T {
	y := make([]T, 0)
	for _, v := range x {
		if p(v) {
			y = append(y, v)
		}
	}
	return y
}

//go:noinline
func FilterTwoPasses(x []T) []T {
	n := 0
	for _, v := range x {
		if p(v) {
			n++
		}
	}
	y := make([]T, 0, n)
	for _, v := range x {
		if p(v) {
			y = append(y, v)
		}
	}
	return y
}

//go:noinline
func FilterTwoPassesOptimized(x []T) []T {
	n := 0
	for _, v := range x {
		if p(v) {
			n++
		}
	}
	y := make([]T, n)
	// This doesn't use append
	j := 0
	for _, v := range x {
		if p(v) {
			y[j] = v
			j++
		}
	}
	return y
}

type T int

func p(v T) bool {
	return v%2 == 0
}

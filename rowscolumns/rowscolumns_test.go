package rowscolumns

import (
	"math/rand"
	"testing"
)

const M = 5000

var (
	a          [][]int
	correctSum int
)

func init() {
	a = make([][]int, M)

	// for i := range a {
	// 	a[i] = make([]int, M)
	// }

	data := make([]int, M*M)
	for i := range a {
		// a[i] = data[i*M : (i+1)*M]
		a[i] = data[i*M : (i+1)*M : (i+1)*M]
	}

	r := rand.New(rand.NewSource(123))
	for i := range a {
		for j := range a[i] {
			x := r.Int()
			a[i][j] = x
			correctSum += x
		}
	}
}

func BenchmarkSumRowsRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := range a {
			for j := range a[i] {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumRowsRange2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, row := range a {
			for _, x := range row {
				sum += x
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < M; i++ {
			for j := 0; j < M; j++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumRowsBCE1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < M; i++ {
			_ = a[i][M-1]
			for j := 0; j < M; j++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumRowsBCE2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		_ = a[M-1][M-1]
		for i := 0; i < M; i++ {
			for j := 0; j < M; j++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumColumns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for j := 0; j < M; j++ {
			for i := 0; i < M; i++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumColumnsBCE1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for j := 0; j < M; j++ {
			_ = a[M-1][j]
			for i := 0; i < M; i++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

func BenchmarkSumColumnsBCE2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		_ = a[M-1][M-1]
		for j := 0; j < M; j++ {
			for i := 0; i < M; i++ {
				sum += a[i][j]
			}
		}
		if sum != correctSum {
			b.Errorf("Expected %d, got %d", correctSum, sum)
		}
	}
}

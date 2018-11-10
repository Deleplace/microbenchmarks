package set1

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func BenchmarkBoolImpl(b *testing.B) {
	type set map[string]bool

	keys := [100]string{}
	for i := range keys {
		keys[i] = strings.Repeat(fmt.Sprintf("%d-", i), 5)
	}

	b.ResetTimer()
	s := make(set)
	count := 0
	for i := 0; i < b.N; i++ {
		// 1 write
		s[keys[i%len(keys)]] = true

		// 3 reads
		for k := 1; k <= 3; k++ {
			if s[keys[(i*k)%len(keys)]] {
				count++
			}
		}
	}
	log.Println("Count =", count)
}

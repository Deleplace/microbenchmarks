package bench

import (
	"regexp"
	"testing"
)

func BenchmarkMatchString(b *testing.B) {
	for range b.N {
		n := 0
		for _, pattern := range stringPatterns {
			for _, s := range testStrings {
				matches, err := regexp.MatchString(pattern, s)
				if err != nil {
					b.Fatal(err)
				}
				if matches {
					// b.Logf("%q matches pattern %q", s, pattern)
					n++
				}
			}
		}
		if want := 8; n != want {
			b.Errorf("want %d, got %d", want, n)
		}
		Sink = n
	}
}

func BenchmarkReMatchString(b *testing.B) {
	for range b.N {
		n := 0
		for _, pattern := range stringPatterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				b.Fatal(err)
			}
			for _, s := range testStrings {
				matches := re.MatchString(s)
				if matches {
					// b.Logf("%q matches pattern %q", s, pattern)
					n++
				}
			}
		}
		if want := 8; n != want {
			b.Errorf("want %d, got %d", want, n)
		}
		Sink = n
	}
}

func BenchmarkCacheMatchString(b *testing.B) {
	for range b.N {
		n := 0
		for _, pattern := range stringPatterns {
			re := cacheCompile(pattern)
			for _, s := range testStrings {
				matches := re.MatchString(s)
				if matches {
					// b.Logf("%q matches pattern %q", s, pattern)
					n++
				}
			}
		}
		if want := 8; n != want {
			b.Errorf("want %d, got %d", want, n)
		}
		Sink = n
	}
}

var cache = map[string]*regexp.Regexp{}

func cacheCompile(pattern string) *regexp.Regexp {
	if re, ok := cache[pattern]; ok {
		return re
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	cache[pattern] = re
	return re
}

var stringPatterns = []string{
	`192\.168\.1\.\d{1,3}`,
	`(\W|^)po[#\-]{0,1}\s{0,1}\d{2}[\s-]{0,1}\d{4}(\W|$)`,
	`192\.168\.\d{1,3}\.1`,
	`192\.\d{1,3}\.1\.1`,
	`\d{1,3}\.168\.1\.1`,
}

var testStrings = []string{
	"192.168.1.1",
	"192.168.1.9",
	"192.168.9.1",
	"192.9.1.1",
	"9.168.1.1",
	"localhost",
	"127.0.0.1",
}

var Sink int

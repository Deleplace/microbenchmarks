package bench

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < 10_000; i++ {
			s += words[i%len(words)] + " "
		}
		Sink = s
	}
}

func BenchmarkStringBuilderWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for i := 0; i < 10_000; i++ {
			sb.Write([]byte(words[i%len(words)]))
			sb.WriteByte(' ')
		}
		Sink = sb.String()
	}
}

func BenchmarkStringBuilderFprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for i := 0; i < 10_000; i++ {
			fmt.Fprint(&sb, words[i%len(words)])
			fmt.Fprint(&sb, " ")
		}
		Sink = sb.String()
	}
}

func BenchmarkBytesBufferWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for i := 0; i < 10_000; i++ {
			buf.Write([]byte(words[i%len(words)]))
			buf.WriteByte(' ')
		}
		Sink = buf.String()
	}
}

func BenchmarkBytesBufferFprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for i := 0; i < 10_000; i++ {
			fmt.Fprint(&buf, words[i%len(words)])
			fmt.Fprint(&buf, " ")
		}
		Sink = buf.String()
	}
}

var words = []string{
	"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit,", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua.",
}

var Sink string

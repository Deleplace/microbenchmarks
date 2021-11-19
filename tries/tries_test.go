package trie

import (
	_ "embed"
	"log"
	"strings"
	"testing"
)

func TestTrieStringLength(t *testing.T) {
	outputs := []string{
		trie1.String(),
		trie2.String(),
		trie3.String(),
		trie4.String(),
		trie5.String(),
		trie6.String(),
	}
	expected := len(outputs[0])
	for i, output := range outputs {
		if len(output) != expected {
			t.Errorf("#%d has length %d, expected %d", i, len(output), expected)
			//t.Errorf("%q", output)
		}
	}
}

func BenchmarkTrie1(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie1.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

func BenchmarkTrie2(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie2.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

func BenchmarkTrie3(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie3.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

func BenchmarkTrie4(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie4.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

func BenchmarkTrie5(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie5.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

func BenchmarkTrie6(b *testing.B) {
	expected := 5121

	var matches int
	for i := 0; i < b.N; i++ {
		matches = 0
		for _, input := range inputs {
			if trie6.hasWord(input) {
				matches++
			}
		}
		if matches != expected {
			b.Errorf("Expected %d, got %d", expected, matches)
		}
	}
}

var (
	trie1 Node1
	trie2 Node2
	trie3 = make(Node3)
	trie4 = make(Node4)
	trie5 = Node5{valid: true}
	trie6 = Node6{valid: true, endOfWord: false}
)

func init() {
	// log.Println(len(words), "words")
	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		n1 := &trie1
		n2 := &trie2
		n3 := trie3
		n4 := trie4
		n5 := &trie5
		n6 := &trie6

		// Sanity check
		for _, c := range word {
			if c < 'a' || c > 'z' {
				log.Fatalf("unexpected rune %c", c)
			}
		}

		// We need an END marker, otherwise prefix word would be lost
		word += "."

		for _, c := range word {
			// 1
			n1 = n1.insert(c)

			// 2
			n2 = n2.insert(c)

			// 3
			n3 = n3.insert(c)

			// 4
			n4.insert(c)
			if c != '.' {
				if n4[c] == nil {
					n4[c] = make(Node4)
				}
				n4 = n4[c]
			}

			// 5
			n5 = n5.insert(c)

			// 6
			n6 = n6.insert(c)
		}
	}

	// Remove uninteresting empty inputs
	tmp := make([]string, 0, len(inputs))
	for _, input := range inputs {
		if len(input) > 0 {
			tmp = append(tmp, input)
		}
	}
	inputs = tmp
}

// List from https://www.mit.edu/~ecprice/wordlist.10000
//
//go:embed words.txt
var englishDict string

var words []string = strings.Split(englishDict, "\n")

//go:embed inputs.txt
var inputList string

var inputs []string = strings.Split(inputList, "\n")

func TestBug6(t *testing.T) {
	t6 := Node6{valid: true, endOfWord: false}
	t.Logf("t6 = %v", &t6)

	n6 := &t6
	n6 = n6.insert('v')
	t.Logf("n6 = %v", n6)
	t.Logf("t6 = %v", &t6)
	n6 = n6.insert('.')
	t.Logf("n6 = %v", n6)
	t.Logf("t6 = %v", &t6)
	t.Logf("t6 = %v", t6)
}

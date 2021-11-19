package trie

import (
	"bytes"
	"fmt"
)

// Node3 is a map-based Trie type.
//
// Access a subsequent letter is direct.
// After a final point '.', a leaf Node3 consists of a nil map.
type Node4 map[rune]Node4

func (n Node4) has(c rune) bool {
	_, exists := n[c]
	return exists
}

func (n Node4) child(c rune) Node4 {
	// This might return nil even if the edge c exists!
	return n[c]
}

func (n Node4) insert(c rune) {
	if !n.has(c) {
		// n[c] = nil does create the transition for letter c
		n[c] = nil
	}
}

func (n Node4) String() string {
	// Exact same impl as Node3.String, as long as we
	// trust that all words properly end with a '.'
	var buf bytes.Buffer
	var rec func(accu string, nn Node4)
	rec = func(accu string, nn Node4) {
		for c, child := range nn {
			if c == '.' {
				fmt.Fprintf(&buf, "%s.\n", accu)
				continue
			}
			accuNext := accu + string(c)
			rec(accuNext, child)
		}
	}
	rec("", n)
	return buf.String()
}

func (n Node4) hasWord(word string) bool {
	if len(word) == 0 {
		return n.has('.')
	}
	c := rune(word[0])
	child := n.child(c)
	if child == nil {
		return false
	}
	return child.hasWord(word[1:])
}

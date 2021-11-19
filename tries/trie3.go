package trie

import (
	"bytes"
	"fmt"
)

// Node3 is a map-based Trie type.
//
// Access a subsequent letter is direct.
// After a final point '.', a leaf Node3 consists of an initialized empty map.
type Node3 map[rune]Node3

func (n Node3) has(c rune) bool {
	_, exists := n[c]
	return exists
}

func (n Node3) child(c rune) Node3 {
	return n[c]
}

func (n Node3) insert(c rune) Node3 {
	if !n.has(c) {
		n[c] = make(Node3)
	}
	return n[c]
}

func (n Node3) String() string {
	var buf bytes.Buffer
	var rec func(accu string, nn Node3)
	rec = func(accu string, nn Node3) {
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

func (n Node3) hasWord(word string) bool {
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

package trie

import (
	"bytes"
	"fmt"
)

// Node2 is another slice-based Trie type.
//
// To access a subsequent letter, one must iterate over all the children.
type Node2 []struct {
	letter rune
	node   Node2
}

func (n Node2) has(c rune) bool {
	//  Simple, legible RANGE loop
	// for _, child := range n {
	// 	if child.letter == c {
	// 		return true
	// 	}
	// }

	// RANGE prevents inlining
	// This is an equivalent loop without RANGE
	for i, length := 0, len(n); i < length; i++ {
		if n[i].letter == c {
			return true
		}
	}
	return false
}

func (n Node2) child(c rune) *Node2 {
	//  Simple, legible RANGE loop
	// for i := range n {
	// 	child := &n[i]
	// 	if child.letter == c {
	// 		return &child.node
	// 	}
	// }
	// return nil

	// RANGE prevents inlining
	// This is an equivalent loop without RANGE
	for i, length := 0, len(n); i < length; i++ {
		child := &n[i]
		if child.letter == c {
			return &child.node
		}
	}
	return nil
}

func (n *Node2) insert(c rune) *Node2 {
	if !n.has(c) {
		*n = append(*n, struct {
			letter rune
			node   Node2
		}{
			letter: c,
		})
	}
	return n.child(c)
}

func (n *Node2) String() string {
	var buf bytes.Buffer
	var rec func(accu string, nn Node2)
	rec = func(accu string, nn Node2) {
		for i := range nn {
			child := &nn[i]
			if child.letter == '.' {
				fmt.Fprintf(&buf, "%s.\n", accu)
				continue
			}
			accuNext := accu + string(child.letter)
			rec(accuNext, child.node)
		}
	}
	rec("", *n)
	return buf.String()
}

func (n Node2) hasWord(word string) bool {
	if len(word) == 0 {
		return n.has('.')
	}
	c := rune(word[0])
	for _, child := range n {
		if child.letter == c {
			return child.node.hasWord(word[1:])
		}
	}
	return false
}

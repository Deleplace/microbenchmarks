package trie

import (
	"bytes"
	"fmt"
)

// Node1 is a slice-based Trie type.
//
// The letter of the very first Node1 of a Trie is ignored.
// To access a subsequent letter, one must iterate over all the children.
type Node1 struct {
	letter   rune
	children []Node1
}

func (n Node1) has(c rune) bool {
	//  Simple, legible RANGE loop
	// 	for _, child := range n.children {
	// 		if child.letter == c {
	// 			return true
	// 		}
	// 	}

	// RANGE prevents inlining
	// This is an equivalent loop without RANGE
	for i, length := 0, len(n.children); i < length; i++ {
		if n.children[i].letter == c {
			return true
		}
	}
	return false
}

func (n Node1) child(c rune) *Node1 {
	//  Simple, legible RANGE loop
	// for i := range n.children {
	// 	child := &n.children[i]
	// 	if child.letter == c {
	// 		return child
	// 	}
	// }
	// return nil

	// RANGE prevents inlining
	// This is an equivalent loop without RANGE
	for i, length := 0, len(n.children); i < length; i++ {
		child := &n.children[i]
		if child.letter == c {
			return child
		}
	}
	return nil
}

func (n *Node1) insert(c rune) *Node1 {
	if !n.has(c) {
		n.children = append(n.children, Node1{
			letter: c,
		})
	}
	return n.child(c)
}

func (n *Node1) String() string {
	var buf bytes.Buffer
	var rec func(accu string, nn *Node1)
	rec = func(accu string, nn *Node1) {
		if nn.letter == '.' {
			fmt.Fprintf(&buf, "%s.\n", accu)
			return
		}
		accuNext := accu + string(nn.letter)
		for i := range nn.children {
			child := &nn.children[i]
			rec(accuNext, child)
		}
	}
	// Note that we ignore the first, empty Node letter
	for i := range n.children {
		child := &n.children[i]
		rec("", child)
	}
	return buf.String()
}

func (n Node1) hasWord(word string) bool {
	// if len(word) == 0 {
	// 	return n.has('.')
	// }
	// c := rune(word[0])
	// for _, child := range n.children {
	// 	if child.letter == c {
	// 		return child.hasWord(word[1:])
	// 	}
	// }
	// return false

	// Non-recursive, inlinable:
	for {
		if len(word) == 0 {
			return n.has('.')
		}
		c := rune(word[0])
		child := n.child(c)
		if child == nil {
			return false
		}

		n = *child
		word = word[1:]
	}
}

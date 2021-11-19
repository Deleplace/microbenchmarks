package trie

import (
	"bytes"
	"fmt"
	"log"
)

// Node5 is another slice-based Trie type.
//
// All the transition letters are implicit: they are slice indices between
// 0 and 26.
// Access a subsequent letter is direct.
// A final point '.' is represented by the index 26.
type Node5 struct {
	valid bool
	// children always has size 0 or size 26+1
	children []Node5
}

func kfor(c rune) int {
	if c == '.' {
		return 26
	}
	if c < 'a' || c > 'z' {
		panic("unexpected rune")
	}
	// something in [0..25]
	return int(c - 'a')
}

func fromk(k int) rune {
	if k == 26 {
		return '.'
	}
	if k < 0 || k > 25 {
		log.Fatalf("unexpected k %d", k)
	}
	return 'a' + rune(k)
}

func (n Node5) has(c rune) bool {
	if len(n.children) == 0 {
		return false
	}
	k := kfor(c)
	return n.children[k].valid
}

func (n Node5) child(c rune) *Node5 {
	if len(n.children) == 0 {
		return nil
	}
	k := kfor(c)
	return &n.children[k]
}

func (n *Node5) insert(c rune) *Node5 {
	if len(n.children) == 0 {
		n.children = make([]Node5, 27)
	}
	k := kfor(c)
	nn := &n.children[k]
	nn.valid = true
	return nn
}

func (n *Node5) String() string {
	var buf bytes.Buffer
	var rec func(accu string, nn *Node5)
	rec = func(accu string, nn *Node5) {
		for k := range nn.children {
			child := &nn.children[k]
			if !child.valid {
				continue
			}
			c := fromk(k)
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

func (n Node5) hasWord(word string) bool {
	// if len(word) == 0 {
	// 	return n.has('.')
	// }
	// c := rune(word[0])
	// child := n.child(c)
	// if child == nil {
	// 	return false
	// }
	// return child.hasWord(word[1:])

	// Non-recursive, but not inlinable:
	// function too complex: cost 128 exceeds budget 80
	// This still saves many function calls, though.
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

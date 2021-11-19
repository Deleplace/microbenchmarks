package trie

import (
	"bytes"
	"fmt"
	"log"
)

// Node6 is another slice-based Trie type.
//
// All the transition letters are implicit: they are slice indices between
// 0 and 25.
// Access a subsequent letter is direct.
// Each Node6 may be an endOfWord or not, and have children.
type Node6 struct {
	valid bool
	// endOfWord acts like a '.'
	// the current node may still have children sharing its prefix, though
	endOfWord bool
	// children always has size 0 or size 26
	children []Node6
}

func qfor(c rune) int {
	if c < 'a' || c > 'z' {
		panic("unexpected rune")
	}
	// something in [0..25]
	return int(c - 'a')
}

func fromq(q int) rune {
	if q < 0 || q > 25 {
		log.Fatalf("unexpected k %d", q)
	}
	return 'a' + rune(q)
}

func (n Node6) has(c rune) bool {
	if len(n.children) == 0 {
		return false
	}
	q := qfor(c)
	return n.children[q].valid
}

func (n Node6) child(c rune) *Node6 {
	if len(n.children) == 0 {
		return nil
	}
	q := qfor(c)
	return &n.children[q]
}

func (n *Node6) insert(c rune) *Node6 {
	if c == '.' {
		n.valid = true
		n.endOfWord = true
		return nil
	}
	if len(n.children) == 0 {
		n.children = make([]Node6, 26)
	}
	q := qfor(c)
	nn := &n.children[q]
	nn.valid = true
	return nn
}

func (n *Node6) String() string {
	var buf bytes.Buffer
	var rec func(accu string, nn *Node6)
	rec = func(accu string, nn *Node6) {
		for q := range nn.children {
			child := &nn.children[q]
			if !child.valid {
				continue
			}
			c := fromq(q)
			accuNext := accu + string(c)
			if child.endOfWord {
				fmt.Fprintf(&buf, "%s.\n", accuNext)
			}
			rec(accuNext, child)
		}
	}
	rec("", n)
	return buf.String()
}

func (n Node6) hasWord(word string) bool {
	// if len(word) == 0 {
	// 	return n.endOfWord
	// }
	// c := rune(word[0])
	// child := n.child(c)
	// if child == nil {
	// 	return false
	// }
	// return child.hasWord(word[1:])

	// Non-recursive, inlinable:
	for {
		if len(word) == 0 {
			return n.endOfWord
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

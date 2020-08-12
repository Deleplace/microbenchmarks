package bitset

type Bitset interface {
	GetBit(i int) bool
	SetBit(i int, value bool)
	Len() int
}

package bitset

// Bitset is the minimum set of features common to my Bitset implementations.
type Bitset interface {
	GetBit(i int) bool
	SetBit(i int, value bool)

	// Len is the length of the bitset, in bits.
	// Its semantics may vary slightly depending on the implementation.
	Len() int
}

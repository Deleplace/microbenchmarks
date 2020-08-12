package bitset

type BitsetBoolDyn []bool

func (b *BitsetBoolDyn) GetBit(i int) bool {
	if i >= len(*b) {
		return false
	}
	return (*b)[i]
}

func (b *BitsetBoolDyn) SetBit(i int, value bool) {
	if i >= len(*b) {
		b.grow(1 + i)
	}
	(*b)[i] = value
}

func (b *BitsetBoolDyn) Len() int {
	return len(*b)
}

func (b *BitsetBoolDyn) grow(size int) {
	b2 := make(BitsetBoolDyn, size)
	copy(b2, *b)
	*b = b2
}

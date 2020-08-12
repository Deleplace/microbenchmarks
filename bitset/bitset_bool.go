package bitset

type BitsetBool []bool

func (b BitsetBool) GetBit(i int) bool {
	return b[i]
}

func (b BitsetBool) SetBit(i int, value bool) {
	b[i] = value
}

func (b BitsetBool) Len() int {
	return len(b)
}

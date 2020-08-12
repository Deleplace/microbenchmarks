package bitset

type BitsetUint64 []uint64

func NewUint64(n int) BitsetUint64 {
	return make(BitsetUint64, (n+63)/64)
}

func (b BitsetUint64) GetBit(index int) bool {
	pos := index / 64
	j := index % 64
	return (b[pos] & (uint64(1) << j)) != 0
}

func (b BitsetUint64) SetBit(index int, value bool) {
	pos := index / 64
	j := index % 64
	if value {
		b[pos] |= (uint64(1) << j)
	} else {
		b[pos] ^= (uint64(1) << j)
	}
}

func (b BitsetUint64) Len() int {
	return 64 * len(b)
}

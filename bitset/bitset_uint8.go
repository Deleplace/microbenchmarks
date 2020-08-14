package bitset

type BitsetUint8 []uint8

func NewUint8(n int) BitsetUint8 {
	return make(BitsetUint8, (n+7)/8)
}

func (b BitsetUint8) GetBit(index int) bool {
	pos := index / 8
	j := uint(index % 8)
	return (b[pos] & (uint8(1) << j)) != 0
}

func (b BitsetUint8) SetBit(index int, value bool) {
	pos := index / 8
	j := uint(index % 8)
	if value {
		b[pos] |= (uint8(1) << j)
	} else {
		b[pos] &= ^(uint8(1) << j)
	}
}

func (b BitsetUint8) Len() int {
	return 8 * len(b)
}

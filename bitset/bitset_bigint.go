package bitset

import "math/big"

type BitsetBigInt struct {
	*big.Int
}

func NewBigInt() BitsetBigInt {
	return BitsetBigInt{
		Int: new(big.Int),
	}
}

func (b BitsetBigInt) GetBit(i int) bool {
	return b.Int.Bit(i) == 1
}

func (b BitsetBigInt) SetBit(i int, value bool) {
	if value {
		b.Int.SetBit(b.Int, i, 1)
	} else {
		b.Int.SetBit(b.Int, i, 0)
	}
}

func (b BitsetBigInt) Len() int {
	return b.Int.BitLen()
}

package bitset

import wb "github.com/willf/bitset"

type BitsetWillf struct {
	wb.BitSet
}

func (b BitsetWillf) GetBit(i int) bool {
	return b.BitSet.Test(uint(i))
}

func (b BitsetWillf) SetBit(i int, value bool) {
	b.BitSet.SetTo(uint(i), value)
}

func (b BitsetWillf) Len() int {
	return int(b.BitSet.Len())
}

package bitset

type BitsetMap map[int]bool

func (b BitsetMap) GetBit(i int) bool {
	return b[i]
}

func (b BitsetMap) SetBit(i int, value bool) {
	b[i] = value
}

func (b BitsetMap) Len() int {
	max := 0
	for index := range b {
		if index > max {
			max = index
		}
	}
	return max
}

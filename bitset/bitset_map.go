package bitset

type BitsetMap map[int]bool

func (b BitsetMap) GetBit(i int) bool {
	return b[i]
}

func (b BitsetMap) SetBit(i int, value bool) {
	if value {
		b[i] = value
	} else {
		// deleting a value may free up space,
		// while setting it to false will probably not.
		delete(b, i)
	}
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

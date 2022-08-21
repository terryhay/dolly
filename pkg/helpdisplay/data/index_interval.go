package data

// IndexInterval contains begin and end indexes
type IndexInterval IntPair

func MakeIndexInterval(begin, end int) IndexInterval {
	return IndexInterval{
		first:  begin,
		second: end,
	}
}

// GetBeginIndex - begin index getter
func (i IndexInterval) GetBeginIndex() int {
	return i.first
}

// GetEndIndex - end index getter
func (i IndexInterval) GetEndIndex() int {
	return i.second
}

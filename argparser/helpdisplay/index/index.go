package index

const Null Index = 0

// Index - position in a text
type Index uint64

// MakeIndex constructs an index object in a stack
func MakeIndex(value int) Index {
	if value < 0 {
		return Null
	}
	return Index(value)
}

// ToInt converts Index to int value
func (i Index) ToInt() int {
	return int(i)
}

// Append adds an int value safely
func Append(i Index, value int) Index {
	return MakeIndex(i.ToInt() + value)
}

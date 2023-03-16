package index

// Zero - default Index value
const Zero Index = 0

// Index - position in a text
type Index uint16

// MakeIndex constructs an index object in a stack
func MakeIndex(v int) Index {
	if v < 0 {
		return Zero
	}

	return Index(v)
}

// Int converts Index to int valueInit
func (i Index) Int() int {
	return int(i)
}

// Inc increments index
func Inc(i Index) Index {
	return i + 1
}

// Append adds an int valueInit safely
func Append(i Index, v int) Index {
	if v < 0 {
		i32 := i.Int() + v
		if i32 < 0 {
			return Zero
		}

		return Index(i32)
	}

	return i + Index(v)
}

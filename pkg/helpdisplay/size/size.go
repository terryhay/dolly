package size

// Width - horizontal size
type Width uint32

// ToInt converts Width to int
func (w Width) ToInt() int {
	return int(w)
}

// Height - vertical size
type Height uint32

// ToInt converts Height to int
func (h Height) ToInt() int {
	return int(h)
}

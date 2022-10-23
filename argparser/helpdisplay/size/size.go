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

// AddInt adds int value safely
func (h Height) AddInt(v int) Height {
	i := h.ToInt() + v
	if i < 0 {
		return 0
	}
	return Height(i)
}

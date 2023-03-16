package size

import (
	"github.com/terryhay/dolly/tools/index"
	"golang.org/x/exp/constraints"
)

// Width - horizontal size
type Width uint

const (
	// WidthZero - default size
	WidthZero Width = 0

	// WidthSpace - space between base columns size
	WidthSpace Width = 1

	// WidthColumn - base column size (fib3)
	WidthColumn Width = 3

	// WidthTab - tab size
	WidthTab = WidthColumn + WidthSpace

	// WidthDescriptionColumnShift - min rune len of title for pasting break line
	WidthDescriptionColumnShift = 2*WidthColumn + 2*WidthSpace
)

// MakeWidth constructs Width object
func MakeWidth[T constraints.Integer](v T) Width {
	if v < 0 {
		return WidthZero
	}

	return Width(v)
}

// Int converts Width to int
func (w Width) Int() int {
	return int(w)
}

// Dif returns positive difference between to same T values
func Dif[T Width | Height](l, r T) T {
	if l < r {
		return 0
	}

	return l - r
}

// Height - vertical size
type Height uint

// HeightZero - default value of Height
const HeightZero Height = 0

// MakeHeight constructs Height object
func MakeHeight[T constraints.Integer](v T) Height {
	if v < 0 {
		return HeightZero
	}

	return Height(v)
}

// Int converts Height to int
func (h Height) Int() int {
	return int(h)
}

// Index converts Height to Index
func (h Height) Index() index.Index {
	return index.MakeIndex(h.Int())
}

// AppendHeight adds int value safely
func AppendHeight(h Height, v int) Height {
	i := h.Int() + v
	if i < 0 {
		return 0
	}

	return Height(i)
}

package paragraph_model

import (
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"math"
)

// badness contains page for determinate value of badness for a row
type badness struct {
	rowLenOptimum            size.Width
	aboveRowBadnessByOptimum int

	rowLen       size.Width
	badByOptimum int
}

// makeBadness constructs badness object in a stack
func makeBadness(rowLenOptimum size.Width, aboveRowBadnessByOptimum int, rowLen size.Width) badness {
	return badness{
		rowLenOptimum: rowLenOptimum,

		aboveRowBadnessByOptimum: aboveRowBadnessByOptimum,
		rowLen:                   rowLen,
		badByOptimum:             rowLen.ToInt() - rowLenOptimum.ToInt(),
	}
}

// dropBadnessByOptimum sets badByOptimum field value as default
func (b *badness) dropBadnessByOptimum() {
	_ = b
	b.badByOptimum = math.MinInt32
}

// isBadnessByOptimumDropped returns if badByOptimum field value is default
func (b *badness) isBadnessByOptimumDropped() bool {
	_ = b
	return b.badByOptimum == math.MinInt32
}

// update calculates badness by optimum by new a rowLen
func (b *badness) update(rowLen size.Width) {
	_ = b
	b.rowLen = rowLen
	b.badByOptimum = rowLen.ToInt() - b.rowLenOptimum.ToInt()
}

// worse compares two badness objects and returns if the receiver badness is worse
func (b *badness) worse(b2 badness) bool {
	_ = b

	if b.aboveRowBadnessByOptimum == 0 {
		return lessOrEqualByAbs(b2.badByOptimum, b.badByOptimum)
	}

	if b.aboveRowBadnessByOptimum < 0 {
		if b.badByOptimum < 0 {
			if b.badByOptimum < b.aboveRowBadnessByOptimum {
				if b2.badByOptimum > 0 {
					return lessByAbs(b2.badByOptimum, b.badByOptimum)
				}
				return true
			}
			return false
		}
		return lessByAbs(b2.badByOptimum, b.badByOptimum)
	}

	if b.badByOptimum <= 0 {
		if equalByAbs(b2.badByOptimum, b.badByOptimum) {
			return true
		}
		if b2.badByOptimum > 0 {
			return b.aboveRowBadnessByOptimum-b.badByOptimum >= absInt(b.aboveRowBadnessByOptimum-b2.badByOptimum)
		}
		return lessByAbs(b2.badByOptimum, b.badByOptimum)
	}
	if b2.badByOptimum > b.aboveRowBadnessByOptimum {
		return false
	}
	return (b.aboveRowBadnessByOptimum - b.badByOptimum) < (b.aboveRowBadnessByOptimum - b2.badByOptimum)
}

// getBadByOptimum returns badByOptimum field value
func (b *badness) getBadByOptimum() int {
	return b.badByOptimum
}

func lessByAbs(v1, v2 int) bool {
	return absInt(v1) < absInt(v2)
}

func equalByAbs(v1, v2 int) bool {
	return absInt(v1) == absInt(v2)
}

func lessOrEqualByAbs(v1, v2 int) bool {
	return absInt(v1) <= absInt(v2)
}

// absInt returns absolutely value of v
func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

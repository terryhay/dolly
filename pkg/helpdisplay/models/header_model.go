package models

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"

	"github.com/nsf/termbox-go"
)

// HeaderModel contains page header source data and get it as cell slice for view
type HeaderModel struct {
	headerCells []termbox.Cell
	outputCells []termbox.Cell

	usingRowLeLimit row_len_limiter.RowLenLimit
	shift           int
}

// MakeHeaderModel constructs HeaderModel in a stack
func MakeHeaderModel(header string, rowLenLimit row_len_limiter.RowLenLimit) HeaderModel {
	hm := HeaderModel{
		headerCells: textToCells(header),
	}
	hm.Update(rowLenLimit)

	return hm
}

// Update processes changed terminal row size data and updates getting header cell slice
func (hm *HeaderModel) Update(rowLenLimit row_len_limiter.RowLenLimit) {
	_ = hm

	if hm.usingRowLeLimit == rowLenLimit {
		return
	}

	hm.usingRowLeLimit = rowLenLimit
	hm.updateShift()
}

// GetHeaderRow returns header cell slice and actionSequence to display text in the page center
func (hm *HeaderModel) GetHeaderRow() (int, []termbox.Cell) {
	return hm.shift, hm.outputCells
}

func (hm *HeaderModel) updateShift() {
	_ = hm

	hm.outputCells = hm.headerCells
	hm.shift = (hm.usingRowLeLimit.Optimum().ToInt() - len(hm.outputCells)) / 2
	if hm.shift < 0 {
		// try to use max terminal size
		hm.shift = (hm.usingRowLeLimit.Max().ToInt() - len(hm.outputCells)) / 2
		if hm.shift < 0 {
			// cut the output header
			hm.outputCells = make([]termbox.Cell, hm.usingRowLeLimit.Max())
			copy(hm.outputCells, hm.headerCells[:hm.usingRowLeLimit.Max().ToInt()])

			counter := 0
			for i := hm.usingRowLeLimit.Max().ToInt() - 1; i >= 0 && counter < 3; i-- {
				hm.outputCells[i].Ch = runes.RuneDot
				counter++
			}
		}

		hm.shift = 0
	}
}

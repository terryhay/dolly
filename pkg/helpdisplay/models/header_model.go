package models

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"

	"github.com/nsf/termbox-go"
)

// HeaderModel contains page header source data and get it as cell slice for view
type HeaderModel struct {
	headerCells []termbox.Cell
	outputCells []termbox.Cell

	usingRowLeLimit rll.RowLenLimit
	shift           size.Width
}

// NewHeaderModel constructs HeaderModel
func NewHeaderModel(header string, size TerminalSize) *HeaderModel {
	hm := &HeaderModel{
		headerCells: textToCells(header),
	}
	hm.Update(size)

	return hm
}

// Update processes changed terminal row size data and updates getting header cell slice
func (hm *HeaderModel) Update(size TerminalSize) {
	if hm == nil {
		return
	}

	if hm.usingRowLeLimit == size.GetWidth() {
		return
	}

	hm.usingRowLeLimit = size.GetWidth()
	hm.updateShift()
}

// GetHeaderRow returns header cell slice and actionSequence to display text in the page center
func (hm *HeaderModel) GetHeaderRow() row.Row {
	if hm == nil {
		return row.Row{}
	}
	return row.MakeRow(hm.shift, hm.outputCells)
}

func (hm *HeaderModel) updateShift() {
	_ = hm

	hm.outputCells = hm.headerCells
	newShift := (hm.usingRowLeLimit.Optimum().ToInt() - len(hm.outputCells)) / 2
	if newShift < 0 {
		// try to use max terminal size
		newShift = (hm.usingRowLeLimit.Max().ToInt() - len(hm.outputCells)) / 2
		if newShift < 0 {
			// cut the output header
			hm.outputCells = make([]termbox.Cell, hm.usingRowLeLimit.Max())
			copy(hm.outputCells, hm.headerCells[:hm.usingRowLeLimit.Max().ToInt()])

			counter := 0
			for i := hm.usingRowLeLimit.Max().ToInt() - 1; i >= 0 && counter < 3; i-- {
				hm.outputCells[i].Ch = runes.RuneDot
				counter++
			}
		}

		newShift = 0
	}

	hm.shift = size.Width(newShift)
}

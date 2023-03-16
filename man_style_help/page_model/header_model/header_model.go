package header_model

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	s2cells "github.com/terryhay/dolly/man_style_help/string_to_cells"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

// HeaderModel contains page header row
type HeaderModel struct {
	cellsSrc []termbox.Cell
	cellsOut []termbox.Cell

	usingRowLenLimit rll.RowLenLimit
	shift            size.Width
}

// NewHeaderModel constructs HeaderModel
func NewHeaderModel(appName coty.NameApp, sizeTerminal ts.TerminalSize) *HeaderModel {
	hm := &HeaderModel{
		cellsSrc: s2cells.StringToCells(appName.String()),
	}

	hm.Update(sizeTerminal)
	return hm
}

// Update processes changed terminal dynamic_row size page and updates getting header cell slice
func (hm *HeaderModel) Update(termSize ts.TerminalSize) {
	if hm == nil {
		return
	}

	if hm.usingRowLenLimit == termSize.GetWidthLimit() {
		return
	}

	hm.usingRowLenLimit = termSize.GetWidthLimit()
	hm.updateShift()
}

// GetViewRow returns header cell slice and actionSequence to display text in the page center
func (hm *HeaderModel) GetViewRow() row.Row {
	if hm == nil {
		return row.Row{}
	}
	return row.MakeRow(hm.shift, hm.cellsOut)
}

func (hm *HeaderModel) updateShift() {
	_ = hm

	hm.cellsOut = hm.cellsSrc
	newShift := (hm.usingRowLenLimit.Optimum().Int() - len(hm.cellsOut)) / 2
	if newShift < 0 {
		// try to use max terminal size
		newShift = (hm.usingRowLenLimit.Max().Int() - len(hm.cellsOut)) / 2
		if newShift < 0 {
			// cut the output header
			hm.cellsOut = make([]termbox.Cell, hm.usingRowLenLimit.Max())
			copy(hm.cellsOut, hm.cellsSrc[:hm.usingRowLenLimit.Max().Int()])

			counter := 0
			for i := hm.usingRowLenLimit.Max().Int() - 1; i >= 0 && counter < 3; i-- {
				hm.cellsOut[i].Ch = runes.RuneDot
				counter++
			}
		}

		newShift = 0
	}

	hm.shift = size.MakeWidth(newShift)
}

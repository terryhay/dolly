package header_model

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/man_style_help/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/man_style_help/size"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// HeaderModel contains page header row
type HeaderModel struct {
	paragraph   page.Paragraph
	outputCells []termbox.Cell

	usingRowLenLimit rll.RowLenLimit
	shift            size.Width
}

// NewHeaderModel constructs HeaderModel
func NewHeaderModel(paragraph page.Paragraph, size ts.TerminalSize) (*HeaderModel, *dollyerr.Error) {
	hm := &HeaderModel{
		paragraph: paragraph,
	}
	err := hm.Update(size)
	if err != nil {
		return nil, dollyerr.Append(err, fmt.Errorf("NewHeaderModel error"))
	}

	return hm, nil
}

// Update processes changed terminal dynamic_row size page and updates getting header cell slice
func (hm *HeaderModel) Update(termSize ts.TerminalSize) *dollyerr.Error {
	if hm == nil {
		return nil
	}

	if err := termSize.IsValid(); err != nil {
		return dollyerr.Append(err, fmt.Errorf("HeaderModel.Update"))
	}

	if hm.usingRowLenLimit == termSize.GetWidthLimit() {
		return nil
	}

	hm.usingRowLenLimit = termSize.GetWidthLimit()
	hm.updateShift()

	return nil
}

// GetViewRow returns header cell slice and actionSequence to display text in the page center
func (hm *HeaderModel) GetViewRow() row.Row {
	if hm == nil {
		return row.Row{}
	}
	return row.MakeRow(hm.shift, hm.outputCells)
}

func (hm *HeaderModel) updateShift() {
	_ = hm

	hm.outputCells = hm.paragraph.GetCells()
	newShift := (hm.usingRowLenLimit.Optimum().ToInt() - len(hm.outputCells)) / 2
	if newShift < 0 {
		// try to use max terminal size
		newShift = (hm.usingRowLenLimit.Max().ToInt() - len(hm.outputCells)) / 2
		if newShift < 0 {
			// cut the output header
			hm.outputCells = make([]termbox.Cell, hm.usingRowLenLimit.Max())
			copy(hm.outputCells, hm.paragraph.GetCells()[:hm.usingRowLenLimit.Max().ToInt()])

			counter := 0
			for i := hm.usingRowLenLimit.Max().ToInt() - 1; i >= 0 && counter < 3; i-- {
				hm.outputCells[i].Ch = runes.RuneDot
				counter++
			}
		}

		newShift = 0
	}

	hm.shift = size.Width(newShift)
}

package models

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
)

type FooterModel struct {
	cells []termbox.Cell

	usingSize TerminalSize
	shift     size.Width
}

// NewFooterModel constructs FooterModel
func NewFooterModel(size TerminalSize) *FooterModel {
	return &FooterModel{
		cells: []termbox.Cell{
			{
				Ch: runes.RuneColon,
			},
		},
		usingSize: size,
	}
}

// Update processes changed terminal row size data and updates getting header cell slice
func (fm *FooterModel) Update(size TerminalSize) {
	if fm == nil {
		return
	}
	fm.usingSize.Width = size.GetWidth()
}

// GetFooterRow returns footer cell slice and a shift
func (fm *FooterModel) GetFooterRow() row.Row {
	if fm == nil {
		return row.Row{}
	}
	return row.MakeRow(fm.shift, fm.cells)
}

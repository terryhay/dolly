package footer_model

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	"github.com/terryhay/dolly/argparser/helpdisplay/runes"
)

// FooterModel contains page footer row
type FooterModel struct {
	rowFooter row.Row
}

// NewFooterModel constructs FooterModel
func NewFooterModel() *FooterModel {
	return &FooterModel{
		rowFooter: row.MakeRow(
			0,
			[]termbox.Cell{
				{
					Ch: runes.RuneColon,
				},
			}),
	}
}

// GetFooterRow returns footer cell slice and a shift
func (fm *FooterModel) GetFooterRow() row.Row {
	if fm == nil {
		return row.Row{}
	}
	return fm.rowFooter
}

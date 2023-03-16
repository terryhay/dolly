package row

import (
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/tools/size"
)

// Row implements an immutable dynamic_row object for rendering in a terminal
type Row struct {
	// shiftIndex contains a left shift characters amount
	shiftIndex size.Width

	// cells contains a dynamic_row content (which is displaying after shifts from shiftIndex)
	cells []termbox.Cell
}

// MakeRow constructs Row object in a stack
func MakeRow(shiftIndex size.Width, cells []termbox.Cell) Row {
	return Row{
		shiftIndex: shiftIndex,
		cells:      cells,
	}
}

// GetShiftIndex gets a shiftIndex field value
func (r Row) GetShiftIndex() size.Width {
	return r.shiftIndex
}

// GetCells gets a cells field value
func (r Row) GetCells() []termbox.Cell {
	return r.cells
}

func (r Row) String() string {
	builder := strings.Builder{}

	for i := size.WidthZero; i < r.GetShiftIndex(); i++ {
		builder.WriteRune(runes.RuneSpace)
	}
	for _, cell := range r.GetCells() {
		builder.WriteRune(cell.Ch)
	}

	return builder.String()
}

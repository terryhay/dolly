package paragraph_model

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"strings"
)

// dynamicRow contains page about a dynamic_row of split source text
type dynamicRow struct {
	paragraphCells []termbox.Cell
	breakFront     *splitter
	breakBack      *splitter
	Badness        badness
}

// makeDynamicRow constructs a dynamic dynamicRow object in a stack
func makeDynamicRow(
	rowLenOptimum size.Width,
	source []termbox.Cell,
	aboveLineByOptimum int,
	breakFront *splitter,
	breakBack *splitter,
) dynamicRow {
	r := dynamicRow{
		paragraphCells: source,
		breakFront:     breakFront,
		breakBack:      breakBack,
	}
	r.Badness = makeBadness(rowLenOptimum, aboveLineByOptimum, r.len())
	return r
}

// getBreakFront returns breakFront field value
func (r *dynamicRow) getBreakFront() *splitter {
	return r.breakFront
}

// getBreakBack returns breakBack field value
func (r *dynamicRow) getBreakBack() *splitter {
	return r.breakBack
}

// setBreakBack updates breakBack field
func (r *dynamicRow) setBreakBack(breakBack *splitter) {
	r.breakBack = breakBack
}

// toRow converts dynamicRow to simple dynamicRow object
func (r *dynamicRow) toRow(shift size.Width) row.Row {
	_ = r
	return row.MakeRow(shift, r.viewCells())
}

// String implements the string interface
func (r *dynamicRow) String() string {
	_ = r

	builder := strings.Builder{}
	cells := r.viewCells()
	for i := range cells {
		builder.WriteString(string(cells[i].Ch))
	}

	return builder.String()
}

// len returns a rune len of the line
func (r *dynamicRow) len() size.Width {
	_ = r

	l := len(r.viewCells())
	res := size.Width(0)
	if l > 0 {
		res = size.Width(l)
	}
	return res
}

// viewCells slices a dynamicRow from a paragraphCells slice and return it
func (r *dynamicRow) viewCells() []termbox.Cell {
	_ = r

	indexFront := 0
	if r.breakFront != nil {
		indexFront = r.breakFront.indexEnd().ToInt()
	}

	indexBack := len(r.paragraphCells)
	if r.breakBack != nil {
		indexBack = r.breakBack.indexBegin().ToInt()
	}

	var res []termbox.Cell
	if indexFront < indexBack {
		res = r.paragraphCells[indexFront:indexBack]
	}
	return res
}

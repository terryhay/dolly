package row_model

import (
	"github.com/nsf/termbox-go"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	s2cells "github.com/terryhay/dolly/man_style_help/string_to_cells"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

// averageWordLen to be precise the average english word length is 5.2
// we are using 6 for avoid reallocation
const averageWordLen = 6

// RowModel - class which is getting paragraph text parts for render in a terminal
type RowModel struct {
	// cellsSrc - paragraph text as cell slice
	cellsSrc []termbox.Cell

	// margin - paragraph content margin
	marginLeft size.Width

	// spaceCellIndexes - sequence which contain indexes of spaces in a cells slice
	spaceCellIndexes []index.Index

	// rowsSplit - interval indexes of cellsSrc for display physical rows.
	// If len(rowsSplit) == 0, ONE row will be displayed, otherwise 1 + len(rowsSplit) rows will be displayed
	rowsSplit []*splitter

	// anchorRowIndex - using for searching begin text dynamic_row for display after a resize of terminal window
	anchorRowIndex index.Index

	// widthTerminal contains using terminal width
	widthTerminal size.Width
}

// NewRowModel constructs a new RowModel object
func NewRowModel(row hp.Row, rowLenLimit rll.RowLenLimit) *RowModel {
	cellsSrc := s2cells.StringToCells(row.GetTextStyled())

	prm := &RowModel{
		cellsSrc:      cellsSrc,
		marginLeft:    row.GetMarginLeft(),
		widthTerminal: rowLenLimit.Max(),

		spaceCellIndexes: func() []index.Index {
			spaceCellIndexes := make([]index.Index, 0, len(cellsSrc)/averageWordLen)

			for i := index.Zero; i.Int() < len(cellsSrc); i++ {
				if cellsSrc[i].Ch == runes.RuneSpace {
					spaceCellIndexes = append(spaceCellIndexes, i)
				}
			}
			return spaceCellIndexes
		}(),

		rowsSplit: func() []*splitter {
			if rowLenLimit.Max() == 0 {
				return nil
			}
			return make([]*splitter, 0, len(cellsSrc)/rowLenLimit.Max().Int())
		}(),
	}
	prm.Update(rowLenLimit)

	return prm
}

// GetAnchorRowIndex returns anchorRowIndex field value
func (rm *RowModel) GetAnchorRowIndex() index.Index {
	if rm == nil {
		return index.Zero
	}
	return rm.anchorRowIndex
}

// SetAnchorRowIndex updates anchorRowIndex field value
func (rm *RowModel) SetAnchorRowIndex(anchorRowIndex index.Index) {
	if rm == nil {
		return
	}
	rm.anchorRowIndex = anchorRowIndex
}

// GetRowCount returns dynamic_row count for rendering in a terminal
func (rm *RowModel) GetRowCount() size.Height {
	if rm == nil {
		return size.HeightZero
	}

	if len(rm.cellsSrc) != 0 {
		return size.MakeHeight(1 + len(rm.rowsSplit))
	}
	return 1
}

// GetTabInCells returns using tabInCells size in runes for current terminal width
func (rm *RowModel) GetTabInCells() size.Width {
	if rm == nil {
		return size.WidthZero
	}

	if rm.widthTerminal <= rll.TerminalMinWidth {
		return 0
	}
	return rm.marginLeft
}

// ShiftAnchorRow does a try to actionSequence anchor dynamic_row and returns if the try is success
func (rm *RowModel) ShiftAnchorRow(shift int) bool {
	if rm == nil {
		return false
	}

	newAnchorRowIndex := rm.anchorRowIndex.Int() + shift
	if newAnchorRowIndex >= rm.GetRowCount().Int() || newAnchorRowIndex < 0 {
		rm.anchorRowIndex = 0
		return false
	}

	rm.anchorRowIndex = index.MakeIndex(newAnchorRowIndex)
	return true
}

// SetBackRowAsAnchor sets index of last dynamic_row as anchor dynamic_row
func (rm *RowModel) SetBackRowAsAnchor() {
	if rm == nil {
		return
	}
	rm.anchorRowIndex = index.MakeIndex(rm.GetRowCount().Int() - 1)
}

// GetRow returns a dynamic_row for rendering in a terminal by index
func (rm *RowModel) GetRow(indexRow index.Index) row.Row {
	if rm == nil {
		return row.Row{}
	}

	if indexRow.Int() >= rm.GetRowCount().Int() {
		return row.MakeRow(0, nil)
	}

	var beginBreakLineIndex index.Index
	endBreakRowIndex := index.MakeIndex(len(rm.cellsSrc))

	if indexRow > 0 {
		beginBreakLineIndex = rm.rowsSplit[indexRow-1].indexEnd()
	}
	if indexRow < index.MakeIndex(len(rm.rowsSplit)) {
		endBreakRowIndex = rm.rowsSplit[indexRow].indexBegin()
	}

	return row.MakeRow(rm.GetTabInCells(), rm.cellsSrc[beginBreakLineIndex:endBreakRowIndex])
}

// Update applies dynamic_row len limit for rebuilding getting display dynamic_row window
func (rm *RowModel) Update(rowLenLimit rll.RowLenLimit) size.Height {
	if rm == nil {
		return size.HeightZero
	}

	rm.widthTerminal = rowLenLimit.Max()

	var anchorSpaceIndex index.Index
	if rm.anchorRowIndex > 0 {
		anchorSpaceIndex = rm.rowsSplit[rm.anchorRowIndex-1].indexBegin()
		rm.anchorRowIndex = 0
	}

	rowLenLimit = rowLenLimit.ApplyTabShift(rm.GetTabInCells())

	if len(rm.rowsSplit) > 0 {
		rm.rowsSplit = rm.rowsSplit[:0]
	}

	if len(rm.cellsSrc) < rowLenLimit.Max().Int() {
		rm.anchorRowIndex = 0
		return rm.GetRowCount()
	}

	if len(rm.spaceCellIndexes) == 0 {
		for i := index.MakeIndex(rowLenLimit.Max().Int()); i.Int() < len(rm.cellsSrc); i = index.Append(i, rowLenLimit.Max().Int()) {
			if anchorSpaceIndex < i {
				rm.anchorRowIndex = index.MakeIndex(len(rm.rowsSplit))
			}
			rm.rowsSplit = append(rm.rowsSplit, newSplitter(i, i))
		}

		return rm.GetRowCount()
	}

	var (
		splitNew = newSplitter(rm.spaceCellIndexes[0], rm.spaceCellIndexes[0]+1)

		rowOptimum        = makeDynamicRow(rowLenLimit.Optimum(), rm.cellsSrc, 0, nil, splitNew)
		optimumSpaceIndex = 0
		rowNew            = rowOptimum
	)
	rowOptimum.Badness.dropBadnessByOptimum()

	for i := 1; i < len(rm.spaceCellIndexes); i++ {
		splitNew = newSplitter(rm.spaceCellIndexes[i], rm.spaceCellIndexes[i]+1)

		rowNew.setBreakBack(splitNew)
		rowNew.Badness.update(rowNew.len())

		if rowNew.len() < rowLenLimit.Min() {
			rowOptimum = rowNew
			optimumSpaceIndex = i

			continue
		}

		if rowNew.len() > rowLenLimit.Max() {
			if anchorSpaceIndex.Int() > i {
				rm.anchorRowIndex = index.MakeIndex(len(rm.rowsSplit)) + 1
			}
			rm.rowsSplit = append(rm.rowsSplit,
				newSplitter(rowOptimum.getBreakBack().indexBegin(), rowOptimum.getBreakBack().indexEnd()))

			i = optimumSpaceIndex
			var breakFront *splitter
			if len(rm.rowsSplit) > 0 {
				breakFront = rowOptimum.getBreakBack()
			}
			splitNew = breakFront
			rowNew = makeDynamicRow(rowLenLimit.Optimum(), rm.cellsSrc, rowOptimum.Badness.getBadByOptimum(), breakFront, splitNew)
			rowOptimum.Badness.dropBadnessByOptimum()

			continue
		}

		if rowOptimum.Badness.worse(rowNew.Badness) {
			rowOptimum = rowNew
			optimumSpaceIndex = i
			continue
		}
	}

	rowNew.setBreakBack(nil)
	if rowNew.len() > rowLenLimit.Max() {
		rm.rowsSplit = append(rm.rowsSplit, rowOptimum.getBreakBack())

		return rm.GetRowCount()
	}

	if len(rm.rowsSplit) > 0 {
		if anchorSpaceIndex > rm.rowsSplit[len(rm.rowsSplit)-1].indexBegin() {
			rm.anchorRowIndex = index.MakeIndex(len(rm.rowsSplit))
		}
	}

	return rm.GetRowCount()
}

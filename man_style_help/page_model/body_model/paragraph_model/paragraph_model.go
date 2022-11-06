package paragraph_model

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/man_style_help/index"
	"github.com/terryhay/dolly/man_style_help/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/man_style_help/size"
)

// averageWordLen to be precise the average english word length is 5.2
// we are using 6 for avoid reallocation
const averageWordLen = 6

// ParagraphModel - class which is getting paragraph text parts for render in a terminal
type ParagraphModel struct {
	paragraph page.Paragraph

	// spaceCellIndexes - sequence which contain indexes of spaces in a cells slice
	spaceCellIndexes []index.Index

	// paragraphSplits contains page for breaking cells slice to some terminal render rows
	paragraphSplits []*splitter

	// anchorRowIndex - using for searching begin text dynamic_row for display after a resize of terminal window
	anchorRowIndex index.Index

	// terminalWidth contains using terminal width
	terminalWidth size.Width
}

// NewParagraphModel constructs a new ParagraphModel object
func NewParagraphModel(rowLenLimit rll.RowLenLimit, paragraph page.Paragraph) *ParagraphModel {
	prm := &ParagraphModel{
		paragraph:     paragraph,
		terminalWidth: rowLenLimit.Max(),

		spaceCellIndexes: func(cells []termbox.Cell) []index.Index {
			spaceCellIndexes := make([]index.Index, 0, len(cells)/averageWordLen)

			for i := index.Null; i.ToInt() < len(cells); i++ {
				if cells[i].Ch == runes.RuneSpace {
					spaceCellIndexes = append(spaceCellIndexes, i)
				}
			}
			return spaceCellIndexes
		}(paragraph.GetCells()),

		paragraphSplits: func(cells []termbox.Cell) []*splitter {
			if rowLenLimit.Max() == 0 {
				return nil
			}
			return make([]*splitter, 0, len(cells)/rowLenLimit.Max().ToInt())
		}(paragraph.GetCells()),
	}
	prm.Update(rowLenLimit)

	return prm
}

// GetAnchorRowIndex returns anchorRowIndex field value
func (prm *ParagraphModel) GetAnchorRowIndex() index.Index {
	if prm == nil {
		return index.Null
	}
	return prm.anchorRowIndex
}

// SetAnchorRowIndex updates anchorRowIndex field value
func (prm *ParagraphModel) SetAnchorRowIndex(anchorRowIndex index.Index) {
	if prm == nil {
		return
	}
	prm.anchorRowIndex = anchorRowIndex
}

// GetRowCount returns dynamic_row count for rendering in a terminal
func (prm *ParagraphModel) GetRowCount() size.Height {
	if prm == nil {
		return size.Height(0)
	}

	if len(prm.paragraph.GetCells()) != 0 {
		return size.Height(1 + len(prm.paragraphSplits))
	}
	return 1
}

// GetTabInCells returns using tabInCells size in runes for current terminal width
func (prm *ParagraphModel) GetTabInCells() size.Width {
	if prm == nil {
		return size.Width(0)
	}

	if prm.terminalWidth <= rll.TerminalMinWidth {
		return 0
	}
	return prm.paragraph.GetTabInCells()
}

// ShiftAnchorRow does a try to actionSequence anchor dynamic_row and returns if the try is success
func (prm *ParagraphModel) ShiftAnchorRow(shift int) bool {
	if prm == nil {
		return false
	}

	newAnchorRowIndex := prm.anchorRowIndex.ToInt() + shift
	if newAnchorRowIndex >= prm.GetRowCount().ToInt() || newAnchorRowIndex < 0 {
		prm.anchorRowIndex = 0
		return false
	}

	prm.anchorRowIndex = index.MakeIndex(newAnchorRowIndex)
	return true
}

// SetBackRowAsAnchor sets index of last dynamic_row as anchor dynamic_row
func (prm *ParagraphModel) SetBackRowAsAnchor() {
	if prm == nil {
		return
	}
	prm.anchorRowIndex = index.MakeIndex(prm.GetRowCount().ToInt() - 1)
}

// GetRow returns a dynamic_row for rendering in a terminal by index
func (prm *ParagraphModel) GetRow(indexRow index.Index) row.Row {
	if prm == nil {
		return row.Row{}
	}

	if indexRow.ToInt() >= prm.GetRowCount().ToInt() {
		return row.MakeRow(0, nil)
	}

	var beginBreakLineIndex index.Index
	endBreakRowIndex := index.MakeIndex(len(prm.paragraph.GetCells()))

	if indexRow > 0 {
		beginBreakLineIndex = prm.paragraphSplits[indexRow-1].indexEnd()
	}
	if indexRow < index.MakeIndex(len(prm.paragraphSplits)) {
		endBreakRowIndex = prm.paragraphSplits[indexRow].indexBegin()
	}

	return row.MakeRow(prm.GetTabInCells(), prm.paragraph.GetCells()[beginBreakLineIndex:endBreakRowIndex])
}

// Update applies dynamic_row len limit for rebuilding getting display dynamic_row window
func (prm *ParagraphModel) Update(rowLenLimit rll.RowLenLimit) size.Height {
	if prm == nil {
		return size.Height(0)
	}

	prm.terminalWidth = rowLenLimit.Max()

	var anchorSpaceIndex index.Index
	if prm.anchorRowIndex > 0 {
		anchorSpaceIndex = prm.paragraphSplits[prm.anchorRowIndex-1].indexBegin()
		prm.anchorRowIndex = 0
	}

	rowLenLimit = rowLenLimit.ApplyTabShift(prm.GetTabInCells())

	if len(prm.paragraphSplits) > 0 {
		prm.paragraphSplits = prm.paragraphSplits[:0]
	}

	if len(prm.paragraph.GetCells()) < rowLenLimit.Max().ToInt() {
		prm.anchorRowIndex = 0
		return prm.GetRowCount()
	}

	if len(prm.spaceCellIndexes) == 0 {
		for i := index.MakeIndex(rowLenLimit.Max().ToInt()); i.ToInt() < len(prm.paragraph.GetCells()); i = index.Append(i, rowLenLimit.Max().ToInt()) {
			if anchorSpaceIndex < i {
				prm.anchorRowIndex = index.MakeIndex(len(prm.paragraphSplits))
			}
			prm.paragraphSplits = append(prm.paragraphSplits, newSplitter(i, i))
		}

		return prm.GetRowCount()
	}

	var (
		splitNew = newSplitter(prm.spaceCellIndexes[0], prm.spaceCellIndexes[0]+1)

		rowOptimum        = makeDynamicRow(rowLenLimit.Optimum(), prm.paragraph.GetCells(), 0, nil, splitNew)
		optimumSpaceIndex = 0
		rowNew            = rowOptimum
	)
	rowOptimum.Badness.dropBadnessByOptimum()

	for i := 1; i < len(prm.spaceCellIndexes); i++ {
		splitNew = newSplitter(prm.spaceCellIndexes[i], prm.spaceCellIndexes[i]+1)

		rowNew.setBreakBack(splitNew)
		rowNew.Badness.update(rowNew.len())

		if rowNew.len() < rowLenLimit.Min() {
			rowOptimum = rowNew
			optimumSpaceIndex = i

			continue
		}

		if rowNew.len() > rowLenLimit.Max() {
			if anchorSpaceIndex.ToInt() > i {
				prm.anchorRowIndex = index.MakeIndex(len(prm.paragraphSplits)) + 1
			}
			prm.paragraphSplits = append(prm.paragraphSplits,
				newSplitter(rowOptimum.getBreakBack().indexBegin(), rowOptimum.getBreakBack().indexEnd()))

			i = optimumSpaceIndex
			var breakFront *splitter
			if len(prm.paragraphSplits) > 0 {
				breakFront = rowOptimum.getBreakBack()
			}
			splitNew = breakFront
			rowNew = makeDynamicRow(rowLenLimit.Optimum(), prm.paragraph.GetCells(), rowOptimum.Badness.getBadByOptimum(), breakFront, splitNew)
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
		prm.paragraphSplits = append(prm.paragraphSplits, rowOptimum.getBreakBack())

		return prm.GetRowCount()
	}

	if len(prm.paragraphSplits) > 0 {
		if anchorSpaceIndex > prm.paragraphSplits[len(prm.paragraphSplits)-1].indexBegin() {
			prm.anchorRowIndex = index.MakeIndex(len(prm.paragraphSplits))
		}
	}

	return prm.GetRowCount()
}

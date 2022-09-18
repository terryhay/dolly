package models

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"

	"github.com/nsf/termbox-go"
)

// averageWordLen to be precise the average english word length is 5.2
// we are using 6 for avoid reallocation
const averageWordLen = 6

// ParagraphModel - class which is getting paragraph text parts for render in a terminal
type ParagraphModel struct {
	tmpSource string

	// cells - solid cell sequence which is created from source paragraph text runes
	cells []termbox.Cell

	// spaceCellIndexes - sequence which contain indexes of spaces in a cells slice
	spaceCellIndexes []int

	// breakLineIndexes contains data for breaking cells slice to some terminal render rows
	breakLineIndexes []data.IndexInterval

	// tab - actionSequence of paragraph rows
	tab size.Width

	// anchorRowIndex - using for searching begin text row for display after a resize of terminal window
	anchorRowIndex size.Height

	// usingTerminalWidth - value of using terminal width
	usingTerminalWidth size.Width
}

// NewParagraphModel constructs a new ParagraphModel object
func NewParagraphModel(rowLenLimit rll.RowLenLimit, source *data.Paragraph) *ParagraphModel {
	cells := textToCells(source.Text)

	prm := &ParagraphModel{
		tmpSource:          source.Text,
		cells:              cells,
		tab:                source.TabCount * rll.TabSize,
		usingTerminalWidth: rowLenLimit.Max(),

		spaceCellIndexes: func(cells []termbox.Cell) []int {
			spaceCellIndexes := make([]int, 0, len(cells)/averageWordLen)

			for index := 0; index < len(cells); index++ {
				if cells[index].Ch == runes.RuneSpace {
					spaceCellIndexes = append(spaceCellIndexes, index)
				}
			}
			return spaceCellIndexes
		}(cells),

		breakLineIndexes: func(cells []termbox.Cell) []data.IndexInterval {
			if rowLenLimit.Max() == 0 {
				return nil
			}
			return make([]data.IndexInterval, 0, len(cells)/rowLenLimit.Max().ToInt())
		}(cells),
	}
	prm.Update(rowLenLimit)

	return prm
}

// GetAnchorRowIndex - anchorRowIndex field getter
func (prm *ParagraphModel) GetAnchorRowIndex() size.Height {
	return prm.anchorRowIndex
}

// GetRowCount returns row count for rendering in a terminal
func (prm *ParagraphModel) GetRowCount() size.Height {
	if len(prm.cells) != 0 {
		return size.Height(1 + len(prm.breakLineIndexes))
	}
	return 1
}

// GetUsingTab returns using tab size in runes for current terminal width
func (prm *ParagraphModel) GetUsingTab() size.Width {
	if prm.usingTerminalWidth <= rll.TerminalMinWidth {
		return 0
	}
	return prm.tab
}

// ShiftAnchorRow does a try to actionSequence anchor row and returns if the try is success
func (prm *ParagraphModel) ShiftAnchorRow(shift int) bool {
	prm.anchorRowIndex = size.Height(prm.anchorRowIndex.ToInt() + shift)
	if prm.anchorRowIndex >= prm.GetRowCount() {
		prm.anchorRowIndex = 0
		return false
	}

	return true
}

// SetBackRowAsAnchor sets index of last row as anchor row
func (prm *ParagraphModel) SetBackRowAsAnchor() {
	prm.anchorRowIndex = size.Height(prm.GetRowCount().ToInt() - 1)
}

// GetRow returns a row for rendering in a terminal by index
func (prm *ParagraphModel) GetRow(index size.Height) row.Row {
	_ = prm

	if index >= prm.GetRowCount() {
		return row.MakeRow(0, nil)
	}

	beginBreakLineIndex := 0
	endBreakRowIndex := len(prm.cells)

	if index > 0 {
		beginBreakLineIndex = prm.breakLineIndexes[index-1].GetEndIndex()
	}
	if index < size.Height(len(prm.breakLineIndexes)) {
		endBreakRowIndex = prm.breakLineIndexes[index].GetBeginIndex()
	}

	return row.MakeRow(prm.GetUsingTab(), prm.cells[beginBreakLineIndex:endBreakRowIndex])
}

// Update applies row len limit for rebuilding getting display row window
func (prm *ParagraphModel) Update(rowLenLimit rll.RowLenLimit) size.Height {
	_ = prm

	prm.usingTerminalWidth = rowLenLimit.Max()

	anchorSpaceIndex := 0
	if prm.anchorRowIndex > 0 {
		anchorSpaceIndex = prm.breakLineIndexes[prm.anchorRowIndex-1].GetBeginIndex()
	}

	rowLenLimit = rowLenLimit.ApplyTabShift(prm.GetUsingTab())

	if len(prm.breakLineIndexes) > 0 {
		prm.breakLineIndexes = prm.breakLineIndexes[:0]
	}

	if len(prm.cells) < rowLenLimit.Max().ToInt() {
		prm.anchorRowIndex = 0
		return prm.GetRowCount()
	}

	if len(prm.spaceCellIndexes) == 0 {
		for index := rowLenLimit.Max().ToInt(); index < len(prm.cells); index += rowLenLimit.Max().ToInt() {
			prm.breakLineIndexes = append(prm.breakLineIndexes, data.MakeIndexInterval(index, index))
		}

		return prm.GetRowCount()
	}

	lastBreakRowIndex := 0

	optimumCandidateIndex := 0
	optimumDistanceCandidateDelta := 0
	optimumDistanceDelta := rowLenLimit.Max().ToInt()

	rowLen := 0
	shift := 1
	spaceIndex := 0

	for index := 0; index <= len(prm.spaceCellIndexes); index++ {
		rowLen = len(prm.cells) - lastBreakRowIndex
		if index < len(prm.spaceCellIndexes) {
			spaceIndex = prm.spaceCellIndexes[index]
			rowLen = spaceIndex - lastBreakRowIndex
		}

		if rowLen < rowLenLimit.Min().ToInt() {
			if index != len(prm.spaceCellIndexes)-1 {
				continue
			}

			rowLen = len(prm.cells) - lastBreakRowIndex
			if rowLen <= rowLenLimit.Max().ToInt() {
				break
			}
		}

		if rowLen <= rowLenLimit.Max().ToInt() {
			optimumDistanceCandidateDelta = absInt(lastBreakRowIndex + rowLenLimit.Optimum().ToInt() - spaceIndex)
			if optimumDistanceCandidateDelta < optimumDistanceDelta {
				optimumCandidateIndex = spaceIndex
				optimumDistanceDelta = optimumDistanceCandidateDelta
			}
		}

		if rowLen > rowLenLimit.Max().ToInt() || optimumDistanceDelta < optimumDistanceCandidateDelta {
			if optimumCandidateIndex == 0 || optimumCandidateIndex == lastBreakRowIndex {
				// need to break a row not by mean spase
				optimumCandidateIndex = lastBreakRowIndex + rowLenLimit.Max().ToInt()
				shift = 0
				index--
			}

			prm.breakLineIndexes = append(prm.breakLineIndexes,
				data.MakeIndexInterval(optimumCandidateIndex, optimumCandidateIndex+shift))

			lastBreakRowIndex = optimumCandidateIndex
			optimumCandidateIndex = 0
			optimumDistanceDelta = rowLenLimit.Max().ToInt()
			shift = 1
			continue
		}
	}

	if anchorSpaceIndex > 0 {
		prm.anchorRowIndex = size.Height(len(prm.breakLineIndexes) - 1)
		for index := range prm.breakLineIndexes {
			if anchorSpaceIndex < prm.breakLineIndexes[index].GetBeginIndex() {
				continue
			}

			prm.anchorRowIndex = size.Height(index)
			break
		}
	}

	return prm.GetRowCount()
}

// absInt returns absolutely value of v
func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

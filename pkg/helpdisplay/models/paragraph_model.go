package models

import (
	data2 "github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"

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
	breakLineIndexes []data2.IndexInterval

	// tab - actionSequence of paragraph rows
	tab row_len_limiter.RowSize

	// anchorRowIndex - using for searching begin text row for display after a resize of terminal window
	anchorRowIndex int

	// usingTerminalWidth - value of using terminal width
	usingTerminalWidth int
}

// Init initializes created ParagraphModel object
func (prm *ParagraphModel) Init(rowLenLimit row_len_limiter.RowLenLimit, source *data2.Paragraph) int {
	_ = prm

	if len(source.Text) == 0 {
		return prm.GetRowCount()
	}

	prm.tmpSource = source.Text
	prm.cells = textToCells(source.Text)
	prm.initSpaceIndexes()
	if rowLenLimit.Max() > 0 {
		prm.breakLineIndexes = make([]data2.IndexInterval, 0, len(prm.cells)/rowLenLimit.Max().ToInt())
	}
	prm.tab = source.TabCount * row_len_limiter.TabSize
	prm.usingTerminalWidth = rowLenLimit.Max().ToInt()

	return prm.Update(rowLenLimit)
}

// GetAnchorRowIndex - anchorRowIndex field getter
func (prm *ParagraphModel) GetAnchorRowIndex() int {
	return prm.anchorRowIndex
}

// GetRowCount returns row count for rendering in a terminal
func (prm *ParagraphModel) GetRowCount() int {
	if len(prm.cells) != 0 {
		return 1 + len(prm.breakLineIndexes)
	}
	return 1
}

// GetUsingTab returns using tab size in runes for current terminal width
func (prm *ParagraphModel) GetUsingTab() int {
	if prm.usingTerminalWidth <= row_len_limiter.TerminalMinWidth {
		return 0
	}
	return prm.tab.ToInt()
}

// ShiftAnchorRow does a try to actionSequence anchor row and returns if the try is success
func (prm *ParagraphModel) ShiftAnchorRow(shift int) bool {
	prm.anchorRowIndex += shift
	if prm.anchorRowIndex < 0 || prm.anchorRowIndex >= prm.GetRowCount() {
		prm.anchorRowIndex = 0
		return false
	}

	return true
}

// SetBackRowAsAnchor sets index of last row as anchor row
func (prm *ParagraphModel) SetBackRowAsAnchor() {
	prm.anchorRowIndex = prm.GetRowCount() - 1
}

// GetRow returns a row for rendering in a terminal by index
func (prm *ParagraphModel) GetRow(index int) (int, []termbox.Cell) {
	_ = prm

	if index >= prm.GetRowCount() {
		return 0, nil
	}

	beginBreakLineIndex := 0
	endBreakRowIndex := len(prm.cells)

	if index > 0 {
		beginBreakLineIndex = prm.breakLineIndexes[index-1].GetEndIndex()
	}
	if index < len(prm.breakLineIndexes) {
		endBreakRowIndex = prm.breakLineIndexes[index].GetBeginIndex()
	}

	return prm.GetUsingTab(), prm.cells[beginBreakLineIndex:endBreakRowIndex]
}

// Update applies row len limit for rebuilding getting display row window
func (prm *ParagraphModel) Update(rowLenLimit row_len_limiter.RowLenLimit) int {
	_ = prm

	prm.usingTerminalWidth = rowLenLimit.Max().ToInt()

	anchorSpaceIndex := 0
	if prm.anchorRowIndex > 0 {
		anchorSpaceIndex = prm.breakLineIndexes[prm.anchorRowIndex-1].GetBeginIndex()
	}

	rowLenLimit = rowLenLimit.ApplyTabShift(row_len_limiter.RowSize(prm.GetUsingTab()))

	if len(prm.breakLineIndexes) > 0 {
		prm.breakLineIndexes = prm.breakLineIndexes[:0]
	}

	if len(prm.cells) < rowLenLimit.Max().ToInt() {
		prm.anchorRowIndex = 0
		return prm.GetRowCount()
	}

	if len(prm.spaceCellIndexes) == 0 {
		for index := rowLenLimit.Max().ToInt(); index < len(prm.cells); index += rowLenLimit.Max().ToInt() {
			prm.breakLineIndexes = append(prm.breakLineIndexes, data2.MakeIndexInterval(index, index))
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
				data2.MakeIndexInterval(optimumCandidateIndex, optimumCandidateIndex+shift))

			lastBreakRowIndex = optimumCandidateIndex
			optimumCandidateIndex = 0
			optimumDistanceDelta = rowLenLimit.Max().ToInt()
			shift = 1
			continue
		}
	}

	if anchorSpaceIndex > 0 {
		prm.anchorRowIndex = len(prm.breakLineIndexes) - 1
		for index := range prm.breakLineIndexes {
			if anchorSpaceIndex < prm.breakLineIndexes[index].GetBeginIndex() {
				continue
			}

			prm.anchorRowIndex = index
			break
		}
	}

	return prm.GetRowCount()
}

// initSpaceIndexes finds space runes and initializes spaceCellIndexes field
func (prm *ParagraphModel) initSpaceIndexes() {
	_ = prm

	prm.spaceCellIndexes = make([]int, 0, len(prm.cells)/averageWordLen)

	for index := 0; index < len(prm.cells); index++ {
		if prm.cells[index].Ch == runes.RuneSpace {
			prm.spaceCellIndexes = append(prm.spaceCellIndexes, index)
		}
	}
}

// absInt returns absolutely value of v
func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

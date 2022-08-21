package models

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
)

// PageModel - class which is getting page text parts for render in a terminal
type PageModel struct {
	headerModel HeaderModel
	paragraphs  []ParagraphModel
	rll         row_len_limiter.RowLenLimiter

	usingRowLeLimit     row_len_limiter.RowLenLimit
	usingTerminalHeight int

	anchorParagraphIndex     int
	anchorRowAbsolutelyIndex int
	rowCount                 int
}

// MakePageModel constructs PageModel object in a stack
func MakePageModel(pageData data.Page, terminalWidth, terminalHeight int) PageModel {
	pgm := PageModel{
		paragraphs: make([]ParagraphModel, len(pageData.Paragraphs)),
		rll:        row_len_limiter.MakeRowLenLimiter(),
	}
	pgm.usingRowLeLimit = pgm.rll.GetRowLenLimit(terminalWidth)
	pgm.usingTerminalHeight = terminalHeight

	pgm.headerModel = MakeHeaderModel(pageData.Header, pgm.usingRowLeLimit)

	pgm.rowCount = 1 // this is a header line
	for index := 0; index < len(pageData.Paragraphs); index++ {
		pgm.rowCount += pgm.paragraphs[index].Init(pgm.usingRowLeLimit, pageData.Paragraphs[index])
	}

	return pgm
}

func (pgm *PageModel) RowBegin() RowIterator {
	reverseCounter := pgm.usingTerminalHeight
	if pgm.rowCount < reverseCounter {
		reverseCounter = pgm.rowCount
	}

	rowNumber := 0
	if pgm.anchorParagraphIndex < len(pgm.paragraphs) {
		rowNumber = pgm.paragraphs[pgm.anchorParagraphIndex].GetAnchorRowIndex()
	}

	shift, cells := pgm.headerModel.GetHeaderRow()
	it := RowIterator{
		ShiftIndex: shift,
		Cells:      cells,

		reverseCounter:  reverseCounter,
		paragraphNumber: pgm.anchorParagraphIndex,
		rowNumber:       rowNumber,
	}

	return it
}

func (pgm *PageModel) RowNext(it RowIterator) RowIterator {
	if it.End() || len(pgm.paragraphs) <= it.paragraphNumber {
		return RowIterator{}
	}

	prm := pgm.paragraphs[it.paragraphNumber]

	it.ShiftIndex, it.Cells = prm.GetRow(it.rowNumber)

	it.reverseCounter--
	if it.reverseCounter == 0 {
		return it
	}

	it.rowNumber++
	if it.rowNumber == prm.GetRowCount() {
		it.rowNumber = 0
		it.paragraphNumber++
	}

	return it
}

// GetAnchorRowAbsolutelyIndex - anchorRowAbsolutelyIndex field getter
func (pgm *PageModel) GetAnchorRowAbsolutelyIndex() int {
	return pgm.anchorRowAbsolutelyIndex
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display row window
func (pgm *PageModel) Update(terminalWidth, terminalHeight, shift int) *dollyerr.Error {
	newRowLenLimit := pgm.rll.GetRowLenLimit(terminalWidth)
	if !newRowLenLimit.IsValid() {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			fmt.Errorf("PageModel.Update: invalid RowLenLimit: %v", newRowLenLimit),
		)
	}

	pgm.headerModel.Update(newRowLenLimit)

	if pgm.usingRowLeLimit != newRowLenLimit {
		pgm.usingRowLeLimit = newRowLenLimit

		oldAnchorRowAbsolutelyIndex := pgm.anchorRowAbsolutelyIndex

		pgm.rowCount = 0
		for index := 0; index < len(pgm.paragraphs); index++ {
			pgm.rowCount += pgm.paragraphs[index].Update(pgm.usingRowLeLimit)
			if index == pgm.anchorParagraphIndex {
				pgm.anchorRowAbsolutelyIndex = pgm.rowCount - pgm.paragraphs[index].GetRowCount() + pgm.paragraphs[index].GetAnchorRowIndex()
			}
		}
		pgm.rowCount++ // add a header line

		if oldAnchorRowAbsolutelyIndex == 0 {
			pgm.anchorParagraphIndex = 0
			pgm.anchorRowAbsolutelyIndex = 0
		}
	}

	return pgm.Shift(terminalHeight, shift)
}

// Shift applies a shift to display row window
func (pgm *PageModel) Shift(terminalHeight, shift int) *dollyerr.Error {
	if terminalHeight < 0 {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayInvalidTerminalHeightArgument,
			fmt.Errorf("PageModel.Shift: terminalHeight argument must be > 0 (passed value is '%v')", terminalHeight))
	}

	pgm.usingTerminalHeight = terminalHeight

	if shift > 0 {
		pgm.shiftForward(shift)
		return nil
	}
	if shift < 0 {
		pgm.shiftBack(shift)
		return nil
	}

	return nil
}

// shiftForward applies forward shift to display row window
func (pgm *PageModel) shiftForward(shift int) {
	if pgm.rowCount-pgm.anchorRowAbsolutelyIndex == pgm.usingTerminalHeight ||
		len(pgm.paragraphs) == 0 {
		return
	}

	if pgm.rowCount-pgm.anchorRowAbsolutelyIndex-shift < pgm.usingTerminalHeight {
		shift = pgm.rowCount - pgm.anchorRowAbsolutelyIndex - pgm.usingTerminalHeight
	}

	for i := 0; i < shift; i++ {
		if !pgm.paragraphs[pgm.anchorParagraphIndex].ShiftAnchorRow(1) {
			pgm.anchorParagraphIndex++
		}
	}

	pgm.anchorRowAbsolutelyIndex += shift
}

// shiftBack applies back shift to display row window
func (pgm *PageModel) shiftBack(shift int) {
	if len(pgm.paragraphs) == 0 {
		return
	}
	if pgm.anchorRowAbsolutelyIndex+shift < 0 {
		pgm.paragraphs[pgm.anchorParagraphIndex].anchorRowIndex = 0
		pgm.anchorParagraphIndex = 0
		pgm.anchorRowAbsolutelyIndex = 0
		return
	}

	for i := 0; i > shift; i-- {
		if !pgm.paragraphs[pgm.anchorParagraphIndex].ShiftAnchorRow(-1) {
			pgm.anchorParagraphIndex--
			pgm.paragraphs[pgm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
	}

	pgm.anchorRowAbsolutelyIndex += shift
}

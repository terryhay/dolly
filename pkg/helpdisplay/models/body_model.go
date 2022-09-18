package models

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
)

// BodyModel implements a text page body model with some paragraphs
type BodyModel struct {
	paragraphs []*ParagraphModel

	usingSize TerminalSize

	anchorParagraphIndex     size.Height
	anchorRowAbsolutelyIndex size.Height
	rowCount                 size.Height
}

// NewBodyModel constructs BodyModel object
func NewBodyModel(pageData data.Page, termSize TerminalSize) *BodyModel {
	paragraphs := make([]*ParagraphModel, 0, len(pageData.Paragraphs))

	rowCount := size.Height(0)
	for index := 0; index < len(pageData.Paragraphs); index++ {
		p := NewParagraphModel(termSize.GetWidth(), pageData.Paragraphs[index])
		rowCount += p.GetRowCount()

		paragraphs = append(paragraphs, p)
	}

	return &BodyModel{
		paragraphs: paragraphs,
		usingSize:  termSize,
		rowCount:   rowCount,
	}
}

func (bm *BodyModel) GetRowCount() size.Height {
	if bm == nil {
		return 0
	}
	return bm.rowCount
}

func (bm *BodyModel) GetAnchorParagraphIndex() size.Height {
	if bm == nil {
		return 0
	}
	return bm.anchorParagraphIndex
}

func (bm *BodyModel) GetAnchorParagraphRowIndex() size.Height {
	if bm == nil {
		return 0
	}

	var index size.Height
	if bm.anchorParagraphIndex.ToInt() < len(bm.paragraphs) {
		index = bm.paragraphs[bm.anchorParagraphIndex].GetAnchorRowIndex()
	}
	return index
}

// GetAnchorRowAbsolutelyIndex - anchorRowAbsolutelyIndex field getter
func (bm *BodyModel) GetAnchorRowAbsolutelyIndex() size.Height {
	if bm == nil {
		return 0
	}
	return bm.anchorRowAbsolutelyIndex
}

func (bm *BodyModel) GetParagraph(index size.Height) (*ParagraphModel, *dollyerr.Error) {
	if bm == nil {
		return nil,
			dollyerr.NewError(
				dollyerr.CodeHelpDisplayReceiverIsNilPointer,
				fmt.Errorf("BodyModel.GetParagraph: receiver is a nil pointer"),
			)
	}

	if len(bm.paragraphs) <= index.ToInt() {
		return nil, nil
	}

	return bm.paragraphs[index], nil
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display row window
func (bm *BodyModel) Update(termSize TerminalSize, shift int) *dollyerr.Error {
	if bm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("BodyModel.GetParagraph: receiver is a nil pointer"),
		)
	}

	width := termSize.GetWidth()
	if !width.IsValid() {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			fmt.Errorf("PageModel.Update: invalid RowLenLimit: %v", width.IsValid()),
		)
	}

	if bm.usingSize.GetWidth() != width {
		bm.usingSize.Width = width

		oldAnchorRowAbsolutelyIndex := bm.anchorRowAbsolutelyIndex

		bm.rowCount = 0
		for index := 0; index < len(bm.paragraphs); index++ {
			bm.rowCount += bm.paragraphs[index].Update(bm.usingSize.GetWidth())
			if index == bm.anchorParagraphIndex.ToInt() {
				bm.anchorRowAbsolutelyIndex = size.Height(bm.rowCount.ToInt() - bm.paragraphs[index].GetRowCount().ToInt() + bm.paragraphs[index].GetAnchorRowIndex().ToInt())
			}
		}

		if oldAnchorRowAbsolutelyIndex == 0 {
			bm.anchorParagraphIndex = 0
			bm.anchorRowAbsolutelyIndex = 0
		}
	}

	return bm.Shift(termSize.GetHeight(), shift)
}

// Shift applies a shift to display row window
func (bm *BodyModel) Shift(terminalHeight size.Height, shift int) *dollyerr.Error {
	if bm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("BodyModel.GetParagraph: receiver is a nil pointer"),
		)
	}

	bm.usingSize.Height = terminalHeight

	if shift > 0 {
		bm.shiftForward(shift)
		return nil
	}
	if shift < 0 {
		bm.shiftBack(shift)
		return nil
	}

	return nil
}

// shiftForward applies forward shift to display row window
func (bm *BodyModel) shiftForward(shift int) {
	_ = bm

	if bm.rowCount <= bm.usingSize.GetHeight() {
		if len(bm.paragraphs) > 0 {
			bm.paragraphs[bm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
		bm.anchorParagraphIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.rowCount.ToInt()-bm.anchorRowAbsolutelyIndex.ToInt() == bm.usingSize.GetHeight().ToInt() ||
		len(bm.paragraphs) == 0 {
		return
	}

	if bm.rowCount.ToInt()-bm.anchorRowAbsolutelyIndex.ToInt()-shift < bm.usingSize.GetHeight().ToInt() {
		shift = bm.rowCount.ToInt() - bm.anchorRowAbsolutelyIndex.ToInt() - bm.usingSize.GetHeight().ToInt()
	}

	for i := 0; i < shift; i++ {
		if !bm.paragraphs[bm.anchorParagraphIndex].ShiftAnchorRow(1) {
			bm.anchorParagraphIndex++
		}
	}

	anchorRowAbsolutelyIndex := bm.anchorRowAbsolutelyIndex.ToInt() + shift
	bm.anchorRowAbsolutelyIndex = size.Height(anchorRowAbsolutelyIndex)
}

// shiftBack applies back shift to display row window
func (bm *BodyModel) shiftBack(shift int) {
	_ = bm

	if bm.rowCount <= bm.usingSize.GetHeight() {
		if len(bm.paragraphs) > 0 {
			bm.paragraphs[bm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
		bm.anchorParagraphIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.anchorRowAbsolutelyIndex.ToInt()+shift < 0 {
		bm.paragraphs[bm.anchorParagraphIndex].anchorRowIndex = 0
		bm.anchorParagraphIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	for i := 0; i > shift; i-- {
		if !bm.paragraphs[bm.anchorParagraphIndex].ShiftAnchorRow(-1) {
			bm.anchorParagraphIndex--
			bm.paragraphs[bm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
	}

	bm.anchorRowAbsolutelyIndex = size.Height(bm.anchorRowAbsolutelyIndex.ToInt() + shift)
}

package body_model

import (
	"fmt"
	"github.com/terryhay/dolly/argparser/helpdisplay/index"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	"github.com/terryhay/dolly/argparser/helpdisplay/page_model/body_model/paragraph_model"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"github.com/terryhay/dolly/argparser/helpdisplay/terminal_size"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// BodyModel implements a text page body model with some paragraphs
type BodyModel struct {
	paragraphs []*paragraph_model.ParagraphModel

	termSize terminal_size.TerminalSize

	anchorParagraphIndex     index.Index
	anchorRowAbsolutelyIndex index.Index
	rowCount                 size.Height
}

// NewBodyModel constructs BodyModel object
func NewBodyModel(pageData page.Page, termSize terminal_size.TerminalSize) *BodyModel {
	paragraphs := make([]*paragraph_model.ParagraphModel, 0, len(pageData.Paragraphs))

	rowCount := size.Height(0)
	for i := 0; i < len(pageData.Paragraphs); i++ {
		p := paragraph_model.NewParagraphModel(termSize.GetWidthLimit(), pageData.Paragraphs[i])
		rowCount += p.GetRowCount()

		paragraphs = append(paragraphs, p)
	}

	return &BodyModel{
		paragraphs: paragraphs,
		termSize:   termSize,
		rowCount:   rowCount,
	}
}

// GetRowCount returns rowCount field value
func (bm *BodyModel) GetRowCount() size.Height {
	if bm == nil {
		return 0
	}
	return bm.rowCount
}

// GetAnchorParagraphIndex returns anchorParagraphIndex field value
func (bm *BodyModel) GetAnchorParagraphIndex() index.Index {
	if bm == nil {
		return 0
	}
	return bm.anchorParagraphIndex
}

func (bm *BodyModel) GetAnchorParagraphRowIndex() index.Index {
	if bm == nil {
		return 0
	}

	var res index.Index
	if bm.anchorParagraphIndex.ToInt() < len(bm.paragraphs) {
		res = bm.paragraphs[bm.anchorParagraphIndex].GetAnchorRowIndex()
	}
	return res
}

// GetAnchorRowAbsolutelyIndex - anchorRowAbsolutelyIndex field getter
func (bm *BodyModel) GetAnchorRowAbsolutelyIndex() index.Index {
	if bm == nil {
		return 0
	}
	return bm.anchorRowAbsolutelyIndex
}

// GetParagraph returns paragraphModel object by index
func (bm *BodyModel) GetParagraph(i index.Index) *paragraph_model.ParagraphModel {
	if bm == nil {
		return nil
	}

	if len(bm.paragraphs) <= i.ToInt() {
		return nil
	}

	return bm.paragraphs[i]
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display dynamic_row window
func (bm *BodyModel) Update(termSize terminal_size.TerminalSize, shift int) *dollyerr.Error {
	if bm == nil {
		return nil
	}

	widthLimit := termSize.GetWidthLimit()
	if !widthLimit.IsValid() {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			fmt.Errorf("PageModel.update: invalid RowLenLimit: %s", widthLimit.String()),
		)
	}

	if bm.termSize.GetWidthLimit() != widthLimit {
		bm.termSize = bm.termSize.CloneWithNewWidthLimit(widthLimit)

		oldAnchorRowAbsolutelyIndex := bm.anchorRowAbsolutelyIndex

		bm.rowCount = 0
		for i := 0; i < len(bm.paragraphs); i++ {
			bm.rowCount += bm.paragraphs[i].Update(bm.termSize.GetWidthLimit())
			if i == bm.anchorParagraphIndex.ToInt() {
				bm.anchorRowAbsolutelyIndex = index.MakeIndex(bm.rowCount.ToInt() - bm.paragraphs[i].GetRowCount().ToInt() + bm.paragraphs[i].GetAnchorRowIndex().ToInt())
			}
		}

		if oldAnchorRowAbsolutelyIndex == 0 {
			bm.anchorParagraphIndex = 0
			bm.anchorRowAbsolutelyIndex = 0
		}
	}

	return bm.Shift(termSize.GetHeight(), shift)
}

// Shift applies a shift to display dynamic_row window
func (bm *BodyModel) Shift(terminalHeight size.Height, shift int) *dollyerr.Error {
	if bm == nil {
		return nil
	}

	bm.termSize = bm.termSize.CloneWithNewHeight(terminalHeight)

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

// shiftForward applies forward shift to display dynamic_row window
func (bm *BodyModel) shiftForward(shift int) {
	_ = bm

	if bm.rowCount <= bm.termSize.GetHeight() {
		if len(bm.paragraphs) > 0 {
			bm.paragraphs[bm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
		bm.anchorParagraphIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.rowCount.ToInt()-bm.anchorRowAbsolutelyIndex.ToInt() == bm.termSize.GetHeight().ToInt() ||
		len(bm.paragraphs) == 0 {
		return
	}

	if bm.rowCount.ToInt()-bm.anchorRowAbsolutelyIndex.ToInt()-shift < bm.termSize.GetHeight().ToInt() {
		shift = bm.rowCount.ToInt() - bm.anchorRowAbsolutelyIndex.ToInt() - bm.termSize.GetHeight().ToInt()
	}

	for i := 0; i < shift; i++ {
		if !bm.paragraphs[bm.anchorParagraphIndex].ShiftAnchorRow(1) {
			bm.anchorParagraphIndex++
		}
	}

	anchorRowAbsolutelyIndex := index.Append(bm.anchorRowAbsolutelyIndex, shift)
	bm.anchorRowAbsolutelyIndex = anchorRowAbsolutelyIndex
}

// shiftBack applies back shift to display dynamic_row window
func (bm *BodyModel) shiftBack(shift int) {
	_ = bm

	if bm.rowCount <= bm.termSize.GetHeight() {
		if len(bm.paragraphs) > 0 {
			bm.paragraphs[bm.anchorParagraphIndex].SetBackRowAsAnchor()
		}
		bm.anchorParagraphIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.anchorRowAbsolutelyIndex.ToInt()+shift < 0 {
		bm.paragraphs[bm.anchorParagraphIndex].SetAnchorRowIndex(index.Null)
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

	bm.anchorRowAbsolutelyIndex = index.Append(bm.anchorRowAbsolutelyIndex, shift)
}

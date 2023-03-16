package body_model

import (
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	prm "github.com/terryhay/dolly/man_style_help/page_model/body_model/row_model"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

// BodyModel implements a text page body model with some modelRows
type BodyModel struct {
	modelRows []*prm.RowModel

	termSize ts.TerminalSize

	anchorRowModelIndex      index.Index
	anchorRowAbsolutelyIndex index.Index
	rowCount                 size.Height
}

// NewBodyModel constructs BodyModel object
func NewBodyModel(pageBody hp.Body, termSize ts.TerminalSize) *BodyModel {
	paragraphs := make([]*prm.RowModel, 0, pageBody.RowCount())

	rowCount := size.HeightZero
	for i := index.Zero; i < pageBody.RowCount(); i++ {
		paragraph := prm.NewRowModel(pageBody.Row(i), termSize.GetWidthLimit())
		rowCount += paragraph.GetRowCount()

		paragraphs = append(paragraphs, paragraph)
	}

	return &BodyModel{
		modelRows: paragraphs,
		termSize:  termSize,
		rowCount:  rowCount,
	}
}

// GetRowCount returns rowCount field value
func (bm *BodyModel) GetRowCount() size.Height {
	if bm == nil {
		return 0
	}
	return bm.rowCount
}

// GetAnchorRowModelIndex gets anchorRowModelIndex field value
func (bm *BodyModel) GetAnchorRowModelIndex() index.Index {
	if bm == nil {
		return index.Zero
	}
	return bm.anchorRowModelIndex
}

// GetAnchorRowIndex finds anchor paragraph by index and returns its anchor row index
func (bm *BodyModel) GetAnchorRowIndex() index.Index {
	if bm == nil || len(bm.modelRows) <= bm.anchorRowModelIndex.Int() {
		return index.Zero
	}

	return bm.modelRows[bm.anchorRowModelIndex].GetAnchorRowIndex()
}

// GetAnchorRowAbsolutelyIndex gets anchorRowAbsolutelyIndex field value
func (bm *BodyModel) GetAnchorRowAbsolutelyIndex() index.Index {
	if bm == nil {
		return 0
	}
	return bm.anchorRowAbsolutelyIndex
}

// GetRowModel gets paragraphModel object by index
func (bm *BodyModel) GetRowModel(i index.Index) *prm.RowModel {
	if bm == nil || len(bm.modelRows) <= i.Int() {
		return nil
	}

	return bm.modelRows[i]
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display dynamic_row window
func (bm *BodyModel) Update(termSize ts.TerminalSize, shiftVertical int) {
	if bm == nil {
		return
	}

	if bm.termSize.GetWidthLimit() != termSize.GetWidthLimit() {
		bm.termSize = ts.MakeTerminalSize(termSize.GetWidthLimit(), bm.termSize.GetHeight())

		oldAnchorRowAbsolutelyIndex := bm.anchorRowAbsolutelyIndex

		bm.rowCount = 0
		for i := index.Zero; i < index.Index(len(bm.modelRows)); i++ {
			bm.rowCount += bm.modelRows[i].Update(bm.termSize.GetWidthLimit())
			if i == bm.anchorRowModelIndex {
				bm.anchorRowAbsolutelyIndex = index.MakeIndex(bm.rowCount.Int() - bm.modelRows[i].GetRowCount().Int() + bm.modelRows[i].GetAnchorRowIndex().Int())
			}
		}

		switch {
		case oldAnchorRowAbsolutelyIndex == 0:
			bm.anchorRowModelIndex = 0
			bm.anchorRowAbsolutelyIndex = 0

		// Big shift and big terminal height were here, now they have collapsed. So need to shift up the text
		case bm.rowCount.Int()-bm.anchorRowAbsolutelyIndex.Int() < bm.termSize.GetHeight().Int():
			shiftVertical += bm.rowCount.Int() - bm.anchorRowAbsolutelyIndex.Int() - bm.termSize.GetHeight().Int()
		}
	}

	bm.Shift(termSize.GetHeight(), shiftVertical)
}

// Shift applies a shift to display dynamic_row window
func (bm *BodyModel) Shift(terminalHeight size.Height, shiftVertical int) {
	if bm == nil {
		return
	}

	bm.termSize = ts.MakeTerminalSize(bm.termSize.GetWidthLimit(), terminalHeight)

	if shiftVertical > 0 {
		bm.shiftDown(shiftVertical)
		return
	}
	if shiftVertical < 0 {
		bm.shiftUp(shiftVertical)
		return
	}
}

// shiftDown applies forward shift to display dynamic_row window
func (bm *BodyModel) shiftDown(shiftVertical int) {
	_ = bm

	if bm.rowCount <= bm.termSize.GetHeight() {
		if len(bm.modelRows) > 0 {
			bm.modelRows[bm.anchorRowModelIndex].SetBackRowAsAnchor()
		}
		bm.anchorRowModelIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.rowCount.Int()-bm.anchorRowAbsolutelyIndex.Int() == bm.termSize.GetHeight().Int() ||
		len(bm.modelRows) == 0 {
		return
	}

	if bm.rowCount.Int()-bm.anchorRowAbsolutelyIndex.Int()-shiftVertical < bm.termSize.GetHeight().Int() {
		shiftVertical = bm.rowCount.Int() - bm.anchorRowAbsolutelyIndex.Int() - bm.termSize.GetHeight().Int()
	}

	for i := 0; i < shiftVertical; i++ {
		if !bm.modelRows[bm.anchorRowModelIndex].ShiftAnchorRow(1) {
			bm.anchorRowModelIndex++
		}
	}

	anchorRowAbsolutelyIndex := index.Append(bm.anchorRowAbsolutelyIndex, shiftVertical)
	bm.anchorRowAbsolutelyIndex = anchorRowAbsolutelyIndex
}

// shiftUp applies back shift to display dynamic_row window
func (bm *BodyModel) shiftUp(shiftVertical int) {
	_ = bm

	if bm.rowCount <= bm.termSize.GetHeight() {
		if len(bm.modelRows) > 0 {
			bm.modelRows[bm.anchorRowModelIndex].SetBackRowAsAnchor()
		}
		bm.anchorRowModelIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	if bm.anchorRowAbsolutelyIndex.Int()+shiftVertical < 0 {
		bm.modelRows[bm.anchorRowModelIndex].SetAnchorRowIndex(index.Zero)
		bm.anchorRowModelIndex = 0
		bm.anchorRowAbsolutelyIndex = 0
		return
	}

	for i := 0; i > shiftVertical; i-- {
		if !bm.modelRows[bm.anchorRowModelIndex].ShiftAnchorRow(-1) {
			bm.anchorRowModelIndex--
			bm.modelRows[bm.anchorRowModelIndex].SetBackRowAsAnchor()
		}
	}

	bm.anchorRowAbsolutelyIndex = index.Append(bm.anchorRowAbsolutelyIndex, shiftVertical)
}

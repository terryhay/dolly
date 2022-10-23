package row_iterator

import (
	"github.com/terryhay/dolly/argparser/helpdisplay/index"
	pgm "github.com/terryhay/dolly/argparser/helpdisplay/page_model"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
)

// RowIterator contains some temp page for iterating by model rows
type RowIterator struct {
	model *pgm.PageModel

	modelRow row.Row

	reverseCounter    size.Height
	paragraphIndex    index.Index
	paragraphRowIndex index.Index
}

// MakeRowIterator constructs a RowIterator object in a stack
func MakeRowIterator(model *pgm.PageModel) RowIterator {
	reverseCounter := model.GetUsingTerminalSize().GetHeight()
	if model.GetRowCount() < reverseCounter {
		reverseCounter = model.GetRowCount()
	}

	return RowIterator{
		model:    model,
		modelRow: model.GetHeaderModel().GetViewRow(),

		reverseCounter:    reverseCounter,
		paragraphIndex:    model.GetBodyModel().GetAnchorParagraphIndex(),
		paragraphRowIndex: model.GetBodyModel().GetAnchorParagraphRowIndex(),
	}
}

// End returns if iterating is ended
func (ri *RowIterator) End() bool {
	return ri.reverseCounter == 0
}

// Row returns a current value of model dynamic_row
func (ri *RowIterator) Row() row.Row {
	return ri.modelRow
}

// Next goes to next model dynamic_row
func (ri *RowIterator) Next() {
	_ = ri

	if ri.End() {
		return
	}

	prm := ri.model.GetBodyModel().GetParagraph(ri.paragraphIndex)
	if prm == nil {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		ri.reverseCounter--
		return
	}

	ri.modelRow = prm.GetRow(ri.paragraphRowIndex)

	ri.reverseCounter--
	if ri.reverseCounter == 1 {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		return
	}

	if ri.reverseCounter == 0 {
		return
	}

	ri.paragraphRowIndex++
	if ri.paragraphRowIndex.ToInt() == prm.GetRowCount().ToInt() {
		ri.paragraphRowIndex = 0
		ri.paragraphIndex++
	}
}

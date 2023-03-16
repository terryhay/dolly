package row_iterator

import (
	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	"github.com/terryhay/dolly/man_style_help/row"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

// RowIterator contains some temp page for iterating by model rows
type RowIterator struct {
	model *pgm.PageModel

	modelRow row.Row

	reverseCounter size.Height
	rowModelIndex  index.Index
	rowIndex       index.Index
}

// MakeRowIterator constructs a RowIterator object in a stack
func MakeRowIterator(model *pgm.PageModel) RowIterator {
	reverseCounter := model.GetUsingTermSize().GetHeight()
	if model.GetRowCount() < reverseCounter {
		reverseCounter = model.GetRowCount()
	}

	return RowIterator{
		model:    model,
		modelRow: model.GetHeaderModel().GetViewRow(),

		reverseCounter: reverseCounter,
		rowModelIndex:  model.GetBodyModel().GetAnchorRowModelIndex(),
		rowIndex:       model.GetBodyModel().GetAnchorRowIndex(),
	}
}

// End returns if iterating is ended
func (ri *RowIterator) End() bool {
	return ri.reverseCounter == 0
}

// RowModel returns a current value of model dynamic_row
func (ri *RowIterator) RowModel() row.Row {
	return ri.modelRow
}

// Next goes to next model dynamic_row
func (ri *RowIterator) Next() {
	_ = ri

	if ri.End() {
		return
	}

	prm := ri.model.GetBodyModel().GetRowModel(ri.rowModelIndex)
	if prm == nil {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		ri.reverseCounter--
		return
	}

	ri.modelRow = prm.GetRow(ri.rowIndex)

	ri.reverseCounter--
	if ri.reverseCounter == 1 {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		return
	}

	if ri.reverseCounter == 0 {
		return
	}

	ri.rowIndex++
	if ri.rowIndex.Int() == prm.GetRowCount().Int() {
		ri.rowIndex = 0
		ri.rowModelIndex++
	}
}

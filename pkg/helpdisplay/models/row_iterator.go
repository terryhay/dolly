package models

import (
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
)

import (
	"fmt"
)

// RowIterator contains some temp data for iterating by model rows
type RowIterator struct {
	model *PageModel

	modelRow row.Row

	ReverseCounter    size.Height
	paragraphIndex    size.Height
	paragraphRowIndex size.Height
}

// MakeRowIterator constructs a RowIterator object in a stack
func MakeRowIterator(model *PageModel) RowIterator {
	reverseCounter := model.GetUsingTerminalSize().GetHeight()
	if model.GetRowCount() < reverseCounter {
		reverseCounter = model.GetRowCount()
	}

	return RowIterator{
		model:    model,
		modelRow: model.GetHeaderModel().GetHeaderRow(),

		ReverseCounter:    reverseCounter,
		paragraphIndex:    model.GetBodyModel().GetAnchorParagraphIndex(),
		paragraphRowIndex: model.GetBodyModel().GetAnchorParagraphRowIndex(),
	}
}

// End returns if iterating is ended
func (ri *RowIterator) End() bool {
	if ri == nil {
		return true
	}
	return ri.ReverseCounter == 0
}

// Row returns a current value of model row
func (ri *RowIterator) Row() row.Row {
	if ri == nil {
		return row.Row{}
	}
	return ri.modelRow
}

// Next goes to next model row
func (ri *RowIterator) Next() *dollyerr.Error {
	if ri == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("RowIterator.Next: receiver is a nil pointer"),
		)
	}
	if ri.End() {
		return dollyerr.NewError(dollyerr.CodeRowIteratorAttemptToIterateFromEndedIterator,
			fmt.Errorf("RowIterator.Next: attempt to iterate from ended iterator"))
	}

	prm, err := ri.model.GetBodyModel().GetParagraph(ri.paragraphIndex)
	if err != nil {
		return err.AppendError(
			fmt.Errorf("RowIterator.Next: can't get paragrpah from BodyModel by index '%d'", ri.paragraphIndex))
	}
	if prm == nil {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		ri.ReverseCounter--
		return nil
	}

	ri.modelRow = prm.GetRow(ri.paragraphRowIndex)

	ri.ReverseCounter--
	if ri.ReverseCounter == 1 {
		ri.modelRow = ri.model.GetFooterModel().GetFooterRow()
		return nil
	}

	if ri.ReverseCounter == 0 {
		return nil
	}

	ri.paragraphRowIndex++
	if ri.paragraphRowIndex == prm.GetRowCount() {
		ri.paragraphRowIndex = 0
		ri.paragraphIndex++
	}

	return nil
}

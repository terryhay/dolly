package models

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
)

// PageModel - class which is getting page text parts for render in a terminal
type PageModel struct {
	headerModel *HeaderModel
	bodyModel   *BodyModel
	footerModel *FooterModel

	usingSize TerminalSize
	rowCount  size.Height
}

// NewPageModel constructs PageModel object
func NewPageModel(pageData data.Page, size TerminalSize) *PageModel {
	headerModel := NewHeaderModel(pageData.Header, size)
	bodyModel := NewBodyModel(pageData, size)
	footerModel := NewFooterModel(size)

	return &PageModel{
		headerModel: headerModel,
		bodyModel:   bodyModel,
		footerModel: footerModel,

		usingSize: size,
		rowCount:  2 + bodyModel.GetRowCount(),
	}
}

// GetAnchorRowAbsolutelyIndex - anchorRowAbsolutelyIndex field getter
func (pgm *PageModel) GetAnchorRowAbsolutelyIndex() size.Height {
	if pgm == nil {
		return 0
	}
	return pgm.bodyModel.GetAnchorRowAbsolutelyIndex()
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display row window
func (pgm *PageModel) Update(size TerminalSize, shift int) *dollyerr.Error {
	if pgm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("PageModel.Shift: receiver is a nil pointer"),
		)
	}
	err := size.IsValid()
	if err != nil {
		return err.AppendError(fmt.Errorf("PageModel.Update: invalid TerminalSize: %v", size))
	}

	pgm.headerModel.Update(size)
	{
		sizeForBodyModel := size.Clone()
		sizeForBodyModel.Height -= 2
		err = pgm.bodyModel.Update(sizeForBodyModel, shift)
		if err != nil {
			return err.AppendError(fmt.Errorf("PageModel.Update: bodyModel updating error"))
		}
	}
	pgm.footerModel.Update(size)

	return nil
}

// Shift applies a shift to display row window
func (pgm *PageModel) Shift(terminalHeight size.Height, shift int) *dollyerr.Error {
	if pgm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("PageModel.Shift: receiver is a nil pointer"),
		)
	}
	return pgm.bodyModel.Shift(terminalHeight, shift)
}

func (pgm *PageModel) GetUsingTerminalSize() TerminalSize {
	if pgm == nil {
		return TerminalSize{}
	}
	return pgm.usingSize
}

func (pgm *PageModel) GetRowCount() size.Height {
	if pgm == nil {
		return 0
	}
	return pgm.rowCount
}

func (pgm *PageModel) GetHeaderModel() *HeaderModel {
	if pgm == nil {
		return nil
	}
	return pgm.headerModel
}

func (pgm *PageModel) GetBodyModel() *BodyModel {
	if pgm == nil {
		return nil
	}
	return pgm.bodyModel
}

func (pgm *PageModel) GetFooterModel() *FooterModel {
	if pgm == nil {
		return nil
	}
	return pgm.footerModel
}

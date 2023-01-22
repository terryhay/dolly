package page_model

import (
	"fmt"
	"github.com/terryhay/dolly/man_style_help/page"
	"github.com/terryhay/dolly/man_style_help/page_model/body_model"
	"github.com/terryhay/dolly/man_style_help/page_model/footer_model"
	"github.com/terryhay/dolly/man_style_help/page_model/header_model"
	"github.com/terryhay/dolly/man_style_help/size"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/utils/dollyerr"
	"github.com/terryhay/dolly/utils/index"
)

// PageModel - class which is getting page text parts for render in a terminal
type PageModel struct {
	modelHeader *header_model.HeaderModel
	modelBody   *body_model.BodyModel
	modelFooter *footer_model.FooterModel

	sizeTerm ts.TerminalSize
	rowCount size.Height
}

// NewPageModel constructs PageModel object
func NewPageModel(pageData page.Page, size ts.TerminalSize) (*PageModel, *dollyerr.Error) {
	headerModel, err := header_model.NewHeaderModel(pageData.Header, size)
	if err != nil {
		return nil, dollyerr.Append(err, fmt.Errorf("NewPageModel error"))
	}

	bodyModel := body_model.NewBodyModel(pageData, size)
	footerModel := footer_model.NewFooterModel()

	return &PageModel{
			modelHeader: headerModel,
			modelBody:   bodyModel,
			modelFooter: footerModel,

			sizeTerm: size,
			rowCount: 2 + bodyModel.GetRowCount(),
		},
		nil
}

// GetAnchorRowAbsolutelyIndex - anchorRowAbsolutelyIndex field getter
func (pgm *PageModel) GetAnchorRowAbsolutelyIndex() index.Index {
	if pgm == nil {
		return index.Null
	}
	return pgm.modelBody.GetAnchorRowAbsolutelyIndex()
}

// Update applies terminal size and display actionSequence to the models and rebuild getting display dynamic_row window
func (pgm *PageModel) Update(sizeTerm ts.TerminalSize, shiftVertical int) *dollyerr.Error {
	if pgm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("PageModel.Shift: receiver is a nil pointer"),
		)
	}
	pgm.sizeTerm = sizeTerm

	err := pgm.modelHeader.Update(sizeTerm)
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf("PageModel.update: modelHeader updating error"))
	}

	const countHeaderAndFooterRows = 2
	err = pgm.modelBody.Update(
		sizeTerm.CloneWithNewHeight(sizeTerm.GetHeight().AddInt(-countHeaderAndFooterRows)),
		shiftVertical,
	)
	return dollyerr.Append(err, fmt.Errorf("PageModel.update: modelBody updating error"))
}

// Shift applies a shift to display dynamic_row window
func (pgm *PageModel) Shift(terminalHeight size.Height, shift int) *dollyerr.Error {
	if pgm == nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayReceiverIsNilPointer,
			fmt.Errorf("PageModel.Shift: receiver is a nil pointer"),
		)
	}
	return pgm.modelBody.Shift(terminalHeight, shift)
}

func (pgm *PageModel) GetUsingTerminalSize() ts.TerminalSize {
	if pgm == nil {
		return ts.TerminalSize{}
	}
	return pgm.sizeTerm
}

func (pgm *PageModel) GetRowCount() size.Height {
	if pgm == nil {
		return 0
	}
	return pgm.rowCount
}

func (pgm *PageModel) GetHeaderModel() *header_model.HeaderModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelHeader
}

func (pgm *PageModel) GetBodyModel() *body_model.BodyModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelBody
}

func (pgm *PageModel) GetFooterModel() *footer_model.FooterModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelFooter
}

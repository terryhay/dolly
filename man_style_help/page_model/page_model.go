package page_model

import (
	"errors"

	"github.com/terryhay/dolly/argparser/help_page/page"
	"github.com/terryhay/dolly/man_style_help/page_model/body_model"
	"github.com/terryhay/dolly/man_style_help/page_model/footer_model"
	"github.com/terryhay/dolly/man_style_help/page_model/header_model"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

// PageModel - class which is getting page text parts for render in a terminal
type PageModel struct {
	modelHeader *header_model.HeaderModel
	modelBody   *body_model.BodyModel
	modelFooter *footer_model.FooterModel

	usingTermSize ts.TerminalSize
	rowCount      size.Height
}

// ErrNewInvalidTerminalSize - TerminalSize.IsValid returned error
var ErrNewInvalidTerminalSize = errors.New(`page_model.New: TerminalSize.IsValid returned error`)

// New constructs PageModel object
func New(appName coty.NameApp, pageBody page.Body, sizeTerminal ts.TerminalSize) (*PageModel, error) {
	if err := sizeTerminal.IsValid(); err != nil {
		return nil, errors.Join(ErrNewInvalidTerminalSize, err)
	}

	bodyModel := body_model.NewBodyModel(pageBody, sizeTerminal)

	return &PageModel{
		modelHeader: header_model.NewHeaderModel(appName, sizeTerminal),
		modelBody:   bodyModel,
		modelFooter: footer_model.NewFooterModel(),

		usingTermSize: sizeTerminal,
		rowCount:      2 + bodyModel.GetRowCount(),
	}, nil
}

// GetAnchorRowAbsolutelyIndex gets anchorRowAbsolutelyIndex field
func (pgm *PageModel) GetAnchorRowAbsolutelyIndex() index.Index {
	if pgm == nil {
		return index.Zero
	}
	return pgm.modelBody.GetAnchorRowAbsolutelyIndex()
}

// ErrUpdateInvalidTerminalSize - TerminalSize.IsValid returned error
var ErrUpdateInvalidTerminalSize = errors.New(`PageModel.Update: TerminalSize.IsValid returned error`)

// Update applies terminal size and display actionSequence to the models and rebuild getting display dynamic_row window
func (pgm *PageModel) Update(sizeTerm ts.TerminalSize, shiftVertical int) error {
	if pgm == nil {
		return nil
	}
	pgm.usingTermSize = sizeTerm

	if err := sizeTerm.IsValid(); err != nil {
		return errors.Join(ErrUpdateInvalidTerminalSize, err)
	}

	pgm.modelHeader.Update(sizeTerm)

	const countHeaderAndFooterRows = 2
	pgm.modelBody.Update(
		ts.MakeTerminalSize(sizeTerm.GetWidthLimit(), size.AppendHeight(sizeTerm.GetHeight(), -countHeaderAndFooterRows)),
		shiftVertical,
	)
	return nil
}

// Shift applies a shift to display dynamic_row window
func (pgm *PageModel) Shift(terminalHeight size.Height, shift int) {
	pgm.GetBodyModel().Shift(terminalHeight, shift)
}

// GetUsingTermSize gets usingTermSize field
func (pgm *PageModel) GetUsingTermSize() ts.TerminalSize {
	if pgm == nil {
		return ts.TerminalSize{}
	}
	return pgm.usingTermSize
}

// GetRowCount gets rowCount field
func (pgm *PageModel) GetRowCount() size.Height {
	if pgm == nil {
		return 0
	}
	return pgm.rowCount
}

// GetHeaderModel gets modelHeader field
func (pgm *PageModel) GetHeaderModel() *header_model.HeaderModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelHeader
}

// GetBodyModel gets modelBody field
func (pgm *PageModel) GetBodyModel() *body_model.BodyModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelBody
}

// GetFooterModel gets modelFooter field
func (pgm *PageModel) GetFooterModel() *footer_model.FooterModel {
	if pgm == nil {
		return nil
	}
	return pgm.modelFooter
}

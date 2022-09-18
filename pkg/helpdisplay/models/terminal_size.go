package models

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
)

// TerminalSize contains terminal size data
type TerminalSize struct {
	Width  rll.RowLenLimit
	Height size.Height
}

// MakeTerminalSize constructs a TerminalSize in a stack
func MakeTerminalSize(width rll.RowLenLimit, height size.Height) TerminalSize {
	return TerminalSize{
		Width:  width,
		Height: height,
	}
}

func (ts TerminalSize) IsValid() *dollyerr.Error {
	if !ts.Width.IsValid() {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			fmt.Errorf("PageModel.Update: invalid RowLenLimit: %v", ts.Width),
		)
	}

	return nil
}

// GetWidth - width field getter
func (ts TerminalSize) GetWidth() rll.RowLenLimit {
	return ts.Width
}

// GetHeight - height field getter
func (ts TerminalSize) GetHeight() size.Height {
	return ts.Height
}

// Clone does full clone of TerminalSizeObject
func (ts TerminalSize) Clone() TerminalSize {
	width := ts.GetWidth()
	return TerminalSize{
		Width:  width.Clone(),
		Height: ts.GetHeight(),
	}
}

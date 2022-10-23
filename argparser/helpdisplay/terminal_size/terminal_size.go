package terminal_size

import (
	"fmt"
	rll "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// TerminalSize contains terminal size page
type TerminalSize struct {
	widthLimit rll.RowLenLimit
	height     size.Height
}

// MakeTerminalSize constructs a TerminalSize in a stack
func MakeTerminalSize(widthLimit rll.RowLenLimit, height size.Height) TerminalSize {
	return TerminalSize{
		widthLimit: widthLimit,
		height:     height,
	}
}

// IsValid returns if the object is initialized
func (ts TerminalSize) IsValid() *dollyerr.Error {
	if !ts.widthLimit.IsValid() {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			fmt.Errorf("PageModel.update: invalid RowLenLimit: %d", ts.widthLimit),
		)
	}

	return nil
}

// GetWidthLimit - width field getter
func (ts TerminalSize) GetWidthLimit() rll.RowLenLimit {
	return ts.widthLimit
}

// GetHeight - height field getter
func (ts TerminalSize) GetHeight() size.Height {
	return ts.height
}

// CloneWithNewWidthLimit does a clone of the object with change widthLimit field
func (ts TerminalSize) CloneWithNewWidthLimit(widthLimit rll.RowLenLimit) TerminalSize {
	return TerminalSize{
		widthLimit: widthLimit,
		height:     ts.GetHeight(),
	}
}

// CloneWithNewHeight does a clone of the object with change height field
func (ts TerminalSize) CloneWithNewHeight(height size.Height) TerminalSize {
	width := ts.GetWidthLimit()
	return TerminalSize{
		widthLimit: width.Clone(),
		height:     height,
	}
}

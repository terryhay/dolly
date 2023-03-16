package terminal_size

import (
	"errors"

	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/tools/size"
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

// ErrIsValidWidthLimit - RowLenLimit.IsValid returned error
var ErrIsValidWidthLimit = errors.New(`TerminalSize.IsValid: RowLenLimit.IsValid returned error`)

// IsValid returns if the object is initialized
func (ts TerminalSize) IsValid() error {
	if err := ts.widthLimit.IsValid(); err != nil {
		return errors.Join(ErrIsValidWidthLimit, err)
	}

	return nil
}

// GetWidthLimit gets width field
func (ts TerminalSize) GetWidthLimit() rll.RowLenLimit {
	return ts.widthLimit
}

// GetHeight gets height field
func (ts TerminalSize) GetHeight() size.Height {
	return ts.height
}

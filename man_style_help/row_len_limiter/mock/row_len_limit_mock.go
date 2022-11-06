package rowLenLimitMock

import (
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/size"
)

// GetRowLenLimit returns RowLenLimit object by terminal width
func GetRowLenLimit(w size.Width) rll.RowLenLimit {
	limiter := rll.MakeRowLenLimiter()
	return limiter.GetRowLenLimit(w)
}

// GetRowLenLimitMin returns minimum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMin() rll.RowLenLimit {
	return GetRowLenLimit(rll.TerminalMinWidth)
}

// GetRowLenLimitForTerminalWidth25 returns RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method creates if terminal width value is 25
func GetRowLenLimitForTerminalWidth25() rll.RowLenLimit {
	return GetRowLenLimit(25)
}

// GetRowLenLimitMax returns maximum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMax() rll.RowLenLimit {
	return GetRowLenLimit(9999)
}

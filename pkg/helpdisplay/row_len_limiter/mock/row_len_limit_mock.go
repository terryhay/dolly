package rowLenLimitMock

import (
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
)

// GetRowLenLimitMin returns minimum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMin() rll.RowLenLimit {
	limiter := rll.MakeRowLenLimiter()
	return limiter.GetRowLenLimit(rll.TerminalMinWidth)
}

// GetRowLenLimitForTerminalWidth25 returns RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method creates if terminal width value is 25
func GetRowLenLimitForTerminalWidth25() rll.RowLenLimit {
	limiter := rll.MakeRowLenLimiter()
	return limiter.GetRowLenLimit(25)
}

// GetRowLenLimitMax returns maximum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMax() rll.RowLenLimit {
	limiter := rll.MakeRowLenLimiter()
	return limiter.GetRowLenLimit(9999)
}

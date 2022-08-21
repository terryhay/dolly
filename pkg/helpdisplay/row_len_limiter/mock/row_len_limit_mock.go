package rowLenLimitMock

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
)

// GetRowLenLimitMin returns minimum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMin() row_len_limiter.RowLenLimit {
	rll := row_len_limiter.MakeRowLenLimiter()
	return rll.GetRowLenLimit(row_len_limiter.TerminalMinWidth)
}

// GetRowLenLimitForTerminalWidth25 returns RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method creates if terminal width value is 25
func GetRowLenLimitForTerminalWidth25() row_len_limiter.RowLenLimit {
	rll := row_len_limiter.MakeRowLenLimiter()
	return rll.GetRowLenLimit(25)
}

// GetRowLenLimitMax returns maximum RowLenLimit object value
// which row_len_limiter.RowLenLimiter.GetRowLenLimit method can get
func GetRowLenLimitMax() row_len_limiter.RowLenLimit {
	rll := row_len_limiter.MakeRowLenLimiter()
	return rll.GetRowLenLimit(9999)
}

package row_len_limiter

import "github.com/terryhay/dolly/argparser/helpdisplay/size"

const (
	// TabSize - horizontal paragraph shift point size (space count)
	TabSize size.Width = 4

	// TerminalMinWidth - expected min size of a terminal
	TerminalMinWidth size.Width = 20

	calcRowLenLimitMinA = 0.5
	calcRowLenLimitMinB = 15. - (20. * 26. / 35.) + 0.5 // += 0.5 for rounding result

	calcRowLenLimitOptimumA = 0.7428571428571429
	calcRowLenLimitOptimumB = -0.8571428571428571 // += 0.5 for rounding result
)

type RowLenLimiter struct {
	lastTerminalWidth  size.Width
	rowLenLimit        RowLenLimit
	defaultRowLenLimit RowLenLimit
}

func MakeRowLenLimiter() RowLenLimiter {
	defaultRowLenLimit := MakeDefaultRowLenLimit()

	return RowLenLimiter{
		lastTerminalWidth: defaultRowLenLimit.Max(),

		rowLenLimit:        defaultRowLenLimit,
		defaultRowLenLimit: defaultRowLenLimit,
	}
}

func (i *RowLenLimiter) GetRowLenLimit(terminalWidth size.Width) RowLenLimit {
	if terminalWidth < defaultRowLenMax {
		if i.lastTerminalWidth == terminalWidth {
			return i.rowLenLimit
		}
		i.lastTerminalWidth = terminalWidth

		i.rowLenLimit.min = calcRowLenLimitMin(terminalWidth)
		i.rowLenLimit.optimum = calcRowLenLimitOptimum(terminalWidth)
		i.rowLenLimit.max = terminalWidth.ToInt()

		return i.rowLenLimit
	}
	return i.defaultRowLenLimit
}

func calcRowLenLimitMin(terminalWith size.Width) int {
	return int(calcRowLenLimitMinA*float64(terminalWith) + calcRowLenLimitMinB)
}

func calcRowLenLimitOptimum(terminalWith size.Width) int {
	return int(calcRowLenLimitOptimumA*float64(terminalWith) + calcRowLenLimitOptimumB)
}

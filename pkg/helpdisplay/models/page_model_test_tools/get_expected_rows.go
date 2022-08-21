package page_model_test_tools

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
)

const (
	TerminalWidth100  = 100
	TerminalWidth20   = row_len_limiter.TerminalMinWidth
	TerminalHeight7   = 7
	TerminalHeight100 = 100
)

type TestAction int

const (
	ActionShiftBackBy30    TestAction = -30
	ActionShiftBackBy7     TestAction = -7
	ActionShiftBackBy1     TestAction = -1
	ActionShiftBy0         TestAction = 0
	ActionShiftForwardBy1  TestAction = 1
	ActionShiftForwardBy5  TestAction = 5
	ActionShiftForwardBy7  TestAction = 7
	ActionShiftForwardBy30 TestAction = 30
	ActionShiftForwardBy50 TestAction = 50

	ActionResizeToMinWeight TestAction = 10000
	ActionResizeToMaxWeight TestAction = 20000
)

func GetExpectedData(action TestAction, terminalWidth, terminalHeight, shift int) (rows []string, newWidth, newHeight, newShift int) {
	allDisplayRows := getExpectedRowsByWidth(terminalWidth)

	if int(action) > 1000 {
		terminalWidth = getNewWidth(action)
		allDisplayRowsForNewWidth := getExpectedRowsByWidth(terminalWidth)
		shift = calcNewShift(allDisplayRows, shift, allDisplayRowsForNewWidth)

		allDisplayRows = allDisplayRowsForNewWidth
	} else {
		shift += int(action)
	}

	if len(allDisplayRows) < terminalHeight {
		return allDisplayRows, terminalWidth, terminalHeight, 0
	}

	checkSize := len(allDisplayRows) - shift
	if checkSize < terminalHeight {
		shift = len(allDisplayRows) - terminalHeight
	}
	if shift < 0 {
		shift = 0
	}

	res := make([]string, 0, terminalHeight)
	res = append(res, allDisplayRows[0])

	i := shift + 1
	for counter := 1; counter < terminalHeight; counter++ {
		res = append(res, allDisplayRows[i])
		i++
	}

	return res, terminalWidth, terminalHeight, shift
}

func getNewWidth(action TestAction) int {
	switch action {
	case ActionResizeToMinWeight:
		return TerminalWidth20
	case ActionResizeToMaxWeight:
		return TerminalWidth100
	default:
		panic(fmt.Sprintf("don't know how process an action %v", action))
	}
}

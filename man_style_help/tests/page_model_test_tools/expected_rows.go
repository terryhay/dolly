package page_model_test_tools

import (
	"fmt"

	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/tools/size"
)

const (
	// TerminalWidth100 - terminal width 100
	TerminalWidth100 = 100

	// TerminalWidth20 - terminal width 20
	TerminalWidth20 = rll.TerminalMinWidth

	// TerminalHeight7 - terminal width 7
	TerminalHeight7 = 7

	// TerminalHeight40 - terminal width 40
	TerminalHeight40 = 40

	// TerminalHeight100 - terminal width 100
	TerminalHeight100 = 100
)

// TestAction - type of test action (terminal window changing or scrolling)
type TestAction int

const (
	// ActionShiftUpBy30 - scroll terminal up by 30
	ActionShiftUpBy30 TestAction = -30

	// ActionShiftUpBy7 - scroll terminal up by 7
	ActionShiftUpBy7 TestAction = -7

	// ActionShiftUpBy1 - scroll terminal up by 1
	ActionShiftUpBy1 TestAction = -1

	// ActionShiftDownBy0 - scroll terminal down by 0
	ActionShiftDownBy0 TestAction = 0

	// ActionShiftDownBy1 - scroll terminal down by 1
	ActionShiftDownBy1 TestAction = 1

	// ActionShiftDownBy5 - scroll terminal down by 5
	ActionShiftDownBy5 TestAction = 5

	// ActionShiftDownBy7 - scroll terminal down by 7
	ActionShiftDownBy7 TestAction = 7

	// ActionShiftDownBy30 - scroll terminal down by 30
	ActionShiftDownBy30 TestAction = 30

	// ActionShiftDownBy50 - scroll terminal down by 50
	ActionShiftDownBy50 TestAction = 50

	// ActionResizeToMinWeight - resize terminal to min weight
	ActionResizeToMinWeight TestAction = 10000

	// ActionResizeToMaxWeight - resize terminal to max weight
	ActionResizeToMaxWeight TestAction = 20000
)

// ExpectedData returns expected rows for tests
func ExpectedData(action TestAction, terminalWidth size.Width, terminalHeight size.Height, shift int) (rows []string, newWidth size.Width, newHeight size.Height, newShift int) {
	allDisplayRows := getExpectedRowsByWidth(terminalWidth)

	if int(action) > 1000 {
		terminalWidth = getNewWidth(action)
		allDisplayRowsForNewWidth := getExpectedRowsByWidth(terminalWidth)
		shift = calcNewShift(allDisplayRows, shift, allDisplayRowsForNewWidth)

		allDisplayRows = allDisplayRowsForNewWidth
	} else {
		shift += int(action)
	}

	if len(allDisplayRows) < terminalHeight.Int() {
		return allDisplayRows, terminalWidth, terminalHeight, 0
	}

	checkSize := len(allDisplayRows) - shift
	if checkSize < terminalHeight.Int() {
		shift = len(allDisplayRows) - terminalHeight.Int()
	}
	if shift < 0 {
		shift = 0
	}

	res := make([]string, 0, terminalHeight)
	res = append(res, allDisplayRows[0])

	i := shift + 1
	for counter := 2; counter < terminalHeight.Int(); counter++ {
		res = append(res, allDisplayRows[i])
		i++
	}
	res = append(res, ":")

	return res, terminalWidth, terminalHeight, shift
}

func getNewWidth(action TestAction) size.Width {
	switch action {
	case ActionResizeToMinWeight:
		return TerminalWidth20
	case ActionResizeToMaxWeight:
		return TerminalWidth100
	default:
		panic(fmt.Sprintf("don't know how process an action %d", action))
	}
}

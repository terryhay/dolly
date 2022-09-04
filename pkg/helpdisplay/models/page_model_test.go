package models

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	tt "github.com/terryhay/dolly/pkg/helpdisplay/models/page_model_test_tools"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyPageViewShifting(t *testing.T) {
	t.Parallel()

	var pageModel PageModel
	require.Nil(t, pageModel.Shift(1, 1))
	require.Nil(t, pageModel.Shift(1, -1))
}

func TestPageViewShifting(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		defaultWidth   int
		defaultHeight  int
		actionSequence []tt.TestAction
	}{
		// check full windows
		{
			caseName: "full_window_with_width_max",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight100,
			actionSequence: []tt.TestAction{
				tt.ActionShiftBy0,
			},
		},
		{
			caseName: "full_window_with_width_min",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight100,
			actionSequence: []tt.TestAction{
				tt.ActionShiftBy0,
			},
		},

		// one shift cases
		{
			caseName: "rows_with_shift_0",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftBy0,
			},
		},
		{
			caseName: "rows_with_shift_0_and_min_width",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftBy0,
			},
		},
		{
			caseName: "rows_with_shift_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy1,
			},
		},
		{
			caseName: "rows_with_shift_minus_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftBackBy1,
			},
		},
		{
			caseName: "rows_with_shift_30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy30,
			},
		},

		// multi shift cases
		{
			caseName: "rows_with_shift_1_minus_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy1,
				tt.ActionShiftBackBy1,
			},
		},
		{
			caseName: "rows_with_shift_30_and_30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy30,
				tt.ActionShiftForwardBy30,
			},
		},
		{
			caseName: "rows_with_min_terminal_width_and_shift_1",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy1,
			},
		},
		{
			caseName: "rows_with_min_terminal_width_and_shift_7",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
			},
		},
		{
			caseName: "rows_with_shift_7_minus_8",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
				tt.ActionShiftBackBy1,
				tt.ActionShiftBackBy7,
			},
		},

		// resizing
		{
			caseName: "rows_with_resize_from_width100_to_width20",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
			},
		},
		{
			caseName: "rows_with_resize_from_width20_to_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "rows_with_resize_from_width20_to_width100_to_width20",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMaxWeight,
				tt.ActionResizeToMaxWeight,
			},
		},

		// multi shift with resizing
		{
			caseName: "rows_with_resize_to_width20_shift_to_forward7",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftForwardBy7,
			},
		},
		{
			caseName: "width10_to_shift_to_forward7_to_resize_width100",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward7_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward5_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy5,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward50_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy50,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "rows_with_resize_to_width20_shift_to_forward7_resize_to_width100_shift_to_back8",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftForwardBy7,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftBackBy7,
				tt.ActionShiftBackBy1,
			},
		},

		{
			caseName: "rows_with_resize_to_width20_shift_to_forward50_resize_to_width100_shift_to_back30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftForwardBy50,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftBackBy30,
				tt.ActionShiftBackBy1,
			},
		},
		{
			caseName: "...",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
				tt.ActionResizeToMinWeight,
				tt.ActionShiftForwardBy50,
			},
		},
		{
			caseName: "...",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftForwardBy7,
				tt.ActionResizeToMinWeight,
				tt.ActionShiftForwardBy50,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftBackBy30,
				tt.ActionShiftBackBy1,
			},
		},
	}

	var (
		expectedRows []string
		shift        int
		absShift     int
	)

	for _, td := range testData {
		t.Run(td.caseName+"_step_by_step_updating", func(t *testing.T) {
			terminalWidth := td.defaultWidth
			terminalHeight := td.defaultHeight

			pageModel := getPageModel(terminalWidth, terminalHeight)
			absShift = 0

			for actionIndex, action := range td.actionSequence {
				shift = 0

				oldTerminalWidth, _ := terminalWidth, terminalHeight
				expectedRows, terminalWidth, terminalHeight, absShift = tt.GetExpectedData(action, terminalWidth, terminalHeight, absShift)
				if action < tt.ActionResizeToMinWeight {
					shift = int(action)
				}

				if oldTerminalWidth != terminalWidth {
					added := 1
					if oldTerminalWidth > terminalWidth {
						added = -1
					}

					for width := oldTerminalWidth + added; width != terminalWidth; width += added {
						require.Nil(t, pageModel.Update(width, terminalHeight, 0))
					}
					require.Nil(t, pageModel.Update(terminalWidth, terminalHeight, 0))
				}

				if shift != 0 {
					added := 1
					if shift < 0 {
						added = -1
					}
					delta := absInt(shift)

					for i := 0; i < delta; i++ {
						require.Nil(t, pageModel.Update(terminalWidth, terminalHeight, added))
					}
				}

				checkRows(t, expectedRows, absShift, terminalWidth, actionIndex, shift, pageModel)
			}
		})

		t.Run(td.caseName+"_fast_updating", func(t *testing.T) {
			terminalWidth := td.defaultWidth
			terminalHeight := td.defaultHeight

			pageModel := getPageModel(terminalWidth, terminalHeight)
			absShift = 0

			for actionIndex, action := range td.actionSequence {
				shift = 0

				expectedRows, terminalWidth, terminalHeight, absShift = tt.GetExpectedData(action, terminalWidth, terminalHeight, absShift)
				if action < tt.ActionResizeToMinWeight {
					shift = int(action)
				}

				require.Nil(t, pageModel.Update(terminalWidth, terminalHeight, shift))

				checkRows(t, expectedRows, absShift, terminalWidth, actionIndex, shift, pageModel)
			}
		})
	}
}

func checkRows(t *testing.T, expectedRows []string, absShift, terminalWidth, actionIndex, shift int, pgm PageModel) {
	//pageModel.Update(terminalWidth, terminalHeight, shift)
	require.Equal(t, absShift, pgm.GetAnchorRowAbsolutelyIndex(),
		fmt.Sprintf("absolutely shift must be equal anchorRowAbsolutelyIndex. Action iter №%d; current shift: %d", actionIndex, shift))

	counter := 0
	for it := pgm.RowBegin(); !it.End(); it = pgm.RowNext(it) {
		require.True(t, counter < len(expectedRows),
			fmt.Sprintf("can't get expected row. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		row := rowToString(it.ShiftIndex, it.Cells)

		require.Equal(t, expectedRows[counter], row,
			fmt.Sprintf("rows must be equal. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		require.True(t, len([]rune(row)) <= terminalWidth, fmt.Sprintf("rune amount must be less or equal to terminal width. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		counter++
	}

	require.Equal(t, counter, len(expectedRows),
		fmt.Sprintf("must be checked all expected rows. Action iter №: %d; shift: %d", actionIndex, shift))
}

func getPageModel(terminalWidth, terminalHeight int) PageModel {
	return MakePageModel(
		data.Page{
			Header: "example",
			Paragraphs: []*data.Paragraph{
				{
					Text: `[1mNAME[0m`,
				},
				{
					Text:     `[1mexample[0m – shows how argtools generator works`,
					TabCount: 1,
				},
				{},

				{
					Text: `[1mSYNOPSIS[0m`,
				},
				{
					Text:     `[1mexample[0m [[1m-fl[0m [4mstr[0m [4m...[0m] [[1m-il[0m [4mstr[0m [4m...[0m] [[1m-sl[0m [4mstr[0m [4m...[0m]`,
					TabCount: 1,
				},
				{
					Text:     `[1mexample print[0m [[1m-checkargs[0m] [[1m-f[0m [4mstr[0m] [[1m-fl[0m [4mstr[0m [4m...[0m] [[1m-i[0m [4mstr[0m] [[1m-il[0m [4mstr[0m [4m...[0m] [[1m-s[0m [4mstr[0m] [[1m-sl[0m [4mstr[0m [4m...[0m]`,
					TabCount: 1,
				},
				{},
				{
					Text: `[1mDESCRIPTION[0m`,
				},
				{
					Text:     `you can write more detailed description here`,
					TabCount: 1,
				},
				{},

				{
					Text:     `and use several paragraphs`,
					TabCount: 1,
				},
				{},

				{
					Text: `The commands are as follows:`,
				},
				{
					Text:     `[1m<empty>[0m	checks arguments types`,
					TabCount: 1,
				},
				{},
				{
					Text:     `[1mprint[0m	print command line arguments with optional checking`,
					TabCount: 1,
				},
				{},

				{
					Text: `The flags are as follows:`,
				},
				{
					Text:     `[1m-checkargs[0m`,
					TabCount: 1,
				},
				{
					Text:     `do arguments checking`,
					TabCount: 2,
				},
				{},

				{
					Text:     `[1m-f[0m	single float`,
					TabCount: 1,
				},
				{},

				{
					Text:     `[1m-fl[0m	float list`,
					TabCount: 1,
				},
				{},

				{
					Text:     `[1m-i[0m	int string`,
					TabCount: 1,
				},
				{},

				{
					Text:     `[1m-il[0m	int list`,
					TabCount: 1,
				},
				{},

				{
					Text:     `[1m-s[0m	single string`,
					TabCount: 1,
				},
				{},

				{
					Text:     `[1m-sl[0m	string list`,
					TabCount: 1,
				},
			},
		},
		terminalWidth,
		terminalHeight,
	)
}

func TestErrors(t *testing.T) {
	t.Parallel()

	pageModel := getPageModel(tt.TerminalWidth20, tt.TerminalHeight100)
	{
		err := pageModel.Update(0, 0, 0)
		require.Equal(t, dollyerr.CodeHelpDisplayTerminalWidthLimitError, err.Code())
	}
	{
		err := pageModel.Shift(-1, 0)
		require.Equal(t, dollyerr.CodeHelpDisplayInvalidTerminalHeightArgument, err.Code())
	}
}

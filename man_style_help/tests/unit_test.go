package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	ri "github.com/terryhay/dolly/man_style_help/row_iterator"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	tt "github.com/terryhay/dolly/man_style_help/tests/page_model_test_tools"
	"github.com/terryhay/dolly/tools/size"
)

func TestPageViewShifting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		defaultWidth   size.Width
		defaultHeight  size.Height
		actionSequence []tt.TestAction
	}{
		// check full windows
		{
			caseName: "full_window_with_width_max",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight100,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy0,
			},
		},
		{
			caseName: "full_window_with_width_min",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight100,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy0,
			},
		},

		// one shift cases
		{
			caseName: "rows_with_shift_0",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy0,
			},
		},
		{
			caseName: "rows_with_shift_0_and_min_width",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy0,
			},
		},
		{
			caseName: "rows_with_shift_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy1,
			},
		},
		{
			caseName: "rows_with_shift_minus_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftUpBy1,
			},
		},
		{
			caseName: "rows_with_shift_30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy30,
			},
		},

		// multi shift cases
		{
			caseName: "rows_with_shift_1_minus_1",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy1,
				tt.ActionShiftUpBy1,
			},
		},
		{
			caseName: "rows_with_shift_30_and_30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy30,
				tt.ActionShiftDownBy30,
			},
		},
		{
			caseName: "rows_with_min_terminal_width_and_shift_1",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy1,
			},
		},
		{
			caseName: "rows_with_min_terminal_width_and_shift_7",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
			},
		},
		{
			caseName: "rows_with_shift_7_minus_8",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
				tt.ActionShiftUpBy1,
				tt.ActionShiftUpBy7,
			},
		},
		{
			caseName: "rows_with_shift_minus_7",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight40,
			actionSequence: []tt.TestAction{
				tt.ActionShiftUpBy7,
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
				tt.ActionResizeToMinWeight,
			},
		},

		// multi shift with resizing
		{
			caseName: "rows_with_resize_to_width20_shift_to_forward7",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftDownBy7,
			},
		},
		{
			caseName: "width10_to_shift_to_forward7_to_resize_width100",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward7_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward5_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy5,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "width20_to_shift_to_forward50_to_resize_width100",

			defaultWidth:  tt.TerminalWidth20,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy50,
				tt.ActionResizeToMaxWeight,
			},
		},
		{
			caseName: "rows_with_resize_to_width20_shift_to_forward7_resize_to_width100_shift_to_back8",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftDownBy7,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftUpBy7,
				tt.ActionShiftUpBy1,
			},
		},

		{
			caseName: "rows_with_resize_to_width20_shift_to_forward50_resize_to_width100_shift_to_back30",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionResizeToMinWeight,
				tt.ActionShiftDownBy50,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftUpBy30,
				tt.ActionShiftUpBy1,
			},
		},
		{
			caseName: "...",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
				tt.ActionResizeToMinWeight,
				tt.ActionShiftDownBy50,
			},
		},
		{
			caseName: "...",

			defaultWidth:  tt.TerminalWidth100,
			defaultHeight: tt.TerminalHeight7,
			actionSequence: []tt.TestAction{
				tt.ActionShiftDownBy7,
				tt.ActionResizeToMinWeight,
				tt.ActionShiftDownBy50,
				tt.ActionResizeToMaxWeight,
				tt.ActionShiftUpBy30,
				tt.ActionShiftUpBy1,
			},
		},
	}

	var (
		expectedRows []string
		shift        int
		absShift     int
	)

	for _, tc := range tests {
		t.Run(tc.caseName+"_step_by_step_updating", func(t *testing.T) {
			terminalWidth := tc.defaultWidth
			terminalHeight := tc.defaultHeight

			rowLenLimiter := rll.MakeRowLenLimiter()
			pageModel := getPageModel(rowLenLimiter.RowLenLimit(terminalWidth), terminalHeight)
			absShift = 0

			for actionIndex, action := range tc.actionSequence {
				shift = 0

				oldTerminalWidth, _ := terminalWidth, terminalHeight
				expectedRows, terminalWidth, terminalHeight, absShift = tt.ExpectedData(action, terminalWidth, terminalHeight, absShift)
				if action < tt.ActionResizeToMinWeight {
					shift = int(action)
				}

				if oldTerminalWidth != terminalWidth {
					added := 1
					if oldTerminalWidth > terminalWidth {
						added = -1
					}

					for width := oldTerminalWidth.Int() + added; width != terminalWidth.Int(); width += added {
						require.Nil(t, pageModel.Update(ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(size.MakeWidth(width)), terminalHeight), 0))
						// tt.PrintTerminalContentState(pageModel) // uncomment for debug
					}
					require.Nil(t, pageModel.Update(ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(terminalWidth), terminalHeight), 0))
				}

				if shift != 0 {
					added := 1
					if shift < 0 {
						added = -1
					}
					delta := absInt(shift)

					for i := 0; i < delta; i++ {
						require.Nil(t, pageModel.Update(ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(terminalWidth), terminalHeight), added))
					}
				}

				// tt.PrintTerminalContentState(pageModel) // uncomment for debug
				checkRows(t, expectedRows, absShift, terminalWidth, actionIndex, shift, pageModel)
			}
		})

		t.Run(tc.caseName+"_fast_updating", func(t *testing.T) {
			terminalWidth := tc.defaultWidth
			terminalHeight := tc.defaultHeight

			rowLenLimiter := rll.MakeRowLenLimiter()
			pageModel := getPageModel(rowLenLimiter.RowLenLimit(terminalWidth), terminalHeight)
			absShift = 0

			for actionIndex, action := range tc.actionSequence {
				shift = 0

				expectedRows, terminalWidth, terminalHeight, absShift = tt.ExpectedData(action, terminalWidth, terminalHeight, absShift)
				if action < tt.ActionResizeToMinWeight {
					shift = int(action)
				}

				require.Nil(t, pageModel.Update(ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(terminalWidth), terminalHeight), shift))

				checkRows(t, expectedRows, absShift, terminalWidth, actionIndex, shift, pageModel)
			}
		})
	}
}

func checkRows(t *testing.T, expectedRows []string, absShift int, terminalWidth size.Width, actionIndex, shift int, pgm *pgm.PageModel) {
	//pageModel.update(terminalWidth, terminalHeight, shift)
	require.Equal(t, absShift, pgm.GetAnchorRowAbsolutelyIndex().Int(),
		fmt.Sprintf("absolutely shift must be equal anchorRowAbsolutelyIndex. Action iter №%d; current shift: %d", actionIndex, shift))

	counter := 0
	for it := ri.MakeRowIterator(pgm); !it.End(); it.Next() {
		require.True(t, counter < len(expectedRows),
			fmt.Sprintf("can't get expected dynamic_row. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		r := it.RowModel().String()

		require.Equal(t, expectedRows[counter], r,
			fmt.Sprintf("rows must be equal. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		require.True(t, len([]rune(r)) <= terminalWidth.Int(), fmt.Sprintf("rune amount must be less or equal to terminal width. Action iter №: %d; shift: %d, iteration: %d", actionIndex, shift, counter))

		counter++
	}

	require.Equal(t, counter, len(expectedRows),
		fmt.Sprintf("must be checked all expected rows. Action iter №: %d; shift: %d", actionIndex, shift))
}

func getPageModel(rowLenLimit rll.RowLenLimit, terminalHeight size.Height) *pgm.PageModel {
	modelPage, _ := pgm.New(
		"example",
		hp.MakeBody([]hp.Row{
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("NAME")),
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk("example", hp.StyleTextBold), hp.MakeRowChunk(" – shows how argtools generator works")),
			{},
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("SYNOPSIS")),
			hp.MakeRow(size.WidthTab,
				hp.MakeRowChunk("example", hp.StyleTextBold),
				hp.MakeRowChunk(" ["), hp.MakeRowChunk("-fl", hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("str", hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("...", hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-il`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("str", hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`...`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk("-sl", hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("str", hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("...", hp.StyleTextUnderlined),
				hp.MakeRowChunk("]"),
			),
			hp.MakeRow(size.WidthTab,
				hp.MakeRowChunk(`example print [--checkargs]`, hp.StyleTextBold),
				hp.MakeRowChunk(" ["), hp.MakeRowChunk(`-f`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`str`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-fl`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`str`, hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`...`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-i`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`str`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-il`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("str", hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`...`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-s`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`str`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("] ["), hp.MakeRowChunk(`-sl`, hp.StyleTextBold),
				hp.MakeRowChunk(" "), hp.MakeRowChunk("str", hp.StyleTextUnderlined),
				hp.MakeRowChunk(" "), hp.MakeRowChunk(`...`, hp.StyleTextUnderlined),
				hp.MakeRowChunk("]"),
			),
			{},
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("DESCRIPTION")),
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk("you can write more detailed description here")),
			{},
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk("and use several paragraphs")),
			{},
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("The commands are as follows:")),
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`<empty>`, hp.StyleTextBold), hp.MakeRowChunk(" checks arguments types")),
			{},
			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`print`, hp.StyleTextBold), hp.MakeRowChunk("   print command line arguments with optional checking")),
			{},
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk(`The flags are as follows:`)),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`--checkargs`, hp.StyleTextBold)),
			hp.MakeRow(size.WidthTab+size.WidthTab+size.WidthTab, hp.MakeRowChunk(`do arguments checking`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-f`), hp.MakeRowChunk(`      single float`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-fl`), hp.MakeRowChunk(`     float list`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-i`), hp.MakeRowChunk(`      int string`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-il`), hp.MakeRowChunk(`     int list`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-s`), hp.MakeRowChunk(`      single string`)),
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk("")),

			hp.MakeRow(size.WidthTab, hp.MakeRowChunk(`-sl`), hp.MakeRowChunk(`     string list`)),
		}),
		ts.MakeTerminalSize(rowLenLimit, terminalHeight),
	)
	return modelPage
}

func TestErrors(t *testing.T) {
	t.Parallel()

	rowLenLimiter := rll.MakeRowLenLimiter()
	pageModel := getPageModel(rowLenLimiter.RowLenLimit(tt.TerminalWidth20), tt.TerminalHeight100)

	err := pageModel.Update(ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(0), 0), 0)
	require.ErrorIs(t, err, pgm.ErrUpdateInvalidTerminalSize)
}

// absInt returns absolutely value of v
func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

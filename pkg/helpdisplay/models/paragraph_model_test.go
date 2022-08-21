package models

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	rowLenLimitMock "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter/mock"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"strings"
	"testing"
)

func TestParagraphViewInit(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		sourceText  string
		tabCount    row_len_limiter.RowSize
		rowLenLimit row_len_limiter.RowLenLimit

		expectedRows []string
	}{
		{
			caseName:    "no_strings",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),
		},

		{
			caseName: "short_string",

			sourceText:  "short string",
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"short string",
			},
		},
		{
			caseName: "short_string_with_tab",

			sourceText:  "short string",
			tabCount:    1,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"    short string",
			},
		},
		{
			caseName: "size_for_terminal_width20_bug",

			sourceText: `and use several paragraphs`,
			// rune ruler:        |10       |20       |30       |40
			tabCount:    0,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),

			expectedRows: []string{
				"and use several",
				"paragraphs",
			},
		},
		{
			caseName: "", // todo: write case name

			sourceText: `and use several paragraphs`,
			// rune ruler:        |10       |20       |30       |40
			tabCount:    0,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),

			expectedRows: []string{
				"and use several",
				"paragraphs",
			},
		},

		// split into two rows
		{
			caseName: "split_text_before_optimum",

			sourceText: `You motherfucker, come on you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfucker,",
				"come on you little ass…",
			},
		},
		{
			caseName: "split_text_by_optimum",

			sourceText: `You motherfuckers, come on you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfuckers,",
				"come on you little ass…",
			},
		},
		{
			caseName: "split_text_after_optimum",

			sourceText: `motherfucker come on you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"motherfucker come",
				"on you little ass…",
			},
		},
		{
			caseName: "split_text_when_we_have_a_lot_of_split_variants",

			sourceText: `You motherfuckers, co e o 4you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfuckers,",
				"co e o 4you little ass…",
			},
		},
		{
			caseName: "split_text_when_we_do_not_have_spaces_in_split_interval",

			sourceText: `You motherfucker, comeonyou little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfucker,",
				"comeonyou little ass…",
			},
		},
		{
			caseName: "split_not_so_big_text_without_any_space",

			sourceText: `Youmotherfucker,comeonyoulittleass…`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"Youmotherfucker,comeonyou",
				"littleass…",
			},
		},

		// split into a lot of rows
		{
			caseName: "split_bit_text_without_any_space",

			sourceText: `Youmotherfucker,comeonyoulittleass…fuckwithme,eh?Youfuckinglittleasshole,dickheadcocksucker…Youfuckin'comeon,comefuckwithme!I'llgetyourass,youjerk!Oh,youfuckheadmotherfucker!Fuckallyouandyourfamily!Comeon,youcocksucker,slimebucket,shitfaceturdball!Comeon,youscumsucker,youfuckingwithme?Comeon,youasshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"Youmotherfucker,comeonyou",
				"littleass…fuckwithme,eh?Y",
				"oufuckinglittleasshole,di",
				"ckheadcocksucker…Youfucki",
				"n'comeon,comefuckwithme!I",
				"'llgetyourass,youjerk!Oh,",
				"youfuckheadmotherfucker!F",
				"uckallyouandyourfamily!Co",
				"meon,youcocksucker,slimeb",
				"ucket,shitfaceturdball!Co",
				"meon,youscumsucker,youfuc",
				"kingwithme?Comeon,youassh",
				"ole!",
			},
		},

		{
			caseName: "split_text_with_max_len_limit",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"You motherfucker, come on you little ass… fuck with me, eh? You",
				"fucking little asshole, dickhead cocksucker… You fuckin' come on,",
				"come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead",
				"motherfucker! Fuck all you and your family! Come on, you cocksucker,",
				"slime bucket, shitface turdball! Come on, you scum sucker, you",
				"fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_one_tab",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			tabCount:    1,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"    You motherfucker, come on you little ass… fuck with me, eh? You",
				"    fucking little asshole, dickhead cocksucker… You fuckin' come",
				"    on, come fuck with me! I'll get your ass, you jerk! Oh, you",
				"    fuckhead motherfucker! Fuck all you and your family! Come on,",
				"    you cocksucker, slime bucket, shitface turdball! Come on, you",
				"    scum sucker, you fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_two_tabs",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			tabCount:    2,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"        You motherfucker, come on you little ass… fuck with me, eh?",
				"        You fucking little asshole, dickhead cocksucker… You fuckin'",
				"        come on, come fuck with me! I'll get your ass, you jerk!",
				"        Oh, you fuckhead motherfucker! Fuck all you and your family!",
				"        Come on, you cocksucker, slime bucket, shitface turdball!",
				"        Come on, you scum sucker, you fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_three_tabs",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			tabCount:    3,
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"            You motherfucker, come on you little ass… fuck with me,",
				"            eh? You fucking little asshole, dickhead cocksucker…",
				"            You fuckin' come on, come fuck with me! I'll get your",
				"            ass, you jerk! Oh, you fuckhead motherfucker! Fuck all",
				"            you and your family! Come on, you cocksucker, slime",
				"            bucket, shitface turdball! Come on, you scum sucker,",
				"            you fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_min_len_limit",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split row ruler:            min|| |max
			rowLenLimit: rowLenLimitMock.GetRowLenLimitMin(),

			expectedRows: []string{
				"You motherfucker,",
				"come on you",
				"little ass…",
				"fuck with me,",
				"eh? You fucking",
				"little asshole,",
				"dickhead cocksucker",
				"… You fuckin'",
				"come on, come",
				"fuck with me!",
				"I'll get your",
				"ass, you jerk!",
				"Oh, you fuckhead",
				"motherfucker! Fuck",
				"all you and",
				"your family!",
				"Come on, you",
				"cocksucker, slime",
				"bucket, shitface",
				"turdball! Come",
				"on, you scum",
				"sucker, you",
				"fucking with",
				"me? Come on,",
				"you asshole!",
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			var p ParagraphModel
			p.Init(
				td.rowLenLimit,
				&data.Paragraph{
					Text:     td.sourceText,
					TabCount: td.tabCount,
				})

			if len(td.expectedRows) == 0 {
				require.Equal(t, 1, p.GetRowCount(),
					"expected rows count must be equal to paragraph row count")
				return
			}

			require.Equal(t, len(td.expectedRows), p.GetRowCount(),
				"expected rows count must be equal to paragraph row count")

			for i := 0; i < p.GetRowCount(); i++ {
				require.True(t, len([]rune(td.expectedRows[i])) < td.rowLenLimit.Max().ToInt()+1,
					fmt.Sprintf("row len is more than max limit = %d", td.rowLenLimit.Max()))

				require.Equal(t, td.expectedRows[i], rowToString(p.GetRow(i)))
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("resizing_with_not_null_anchor", func(t *testing.T) {
		rowLenLimit := rowLenLimitMock.GetRowLenLimitMin()
		var prm ParagraphModel
		prm.Init(
			rowLenLimit,
			&data.Paragraph{
				Text: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…`,
			})

		require.Equal(t, 8, prm.GetRowCount())
		require.True(t, prm.ShiftAnchorRow(4))

		prm.Update(rowLenLimitMock.GetRowLenLimitMax())
	})

	t.Run("resizing_with_not_null_anchor", func(t *testing.T) {
		rowLenLimit := rowLenLimitMock.GetRowLenLimitMin()
		var prm ParagraphModel
		prm.Init(
			rowLenLimit,
			&data.Paragraph{
				Text: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…`,
			})

		require.Equal(t, 8, prm.GetRowCount())
		require.True(t, prm.ShiftAnchorRow(4))

		prm.Update(rowLenLimitMock.GetRowLenLimitMin())
	})
}

func TestSecondUpdate(t *testing.T) {
	t.Parallel()

	// usingRowLeLimit contains hardcode values
	// which creates row_len_limiter.RowLenLimiter.GetRowLenLimit method
	// if terminal width value is 25
	rowLenLimit := rowLenLimitMock.GetRowLenLimitMin()

	var pr ParagraphModel
	pr.Init(
		rowLenLimit,
		&data.Paragraph{Text: `[1mexample[0m – shows how argtools generator works`})

	rowLenLimit = row_len_limiter.MakeRowLenLimit(25, 29, 33)

	pr.Update(rowLenLimit)
	updatedPr := pr

	pr.Update(rowLenLimit)
	require.Equal(t, updatedPr, pr)
}

func TestGetRow(t *testing.T) {
	t.Parallel()

	t.Run("call_with_invalid_index", func(t *testing.T) {
		var prm ParagraphModel
		prm.Init(
			rowLenLimitMock.GetRowLenLimitMax(),
			&data.Paragraph{Text: `[1mexample[0m – shows how argtools generator works`})

		row, cells := prm.GetRow(prm.GetRowCount())
		require.Equal(t, 0, row)
		require.Nil(t, cells)
	})
}

func rowToString(shiftIndex int, cells []termbox.Cell) string {
	builder := strings.Builder{}
	builder.Reset()

	for i := 0; i < shiftIndex; i++ {
		builder.WriteRune(runes.RuneSpace)
	}
	for _, cell := range cells {
		builder.WriteRune(cell.Ch)
	}

	return builder.String()
}

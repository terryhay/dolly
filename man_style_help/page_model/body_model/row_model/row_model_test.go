package row_model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

func TestRowModel(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var prm *RowModel
		require.Equal(t, index.Zero, prm.GetAnchorRowIndex())
		prm.SetAnchorRowIndex(index.Zero)
		require.Equal(t, size.HeightZero, prm.GetRowCount())
		require.Equal(t, size.WidthZero, prm.GetTabInCells())
		require.False(t, prm.ShiftAnchorRow(int(gofakeit.Int32())))
		prm.SetBackRowAsAnchor()
		require.Equal(t, row.Row{}, prm.GetRow(index.Zero))
		require.Equal(t, size.HeightZero, prm.Update(rllMock.GetRowLenLimitMin()))
	})
}

func TestRowModelInit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		sourceText  string
		marginLeft  size.Width
		rowLenLimit rll.RowLenLimit

		expectedRows []string
	}{
		{
			caseName:    "no_strings_no_limit",
			rowLenLimit: rll.RowLenLimit{},
		},
		{
			caseName:    "no_strings",
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),
		},

		{
			caseName: "short_string",

			sourceText:  "short string",
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"short string",
			},
		},
		{
			caseName: "short_string_with_tab",

			sourceText:  "short string",
			marginLeft:  size.WidthTab,
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"    short string",
			},
		},
		{
			caseName: "size_for_terminal_width20_bug",

			sourceText: `and use several paragraphs`,
			// rune ruler:        |10       |20       |30       |40
			rowLenLimit: rllMock.GetRowLenLimitMin(),

			expectedRows: []string{
				"and use several",
				"paragraphs",
			},
		},
		{
			caseName: "", // todo: write case name

			sourceText: `and use several paragraphs`,
			// rune ruler:        |10       |20       |30       |40
			rowLenLimit: rllMock.GetRowLenLimitMin(),

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
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfucker,",
				"come on you little ass…",
			},
		},
		{
			caseName: "split_text_by_optimum",

			sourceText: `You motherfuckers, come on you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfuckers,",
				"come on you little ass…",
			},
		},
		{
			caseName: "split_text_after_optimum",

			sourceText: `motherfucker come on you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"motherfucker come",
				"on you little ass…",
			},
		},
		{
			caseName: "split_text_when_we_have_a_lot_of_split_variants",

			sourceText: `You motherfuckers, co e o 4you little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfuckers,",
				"co e o 4you little ass…",
			},
		},
		{
			caseName: "split_text_when_we_do_not_have_spaces_in_split_interval",

			sourceText: `You motherfucker, comeonyou little ass…`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

			expectedRows: []string{
				"You motherfucker,",
				"comeonyou little ass…",
			},
		},
		{
			caseName: "split_not_so_big_text_without_any_space",

			sourceText: `Youmotherfucker,comeonyoulittleass…`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

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
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitForTerminalWidth25(),

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
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"You motherfucker, come on you little ass… fuck with me, eh? You",
				"fucking little asshole, dickhead cocksucker… You fuckin' come on,",
				"come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead",
				"motherfucker! Fuck all you and your family! Come on, you cocksucker,",
				"slime bucket, shitface turdball! Come on, you scum sucker, you fucking",
				"with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_one_tab",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			marginLeft:  size.WidthTab,
			rowLenLimit: rllMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"    You motherfucker, come on you little ass… fuck with me, eh? You",
				"    fucking little asshole, dickhead cocksucker… You fuckin' come on,",
				"    come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead",
				"    motherfucker! Fuck all you and your family! Come on, you cocksucker,",
				"    slime bucket, shitface turdball! Come on, you scum sucker, you fucking",
				"    with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_two_tabs",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			marginLeft:  size.WidthTab + size.WidthTab,
			rowLenLimit: rllMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"        You motherfucker, come on you little ass… fuck with me, eh?",
				"        You fucking little asshole, dickhead cocksucker… You fuckin'",
				"        come on, come fuck with me! I'll get your ass, you jerk! Oh,",
				"        you fuckhead motherfucker! Fuck all you and your family! Come",
				"        on, you cocksucker, slime bucket, shitface turdball! Come on,",
				"        you scum sucker, you fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_max_len_limit_and_three_tabs",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			marginLeft:  size.WidthTab + size.WidthTab + size.WidthTab,
			rowLenLimit: rllMock.GetRowLenLimitMax(),

			expectedRows: []string{
				"            You motherfucker, come on you little ass… fuck with me,",
				"            eh? You fucking little asshole, dickhead cocksucker… You",
				"            fuckin' come on, come fuck with me! I'll get your ass, you",
				"            jerk! Oh, you fuckhead motherfucker! Fuck all you and your",
				"            family! Come on, you cocksucker, slime bucket, shitface",
				"            turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!",
			},
		},
		{
			caseName: "split_text_with_min_len_limit",

			sourceText: `You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker… You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker! Fuck all you and your family! Come on, you cocksucker, slime bucket, shitface turdball! Come on, you scum sucker, you fucking with me? Come on, you asshole!`,
			// rune ruler:        |10       |20       |30       |40
			// split dynamic_row ruler:            min|| |max
			rowLenLimit: rllMock.GetRowLenLimitMin(),

			expectedRows: []string{
				"You motherfucker,",
				"come on you little",
				"ass… fuck with me,",
				"eh? You fucking",
				"little asshole,",
				"dickhead cocksucker…",
				"You fuckin' come",
				"on, come fuck with",
				"me! I'll get your",
				"ass, you jerk! Oh,",
				"you fuckhead",
				"motherfucker!",
				"Fuck all you",
				"and your family!",
				"Come on, you",
				"cocksucker,",
				"slime bucket,",
				"shitface turdball!",
				"Come on, you scum",
				"sucker, you fucking",
				"with me? Come on,",
				"you asshole!",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			prm := NewRowModel(
				hp.MakeRow(tc.marginLeft, hp.MakeRowChunk(tc.sourceText)),
				tc.rowLenLimit,
			)

			if len(tc.expectedRows) == 0 {
				require.Equal(t, 1, prm.GetRowCount().Int(),
					"expected rows count must be equal to paragraph dynamic_row count")
				return
			}

			require.Equal(t, len(tc.expectedRows), prm.GetRowCount().Int(),
				"expected rows count must be equal to paragraph dynamic_row count")

			for i := index.Zero; i.Int() < prm.GetRowCount().Int(); i++ {
				require.True(t, len([]rune(tc.expectedRows[i])) < tc.rowLenLimit.Max().Int()+1,
					fmt.Sprintf("dynamic_row len is more than max limit = %d", tc.rowLenLimit.Max()))

				require.Equal(t, tc.expectedRows[i], rowToString(prm.GetRow(i)))
			}
		})
	}
}

func TestRowModelUpdate(t *testing.T) {
	t.Parallel()

	t.Run("resizing_with_anchor_4", func(t *testing.T) {
		t.Parallel()

		prm := NewRowModel(
			hp.MakeRow(size.WidthZero, hp.MakeRowChunk(`You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…`)),
			rllMock.GetRowLenLimitMin(),
		)

		expRowCountBefore := size.MakeHeight(6)
		expRowCountAfter := size.MakeHeight(2)

		require.Equal(t, expRowCountBefore, prm.GetRowCount())
		require.True(t, prm.ShiftAnchorRow(4))

		prm.Update(rllMock.GetRowLenLimitMax())
		require.Equal(t, expRowCountAfter, prm.GetRowCount())

		require.LessOrEqual(t, expRowCountAfter, expRowCountBefore)
	})

	t.Run("resizing_loop_error_case", func(t *testing.T) {
		t.Parallel()

		rowLenLimitBefore := rllMock.GetRowLenLimit(23)

		expRowCountBefore := size.MakeHeight(4)
		expRowCountAfter := size.MakeHeight(1)

		prm := NewRowModel(hp.MakeRow(size.WidthTab,
			hp.MakeRowChunk(`print`, hp.StyleTextBold),
			hp.MakeRowChunk(`	print command line arguments with optional checking`),
		), rowLenLimitBefore)

		require.Equal(t, expRowCountBefore, prm.GetRowCount())
		require.True(t, prm.ShiftAnchorRow(0))

		prm.Update(rllMock.GetRowLenLimitMax())
		require.Equal(t, expRowCountAfter, prm.GetRowCount())

		require.LessOrEqual(t, expRowCountAfter, expRowCountBefore)
	})
}

func TestSecondUpdate(t *testing.T) {
	t.Parallel()

	// usingRowLenLimit contains hardcode values
	// which creates row_len_limiter.RowLenLimiter.RowLenLimit method
	// if terminal width value is 25
	rowLenLimit := rllMock.GetRowLenLimitMin()

	prm := NewRowModel(
		hp.MakeRow(size.WidthZero,
			hp.MakeRowChunk(`example`, hp.StyleTextBold),
			hp.MakeRowChunk(` – shows how argtools generator works`),
		),
		rowLenLimit,
	)

	rowLenLimit = rll.MakeRowLenLimit(25, 29, 33)

	prm.Update(rowLenLimit)
	updatedPr := prm

	prm.Update(rowLenLimit)
	require.Equal(t, updatedPr, prm)
}

func TestGetRow(t *testing.T) {
	t.Parallel()

	t.Run("call_with_invalid_index", func(t *testing.T) {
		prm := NewRowModel(
			hp.MakeRow(size.WidthZero,
				hp.MakeRowChunk(`example`, hp.StyleTextBold),
				hp.MakeRowChunk(` – shows how argtools generator works`),
			),
			rllMock.GetRowLenLimitMax(),
		)

		r := prm.GetRow(index.Index(prm.GetRowCount()))
		require.Equal(t, 0, r.GetShiftIndex().Int())
		require.Nil(t, r.GetCells())
	})

	t.Run("anchor", func(t *testing.T) {
		prm := NewRowModel(
			hp.MakeRow(size.WidthZero,
				hp.MakeRowChunk(`example`, hp.StyleTextBold),
				hp.MakeRowChunk(` – shows how argtools generator works`),
			),
			rllMock.GetRowLenLimitMin(),
		)

		prm.SetAnchorRowIndex(1)
		require.Equal(t, index.MakeIndex(1), prm.GetAnchorRowIndex())
		prm.SetBackRowAsAnchor()
		prm.ShiftAnchorRow(1)
		require.Equal(t, index.Zero, prm.GetAnchorRowIndex())
	})
}

func rowToString(r row.Row) string {
	builder := strings.Builder{}
	builder.Reset()

	for i := size.WidthZero; i < r.GetShiftIndex(); i++ {
		builder.WriteRune(runes.RuneSpace)
	}
	for _, cell := range r.GetCells() {
		builder.WriteRune(cell.Ch)
	}

	return builder.String()
}

package body_model

import (
	"testing"

	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

func TestBodyModel(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var bdm *BodyModel

		require.Equal(t, size.MakeHeight(0), bdm.GetRowCount())
		require.Equal(t, index.Zero, bdm.GetAnchorRowModelIndex())
		require.Equal(t, index.Zero, bdm.GetAnchorRowIndex())
		require.Equal(t, index.Zero, bdm.GetAnchorRowAbsolutelyIndex())
		require.Nil(t, bdm.GetRowModel(index.Zero))
		require.NotPanics(t, func() { bdm.Update(ts.TerminalSize{}, 0) })
		require.NotPanics(t, func() { bdm.Shift(size.MakeHeight(0), 0) })
	})

	t.Run("errors", func(t *testing.T) {
		bm := NewBodyModel(hp.Body{}, ts.TerminalSize{})
		require.Nil(t, bm.GetRowModel(0))
		//require.NotNil(t, bm.Update(ts.TerminalSize{}, 0))
	})

	t.Run("simple", func(t *testing.T) {
		bdm := NewBodyModel(
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
				),
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				),
			}),
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(10)),
		)

		require.Equal(t, size.MakeHeight(12), bdm.GetRowCount())
		require.Equal(t, index.Zero, bdm.GetAnchorRowModelIndex())
		require.Equal(t, index.Zero, bdm.GetAnchorRowIndex())
		require.Equal(t, index.Zero, bdm.GetAnchorRowAbsolutelyIndex())
		require.NotNil(t, bdm.GetRowModel(index.Zero))
		require.Nil(t, bdm.GetRowModel(index.MakeIndex(bdm.GetRowCount().Int()+1)))

		//require.Nil(t, bdm.Shift(size.MakeHeight(10), 0))
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("with_changed_terminal_size", func(t *testing.T) {
		bm := NewBodyModel(
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
				),
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				),
			}),
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(10)),
		)

		bm.Update(ts.MakeTerminalSize(rllMock.GetRowLenLimitMax(), size.MakeHeight(10)), 10)

		expectedRows := [...]string{
			"    You motherfucker, come on you little ass… fuck with me, eh? You",
			"    fucking little asshole, dickhead cocksucker…",
			"    You fuckin' come on, come fuck with me! I'll get your ass, you",
			"    jerk! Oh, you fuckhead motherfucker!",
		}
		ei := 0

		for p := 0; p < bm.GetRowCount().Int(); p++ {
			prm := bm.GetRowModel(index.MakeIndex(p))
			for i := index.Zero; i < prm.GetRowCount().Index(); i++ {
				require.Equal(t, expectedRows[ei], prm.GetRow(i).String())
				ei++
			}
		}
	})

	t.Run("update_after_shifting", func(t *testing.T) {
		t.Parallel()

		bm := NewBodyModel(
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
				),
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				),
			}),
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(50)),
		)

		bm.Shift(size.MakeHeight(3), 50)
		bm.Update(ts.MakeTerminalSize(rllMock.GetRowLenLimitMax(), size.MakeHeight(3)), 0)

		expectedRows := [...]string{
			"    You motherfucker, come on you little ass… fuck with me, eh? You",
			"    fucking little asshole, dickhead cocksucker…",
			"    You fuckin' come on, come fuck with me! I'll get your ass, you",
			"    jerk! Oh, you fuckhead motherfucker!",
		}
		ei := 0

		for p := 0; p < bm.GetRowCount().Int(); p++ {
			prm := bm.GetRowModel(index.MakeIndex(p))
			for i := index.Zero; i < prm.GetRowCount().Index(); i++ {
				require.Equal(t, expectedRows[ei], prm.GetRow(i).String())
				ei++
			}
		}
	})
}

func TestShift(t *testing.T) {
	t.Parallel()

	t.Run("short_text_shifting", func(t *testing.T) {
		terminalSize := ts.MakeTerminalSize(
			rll.MakeDefaultRowLenLimit(),
			10,
		)
		bdm := NewBodyModel(hp.MakeBody([]hp.Row{
			hp.MakeRow(size.WidthZero,
				hp.MakeRowChunk("1"),
			),
			hp.MakeRow(size.WidthZero,
				hp.MakeRowChunk("2"),
			),
		}),
			terminalSize,
		)
		{
			bdm.Shift(terminalSize.GetHeight(), 2)
			checkEqual(t, []string{"1", "2"}, bdm)
		}
		{
			bdm.Shift(terminalSize.GetHeight(), -2)
			checkEqual(t, []string{"1", "2"}, bdm)
		}
		{
			bdm.Shift(terminalSize.GetHeight(), 0)
			checkEqual(t, []string{"1", "2"}, bdm)
		}
	})

	t.Run("long_text_shifting", func(t *testing.T) {
		bdm := NewBodyModel(
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
				),
				hp.MakeRow(size.WidthTab,
					hp.MakeRowChunk("You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				),
			}),
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(1)),
		)

		checkStrings := []string{
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
		}
		{
			bdm.Shift(size.MakeHeight(1), 15)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), 2)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), 2)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), -2)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), 15)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), -1)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), -5)
			checkEqual(t, checkStrings, bdm)
		}
		{
			bdm.Shift(size.MakeHeight(1), -15)
			checkEqual(t, checkStrings, bdm)
		}
	})
}

func checkEqual(t *testing.T, expected []string, bdm *BodyModel) {
	i := 0
	for p := 0; p < bdm.GetRowCount().Int(); p++ {
		prm := bdm.GetRowModel(index.MakeIndex(p))

		for r := 0; r < prm.GetRowCount().Int(); r++ {
			require.Equal(t, expected[i], prm.GetRow(index.MakeIndex(r)).String())
			i++
		}
	}
}

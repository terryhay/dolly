package body_model

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/man_style_help/index"
	"github.com/terryhay/dolly/man_style_help/page"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	"github.com/terryhay/dolly/man_style_help/size"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"testing"
)

func TestBodyModel(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var bdm *BodyModel

		require.Equal(t, size.Height(0), bdm.GetRowCount())
		require.Equal(t, index.Null, bdm.GetAnchorParagraphIndex())
		require.Equal(t, index.Null, bdm.GetAnchorParagraphRowIndex())
		require.Equal(t, index.Null, bdm.GetAnchorRowAbsolutelyIndex())
		require.Nil(t, bdm.GetParagraph(index.Null))
		require.Nil(t, bdm.Update(ts.TerminalSize{}, 0))
		require.Nil(t, bdm.Shift(size.Height(0), 0))
	})

	t.Run("errors", func(t *testing.T) {
		bm := NewBodyModel(page.Page{}, ts.TerminalSize{})
		require.Nil(t, bm.GetParagraph(0))
		require.NotNil(t, bm.Update(ts.TerminalSize{}, 0))
	})

	t.Run("simple", func(t *testing.T) {
		bdm := NewBodyModel(
			page.Page{
				Header: page.MakeParagraph(0, "header"),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(1, "You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
					page.MakeParagraph(1, "You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				},
			},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(10)))
		require.Equal(t, size.Height(12), bdm.GetRowCount())
		require.Equal(t, index.Null, bdm.GetAnchorParagraphIndex())
		require.Equal(t, index.Null, bdm.GetAnchorParagraphRowIndex())
		require.Equal(t, index.Null, bdm.GetAnchorRowAbsolutelyIndex())
		require.NotNil(t, bdm.GetParagraph(index.Null))
		require.Nil(t, bdm.GetParagraph(index.MakeIndex(bdm.GetRowCount().ToInt()+1)))

		require.Nil(t, bdm.Shift(size.Height(10), 0))
	})

	t.Run("updating", func(t *testing.T) {
		bdm := NewBodyModel(
			page.Page{
				Header: page.MakeParagraph(0, "header"),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(1, "You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
					page.MakeParagraph(1, "You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				},
			},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(10)),
		)

		require.Nil(t, bdm.Update(ts.MakeTerminalSize(rllMock.GetRowLenLimitMax(), size.Height(10)), 10))

		expectedRows := []string{
			"    You motherfucker, come on you little ass… fuck with me, eh? You",
			"    fucking little asshole, dickhead cocksucker…",
			"    You fuckin' come on, come fuck with me! I'll get your ass, you",
			"    jerk! Oh, you fuckhead motherfucker!",
		}
		ei := 0

		for p := 0; p < bdm.GetRowCount().ToInt(); p++ {
			prm := bdm.GetParagraph(index.MakeIndex(p))
			for r := 0; r < prm.GetRowCount().ToInt(); r++ {
				require.Equal(t, expectedRows[ei], prm.GetRow(index.MakeIndex(r)).String())
				ei++
			}
		}
	})
}

func TestBodyModelShifting(t *testing.T) {
	t.Parallel()

	t.Run("short_text_shifting", func(t *testing.T) {
		terminalSize := ts.MakeTerminalSize(
			rll.MakeDefaultRowLenLimit(),
			10,
		)
		bdm := NewBodyModel(page.Page{
			Paragraphs: []page.Paragraph{
				page.MakeParagraph(0, "1"),
				page.MakeParagraph(0, "2"),
			},
		},
			terminalSize,
		)
		{
			require.Nil(t, bdm.Shift(terminalSize.GetHeight(), 2))
			checkEqual(t, []string{"1", "2"}, bdm)
		}
		{
			require.Nil(t, bdm.Shift(terminalSize.GetHeight(), -2))
			checkEqual(t, []string{"1", "2"}, bdm)
		}
	})

	t.Run("long_text_shifting", func(t *testing.T) {
		bdm := NewBodyModel(
			page.Page{
				Header: page.MakeParagraph(0, "header"),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(1, "You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
					page.MakeParagraph(1, "You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				},
			},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(1)),
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
			require.Nil(t, bdm.Shift(size.Height(1), 15))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), 2))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), 2))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), -2))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), 15))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), -1))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), -5))
			checkEqual(t, checkStrings, bdm)
		}
		{
			require.Nil(t, bdm.Shift(size.Height(1), -15))
			checkEqual(t, checkStrings, bdm)
		}
	})
}

func checkEqual(t *testing.T, expected []string, bdm *BodyModel) {
	i := 0
	for p := 0; p < bdm.GetRowCount().ToInt(); p++ {
		prm := bdm.GetParagraph(index.MakeIndex(p))

		for r := 0; r < prm.GetRowCount().ToInt(); r++ {
			require.Equal(t, expected[i], prm.GetRow(index.MakeIndex(r)).String())
			i++
		}
	}
}

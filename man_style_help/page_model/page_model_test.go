package page_model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/man_style_help/index"
	"github.com/terryhay/dolly/man_style_help/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	"github.com/terryhay/dolly/man_style_help/size"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestPageModel(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointers", func(t *testing.T) {
		var pageModel *PageModel

		require.Equal(t, 0, pageModel.GetAnchorRowAbsolutelyIndex().ToInt())
		require.Error(t, pageModel.Update(ts.TerminalSize{}, 0).Error())
		require.Error(t, pageModel.Shift(0, 0).Error())
		require.Equal(t, ts.TerminalSize{}, pageModel.GetUsingTerminalSize())

		require.Equal(t, 0, pageModel.GetRowCount().ToInt())

		err := pageModel.GetHeaderModel().Update(ts.MakeTerminalSize(rll.MakeRowLenLimit(1, 1, 1), 1))
		require.Nil(t, err)
		require.Equal(t, row.Row{}, pageModel.GetHeaderModel().GetViewRow())

		require.Equal(t, 0, pageModel.GetBodyModel().GetRowCount().ToInt())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorParagraphIndex().ToInt())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorParagraphRowIndex().ToInt())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorRowAbsolutelyIndex().ToInt())
		require.Nil(t, pageModel.GetBodyModel().GetParagraph(0))
		require.Nil(t, err)
		require.Nil(t, pageModel.GetBodyModel().Update(ts.TerminalSize{}, 0))
		require.Nil(t, pageModel.GetBodyModel().Shift(0, 0))

		require.Equal(t, row.Row{}, pageModel.GetFooterModel().GetFooterRow())
	})

	t.Run("initialization_errors", func(t *testing.T) {
		pgm, err := NewPageModel(page.Page{}, ts.TerminalSize{})
		require.Nil(t, pgm)
		require.NotNil(t, err)
	})

	t.Run("simple", func(t *testing.T) {
		pgm, err := NewPageModel(
			page.Page{
				Header: page.MakeParagraph(0, "header"),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(1, "You motherfucker, come on you little ass… fuck with me, eh? You fucking little asshole, dickhead cocksucker…"),
					page.MakeParagraph(1, "You fuckin' come on, come fuck with me! I'll get your ass, you jerk! Oh, you fuckhead motherfucker!"),
				},
			},
			ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(10)),
		)
		require.Nil(t, err)

		require.Equal(t, index.Null, pgm.GetAnchorRowAbsolutelyIndex())
		require.Equal(t, dollyerr.CodeHelpDisplayTerminalWidthLimitError,
			pgm.Update(ts.TerminalSize{}, 0).Code())
		require.NotNil(t, pgm.Update(ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(2)), 0).Code())

		assert.Equal(t, ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.Height(2)), pgm.GetUsingTerminalSize())
		require.Nil(t, pgm.Shift(size.Height(20), 0))

		require.Equal(t, size.Height(14), pgm.GetRowCount())
		require.NotNil(t, pgm.GetHeaderModel())
		require.NotNil(t, pgm.GetBodyModel())
		require.NotNil(t, pgm.GetFooterModel())
	})
}

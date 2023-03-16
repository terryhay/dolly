package page_model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	"github.com/terryhay/dolly/man_style_help/row"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	"github.com/terryhay/dolly/tools/index"
	"github.com/terryhay/dolly/tools/size"
)

func TestPageModel(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointers", func(t *testing.T) {
		var pageModel *PageModel

		require.Equal(t, 0, pageModel.GetAnchorRowAbsolutelyIndex().Int())
		require.NoError(t, pageModel.Update(ts.TerminalSize{}, 0))
		pageModel.Shift(0, 0)
		require.Equal(t, ts.TerminalSize{}, pageModel.GetUsingTermSize())

		require.Equal(t, 0, pageModel.GetRowCount().Int())

		pageModel.GetHeaderModel().Update(ts.MakeTerminalSize(rll.MakeRowLenLimit(1, 1, 1), 1))
		require.Equal(t, row.Row{}, pageModel.GetHeaderModel().GetViewRow())

		require.Equal(t, 0, pageModel.GetBodyModel().GetRowCount().Int())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorRowModelIndex().Int())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorRowIndex().Int())
		require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorRowAbsolutelyIndex().Int())
		require.Nil(t, pageModel.GetBodyModel().GetRowModel(0))

		pageModel.GetBodyModel().Update(ts.TerminalSize{}, 0)
		pageModel.GetBodyModel().Shift(0, 0)

		require.Equal(t, row.Row{}, pageModel.GetFooterModel().GetFooterRow())
	})

	t.Run("initialization_errors", func(t *testing.T) {
		pgm, err := New("", hp.Body{}, ts.TerminalSize{})
		require.Nil(t, pgm)
		require.NotNil(t, err)
	})

	t.Run("simple", func(t *testing.T) {
		pgm, err := New(
			"header",
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
		require.Nil(t, err)

		require.Equal(t, index.Zero, pgm.GetAnchorRowAbsolutelyIndex())
		require.ErrorIs(t, pgm.Update(ts.TerminalSize{}, 0), ErrUpdateInvalidTerminalSize)
		require.Nil(t, pgm.Update(ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(2)), 0))

		assert.Equal(t, ts.MakeTerminalSize(rllMock.GetRowLenLimitMin(), size.MakeHeight(2)), pgm.GetUsingTermSize())
		pgm.Shift(size.MakeHeight(20), 0)

		require.Equal(t, size.MakeHeight(14), pgm.GetRowCount())
		require.NotNil(t, pgm.GetHeaderModel())
		require.NotNil(t, pgm.GetBodyModel())
		require.NotNil(t, pgm.GetFooterModel())
	})
}

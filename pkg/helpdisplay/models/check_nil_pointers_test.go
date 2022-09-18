package models

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/helpdisplay/row"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"testing"
)

func TestCheckNilPointers(t *testing.T) {
	t.Parallel()

	var pageModel *PageModel

	require.Equal(t, 0, pageModel.GetAnchorRowAbsolutelyIndex().ToInt())
	require.Error(t, pageModel.Update(TerminalSize{}, 0))
	require.Error(t, pageModel.Shift(0, 0))
	require.Equal(t, TerminalSize{}, pageModel.GetUsingTerminalSize())

	require.Equal(t, 0, pageModel.GetRowCount().ToInt())

	pageModel.GetHeaderModel().Update(TerminalSize{rll.MakeRowLenLimit(1, 1, 1), 1})
	require.Equal(t, row.Row{}, pageModel.GetHeaderModel().GetHeaderRow())

	require.Equal(t, 0, pageModel.GetBodyModel().GetRowCount().ToInt())
	require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorParagraphIndex().ToInt())
	require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorParagraphRowIndex().ToInt())
	require.Equal(t, 0, pageModel.GetBodyModel().GetAnchorRowAbsolutelyIndex().ToInt())
	_, err := pageModel.GetBodyModel().GetParagraph(0)
	require.Error(t, err)
	require.Error(t, pageModel.GetBodyModel().Update(TerminalSize{}, 0))
	require.Error(t, pageModel.GetBodyModel().Shift(0, 0))

	pageModel.GetFooterModel().Update(TerminalSize{rll.MakeRowLenLimit(1, 1, 1), 1})
	require.Equal(t, row.Row{}, pageModel.GetFooterModel().GetFooterRow())
}

package paragraph_model

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	rllMock "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter/mock"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"testing"
)

func TestDynamicRow(t *testing.T) {
	t.Parallel()

	splitFront, splitBack := newSplitter(1, 2), newSplitter(5, 6)
	limit := rllMock.GetRowLenLimitMin()
	r := makeDynamicRow(limit.Max(), page.TextToCells("example"), 1, splitFront, splitBack)

	require.Equal(t, splitFront, r.getBreakFront())
	require.Equal(t, splitBack, r.getBreakBack())

	require.Equal(t, row.MakeRow(0, page.TextToCells("amp")), r.toRow(0))
	require.Equal(t, size.Width(3), r.len())

	splitBack = newSplitter(6, 7)
	r.setBreakBack(splitBack)
	require.Equal(t, "ampl", r.String())
}

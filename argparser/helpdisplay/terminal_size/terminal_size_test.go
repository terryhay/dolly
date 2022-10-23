package terminal_size

import (
	"github.com/stretchr/testify/require"
	rllMock "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter/mock"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	"testing"
)

func TestTerminalSize(t *testing.T) {
	t.Parallel()

	ts := TerminalSize{}
	require.NotNil(t, ts.IsValid())

	widthLimit := rllMock.GetRowLenLimitMin()
	height := size.Height(10)
	ts = MakeTerminalSize(widthLimit, height)
	require.Nil(t, ts.IsValid())

	require.Equal(t, widthLimit, ts.GetWidthLimit())
	require.Equal(t, height, ts.GetHeight())

	tsClone := ts.CloneWithNewWidthLimit(ts.GetWidthLimit())
	require.Equal(t, ts, tsClone)

	tsClone = ts.CloneWithNewHeight(ts.GetHeight())
	require.Equal(t, ts, tsClone)
}

package terminal_size

import (
	"testing"

	"github.com/stretchr/testify/require"
	rllMock "github.com/terryhay/dolly/man_style_help/row_len_limiter/mock"
	"github.com/terryhay/dolly/tools/size"
)

func TestTerminalSize(t *testing.T) {
	t.Parallel()

	ts := TerminalSize{}
	require.NotNil(t, ts.IsValid())

	widthLimit := rllMock.GetRowLenLimitMin()
	height := size.MakeHeight(10)
	ts = MakeTerminalSize(widthLimit, height)
	require.Nil(t, ts.IsValid())

	require.Equal(t, widthLimit, ts.GetWidthLimit())
	require.Equal(t, height, ts.GetHeight())
}

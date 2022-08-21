package termbox_decorator_mock

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"
)

func TestTermBoxDecoratorMock(t *testing.T) {
	t.Parallel()

	funcClearRes := dollyerr.NewError(
		dollyerr.CodeTermBoxDecoratorClearError,
		fmt.Errorf("TermBoxDecorator: Clear call error"),
	)
	funcCloseRes := false
	funcFlushRes := dollyerr.NewError(
		dollyerr.CodeTermBoxDecoratorFlushError,
		fmt.Errorf("TermBoxDecorator: Flush call error"),
	)
	funcInitRes := dollyerr.NewError(
		dollyerr.CodeTermBoxDecoratorInitError,
		fmt.Errorf("TermBoxDecorator: Init call error"),
	)
	funcPollEventRes := termbox.Event{
		Width: int(gofakeit.Int32()),
	}
	funcSetCellRes := false
	funcSetRuneRes := false
	funcSizeWidthRes := int(gofakeit.Int32())
	funcSizeHeightRes := int(gofakeit.Int32())

	tbDecMock := NewTermBoxDecoratorMock(TermBoxDecoratorMockInit{
		FuncClear: func() error {
			return funcClearRes
		},
		FuncClose: func() {
			funcCloseRes = true
		},
		FuncFlush: func() error {
			return funcFlushRes
		},
		FuncInit: func() error {
			return funcInitRes
		},
		FuncPollEvent: func() termbox.Event {
			return funcPollEventRes
		},
		FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
			funcSetCellRes = true
		},
		FuncSetRune: func(x, y int, ch rune) {
			funcSetRuneRes = true
		},
		FuncSize: func() (width int, height int) {
			return funcSizeWidthRes, funcSizeHeightRes
		},
	})

	require.NotNil(t, tbDecMock)

	require.Equal(t, tbDecMock.Clear(), funcClearRes)

	tbDecMock.Close()
	require.True(t, funcCloseRes)

	require.Equal(t, tbDecMock.Flush(), funcFlushRes)
	require.Equal(t, tbDecMock.Init(), funcInitRes)
	require.Equal(t, tbDecMock.PollEvent(), funcPollEventRes)

	tbDecMock.SetCell(0, 0, rune(0), 0, 0)
	require.True(t, funcSetCellRes)

	tbDecMock.SetRune(0, 0, rune(0))
	require.True(t, funcSetRuneRes)

	w, h := tbDecMock.Size()
	require.Equal(t, funcSizeWidthRes, w)
	require.Equal(t, funcSizeHeightRes, h)
}

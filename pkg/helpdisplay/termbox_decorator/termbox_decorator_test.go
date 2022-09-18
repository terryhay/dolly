package termbox_decorator

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"
)

func TestTermBoxDecorator(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var termBoxDecor *termBoxDecoratorImpl

		require.Nil(t, termBoxDecor.Clear())
		termBoxDecor.Close()
		require.Nil(t, termBoxDecor.Flush())
		require.Nil(t, termBoxDecor.Init())
		require.Equal(t, termbox.Event{}, termBoxDecor.PollEvent())
		termBoxDecor.SetCell(0, 0, rune(1), termbox.AttrBlink, termbox.AttrBlink)
		termBoxDecor.SetRune(0, 0, rune(1))
		w, h := termBoxDecor.Size()
		require.Equal(t, 0, w)
		require.Equal(t, 0, h)
	})

	t.Run("methods", func(t *testing.T) {
		funcClearRes := dollyerr.NewErrorIfItIs(
			dollyerr.CodeTermBoxDecoratorClearError,
			"termBoxDecorator.Clear",
			fmt.Errorf("funcClearRes"))
		funcCloseRes := false
		funcFlushRes := dollyerr.NewErrorIfItIs(dollyerr.CodeTermBoxDecoratorFlushError, "termBoxDecorator.Flush", fmt.Errorf("funcFlushRes"))
		funcInitRes := dollyerr.NewErrorIfItIs(dollyerr.CodeTermBoxDecoratorInitError, "termBoxDecorator.Init", fmt.Errorf("funcInitRes"))
		funcPollEventRes := termbox.Event{Type: termbox.EventKey}
		funcSetCellRes := false
		funcSetCharRes := false
		funcSizeResW, funcSizeResH := 0, 0

		termBoxDecor := NewTermBoxDecorator(
			&Mock{
				FuncClear: func(_, _ termbox.Attribute) error {
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
				FuncSetChar: func(x, y int, ch rune) {
					funcSetCharRes = true
				},
				FuncSize: func() (width int, height int) {
					return funcSizeResW, funcSizeResH
				},
			})

		require.Equal(t, funcClearRes.Code(), termBoxDecor.Clear().Code())

		termBoxDecor.Close()
		require.True(t, funcCloseRes)

		require.Equal(t, funcFlushRes.Code(), termBoxDecor.Flush().Code())

		require.Equal(t, funcInitRes.Code(), termBoxDecor.Init().Code())

		require.Equal(t, funcPollEventRes, termBoxDecor.PollEvent())

		termBoxDecor.SetCell(0, 0, rune(1), termbox.AttrBlink, termbox.AttrBlink)
		require.True(t, funcSetCellRes)

		termBoxDecor.SetRune(0, 0, rune(1))
		require.True(t, funcSetCharRes)

		w, h := termBoxDecor.Size()
		require.Equal(t, funcSizeResW, w)
		require.Equal(t, funcSizeResH, h)
	})
}

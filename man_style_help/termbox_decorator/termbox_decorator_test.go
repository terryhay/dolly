package termbox_decorator

import (
	"testing"

	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
)

func TestTermBoxDecorator(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var termBoxDecor *impl

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
		funcClearRes := ErrTermBoxDecoratorClear
		funcCloseRes := false
		funcFlushRes := ErrTermBoxDecoratorFlush
		funcInitRes := ErrTermBoxDecoratorInit
		funcPollEventRes := termbox.Event{Type: termbox.EventKey}
		funcSetCellRes := false
		funcSetCharRes := false
		funcSizeResW, funcSizeResH := 0, 0

		decTermBox := NewTermBoxDecorator()
		require.NotNil(t, decTermBox)

		decTermBox = TermBoxDecoratorMock{
			FuncClear: func(_, _ termbox.Attribute) error {
				return ErrTermBoxDecoratorClear
			},
			FuncClose: func() {
				funcCloseRes = true
			},
			FuncFlush: func() error {
				return ErrTermBoxDecoratorFlush
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
		}.Create()

		require.ErrorIs(t, decTermBox.Clear(), funcClearRes)

		decTermBox.Close()
		require.True(t, funcCloseRes)

		require.ErrorIs(t, decTermBox.Flush(), funcFlushRes)

		require.ErrorIs(t, decTermBox.Init(), funcInitRes)

		require.Equal(t, funcPollEventRes, decTermBox.PollEvent())

		decTermBox.SetCell(0, 0, rune(1), termbox.AttrBlink, termbox.AttrBlink)
		require.True(t, funcSetCellRes)

		decTermBox.SetRune(0, 0, rune(1))
		require.True(t, funcSetCharRes)

		w, h := decTermBox.Size()
		require.Equal(t, funcSizeResW, w)
		require.Equal(t, funcSizeResH, h)
	})
}

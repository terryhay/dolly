package page_view

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	tbd "github.com/terryhay/dolly/man_style_help/termbox_decorator"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

func TestPageView(t *testing.T) {
	t.Parallel()

	t.Run("error_initialization", func(t *testing.T) {
		_, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncInit: func() error {
					return fmt.Errorf("TermBoxDecorator: Init call error")
				},
			}.Create(),
			coty.AppNameUndefined,
			hp.Body{},
		)
		require.Error(t, err)

		_, err = NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncInit: func() error {
					return nil
				},
				FuncSize: func() (_ int, _ int) {
					return 0, 0
				},
			}.Create(),
			coty.AppNameUndefined,
			hp.Body{},
		)
		require.Error(t, err)
	})

	t.Run("running", func(t *testing.T) {
		ev := make(chan termbox.Event, 1)
		go func() {
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrl8,
			}
			ev <- termbox.Event{
				Type: termbox.EventResize,
			}
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyArrowDown,
			}
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyArrowUp,
			}
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
			}
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-ev
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.Nil(t, err)
	})

	t.Run("running2", func(t *testing.T) {
		ev := make(chan termbox.Event, 1)
		go func() {
			ev <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		_, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-ev
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)
	})
}

func TestErrorsBeforeEventLoop(t *testing.T) {
	t.Parallel()

	chanWidths := make(chan int, 1)
	go func() {
		chanWidths <- rll.TerminalMinWidth.Int()
		chanWidths <- 0

		close(chanWidths)
	}()

	tests := []struct {
		caseName string

		termBoxMockSrc tbd.TermBoxDecoratorMock
		expErr         error
	}{
		{
			caseName: "clear_call_error",

			termBoxMockSrc: tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return coty.RandError()
				},
				FuncClose: func() {
				},
				FuncInit: func() error {
					return nil
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			},
			expErr: coty.RandError(),
		},
		{
			caseName: "flush_call_error",

			termBoxMockSrc: tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return coty.RandError()
				},
				FuncInit: func() error {
					return nil
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			},
			expErr: coty.RandError(),
		},
		{
			caseName: "page_model_update_error",

			termBoxMockSrc: tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return coty.RandError()
				},
				FuncInit: func() error {
					return nil
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSize: func() (width int, height int) {
					return <-chanWidths, 7
				},
			},
			expErr: rll.ErrIsValidMin,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			pageView, err := NewPageView(
				tc.termBoxMockSrc.Create(),
				coty.RandNameApp(),
				hp.MakeBody([]hp.Row{
					hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
				}),
			)
			require.Nil(t, err)

			err = pageView.Run()
			require.NotNil(t, err)
			require.ErrorIs(t, err, tc.expErr)
		})
	}
}

func TestErrorInsideEventLoop(t *testing.T) {
	t.Parallel()

	t.Run("error_in_KeyArrowDown_event", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyArrowDown,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		funcClearRes := fmt.Errorf("TermBoxDecorator: Clear call error")
		clearResChan := make(chan error, 1)
		go func() {
			clearResChan <- nil
			clearResChan <- funcClearRes
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return <-clearResChan
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})

	t.Run("error_in_KeyArrowUp_event", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyArrowUp,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		funcClearRes := fmt.Errorf("TermBoxDecorator: Clear call error")
		clearResChan := make(chan error, 1)
		go func() {
			clearResChan <- nil
			clearResChan <- funcClearRes
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return <-clearResChan
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})

	t.Run("error_in_KeyArrowUp_event", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		funcClearRes := fmt.Errorf("TermBoxDecorator: Clear call error")
		clearResChan := make(chan error, 1)
		go func() {
			clearResChan <- nil
			clearResChan <- funcClearRes
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return <-clearResChan
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})

	t.Run("error_in_default_event", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlUnderscore,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventKey,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		funcClearRes := fmt.Errorf("TermBoxDecorator: Clear call error")
		clearResChan := make(chan error, 1)
		go func() {
			clearResChan <- nil
			clearResChan <- funcClearRes
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return <-clearResChan
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})

	t.Run("error_in_default_event_key", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventNone,
				Key:  termbox.KeyCtrlUnderscore,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventNone,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		funcClearRes := fmt.Errorf("TermBoxDecorator: Clear call error")
		clearResChan := make(chan error, 1)
		go func() {
			clearResChan <- nil
			clearResChan <- funcClearRes
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return <-clearResChan
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return nil
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})

	t.Run("error_in_flush", func(t *testing.T) {
		eventChan := make(chan termbox.Event, 1)
		go func() {
			eventChan <- termbox.Event{
				Type: termbox.EventNone,
				Key:  termbox.KeyCtrlUnderscore,
			}
			eventChan <- termbox.Event{
				Type: termbox.EventNone,
				Key:  termbox.KeyCtrlTilde,
				Ch:   runes.RuneLwQ,
			}
		}()

		flushResChan := make(chan error, 1)
		go func() {
			flushResChan <- nil
			flushResChan <- fmt.Errorf("TermBoxDecorator: Clear call error")
		}()

		pageView, err := NewPageView(
			tbd.TermBoxDecoratorMock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return <-flushResChan
				},
				FuncInit: func() error {
					return nil
				},
				FuncPollEvent: func() termbox.Event {
					return <-eventChan
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSetChar: func(x, y int, ch rune) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.Int(), 7
				},
			}.Create(),
			coty.RandNameApp(),
			hp.MakeBody([]hp.Row{
				hp.MakeRow(size.WidthZero, hp.MakeRowChunk(gofakeit.Name())),
			}),
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})
}

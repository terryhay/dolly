package page_view

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	rll "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/argparser/helpdisplay/runes"
	tbd "github.com/terryhay/dolly/argparser/helpdisplay/termbox_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestPageView(t *testing.T) {
	t.Parallel()

	t.Run("error_initialization", func(t *testing.T) {
		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
					FuncInit: func() error {
						return fmt.Errorf("TermBoxDecorator: Init call error")
					},
				}),
			page.Page{},
		)
		require.Error(t, err.Error())

		err = pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
					FuncInit: func() error {
						return nil
					},
					FuncSize: func() (_ int, _ int) {
						return 0, 0
					},
				}),
			page.Page{},
		)
		require.Error(t, err.Error())
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
		)
		require.Nil(t, err)
	})
}

func TestErrorsBeforeEventLoop(t *testing.T) {
	t.Parallel()

	pageData := page.Page{
		Header: page.MakeParagraph(0, gofakeit.Name()),
		Paragraphs: []page.Paragraph{
			page.MakeParagraph(0, gofakeit.Name()),
		},
	}
	funcClearRes := dollyerr.NewError(
		dollyerr.CodeTermBoxDecoratorClearError,
		fmt.Errorf("TermBoxDecorator: Clear call error"),
	)
	funcFlushRes := dollyerr.NewError(
		dollyerr.CodeTermBoxDecoratorFlushError,
		fmt.Errorf("TermBoxDecorator: Flush call error"),
	)

	widthsChan := make(chan int, 1)
	go func() {
		widthsChan <- rll.TerminalMinWidth.ToInt()
		widthsChan <- 0
	}()

	testData := []struct {
		caseName string

		initData     tbd.Mock
		expectedCode dollyerr.Code
	}{
		{
			caseName: "clear_call_error",

			initData: tbd.Mock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return funcClearRes.Error()
				},
				FuncClose: func() {
				},
				FuncInit: func() error {
					return nil
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.ToInt(), 7
				},
			},
			expectedCode: dollyerr.CodeTermBoxDecoratorClearError,
		},
		{
			caseName: "flush_call_error",

			initData: tbd.Mock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return funcFlushRes.Error()
				},
				FuncInit: func() error {
					return nil
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSize: func() (width int, height int) {
					return rll.TerminalMinWidth.ToInt(), 7
				},
			},
			expectedCode: dollyerr.CodeTermBoxDecoratorFlushError,
		},
		{
			caseName: "page_model_update_error",

			initData: tbd.Mock{
				FuncClear: func(_, _ termbox.Attribute) error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return funcFlushRes.Error()
				},
				FuncInit: func() error {
					return nil
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSize: func() (width int, height int) {
					return <-widthsChan, 7
				},
			},
			expectedCode: dollyerr.CodeHelpDisplayTerminalWidthLimitError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			var pageView PageView
			err := pageView.Init(tbd.NewTermBoxDecorator(&td.initData), pageData)
			require.Nil(t, err)

			err = pageView.Run()
			require.NotNil(t, err)
			require.Equal(t, td.expectedCode, err.Code())
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
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

		var pageView PageView
		err := pageView.Init(
			tbd.NewTermBoxDecorator(
				&tbd.Mock{
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
						return rll.TerminalMinWidth.ToInt(), 7
					},
				}),
			page.Page{
				Header: page.MakeParagraph(0, gofakeit.Name()),
				Paragraphs: []page.Paragraph{
					page.MakeParagraph(0, gofakeit.Name()),
				},
			},
		)
		require.Nil(t, err)

		err = pageView.Run()
		require.NotNil(t, err)
	})
}

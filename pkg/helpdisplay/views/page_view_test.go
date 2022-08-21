package views

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	rowLenLimit "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	tbDecorMock "github.com/terryhay/dolly/pkg/helpdisplay/termbox_decorator/mock"
	"testing"
)

func TestPageView(t *testing.T) {
	t.Parallel()

	t.Run("error_initialization", func(t *testing.T) {
		var pageView PageView
		err := pageView.Init(
			tbDecorMock.NewTermBoxDecoratorMock(
				tbDecorMock.TermBoxDecoratorMockInit{
					FuncInit: func() error {
						return dollyerr.NewError(
							dollyerr.CodeTermBoxDecoratorInitError,
							fmt.Errorf("TermBoxDecorator: Init call error"),
						)
					},
				}),
			data.Page{},
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

		var pageView PageView
		err := pageView.Init(
			tbDecorMock.NewTermBoxDecoratorMock(
				tbDecorMock.TermBoxDecoratorMockInit{
					FuncClear: func() error {
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
					FuncSetRune: func(x, y int, ch rune) {
					},
					FuncSize: func() (width int, height int) {
						return rowLenLimit.TerminalMinWidth, 7
					},
				}),
			data.Page{
				Header: gofakeit.Name(),
				Paragraphs: []*data.Paragraph{
					{
						Text: gofakeit.Name(),
					},
				},
			},
		)
		require.Nil(t, err)
	})
}

func TestErrorsBeforeEventLoop(t *testing.T) {
	t.Parallel()

	pageData := data.Page{
		Header: gofakeit.Name(),
		Paragraphs: []*data.Paragraph{
			{
				Text: gofakeit.Name(),
			},
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
		widthsChan <- rowLenLimit.TerminalMinWidth
		widthsChan <- 0
	}()

	testData := []struct {
		caseName string

		initData     tbDecorMock.TermBoxDecoratorMockInit
		expectedCode dollyerr.Code
	}{
		{
			caseName: "clear_call_error",

			initData: tbDecorMock.TermBoxDecoratorMockInit{
				FuncClear: func() error {
					return funcClearRes
				},
				FuncClose: func() {
				},
				FuncInit: func() error {
					return nil
				},
				FuncSize: func() (width int, height int) {
					return rowLenLimit.TerminalMinWidth, 7
				},
			},
			expectedCode: dollyerr.CodeHelpDisplayRenderError,
		},
		{
			caseName: "flush_call_error",

			initData: tbDecorMock.TermBoxDecoratorMockInit{
				FuncClear: func() error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return funcFlushRes
				},
				FuncInit: func() error {
					return nil
				},
				FuncSetCell: func(x, y int, ch rune, fg, bg termbox.Attribute) {
				},
				FuncSize: func() (width int, height int) {
					return rowLenLimit.TerminalMinWidth, 7
				},
			},
			expectedCode: dollyerr.CodeHelpDisplayRunError,
		},
		{
			caseName: "page_model_update_error",

			initData: tbDecorMock.TermBoxDecoratorMockInit{
				FuncClear: func() error {
					return nil
				},
				FuncClose: func() {
				},
				FuncFlush: func() error {
					return funcFlushRes
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
			expectedCode: dollyerr.CodeHelpDisplayRenderError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			var pageView PageView
			err := pageView.Init(tbDecorMock.NewTermBoxDecoratorMock(td.initData), pageData)
			require.Nil(t, err)

			err = pageView.Run()
			require.NotNil(t, err)
			require.Equal(t, td.expectedCode, err.Code())
		})
	}
}

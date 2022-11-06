package termbox_decorator

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// TermBoxDecorator decorates some methods of termbox module
type TermBoxDecorator interface {
	Clear() *dollyerr.Error
	Close()
	Flush() *dollyerr.Error
	Init() *dollyerr.Error
	PollEvent() termbox.Event
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	SetRune(x, y int, ch rune)
	Size() (width int, height int)
}

// Mock contains mock methods of initialize mocked decorator object
type Mock struct {
	FuncClear     func(fg, bg termbox.Attribute) error
	FuncClose     func()
	FuncFlush     func() error
	FuncInit      func() error
	FuncPollEvent func() termbox.Event
	FuncSetCell   func(x, y int, ch rune, fg, bg termbox.Attribute)
	FuncSetChar   func(x, y int, ch rune)
	FuncSize      func() (width int, height int)
}

// NewTermBoxDecorator constructs a TermBoxDecorator object.
// You can mock it by mean not nil mock argument
func NewTermBoxDecorator(mock *Mock) TermBoxDecorator {
	termBoxDecor := &termBoxDecoratorImpl{
		funcClear:     termbox.Clear,
		funcClose:     termbox.Close,
		funcFlush:     termbox.Flush,
		funcInit:      termbox.Init,
		funcPollEvent: termbox.PollEvent,
		funcSetCell:   termbox.SetCell,
		funcSetChar:   termbox.SetChar,
		funcSize:      termbox.Size,
	}
	if mock != nil {
		termBoxDecor = &termBoxDecoratorImpl{
			funcClear:     mock.FuncClear,
			funcClose:     mock.FuncClose,
			funcFlush:     mock.FuncFlush,
			funcInit:      mock.FuncInit,
			funcPollEvent: mock.FuncPollEvent,
			funcSetCell:   mock.FuncSetCell,
			funcSetChar:   mock.FuncSetChar,
			funcSize:      mock.FuncSize,
		}
	}
	return termBoxDecor
}

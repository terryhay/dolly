package termbox_decorator

import "github.com/nsf/termbox-go"

// TermBoxDecoratorMock contains mock methods of initialize mocked decorator object
type TermBoxDecoratorMock struct {
	FuncClear     func(fg, bg termbox.Attribute) error
	FuncClose     func()
	FuncFlush     func() error
	FuncInit      func() error
	FuncPollEvent func() termbox.Event
	FuncSetCell   func(x, y int, ch rune, fg, bg termbox.Attribute)
	FuncSetChar   func(x, y int, ch rune)
	FuncSize      func() (width int, height int)
}

// Create creates mocked os decorator object
func (src TermBoxDecoratorMock) Create() TermBoxDecorator {
	return &impl{
		funcClear:     src.FuncClear,
		funcClose:     src.FuncClose,
		funcFlush:     src.FuncFlush,
		funcInit:      src.FuncInit,
		funcPollEvent: src.FuncPollEvent,
		funcSetCell:   src.FuncSetCell,
		funcSetChar:   src.FuncSetChar,
		funcSize:      src.FuncSize,
	}
}

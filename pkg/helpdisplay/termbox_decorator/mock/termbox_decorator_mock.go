package termbox_decorator_mock

import (
	"github.com/nsf/termbox-go"
	tbd "github.com/terryhay/dolly/pkg/helpdisplay/termbox_decorator"
)

type TermBoxDecoratorMockInit struct {
	FuncClear     func() error
	FuncClose     func()
	FuncFlush     func() error
	FuncInit      func() error
	FuncPollEvent func() termbox.Event
	FuncSetCell   func(x, y int, ch rune, fg, bg termbox.Attribute)
	FuncSetRune   func(x, y int, ch rune)
	FuncSize      func() (width int, height int)
}

func NewTermBoxDecoratorMock(init TermBoxDecoratorMockInit) tbd.TermBoxDecorator {
	return &init
}

func (i *TermBoxDecoratorMockInit) Clear() error {
	return i.FuncClear()
}

func (i *TermBoxDecoratorMockInit) Close() {
	i.FuncClose()
}

func (i *TermBoxDecoratorMockInit) Flush() error {
	return i.FuncFlush()
}

func (i *TermBoxDecoratorMockInit) Init() error {
	return i.FuncInit()
}

func (i *TermBoxDecoratorMockInit) PollEvent() termbox.Event {
	return i.FuncPollEvent()
}

func (i *TermBoxDecoratorMockInit) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	i.FuncSetCell(x, y, ch, fg, bg)
}

func (i *TermBoxDecoratorMockInit) SetRune(x, y int, ch rune) {
	i.FuncSetRune(x, y, ch)
}

func (i *TermBoxDecoratorMockInit) Size() (width int, height int) {
	return i.FuncSize()
}

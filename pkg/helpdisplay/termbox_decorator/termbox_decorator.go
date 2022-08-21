package termbox_decorator

import (
	"github.com/nsf/termbox-go"
)

type TermBoxDecorator interface {
	Clear() error
	Close()
	Flush() error
	Init() error
	PollEvent() termbox.Event
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	SetRune(x, y int, ch rune)
	Size() (width int, height int)
}

func NewTermBoxDecorator() TermBoxDecorator {
	return &termBoxDecoratorImpl{}
}

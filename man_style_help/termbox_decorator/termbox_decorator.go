package termbox_decorator

import (
	"github.com/nsf/termbox-go"
)

// TermBoxDecorator decorates some methods of termbox module
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

// NewTermBoxDecorator constructs a TermBoxDecorator object.
// You can mock it by mean not nil mock argument
func NewTermBoxDecorator() TermBoxDecorator {
	return &impl{
		funcClear:     termbox.Clear,
		funcClose:     termbox.Close,
		funcFlush:     termbox.Flush,
		funcInit:      termbox.Init,
		funcPollEvent: termbox.PollEvent,
		funcSetCell:   termbox.SetCell,
		funcSetChar:   termbox.SetChar,
		funcSize:      termbox.Size,
	}
}

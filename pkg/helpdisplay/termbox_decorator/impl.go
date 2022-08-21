package termbox_decorator

import (
	"github.com/nsf/termbox-go"
)

type termBoxDecoratorImpl struct {
}

func (*termBoxDecoratorImpl) Clear() error {
	return termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (*termBoxDecoratorImpl) Close() {
	termbox.Close()
}

func (*termBoxDecoratorImpl) Flush() error {
	return termbox.Flush()
}

func (*termBoxDecoratorImpl) Init() error {
	return termbox.Init()
}

func (*termBoxDecoratorImpl) PollEvent() termbox.Event {
	return termbox.PollEvent()
}

func (*termBoxDecoratorImpl) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (*termBoxDecoratorImpl) SetRune(x, y int, r rune) {
	termbox.SetChar(x, y, r)
}

func (*termBoxDecoratorImpl) Size() (int, int) {
	return termbox.Size()
}

package termbox_decorator

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/utils/dollyerr"
)

type termBoxDecoratorImpl struct {
	funcClear     func(fg, bg termbox.Attribute) error
	funcClose     func()
	funcFlush     func() error
	funcInit      func() error
	funcPollEvent func() termbox.Event
	funcSetCell   func(x, y int, ch rune, fg, bg termbox.Attribute)
	funcSetChar   func(x, y int, ch rune)
	funcSize      func() (width int, height int)
}

func (tbd *termBoxDecoratorImpl) Clear() *dollyerr.Error {
	if tbd == nil {
		return nil
	}
	err := tbd.funcClear(termbox.ColorDefault, termbox.ColorDefault)
	return dollyerr.NewErrorIfItIs(dollyerr.CodeTermBoxDecoratorClearError, "termBoxDecorator.Clear", err)
}

func (tbd *termBoxDecoratorImpl) Close() {
	if tbd == nil {
		return
	}
	tbd.funcClose()
}

func (tbd *termBoxDecoratorImpl) Flush() *dollyerr.Error {
	if tbd == nil {
		return nil
	}
	err := tbd.funcFlush()
	return dollyerr.NewErrorIfItIs(dollyerr.CodeTermBoxDecoratorFlushError, "termBoxDecorator.Flush", err)
}

func (tbd *termBoxDecoratorImpl) Init() *dollyerr.Error {
	if tbd == nil {
		return nil
	}
	err := tbd.funcInit()
	return dollyerr.NewErrorIfItIs(dollyerr.CodeTermBoxDecoratorInitError, "termBoxDecorator.Init", err)
}

func (tbd *termBoxDecoratorImpl) PollEvent() termbox.Event {
	if tbd == nil {
		return termbox.Event{}
	}
	return tbd.funcPollEvent()
}

func (tbd *termBoxDecoratorImpl) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if tbd == nil {
		return
	}
	tbd.funcSetCell(x, y, ch, fg, bg)
}

func (tbd *termBoxDecoratorImpl) SetRune(x, y int, r rune) {
	if tbd == nil {
		return
	}
	tbd.funcSetChar(x, y, r)
}

func (tbd *termBoxDecoratorImpl) Size() (int, int) {
	if tbd == nil {
		return 0, 0
	}
	return tbd.funcSize()
}

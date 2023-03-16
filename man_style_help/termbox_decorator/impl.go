package termbox_decorator

import (
	"errors"
	"fmt"

	"github.com/nsf/termbox-go"
)

type impl struct {
	funcClear     func(fg, bg termbox.Attribute) error
	funcClose     func()
	funcFlush     func() error
	funcInit      func() error
	funcPollEvent func() termbox.Event
	funcSetCell   func(x, y int, ch rune, fg, bg termbox.Attribute)
	funcSetChar   func(x, y int, ch rune)
	funcSize      func() (width int, height int)
}

// ErrTermBoxDecoratorClear - termbox.Clear returned error
var ErrTermBoxDecoratorClear = errors.New(`TermBoxDecorator.Clear: termbox.Clear returned error`)

func (i *impl) Clear() (err error) {
	if i == nil {
		return nil
	}

	if errClear := i.funcClear(termbox.ColorDefault, termbox.ColorDefault); errClear != nil {
		err = errors.Join(
			fmt.Errorf(`%w: foreground color "%v", background color "%v"`,
				ErrTermBoxDecoratorClear, termbox.ColorDefault, termbox.ColorDefault),
			errClear,
		)
	}
	return err
}

func (i *impl) Close() {
	if i == nil {
		return
	}
	i.funcClose()
}

// ErrTermBoxDecoratorFlush - termbox.Flush returned error
var ErrTermBoxDecoratorFlush = errors.New(`TermBoxDecorator.Flush: termbox.Flush returned error`)

func (i *impl) Flush() (err error) {
	if i == nil {
		return nil
	}

	if errFlush := i.funcFlush(); errFlush != nil {
		err = errors.Join(ErrTermBoxDecoratorFlush, errFlush)
	}
	return err
}

// ErrTermBoxDecoratorInit - termbox.Init returned error
var ErrTermBoxDecoratorInit = errors.New(`TermBoxDecorator.Flush: termbox.Init returned error`)

func (i *impl) Init() (err error) {
	if i == nil {
		return nil
	}
	if errInit := i.funcInit(); errInit != nil {
		err = errors.Join(ErrTermBoxDecoratorInit, errInit)
	}
	return err
}

func (i *impl) PollEvent() termbox.Event {
	if i == nil {
		return termbox.Event{}
	}
	return i.funcPollEvent()
}

func (i *impl) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if i == nil {
		return
	}
	i.funcSetCell(x, y, ch, fg, bg)
}

func (i *impl) SetRune(x, y int, r rune) {
	if i == nil {
		return
	}
	i.funcSetChar(x, y, r)
}

func (i *impl) Size() (int, int) {
	if i == nil {
		return 0, 0
	}
	return i.funcSize()
}

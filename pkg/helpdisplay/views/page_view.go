package views

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/models"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	tbd "github.com/terryhay/dolly/pkg/helpdisplay/termbox_decorator"
)

// PageView renders page of text
type PageView struct {
	termBoxDecor tbd.TermBoxDecorator
	pageModel    models.PageModel
	exitCodes    map[rune]bool
}

// Init initializes the instance
func (pv *PageView) Init(termBoxDecor tbd.TermBoxDecorator, pageData data.Page) *dollyerr.Error {
	err := termBoxDecor.Init()
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeTermBoxDecoratorInitError,
			multierror.Append(err, fmt.Errorf("PageView.Init: can't create a termbox decorator")),
		)
	}

	terminalWidth, terminalHeight := termBoxDecor.Size()

	pv.termBoxDecor = termBoxDecor
	pv.pageModel = models.MakePageModel(pageData, terminalWidth, terminalHeight)
	pv.exitCodes = map[rune]bool{
		runes.RuneLwQ:   true,
		runes.RuneLwQRu: true,
		runes.RuneUpQ:   true,
		runes.RuneUpQRu: true,
	}

	return nil
}

func (pv *PageView) Run() *dollyerr.Error {
	defer pv.termBoxDecor.Close()

	{
		err := pv.render(0)
		if err != nil {
			return err.AppendError(fmt.Errorf("PageView.Run: render call error"))
		}
	}
	{
		err := pv.termBoxDecor.Flush()
		if err != nil {
			return dollyerr.NewError(
				dollyerr.CodeHelpDisplayRunError,
				multierror.Append(err, fmt.Errorf("PageView.Run: termbox.Flush call error")),
			)
		}
	}

	for {
		switch ev := pv.termBoxDecor.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				err := pv.render(1)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: render call error for termbox event '%v'", ev))
				}

			case termbox.KeyArrowUp:
				err := pv.render(-1)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: render call error for termbox event '%v'", ev))
				}

			case termbox.KeyCtrlTilde:
				if _, contain := pv.exitCodes[ev.Ch]; contain {
					return nil
				}
				err := pv.render(0)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: render call error for termbox event '%v'", ev))
				}

			default:
				err := pv.render(0)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: render call error for termbox event '%v'", ev))
				}
			}
		default:
			err := pv.render(0)
			if err != nil {
				return err.AppendError(fmt.Errorf("PageView.Run: render call error for termbox event '%v'", ev))
			}
		}

		err := pv.termBoxDecor.Flush()
		if err != nil {
			return dollyerr.NewError(
				dollyerr.CodeHelpDisplayRunError,
				multierror.Append(err, fmt.Errorf("PageView.Run: termbox.Flush call error")),
			)
		}
	}
}

func (pv *PageView) render(shift int) *dollyerr.Error {
	errStd := pv.termBoxDecor.Clear()
	if errStd != nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayRenderError,
			multierror.Append(errStd, fmt.Errorf("PageView.render: termbox.Clear call error")),
		)
	}

	w, h := pv.termBoxDecor.Size()
	err := pv.pageModel.Update(w, h, shift)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayRenderError,
			multierror.Append(err, fmt.Errorf("PageView.render: pageModel.Update call error")),
		)
	}

	var (
		x, y int
		cell termbox.Cell
	)
	for it := pv.pageModel.RowBegin(); !it.End(); it = pv.pageModel.RowNext(it) {
		for x, cell = range it.Cells {
			pv.termBoxDecor.SetCell(it.ShiftIndex+x, y, cell.Ch, cell.Fg, cell.Bg)
		}
		y++
	}

	return nil
}

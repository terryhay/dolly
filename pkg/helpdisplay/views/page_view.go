package views

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	"github.com/terryhay/dolly/pkg/helpdisplay/models"
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
	tbd "github.com/terryhay/dolly/pkg/helpdisplay/termbox_decorator"
)

// PageView renders page of text
type PageView struct {
	termBoxDecor  tbd.TermBoxDecorator
	rowLenLimiter rll.RowLenLimiter
	pageModel     *models.PageModel
	exitCodes     map[rune]bool
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
	pv.rowLenLimiter = rll.MakeRowLenLimiter()
	pv.pageModel = models.NewPageModel(
		pageData,
		models.MakeTerminalSize(
			pv.rowLenLimiter.GetRowLenLimit(size.Width(terminalWidth)),
			size.Height(terminalHeight),
		),
	)
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
		err := pv.process(0)
		if err != nil {
			return err.AppendError(fmt.Errorf("PageView.Run: process call error"))
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
				err := pv.process(1)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			case termbox.KeyArrowUp:
				err := pv.process(-1)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			case termbox.KeyCtrlTilde:
				if _, contain := pv.exitCodes[ev.Ch]; contain {
					return nil
				}
				err := pv.process(0)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			default:
				err := pv.process(0)
				if err != nil {
					return err.AppendError(fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}
			}
		default:
			err := pv.process(0)
			if err != nil {
				return err.AppendError(fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
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

func (pv *PageView) process(shift int) *dollyerr.Error {
	err := update(pv.termBoxDecor, pv.pageModel, pv.rowLenLimiter, shift)
	if err != nil {
		return err
	}

	return render(pv.termBoxDecor, models.MakeRowIterator(pv.pageModel))
}

func update(termBoxDecor tbd.TermBoxDecorator, pageModel *models.PageModel, rowLenLimiter rll.RowLenLimiter, shift int) *dollyerr.Error {
	{
		err := termBoxDecor.Clear()
		if err != nil {
			return dollyerr.NewError(
				dollyerr.CodeHelpDisplayRenderError,
				multierror.Append(err, fmt.Errorf("PageView.process: termbox.Clear call error")),
			)
		}
	}

	w, h := termBoxDecor.Size()

	err := pageModel.Update(
		models.MakeTerminalSize(rowLenLimiter.GetRowLenLimit(size.Width(w)), size.Height(h)),
		shift,
	)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeHelpDisplayRenderError,
			multierror.Append(err, fmt.Errorf("PageView.process: pageModel.Update call error")),
		)
	}

	return nil
}

func render(termBoxDecor tbd.TermBoxDecorator, it models.RowIterator) *dollyerr.Error {
	var (
		x, y int
		cell termbox.Cell
		err  *dollyerr.Error
	)
	for ; !it.End(); err = it.Next() {
		if err != nil {
			return err.AppendError(fmt.Errorf("process: RowIterator.Next call error"))
		}

		for x, cell = range it.Row().GetCells() {
			termBoxDecor.SetCell(it.Row().GetShiftIndex().ToInt()+x, y, cell.Ch, cell.Fg, cell.Bg)
		}
		y++
	}

	return nil
}

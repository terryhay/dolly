package page_view

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	pgm "github.com/terryhay/dolly/argparser/helpdisplay/page_model"
	ri "github.com/terryhay/dolly/argparser/helpdisplay/row_iterator"
	rll "github.com/terryhay/dolly/argparser/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/argparser/helpdisplay/runes"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
	tbd "github.com/terryhay/dolly/argparser/helpdisplay/termbox_decorator"
	ts "github.com/terryhay/dolly/argparser/helpdisplay/terminal_size"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// PageView renders page of text
type PageView struct {
	termBoxDecor  tbd.TermBoxDecorator
	rowLenLimiter rll.RowLenLimiter
	pageModel     *pgm.PageModel
	exitCodes     map[rune]bool
}

// Init initializes the instance
func (pv *PageView) Init(termBoxDecor tbd.TermBoxDecorator, pageData page.Page) *dollyerr.Error {
	err := termBoxDecor.Init()
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf("PageView.Init: can't create a termbox decorator"))
	}

	terminalWidth, terminalHeight := termBoxDecor.Size()

	pv.termBoxDecor = termBoxDecor
	pv.rowLenLimiter = rll.MakeRowLenLimiter()
	pv.pageModel, err = pgm.NewPageModel(
		pageData,
		ts.MakeTerminalSize(
			pv.rowLenLimiter.GetRowLenLimit(size.Width(terminalWidth)),
			size.Height(terminalHeight),
		),
	)
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf("PageView.Init: can't create a pageModel"))
	}

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
			return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error"))
		}
	}
	{
		err := pv.termBoxDecor.Flush()
		if err != nil {
			return dollyerr.Append(err, fmt.Errorf("PageView.Run: termbox.Flush call error"))
		}
	}

	for {
		switch ev := pv.termBoxDecor.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				err := pv.process(1)
				if err != nil {
					return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			case termbox.KeyArrowUp:
				err := pv.process(-1)
				if err != nil {
					return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			case termbox.KeyCtrlTilde:
				if _, contain := pv.exitCodes[ev.Ch]; contain {
					return nil
				}
				err := pv.process(0)
				if err != nil {
					return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}

			default:
				err := pv.process(0)
				if err != nil {
					return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
				}
			}
		default:
			err := pv.process(0)
			if err != nil {
				return dollyerr.Append(err, fmt.Errorf("PageView.Run: process call error for termbox event '%v'", ev))
			}
		}

		err := pv.termBoxDecor.Flush()
		if err != nil {
			return dollyerr.Append(err, fmt.Errorf("PageView.Run: termbox.Flush call error"))
		}
	}
}

func (pv *PageView) process(shift int) *dollyerr.Error {
	err := update(pv.termBoxDecor, pv.pageModel, pv.rowLenLimiter, shift)
	if err != nil {
		return err
	}

	render(pv.termBoxDecor, ri.MakeRowIterator(pv.pageModel))
	return nil
}

func update(termBoxDecor tbd.TermBoxDecorator, pageModel *pgm.PageModel, rowLenLimiter rll.RowLenLimiter, shift int) *dollyerr.Error {
	{
		err := termBoxDecor.Clear()
		if err != nil {
			return dollyerr.Append(err, fmt.Errorf("PageView.process: termbox.Clear call error"))
		}
	}

	w, h := termBoxDecor.Size()

	err := pageModel.Update(
		ts.MakeTerminalSize(rowLenLimiter.GetRowLenLimit(size.Width(w)), size.Height(h)),
		shift,
	)
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf("PageView.process: pageModel.update call error"))
	}

	return nil
}

func render(termBoxDecor tbd.TermBoxDecorator, it ri.RowIterator) {
	var (
		x, y int
		cell termbox.Cell
	)
	for ; !it.End(); it.Next() {
		for x, cell = range it.Row().GetCells() {
			termBoxDecor.SetCell(it.Row().GetShiftIndex().ToInt()+x, y, cell.Ch, cell.Fg, cell.Bg)
		}
		y++
	}
}

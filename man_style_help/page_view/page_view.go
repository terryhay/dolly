package page_view

import (
	"errors"
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/argparser/help_page/page"
	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	ri "github.com/terryhay/dolly/man_style_help/row_iterator"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	tbd "github.com/terryhay/dolly/man_style_help/termbox_decorator"
	ts "github.com/terryhay/dolly/man_style_help/terminal_size"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

// PageView renders page of text
type PageView struct {
	decTermBox    tbd.TermBoxDecorator
	rowLenLimiter rll.RowLenLimiter
	pageModel     *pgm.PageModel
	exitCodes     map[rune]struct{}
}

var (
	// ErrNewPageViewPageModel - page model construct error
	ErrNewPageViewPageModel = errors.New(`page_view.Init: page model "New" method returned error`)

	// ErrNewPageViewTermBoxDecorator - term box decorator init error
	ErrNewPageViewTermBoxDecorator = errors.New(`page_view.Init: termbox "Init" method returned error`)
)

// NewPageView constructs PageView object
func NewPageView(termBoxDecor tbd.TermBoxDecorator, appName coty.NameApp, pageBody page.Body) (*PageView, error) {
	if err := termBoxDecor.Init(); err != nil {
		return nil, errors.Join(ErrNewPageViewTermBoxDecorator, err)
	}

	terminalWidth, terminalHeight := termBoxDecor.Size()

	rowLenLimiter := rll.MakeRowLenLimiter()

	pageModel, err := pgm.New(
		appName,
		pageBody,
		ts.MakeTerminalSize(
			rowLenLimiter.RowLenLimit(size.MakeWidth(terminalWidth)),
			size.MakeHeight(terminalHeight),
		),
	)
	if err != nil {
		return nil, errors.Join(ErrNewPageViewPageModel, err)
	}

	return &PageView{
		decTermBox:    termBoxDecor,
		rowLenLimiter: rowLenLimiter,
		pageModel:     pageModel,
		exitCodes: map[rune]struct{}{
			runes.RuneLwQ:   {},
			runes.RuneLwQRu: {},
			runes.RuneUpQ:   {},
			runes.RuneUpQRu: {},
		},
	}, nil
}

var (
	// ErrRunProcess - PageView.process method returned error
	ErrRunProcess = errors.New(`PageView.Run.process call error`)

	// ErrRunTermBoxFlush - term box Flush method returned error
	ErrRunTermBoxFlush = errors.New(`PageView.Run: termbox.Flush call error`)
)

// Run starts display help page session
func (pv *PageView) Run() error {
	defer pv.decTermBox.Close()

	if err := pv.process(0); err != nil {
		return errors.Join(ErrRunProcess, err)
	}

	if err := pv.decTermBox.Flush(); err != nil {
		return errors.Join(ErrRunTermBoxFlush, err)
	}

	for {
		eventTB := pv.decTermBox.PollEvent()
		switch {
		case eventTB.Type == termbox.EventKey:
			switch {
			case eventTB.Key == termbox.KeyArrowDown:
				if err := pv.process(1); err != nil {
					return errors.Join(
						fmt.Errorf(`%w: termbox event type "%v; event key "%v""`,
							ErrRunProcess, eventTB.Type, eventTB.Key),
						err,
					)
				}

			case eventTB.Key == termbox.KeyArrowUp:
				if err := pv.process(-1); err != nil {
					return errors.Join(
						fmt.Errorf(`%w: termbox event type "%v; event key "%v""`,
							ErrRunProcess, eventTB.Type, eventTB.Key),
						err,
					)
				}

			case eventTB.Key == termbox.KeyCtrlTilde:
				if _, contain := pv.exitCodes[eventTB.Ch]; contain {
					return nil
				}
				if err := pv.process(0); err != nil {
					return errors.Join(
						fmt.Errorf(`%w: termbox event type "%v; event key "%v""`,
							ErrRunProcess, eventTB.Type, eventTB.Key),
						err,
					)
				}

			default:
				if err := pv.process(0); err != nil {
					return errors.Join(
						fmt.Errorf(`%w: termbox event type "%v; event key "%v""`,
							ErrRunProcess, eventTB.Type, eventTB.Key),
						err,
					)
				}
			}
		default:
			if err := pv.process(0); err != nil {
				return errors.Join(
					fmt.Errorf(`%w: termbox event type "%v; event key "%v""`,
						ErrRunProcess, eventTB.Type, eventTB.Key),
					err,
				)
			}
		}

		if err := pv.decTermBox.Flush(); err != nil {
			return errors.Join(ErrRunTermBoxFlush, err)
		}
	}
}

// ErrProcessUpdate - update method returned error
var ErrProcessUpdate = errors.New(`PageView.process: update call error`)

func (pv *PageView) process(shift int) error {
	if err := update(pv.decTermBox, pv.pageModel, pv.rowLenLimiter, shift); err != nil {
		return errors.Join(ErrProcessUpdate, err)
	}

	render(pv.decTermBox, ri.MakeRowIterator(pv.pageModel))
	return nil
}

var (
	// ErrUpdateTermBoxClear - TermBoxDecorator.Clear method returned error
	ErrUpdateTermBoxClear = errors.New(`update: TermBoxDecorator.Clear method returned error`)

	// ErrUpdatePageModelUpdate - PageModel.Update method returned error
	ErrUpdatePageModelUpdate = errors.New(`update: PageModel.Update method returned error`)
)

func update(decTermBox tbd.TermBoxDecorator, pageModel *pgm.PageModel, rowLenLimiter rll.RowLenLimiter, shift int) error {
	if err := decTermBox.Clear(); err != nil {
		return errors.Join(ErrUpdateTermBoxClear, err)
	}

	w, h := decTermBox.Size()
	if err := pageModel.Update(
		ts.MakeTerminalSize(rowLenLimiter.RowLenLimit(size.MakeWidth(w)), size.MakeHeight(h)),
		shift,
	); err != nil {
		return errors.Join(ErrUpdatePageModelUpdate, err)
	}

	return nil
}

func render(termBoxDecor tbd.TermBoxDecorator, itRow ri.RowIterator) {
	var (
		x, y int
		cell termbox.Cell
	)
	for ; !itRow.End(); itRow.Next() {
		for x, cell = range itRow.RowModel().GetCells() {
			termBoxDecor.SetCell(itRow.RowModel().GetShiftIndex().Int()+x, y, cell.Ch, cell.Fg, cell.Bg)
		}
		y++
	}
}

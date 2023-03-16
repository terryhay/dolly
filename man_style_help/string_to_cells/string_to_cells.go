package string_to_cells

import (
	"github.com/nsf/termbox-go"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/tools/index"
)

type foreground struct {
	stackFG []termbox.Attribute
	fg      termbox.Attribute
}

func makeForeground(stackSize index.Index) foreground {
	return foreground{
		stackFG: make([]termbox.Attribute, 0, stackSize),
	}
}

func (fg *foreground) Get() termbox.Attribute {
	if fg == nil {
		return termbox.Attribute(0)
	}
	return fg.fg
}

func (fg *foreground) Set(attr termbox.Attribute) {
	if fg == nil {
		return
	}
	fg.stackFG = append(fg.stackFG, attr)
	fg.fg |= attr
}

func (fg *foreground) DropBack() {
	if fg == nil || len(fg.stackFG) == 0 {
		return
	}

	lastFG := fg.stackFG[len(fg.stackFG)-1]
	fg.fg &= ^lastFG
	fg.stackFG = fg.stackFG[:len(fg.stackFG)-1]
}

// StringToCells converts styled string to termbox cells
func StringToCells(fromString string) []termbox.Cell {
	if len(fromString) == 0 {
		return nil
	}

	const countStyleMax = 2
	fg := makeForeground(countStyleMax)

	from := []rune(fromString)
	to := make([]termbox.Cell, 0, len(from))

	for i := 0; i < len(from); i++ {
		r := from[i]

		if r == runes.RuneEsc {
			fg, i = getStyle(fg, i, from)
			continue
		}

		to = append(to, termbox.Cell{
			Ch: r,
			Fg: fg.Get(),
		})
	}

	return to
}

func getStyle(fg foreground, indexBeginStyleSeq int, fromRunes []rune) (foreground, int) {

	indexEndStyleSeq := indexBeginStyleSeq + 1
	for ; indexEndStyleSeq < len(fromRunes); indexEndStyleSeq++ {
		if fromRunes[indexEndStyleSeq] == runes.RuneLwM {
			break
		}
	}

	mark := string(fromRunes[indexBeginStyleSeq+1 : indexEndStyleSeq+1])
	switch {
	case mark == "[1m":
		fg.Set(termbox.AttrBold)
		return fg, indexEndStyleSeq

	case mark == "[0m":
		fg.DropBack()
		return fg, indexEndStyleSeq

	case mark == "[4m":
		fg.Set(termbox.AttrUnderline)
		return fg, indexEndStyleSeq

	default:
		return fg, indexBeginStyleSeq
	}
}

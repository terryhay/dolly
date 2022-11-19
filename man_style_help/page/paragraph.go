package page

import (
	"github.com/nsf/termbox-go"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/runes"
	"github.com/terryhay/dolly/man_style_help/size"
	"strings"
)

// Paragraph contains source paragraph text
type Paragraph struct {
	// text contains paragraph text as string
	text string

	// cells contains paragraph text as termbox cells
	cells []termbox.Cell

	// tabInCells contains amount of empty cells needed to indent a paragraph
	tabInCells size.Width
}

func MakeParagraph(tabCount size.Width, text string) Paragraph {
	return Paragraph{
		text:       text,
		cells:      TextToCells(text),
		tabInCells: tabCount * rll.TabSize,
	}
}

// GetCells returns cells field page
func (p *Paragraph) GetCells() []termbox.Cell {
	return p.cells
}

// GetTabInCells returns tabInCells field page
func (p *Paragraph) GetTabInCells() size.Width {
	return p.tabInCells
}

func (p *Paragraph) String() string {
	return strings.Repeat(" ", int(p.tabInCells)) + p.text
}

func TextToCells(sourceText string) []termbox.Cell {
	if len(sourceText) == 0 {
		return nil
	}

	var (
		r     rune
		style termbox.Attribute
	)

	sourceRunes := []rune(sourceText)
	cells := make([]termbox.Cell, 0, len(sourceRunes))

	for index := 0; index < len(sourceRunes); index++ {
		r = sourceRunes[index]

		if r == runes.RuneEsc {
			style, index = getStyle(sourceRunes, index)
			continue
		}

		cells = append(cells, termbox.Cell{
			Ch: r,
			Fg: style,
		})
	}

	return cells
}

func getStyle(sourceRunes []rune, indexBeginStyleSeq int) (termbox.Attribute, int) {
	indexEndStyleSeq := indexBeginStyleSeq + 1
	for ; indexEndStyleSeq < len(sourceRunes); indexEndStyleSeq++ {
		if sourceRunes[indexEndStyleSeq] == runes.RuneLwM {
			break
		}
	}

	switch string(sourceRunes[indexBeginStyleSeq+1 : indexEndStyleSeq+1]) {
	case "[1m":
		return termbox.AttrBold, indexEndStyleSeq
	case "[0m":
		return 0, indexEndStyleSeq
	case "[4m":
		return termbox.AttrUnderline, indexEndStyleSeq
	}

	return 0, indexEndStyleSeq
}

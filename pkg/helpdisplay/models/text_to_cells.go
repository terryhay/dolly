package models

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/runes"

	"github.com/nsf/termbox-go"
)

func textToCells(sourceText string) []termbox.Cell {
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
	}

	return 0, indexEndStyleSeq
}

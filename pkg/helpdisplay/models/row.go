package models

import "github.com/nsf/termbox-go"

type RowIterator struct {
	ShiftIndex int
	Cells      []termbox.Cell

	reverseCounter  int
	paragraphNumber int
	rowNumber       int
}

func (di *RowIterator) End() bool {
	return di.reverseCounter == 0
}

package data

import (
	"github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"strings"
)

// Paragraph contains source paragraph text
type Paragraph struct {
	// Text - paragraph text
	Text string

	// TabCount - amount of tabs which will be added before every paragraph row
	// How it works: TabCount*tabSize is a rune actionSequence of every paragraph row
	TabCount row_len_limiter.RowSize
}

func (p *Paragraph) String() string {
	return strings.Repeat(" ", int(p.TabCount*row_len_limiter.TabSize)) + p.Text
}

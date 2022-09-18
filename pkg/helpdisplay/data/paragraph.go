package data

import (
	rll "github.com/terryhay/dolly/pkg/helpdisplay/row_len_limiter"
	"github.com/terryhay/dolly/pkg/helpdisplay/size"
	"strings"
)

// Paragraph contains source paragraph text
type Paragraph struct {
	// Text - paragraph text
	Text string

	// TabCount - amount of tabs which will be added before every paragraph row
	// How it works: TabCount*tabSize is a rune actionSequence of every paragraph row
	TabCount size.Width
}

func (p *Paragraph) String() string {
	return strings.Repeat(" ", int(p.TabCount*rll.TabSize)) + p.Text
}

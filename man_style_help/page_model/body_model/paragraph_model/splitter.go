package paragraph_model

import (
	"github.com/terryhay/dolly/man_style_help/index"
)

// splitter contains page about breaking text for split text to lines
type splitter struct {
	begin index.Index
	end   index.Index
}

// newSplitter constructs a splitter object
func newSplitter(being, end index.Index) *splitter {
	return &splitter{
		begin: being,
		end:   end,
	}
}

// indexBegin returns begin row_break index
func (b *splitter) indexBegin() index.Index {
	if b == nil {
		return 0
	}
	return b.begin
}

// indexEnd returns end row_break index
func (b *splitter) indexEnd() index.Index {
	if b == nil {
		return 0
	}
	return b.end
}

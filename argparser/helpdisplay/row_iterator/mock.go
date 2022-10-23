package row_iterator

import (
	"github.com/terryhay/dolly/argparser/helpdisplay/index"
	"github.com/terryhay/dolly/argparser/helpdisplay/page_model"
	"github.com/terryhay/dolly/argparser/helpdisplay/row"
	"github.com/terryhay/dolly/argparser/helpdisplay/size"
)

// Mock contains mock data of initialize mocked iterator object
type Mock struct {
	Model             *page_model.PageModel
	ModelRow          row.Row
	ReverseCounter    size.Height
	ParagraphIndex    index.Index
	ParagraphRowIndex index.Index
}

// MockRowIterator constructs Mocked RowIterator
func MockRowIterator(m Mock) RowIterator {
	return RowIterator{
		model: m.Model,

		modelRow: m.ModelRow,

		reverseCounter:    m.ReverseCounter,
		paragraphIndex:    m.ParagraphIndex,
		paragraphRowIndex: m.ParagraphRowIndex,
	}
}

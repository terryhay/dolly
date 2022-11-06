package row_iterator

import (
	"github.com/terryhay/dolly/man_style_help/index"
	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	"github.com/terryhay/dolly/man_style_help/row"
	"github.com/terryhay/dolly/man_style_help/size"
)

// Mock contains mock data of initialize mocked iterator object
type Mock struct {
	Model             *pgm.PageModel
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

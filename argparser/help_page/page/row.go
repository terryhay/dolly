package page

import (
	"strings"

	"github.com/terryhay/dolly/tools/size"
)

// Row - paragraph row
type Row struct {
	textStyled    string
	textRuneCount size.Width
	marginLeft    size.Width
}

// MakeRow constructs Row object on stack
func MakeRow(marginLeft size.Width, chunks ...RowChunk) Row {
	if len(chunks) == 0 {
		return Row{}
	}

	textRuneCount := size.WidthZero
	builder := strings.Builder{}

	for _, chunk := range chunks {
		textRuneCount += size.MakeWidth(len(chunk.GetText()))
		builder.WriteString(CreateStyledText(chunk))
	}

	return Row{
		textStyled:    builder.String(),
		textRuneCount: textRuneCount,
		marginLeft:    marginLeft,
	}
}

// GetTextStyled gets textStyled field
func (r *Row) GetTextStyled() string {
	if r == nil {
		return ""
	}
	return r.textStyled
}

// GetTextRuneCount gets textRuneCount field
func (r *Row) GetTextRuneCount() size.Width {
	if r == nil {
		return size.WidthZero
	}
	return r.textRuneCount
}

// GetMarginLeft gets marginLeft field
func (r *Row) GetMarginLeft() size.Width {
	if r == nil {
		return size.WidthZero
	}
	return r.marginLeft
}

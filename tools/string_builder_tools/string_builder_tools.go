package string_builder_tools

import (
	"strings"

	"github.com/terryhay/dolly/tools/size"
)

const (
	// ChunkSpace - space
	ChunkSpace = " "

	// ChunkBreakRow - break row
	ChunkBreakRow = "\n"
)

// Append appends the text without new line pasting
func Append(builder *strings.Builder, textChunks ...string) *strings.Builder {
	if builder == nil {
		return nil
	}

	for _, chunk := range textChunks {
		builder.WriteString(chunk)
	}

	return builder
}

// BreakRow pastes a row break
func BreakRow(builder *strings.Builder) *strings.Builder {
	if builder == nil {
		return builder
	}

	builder.WriteString(ChunkBreakRow)
	return builder
}

// NewRow paste the row with insert a break line before
func NewRow(builder *strings.Builder, marginLeft size.Width, chunks ...string) *strings.Builder {
	if builder == nil {
		return builder
	}

	builder.WriteString(ChunkBreakRow)

	if marginLeft > size.WidthZero {
		builder.WriteString(strings.Repeat(ChunkSpace, marginLeft.Int()))
	}

	for _, chunk := range chunks {
		builder.WriteString(chunk)
	}

	return builder
}

package page

import (
	"regexp"
	"strings"

	"github.com/terryhay/dolly/tools/size"
)

// StyleTextMask - style of text byte mask
type StyleTextMask int

const (
	// StyleTextDefault - default value of StyleTextMask
	StyleTextDefault StyleTextMask = 0

	// StyleTextBold - use bold text
	StyleTextBold StyleTextMask = 0x1

	// StyleTextUnderlined - use underlined text
	StyleTextUnderlined StyleTextMask = 0x2
)

// Has returns if StyleTextMask has set StyleTextMask
func (m StyleTextMask) Has(check StyleTextMask) bool {
	return m&check == check
}

// RowChunk contains data for display styled text part of row
type RowChunk struct {
	text  string
	style StyleTextMask
}

// MakeRowChunk creates RowChunk
func MakeRowChunk(text string, styles ...StyleTextMask) RowChunk {
	return RowChunk{
		text: text,
		style: func() StyleTextMask {
			if len(styles) == 0 {
				return StyleTextDefault
			}

			textStyle := StyleTextDefault
			for _, style := range styles {
				textStyle |= style
			}

			return textStyle
		}(),
	}
}

// MakeRowChunkSpaces creates RowChunk with countSpaces spaces
func MakeRowChunkSpaces(countSpaces size.Width) RowChunk {
	const space = " "
	return RowChunk{
		text: strings.Repeat(space, countSpaces.Int()),
	}
}

// GetStyle gets style field
func (tc *RowChunk) GetStyle() StyleTextMask {
	if tc == nil {
		return StyleTextDefault
	}
	return tc.style
}

// GetText gets text field
func (tc *RowChunk) GetText() string {
	if tc == nil {
		return ""
	}
	return tc.text
}

// CountRunes returns size text field as rune slice
func (tc *RowChunk) CountRunes() size.Width {
	if tc == nil {
		return size.WidthZero
	}
	return size.MakeWidth(len([]rune(tc.text)))
}

const (
	textStyleBoldOpen       = `[1m`
	textStyleUnderlinedOpen = `[4m`

	textStyleClose = `[0m`
)

// CreateStyledText returns text field with applied style
func CreateStyledText(tc RowChunk) string {
	if tc.GetStyle() == StyleTextDefault {
		return tc.text
	}

	builder := strings.Builder{}
	if tc.GetStyle().Has(StyleTextBold) {
		builder.WriteString(textStyleBoldOpen)
	}
	if tc.GetStyle().Has(StyleTextUnderlined) {
		builder.WriteString(textStyleUnderlinedOpen)
	}

	builder.WriteString(tc.text)

	if tc.GetStyle().Has(StyleTextBold) {
		builder.WriteString(textStyleClose)
	}
	if tc.GetStyle().Has(StyleTextUnderlined) {
		builder.WriteString(textStyleClose)
	}

	return builder.String()
}

var styleTextMarkersRemovePattern = regexp.MustCompile(`(\[1m)|(\[0m)|(\[4m)`)

// RemoveStyleTextMarkers removes styled markers from string
func RemoveStyleTextMarkers(text string) string {
	return styleTextMarkersRemovePattern.ReplaceAllString(text, "")
}

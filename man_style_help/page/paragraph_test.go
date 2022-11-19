package page

import (
	"github.com/stretchr/testify/require"
	rll "github.com/terryhay/dolly/man_style_help/row_len_limiter"
	"github.com/terryhay/dolly/man_style_help/size"
	"testing"
)

func TestParagraph(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		caseName string

		text              string
		tabCount          size.Width
		expectedWidth     size.Width
		expectedCellCount int
		expectedString    string
	}{
		{
			caseName: "paragraph_with_defined_styles",

			text:              `[1mappname command[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]`,
			tabCount:          1,
			expectedWidth:     1 * rll.TabSize,
			expectedCellCount: 49,
			expectedString:    "    \x1b[1mappname command\x1b[0m \x1b[1m-sa\x1b[0m \x1b[4marg\x1b[0m [\x1b[1m-la\x1b[0m \x1b[4mstr\x1b[0m=val1 [val2] \x1b[4m...\x1b[0m]",
		},
		{
			caseName: "paragraph_with_undefined_styles",

			text:              `[9mmappname command -sa arg [-la str=val1 [val2] ...][0m`,
			expectedWidth:     1 * rll.TabSize,
			expectedCellCount: 50,
			expectedString:    "\x1b[9mmappname command -sa arg [-la str=val1 [val2] ...]\u001B[0m",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			prg := MakeParagraph(tc.tabCount, tc.text)
			require.Equal(t, tc.tabCount*rll.TabSize, prg.GetTabInCells())
			require.Equal(t, tc.expectedCellCount, len(prg.GetCells()))
			require.Equal(t, tc.expectedString, prg.String())
		})
	}
}

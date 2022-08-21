package data

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		config   dollyconf.ArgParserConfig
		expected Page
	}{
		{
			caseName: "no_data",
			expected: Page{
				Paragraphs: []*Paragraph{
					{
						Text: nameChapterTitle,
					},
					{
						TabCount: 1,
						Text:     "\u001B[1m\u001B[0m – ",
					},
					{
						Text: synopsisChapterTitle,
					},
				},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			hpd := MakePage(td.config)

			require.Equal(t, len(td.expected.Paragraphs), len(hpd.Paragraphs))
			assert.Equal(t, td.expected, hpd)
		})
	}
}

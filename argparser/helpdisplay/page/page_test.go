package page

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"testing"
)

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		config   apConf.ArgParserConfig
		expected Page
	}{
		{
			caseName: "no_data",
			expected: Page{
				Paragraphs: []Paragraph{
					MakeParagraph(0, nameChapterTitle),
					MakeParagraph(1, "\u001B[1m\u001B[0m – "),
					MakeParagraph(0, synopsisChapterTitle),
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

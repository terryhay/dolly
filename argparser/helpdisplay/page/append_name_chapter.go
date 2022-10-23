package page

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
)

const nameChapterTitle = `[1mNAME[0m`
const nameChapterParagraphTemplate = `[1m%s[0m – %s`

// appendNameChapterParagraphs creates and appends NAME chapter paragraphs
func appendNameChapterParagraphs(paragraphs []Paragraph, appDescription apConf.ApplicationDescription) []Paragraph {
	return append(paragraphs,
		MakeParagraph(0, nameChapterTitle),
		MakeParagraph(
			1,
			fmt.Sprintf(nameChapterParagraphTemplate,
				appDescription.GetAppName(),
				appDescription.GetNameHelpInfo()),
		),
	)
}

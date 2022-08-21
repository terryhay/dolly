package data

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
)

const nameChapterTitle = `[1mNAME[0m`
const nameChapterParagraphTemplate = `[1m%s[0m – %s`

// appendNameChapterParagraphs creates and appends NAME chapter paragraphs
func appendNameChapterParagraphs(paragraphs []*Paragraph, appDescription dollyconf.ApplicationDescription) []*Paragraph {
	return append(paragraphs,
		&Paragraph{
			Text: nameChapterTitle,
		},
		&Paragraph{
			TabCount: 1,
			Text: fmt.Sprintf(nameChapterParagraphTemplate,
				appDescription.GetAppName(),
				appDescription.GetNameHelpInfo()),
		},
	)
}

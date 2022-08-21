package data

import (
	"github.com/terryhay/dolly/pkg/dollyconf"
)

const paragraphDefaultCount = 100

type Page struct {
	Header     string
	Paragraphs []*Paragraph
}

// MakePage constructs Page object in a stack
func MakePage(config dollyconf.ArgParserConfig) Page {
	appDescription := config.GetAppDescription()
	namelessCmdDescription := config.GetNamelessCommandDescription()
	cmdDescriptions := config.GetCommandDescriptions()
	flagDescriptions := config.GetFlagDescriptions()

	paragraphs := make([]*Paragraph, 0, paragraphDefaultCount)

	paragraphs = appendNameChapterParagraphs(paragraphs,
		appDescription)
	paragraphs = appendSynopsisChapterParagraphs(paragraphs,
		appDescription,
		namelessCmdDescription,
		cmdDescriptions,
		flagDescriptions)
	paragraphs = appendDescriptionChapter(paragraphs,
		appDescription,
		namelessCmdDescription,
		cmdDescriptions,
		flagDescriptions)

	return Page{
		Header:     config.GetAppDescription().GetAppName(),
		Paragraphs: paragraphs,
	}

}

package page

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
)

const paragraphDefaultCount = 100

type Page struct {
	Header     Paragraph
	Paragraphs []Paragraph
}

// MakePage constructs Page object in a stack
func MakePage(config apConf.ArgParserConfig) Page {
	appDescription := config.GetAppDescription()
	namelessCmdDescription := config.GetNamelessCommandDescription()
	cmdDescriptions := config.GetCommandDescriptions()
	flagDescriptions := config.GetFlagDescriptions()

	paragraphs := make([]Paragraph, 0, paragraphDefaultCount)

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
		Header:     MakeParagraph(0, config.GetAppDescription().GetAppName()),
		Paragraphs: paragraphs,
	}
}

package page

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
)

const (
	descriptionChapterTitle = "\u001B[1mDESCRIPTION\u001B[0m"

	commandDescriptionsSubtitle = "The commands are as follows:"
	descriptionLine             = "\u001B[1m%s\u001B[0m %s"

	descriptionTwoLines = "\u001B[1m%s\u001B[0m"

	flagDescriptionsSubtitle = "The flags are as follows:"

	namelessCommandDescriptionName = "<empty>"
)

const tabLen = 7

// appendDescriptionChapter creates and appends DESCRIPTION chapter paragraphs
func appendDescriptionChapter(
	paragraphs []Paragraph,
	appDescription apConf.ApplicationDescription,
	namelessCmdDescription apConf.NamelessCommandDescription,
	cmdDescriptions []*apConf.CommandDescription,
	flagDescriptions map[apConf.Flag]*apConf.FlagDescription,
) []Paragraph {

	if len(appDescription.GetDescriptionHelpInfo()) == 0 &&
		namelessCmdDescription == nil &&
		len(cmdDescriptions) == 0 &&
		len(flagDescriptions) == 0 {

		return paragraphs
	}

	var (
		callNames       string
		flagDescription *apConf.FlagDescription
	)

	paragraphs = append(paragraphs,
		MakeParagraph(0, ""),
		MakeParagraph(0, descriptionChapterTitle),
	)

	if descriptionHelpInfo := appDescription.GetDescriptionHelpInfo(); len(descriptionHelpInfo) > 0 {
		for i := range descriptionHelpInfo {
			if i != 0 {
				paragraphs = append(paragraphs, MakeParagraph(0, ""))
			}

			paragraphs = append(paragraphs,
				MakeParagraph(
					1,
					descriptionHelpInfo[i],
				),
			)
		}
	}

	if commandDescriptions := cmdDescriptions; len(commandDescriptions) > 0 {
		paragraphs = append(paragraphs,
			MakeParagraph(0, ""),
			MakeParagraph(0, commandDescriptionsSubtitle),
		)

		if namelessCommandDescription := namelessCmdDescription; namelessCommandDescription != nil {
			paragraphs = append(paragraphs,
				MakeParagraph(
					1,
					fmt.Sprintf(descriptionLine,
						namelessCommandDescriptionName,
						namelessCommandDescription.GetDescriptionHelpInfo()),
				),
			)
		}

		for i := range commandDescriptions {
			callNames = strings.Join(getSortedCommands(commandDescriptions[i].GetCommands()), ", ")

			if i != 0 || namelessCmdDescription != nil {
				// append empty paragraph
				paragraphs = append(paragraphs, MakeParagraph(0, ""))
			}

			if len(callNames) > tabLen {
				paragraphs = append(paragraphs,
					MakeParagraph(
						1,
						fmt.Sprintf(descriptionTwoLines, callNames),
					),
					MakeParagraph(
						2,
						commandDescriptions[i].GetDescriptionHelpInfo(),
					),
				)
				continue
			}

			paragraphs = append(paragraphs,
				MakeParagraph(
					1,
					fmt.Sprintf(descriptionLine,
						callNames,
						commandDescriptions[i].GetDescriptionHelpInfo()),
				),
			)
		}
	}

	if len(flagDescriptions) > 0 {
		paragraphs = append(paragraphs,
			MakeParagraph(0, ""),
			MakeParagraph(0, flagDescriptionsSubtitle),
		)

		callNamesSlice := getSortedFlagsForDescription(flagDescriptions)
		for i := range callNamesSlice {
			callNames = callNamesSlice[i]

			if i != 0 {
				// append empty paragraph
				paragraphs = append(paragraphs, MakeParagraph(0, ""))
			}

			flagDescription = flagDescriptions[apConf.Flag(callNames)]

			if len(callNames) > tabLen {
				paragraphs = append(paragraphs,
					MakeParagraph(
						1,
						fmt.Sprintf(descriptionTwoLines, callNames),
					),
					MakeParagraph(
						2,
						flagDescription.GetDescriptionHelpInfo(),
					),
				)

				continue
			}

			paragraphs = append(paragraphs,
				MakeParagraph(
					1,
					fmt.Sprintf(descriptionLine, callNames, flagDescription.GetDescriptionHelpInfo()),
				),
			)
		}
	}

	return paragraphs
}

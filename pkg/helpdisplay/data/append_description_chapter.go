package data

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
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
	paragraphs []*Paragraph,
	appDescription dollyconf.ApplicationDescription,
	namelessCmdDescription dollyconf.NamelessCommandDescription,
	cmdDescriptions []*dollyconf.CommandDescription,
	flagDescriptions map[dollyconf.Flag]*dollyconf.FlagDescription,
) []*Paragraph {

	if len(appDescription.GetDescriptionHelpInfo()) == 0 &&
		namelessCmdDescription == nil &&
		len(cmdDescriptions) == 0 &&
		len(flagDescriptions) == 0 {

		return paragraphs
	}

	var (
		callNames       string
		flagDescription *dollyconf.FlagDescription
	)

	paragraphs = append(paragraphs,
		&Paragraph{
			Text: descriptionChapterTitle,
		},
	)

	if descriptionHelpInfo := appDescription.GetDescriptionHelpInfo(); len(descriptionHelpInfo) > 0 {
		for i := range descriptionHelpInfo {
			if i != 0 {
				paragraphs = append(paragraphs, &Paragraph{})
			}

			paragraphs = append(paragraphs,
				&Paragraph{
					TabCount: 1,
					Text:     descriptionHelpInfo[i],
				},
			)
		}
	}

	if commandDescriptions := cmdDescriptions; len(commandDescriptions) > 0 {
		paragraphs = append(paragraphs,
			&Paragraph{},
			&Paragraph{
				Text: commandDescriptionsSubtitle,
			},
		)

		if namelessCommandDescription := namelessCmdDescription; namelessCommandDescription != nil {
			paragraphs = append(paragraphs,
				&Paragraph{
					TabCount: 1,
					Text: fmt.Sprintf(descriptionLine,
						namelessCommandDescriptionName,
						namelessCommandDescription.GetDescriptionHelpInfo()),
				},
			)
		}

		for i := range commandDescriptions {
			callNames = strings.Join(getSortedCommands(commandDescriptions[i].GetCommands()), ", ")

			if i != 0 || namelessCmdDescription != nil {
				// append empty paragraph
				paragraphs = append(paragraphs, &Paragraph{})
			}

			if len(callNames) > tabLen {
				paragraphs = append(paragraphs,
					&Paragraph{
						TabCount: 1,
						Text:     fmt.Sprintf(descriptionTwoLines, callNames),
					},
					&Paragraph{
						TabCount: 2,
						Text:     commandDescriptions[i].GetDescriptionHelpInfo(),
					},
				)
				continue
			}

			paragraphs = append(paragraphs,
				&Paragraph{
					TabCount: 1,
					Text: fmt.Sprintf(descriptionLine,
						callNames,
						commandDescriptions[i].GetDescriptionHelpInfo()),
				},
			)
		}
	}

	if len(flagDescriptions) > 0 {
		paragraphs = append(paragraphs,
			&Paragraph{},
			&Paragraph{
				Text: flagDescriptionsSubtitle,
			},
		)

		callNamesSlice := getSortedFlagsForDescription(flagDescriptions)
		for i := range callNamesSlice {
			callNames = callNamesSlice[i]

			if i != 0 {
				// append empty paragraph
				paragraphs = append(paragraphs, &Paragraph{})
			}

			flagDescription = flagDescriptions[dollyconf.Flag(callNames)]

			if len(callNames) > tabLen {
				paragraphs = append(paragraphs,
					&Paragraph{
						TabCount: 1,
						Text:     fmt.Sprintf(descriptionTwoLines, callNames),
					},
					&Paragraph{
						TabCount: 2,
						Text:     flagDescription.GetDescriptionHelpInfo(),
					},
				)

				continue
			}

			paragraphs = append(paragraphs,
				&Paragraph{
					TabCount: 1,
					Text:     fmt.Sprintf(descriptionLine, callNames, flagDescription.GetDescriptionHelpInfo()),
				},
			)
		}
	}

	return paragraphs
}

package helpprinter

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
)

const (
	descriptionChapterTitle = "\u001B[1mDESCRIPTION\u001B[0m\n"

	commonDescriptionParagraphs = "\t%s\n\n"

	commandDescriptionsSubtitle = "The commands are as follows:"
	descriptionLine             = "\n\t\u001B[1m%s\u001B[0m\t%s\n"
	descriptionTwoLines         = "\n\t\u001B[1m%s\u001B[0m\n\t\t%s\n"

	flagDescriptionsSubtitle = "The flags are as follows:"

	namelessCommandDescriptionName = "<empty>"
)

const tabLen = 7

// CreateDescriptionChapter - create s description help chapter
func CreateDescriptionChapter(
	descriptionHelpInfo []string,
	namelessCommandDescription dollyconf.NamelessCommandDescription,
	commandDescriptions []*dollyconf.CommandDescription,
	flagDescriptions map[dollyconf.Flag]*dollyconf.FlagDescription,
) string {

	var (
		builder         strings.Builder
		callNames       string
		flagDescription *dollyconf.FlagDescription
		usingPattern    string
	)

	builder.WriteString(descriptionChapterTitle)
	if len(descriptionHelpInfo) == 0 &&
		namelessCommandDescription == nil &&
		len(commandDescriptions) == 0 &&
		len(flagDescriptions) == 0 {

		return builder.String()
	}

	commonParagraphPart := "\n"
	if len(descriptionHelpInfo) > 0 {
		commonParagraphPart = fmt.Sprintf(commonDescriptionParagraphs, strings.Join(descriptionHelpInfo, "\n\n\t"))
	}
	builder.WriteString(commonParagraphPart)

	if len(commandDescriptions) > 0 {
		builder.WriteString(commandDescriptionsSubtitle)

		if namelessCommandDescription != nil {
			builder.WriteString(fmt.Sprintf(descriptionLine,
				namelessCommandDescriptionName,
				namelessCommandDescription.GetDescriptionHelpInfo()))
		}

		for i := range commandDescriptions {
			callNames = strings.Join(getSortedCommands(commandDescriptions[i].GetCommands()), ", ")

			usingPattern = descriptionLine
			if len(callNames) > tabLen {
				usingPattern = descriptionTwoLines
			}

			builder.WriteString(fmt.Sprintf(usingPattern,
				callNames,
				commandDescriptions[i].GetDescriptionHelpInfo()))
		}
	}

	if len(flagDescriptions) > 0 {
		builder.WriteString(flagDescriptionsSubtitle)

		for _, callNames = range getSortedFlagsForDescription(flagDescriptions) {
			usingPattern = descriptionLine
			if len(callNames) > tabLen {
				usingPattern = descriptionTwoLines
			}

			flagDescription = flagDescriptions[dollyconf.Flag(callNames)]
			builder.WriteString(fmt.Sprintf(usingPattern, callNames, flagDescription.GetDescriptionHelpInfo()))
		}
	}

	return builder.String()
}

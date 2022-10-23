package page

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
)

const (
	synopsisChapterTitle            = "\u001B[1mSYNOPSIS\u001B[0m"
	synopsisLineCommandPart         = "\u001B[1m%s %s\u001B[0m%s"
	synopsisLineNamelessCommandPart = "\u001B[1m%s\u001B[0m%s"
)

// appendSynopsisChapterParagraphs creates and appends SYNOPSIS chapter paragraphs
func appendSynopsisChapterParagraphs(
	paragraphs []Paragraph,
	appDescription apConf.ApplicationDescription,
	namelessCmdDescription apConf.NamelessCommandDescription,
	cmdDescriptions []*apConf.CommandDescription,
	flagDescriptions map[apConf.Flag]*apConf.FlagDescription,
) []Paragraph {

	paragraphs = append(paragraphs,
		MakeParagraph(0, synopsisChapterTitle))

	if namelessCmdDescription != nil {
		paragraphs = append(paragraphs,
			MakeParagraph(
				1,
				createCommandSynopsisParagraph(
					appDescription.GetAppName(),
					namelessCmdDescription.(*apConf.CommandDescription),
					flagDescriptions),
			))
	}

	for _, commandDescription := range cmdDescriptions {
		paragraphs = append(paragraphs,
			MakeParagraph(
				1,
				createCommandSynopsisParagraph(
					appDescription.GetAppName(),
					commandDescription,
					flagDescriptions,
				),
			),
		)
	}

	return paragraphs
}

func createCommandSynopsisParagraph(
	appName string,
	commandDescription *apConf.CommandDescription,
	flagDescriptions map[apConf.Flag]*apConf.FlagDescription) string {

	builder := strings.Builder{}

	if len(commandDescription.GetCommands()) > 0 {
		builder.WriteString(fmt.Sprintf(synopsisLineCommandPart,
			appName,
			strings.Join(getSortedCommands(commandDescription.GetCommands()), ", "),
			createArgumentsPart(commandDescription.GetArgDescription())))
	} else {
		builder.WriteString(fmt.Sprintf(synopsisLineNamelessCommandPart,
			appName,
			strings.Join(getSortedCommands(commandDescription.GetCommands()), ", ")))
	}

	var (
		flag            string
		flagDescription *apConf.FlagDescription
	)

	// required flags part
	for _, flag = range getSortedFlags(commandDescription.GetRequiredFlags()) {
		flagDescription = flagDescriptions[apConf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))
	}

	// optional flags part
	for _, flag = range getSortedFlags(commandDescription.GetOptionalFlags()) {
		flagDescription = flagDescriptions[apConf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" [\u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))

		builder.WriteString("]")
	}

	return builder.String()
}

func createArgumentsPart(argDescription *apConf.ArgumentsDescription) string {
	if argDescription == nil {
		return ""
	}

	var builder strings.Builder

	defaultValuesTemplatePart := ""
	if len(argDescription.GetDefaultValues()) > 0 {
		defaultValuesTemplatePart = fmt.Sprintf(`=%s`, strings.Join(argDescription.GetDefaultValues(), ", "))
	}

	allowedValuesTemplatePart := ""
	joinedString := strings.Join(getSortedStrings(argDescription.GetAllowedValues()), ", ")
	if len(joinedString) > 0 {
		allowedValuesTemplatePart = fmt.Sprintf(` [%s]`, joinedString)
	}

	switch argDescription.GetAmountType() {
	case apConf.ArgAmountTypeSingle:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))

	case apConf.ArgAmountTypeList:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s [4m...[0m`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))

	default:
		return ""
	}

	return builder.String()
}

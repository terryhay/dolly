package plain_help_out

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
)

const (
	synopsisChapterTitle            = "\u001B[1mSYNOPSIS\u001B[0m\n"
	synopsisLineCommandPart         = "\t\u001B[1m%s %s\u001B[0m%s"
	synopsisLineNamelessCommandPart = "\t\u001B[1m%s\u001B[0m%s"
)

// createSynopsisChapter creates synopsis help chapter
func createSynopsisChapter(
	appName string,
	namelessCommandDescription apConf.NamelessCommandDescription,
	commandDescriptions []*apConf.CommandDescription,
	flagDescriptions map[apConf.Flag]*apConf.FlagDescription,
) string {

	var builder strings.Builder
	builder.WriteString(synopsisChapterTitle)

	if namelessCommandDescription != nil {
		addCommandSynopsisLine(
			&builder,
			appName,
			namelessCommandDescription.(*apConf.CommandDescription),
			flagDescriptions,
		)
	}

	for _, commandDescription := range commandDescriptions {
		addCommandSynopsisLine(
			&builder,
			appName,
			commandDescription,
			flagDescriptions,
		)
	}
	builder.WriteString("\n")

	return builder.String()
}

func addCommandSynopsisLine(
	builder *strings.Builder,
	appName string,
	descriptionCommand *apConf.CommandDescription,
	descriptionFlags map[apConf.Flag]*apConf.FlagDescription,
) {
	if len(descriptionCommand.GetCommands()) > 0 {
		builder.WriteString(fmt.Sprintf(synopsisLineCommandPart,
			appName,
			strings.Join(getSortedCommands(descriptionCommand.GetCommands()), ", "),
			createArgumentsPart(descriptionCommand.GetArgDescription())))
	} else {
		builder.WriteString(fmt.Sprintf(synopsisLineNamelessCommandPart,
			appName,
			strings.Join(getSortedCommands(descriptionCommand.GetCommands()), ", ")))
	}

	// required flags part
	for _, flag := range getSortedFlags(descriptionCommand.GetRequiredFlags()) {
		flagDescription := descriptionFlags[apConf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))
	}

	// optional flags part
	for _, flag := range getSortedFlags(descriptionCommand.GetOptionalFlags()) {
		flagDescription := descriptionFlags[apConf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" [\u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))

		builder.WriteString("]")
	}

	builder.WriteString("\n")
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

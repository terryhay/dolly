package helpprinter

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
)

const (
	synopsisChapterTitle            = "\u001B[1mSYNOPSIS\u001B[0m\n"
	synopsisLineCommandPart         = "\t\u001B[1m%s %s\u001B[0m%s"
	synopsisLineNamelessCommandPart = "\t\u001B[1m%s\u001B[0m%s"
)

// CreateSynopsisChapter creates synopsis help chapter
func CreateSynopsisChapter(
	appName string,
	namelessCommandDescription dollyconf.NamelessCommandDescription,
	commandDescriptions []*dollyconf.CommandDescription,
	flagDescriptions map[dollyconf.Flag]*dollyconf.FlagDescription,
) string {

	var builder strings.Builder
	builder.WriteString(synopsisChapterTitle)

	if namelessCommandDescription != nil {
		addCommandSynopsisLine(
			&builder,
			appName,
			namelessCommandDescription.(*dollyconf.CommandDescription),
			flagDescriptions)
	}

	for _, commandDescription := range commandDescriptions {
		addCommandSynopsisLine(
			&builder,
			appName,
			commandDescription,
			flagDescriptions)
	}
	builder.WriteString("\n")

	return builder.String()
}

func addCommandSynopsisLine(
	builder *strings.Builder,
	appName string,
	commandDescription *dollyconf.CommandDescription,
	flagDescriptions map[dollyconf.Flag]*dollyconf.FlagDescription) {

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
		flagDescription *dollyconf.FlagDescription
	)

	// required flags part
	for _, flag = range getSortedFlags(commandDescription.GetRequiredFlags()) {
		flagDescription = flagDescriptions[dollyconf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))
	}

	// optional flags part
	for _, flag = range getSortedFlags(commandDescription.GetOptionalFlags()) {
		flagDescription = flagDescriptions[dollyconf.Flag(flag)]

		builder.WriteString(fmt.Sprintf(" [\u001B[1m%s\u001B[0m", flag))
		builder.WriteString(createArgumentsPart(flagDescription.GetArgDescription()))

		builder.WriteString("]")
	}

	builder.WriteString("\n")
}

func createArgumentsPart(argDescription *dollyconf.ArgumentsDescription) string {
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
	case dollyconf.ArgAmountTypeSingle:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))
	case dollyconf.ArgAmountTypeList:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s [4m...[0m`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))
	default:
		return ""
	}

	return builder.String()
}

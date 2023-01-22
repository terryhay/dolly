package id_template_data_creator

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"regexp"
	"unicode"
)

const (
	// PrefixCommandID - using prefix for command id variables naming
	PrefixCommandID = "CommandID"

	// PrefixCommandStringID - using prefix for command variables naming
	PrefixCommandStringID = "Command"

	// PrefixFlagStringID - using prefix for flag variables naming
	PrefixFlagStringID = "Flag"

	// NamelessCommandIDPostfix - using postfix for nameless command naming
	NamelessCommandIDPostfix = "NamelessCommand"
)

const (
	helpCommandIDStr   = "PrintHelpInfo"
	helpCommandComment = "print help info"
)

// IDTemplateDataCreator creates slices of id template page
type IDTemplateDataCreator struct {
	dashRemover *regexp.Regexp
}

// NewIDTemplateCreator constructs IDTemplateDataCreator object
func NewIDTemplateCreator() IDTemplateDataCreator {
	return IDTemplateDataCreator{dashRemover: regexp.MustCompile("-+")}
}

// RemoveDashes removes all dashes from a string
func (i IDTemplateDataCreator) RemoveDashes(str string) string {
	return i.dashRemover.ReplaceAllString(str, "")
}

// CreateID creates ID string by call name
func (i IDTemplateDataCreator) CreateID(prefix string, callNameStr string) string {
	callNameRunes := []rune(i.RemoveDashes(callNameStr))

	if len(callNameRunes) == 0 {
		return ""
	}
	if len(callNameRunes) == 1 {
		const (
			postfixUp = "Up"
			postfixLw = "Lw"
		)
		id := prefix + string(unicode.ToUpper(callNameRunes[0])) + postfixLw
		if unicode.IsUpper(callNameRunes[0]) {
			id = prefix + string(unicode.ToUpper(callNameRunes[0])) + postfixUp
		}
		return id
	}

	return prefix + string(unicode.ToUpper(callNameRunes[0])) + string(callNameRunes[1:])
}

// CreateIDTemplateData creates IDTemplateData slices for commands and flags
func (i IDTemplateDataCreator) CreateIDTemplateData(
	commandDescriptions []*confYML.CommandDescription,
	helpCommandDescription *confYML.HelpCommandDescription,
	nullCommandDescription *confYML.NamelessCommandDescription,
	flagDescriptionMap map[string]*confYML.FlagDescription,
) (
	commandsIDTemplateData map[string]*IDTemplateData,
	nullCommandIDTemplateData *IDTemplateData,
	flagsIDTemplateData map[string]*IDTemplateData,
) {
	var commandDescription *confYML.CommandDescription

	commandsIDTemplateData = make(map[string]*IDTemplateData, len(commandDescriptions))
	flagsIDTemplateData = make(map[string]*IDTemplateData, len(flagDescriptionMap))

	// standard commands
	for j := range commandDescriptions {
		commandDescription = commandDescriptions[j]

		callName := commandDescription.GetCommand()
		commandId := i.CreateID(PrefixCommandID, callName)

		commandsIDTemplateData[commandDescription.GetCommand()] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			commandDescription.GetDescriptionHelpInfo(),
		)

		for k := range commandDescription.GetAdditionalCommands() {
			callName = commandDescription.GetAdditionalCommands()[k]
			commandsIDTemplateData[commandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
				commandId,
				i.CreateID(PrefixCommandStringID, callName),
				callName,
				commandDescription.GetDescriptionHelpInfo(),
			)
		}
	}

	// help command
	if helpCommandDescription != nil {
		commandId := i.CreateID(PrefixCommandID, helpCommandIDStr)
		callName := helpCommandDescription.GetCommand()

		commandsIDTemplateData[helpCommandDescription.GetCommand()] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			helpCommandComment,
		)

		for k := range helpCommandDescription.GetAdditionalCommands() {
			callName = helpCommandDescription.GetAdditionalCommands()[k]
			commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
				commandId,
				i.CreateID(PrefixCommandStringID, callName),
				callName,
				helpCommandComment,
			)
		}
	}

	// null command
	if nullCommandDescription != nil {
		nullCommandIDTemplateData = NewIDTemplateData(
			i.CreateID(PrefixCommandID, NamelessCommandIDPostfix),
			"",
			"",
			nullCommandDescription.GetDescriptionHelpInfo(),
		)
	}

	// flags
	for _, flagDescription := range flagDescriptionMap {
		callName := flagDescription.GetFlag()
		flagsIDTemplateData[flagDescription.GetFlag()] = NewIDTemplateData(
			"",
			i.CreateID(PrefixFlagStringID, callName),
			callName,
			flagDescription.GetDescriptionHelpInfo(),
		)
	}

	return commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData
}

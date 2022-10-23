package id_template_data_creator

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"regexp"
	"unicode"
)

const (
	PrefixCommandID       = "CommandID"
	PrefixCommandStringID = "Command"
	PrefixFlagStringID    = "Flag"

	NamelessCommandIDPostfix = "NamelessCommand"
)

const (
	helpCommandIDStr   = "PrintHelpInfo"
	helpCommandComment = "print help info"
)

// IDTemplateDataCreator - creates slices of id template page
type IDTemplateDataCreator struct {
	dashRemover *regexp.Regexp
}

// NewIDTemplateCreator - IDTemplateDataCreator object constructor
func NewIDTemplateCreator() IDTemplateDataCreator {
	return IDTemplateDataCreator{dashRemover: regexp.MustCompile("-+")}
}

// RemoveDashes - removes all dashes from a string
func (i IDTemplateDataCreator) RemoveDashes(str string) string {
	return i.dashRemover.ReplaceAllString(str, "")
}

// CreateID - creates ID string by call name
func (i IDTemplateDataCreator) CreateID(prefix string, callName string) string {
	callName = i.RemoveDashes(callName)

	callNameRunes := []rune(callName)
	callNameRuneCount := len(callNameRunes)

	if callNameRuneCount == 0 {
		return ""
	}

	res := prefix + string(unicode.ToUpper(callNameRunes[0]))
	if callNameRuneCount > 1 {
		res += string(callNameRunes[1:])
	}

	return res
}

// CreateIDTemplateData - creates IDTemplateData slices for commands and flags
func (i IDTemplateDataCreator) CreateIDTemplateData(
	commandDescriptions []*confYML.CommandDescription,
	helpCommandDescription *confYML.HelpCommandDescription,
	nullCommandDescription *confYML.NamelessCommandDescription,
	flagDescriptionMap map[string]*confYML.FlagDescription,
) (
	commandsIDTemplateData map[string]*IDTemplateData,
	nullCommandIDTemplateData *IDTemplateData,
	flagsIDTemplateData map[string]*IDTemplateData) {

	var (
		j, k               int
		callName           string
		commandId          string
		commandDescription *confYML.CommandDescription
	)

	commandsIDTemplateData = make(map[string]*IDTemplateData, len(commandDescriptions))
	flagsIDTemplateData = make(map[string]*IDTemplateData, len(flagDescriptionMap))

	// standard commands
	for j = range commandDescriptions {
		commandDescription = commandDescriptions[j]

		callName = commandDescription.GetCommand()
		commandId = i.CreateID(PrefixCommandID, callName)

		commandsIDTemplateData[commandDescription.GetCommand()] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			commandDescription.GetDescriptionHelpInfo())

		for k = range commandDescription.GetAdditionalCommands() {
			callName = commandDescription.GetAdditionalCommands()[k]
			commandsIDTemplateData[commandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
				commandId,
				i.CreateID(PrefixCommandStringID, callName),
				callName,
				commandDescription.GetDescriptionHelpInfo())
		}
	}

	// help command
	if helpCommandDescription != nil {
		commandId = i.CreateID(PrefixCommandID, helpCommandIDStr)
		callName = helpCommandDescription.GetCommand()

		commandsIDTemplateData[helpCommandDescription.GetCommand()] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			helpCommandComment)

		for k = range helpCommandDescription.GetAdditionalCommands() {
			callName = helpCommandDescription.GetAdditionalCommands()[k]
			commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
				commandId,
				i.CreateID(PrefixCommandStringID, callName),
				callName,
				helpCommandComment)
		}
	}

	// null command
	if nullCommandDescription != nil {
		nullCommandIDTemplateData = NewIDTemplateData(
			i.CreateID(PrefixCommandID, NamelessCommandIDPostfix),
			"",
			"",
			nullCommandDescription.GetDescriptionHelpInfo())
	}

	// flags
	for _, flagDescription := range flagDescriptionMap {
		callName = flagDescription.GetFlag()
		flagsIDTemplateData[flagDescription.GetFlag()] = NewIDTemplateData(
			"",
			i.CreateID(PrefixFlagStringID, callName),
			callName,
			flagDescription.GetDescriptionHelpInfo())
	}

	return commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData
}

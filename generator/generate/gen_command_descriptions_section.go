package generate

import (
	"fmt"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	commandDescriptionSliceNilPart = `
		CommandDescriptions: nil`

	commandDescriptionSliceElementPrefixPart = `
		CommandDescriptions: []*apConf.CommandDescription{`
	commandDescriptionSliceElementRequiredPart = `
			apConf.CommandDescriptionSrc{
				ID:                  %s,
				DescriptionHelpInfo: "%s",`
	commandDescriptionSliceElementCommandsPart = `
				Commands: map[apConf.Command]bool{%s
				},`
	commandDescriptionSliceElementRequiredFlagsPart = `
				RequiredFlags: map[apConf.Flag]bool{%s
				},`
	commandDescriptionSliceElementOptionalFlagsPart = `
				OptionalFlags: map[apConf.Flag]bool{%s
				},`
	commandDescriptionSliceElementPostfix = `
			}.ToConstPtr(),`
	commandDescriptionSlicePostfix = `
		}`
)

// sectionCommandDescriptions - string with command constant definitions list
type sectionCommandDescriptions string

// genCommandDescriptionsSection creates a paste section with command descriptions
func genCommandDescriptionsSection(
	commandDescriptions []*confYML.CommandDescription,
	commandsIDTemplateData map[string]*creator.IDTemplateData,
	flagsIDTemplateData map[string]*creator.IDTemplateData,
) sectionCommandDescriptions {

	if len(commandDescriptions) == 0 {
		return commandDescriptionSliceNilPart
	}

	builder := strings.Builder{}
	builder.WriteString(commandDescriptionSliceElementPrefixPart)

	var (
		commandDescription *confYML.CommandDescription
		i, j               int
	)

	idTemplateDataSlice := make([]*creator.IDTemplateData, 0, 8)
	for i = range commandDescriptions {
		commandDescription = commandDescriptions[i]

		idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[commandDescription.GetCommand()])
		for j = range commandDescription.GetAdditionalCommands() {
			idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[commandDescription.GetAdditionalCommands()[j]])
		}

		builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementRequiredPart,
			commandsIDTemplateData[commandDescription.GetCommand()].GetID(),
			commandDescription.GetDescriptionHelpInfo()))
		builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementCommandsPart, joinCallNames(idTemplateDataSlice)))
		idTemplateDataSlice = []*creator.IDTemplateData{}

		if len(commandDescription.GetRequiredFlags()) > 0 {
			for j = range commandDescription.GetRequiredFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetRequiredFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementRequiredFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*creator.IDTemplateData{}
		}

		if len(commandDescription.GetOptionalFlags()) > 0 {
			for j = range commandDescription.GetOptionalFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetOptionalFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementOptionalFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*creator.IDTemplateData{}
		}

		builder.WriteString(commandDescriptionSliceElementPostfix)
	}

	builder.WriteString(commandDescriptionSlicePostfix)

	return sectionCommandDescriptions(builder.String())
}

func joinCallNames(nameAndIDSlice []*creator.IDTemplateData) (res string) {
	for i := range nameAndIDSlice {
		res += fmt.Sprintf("\n\t\t\t\t\t%s: true,", nameAndIDSlice[i].GetNameID())
	}
	return res
}

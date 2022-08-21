package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	commandDescriptionSliceNilPart = `
		// commandDescriptions
		nil`

	commandDescriptionSliceElementPrefixPart = `
		// commandDescriptions
		[]*dollyconf.CommandDescription{`
	commandDescriptionSliceElementRequiredPart = `
			{
				ID:                  %s,
				DescriptionHelpInfo: "%s",`
	commandDescriptionSliceElementCommandsPart = `
				Commands: map[dollyconf.Command]bool{%s
				},`
	commandDescriptionSliceElementRequiredFlagsPart = `
				RequiredFlags: map[dollyconf.Flag]bool{%s
				},`
	commandDescriptionSliceElementOptionalFlagsPart = `
				OptionalFlags: map[dollyconf.Flag]bool{%s
				},`
	commandDescriptionSliceElementPostfix = `
			},`
	commandDescriptionSlicePostfix = `
		}`
)

// CommandDescriptionsSection - string with command constant definitions list
type CommandDescriptionsSection string

// GenCommandDescriptionsSection creates a paste section with command descriptions
func GenCommandDescriptionsSection(
	commandDescriptions []*config_yaml.CommandDescription,
	commandsIDTemplateData map[string]*id_template_data_creator.IDTemplateData,
	flagsIDTemplateData map[string]*id_template_data_creator.IDTemplateData,
) CommandDescriptionsSection {

	if len(commandDescriptions) == 0 {
		return commandDescriptionSliceNilPart
	}

	builder := strings.Builder{}
	builder.WriteString(commandDescriptionSliceElementPrefixPart)

	var (
		commandDescription *config_yaml.CommandDescription
		i, j               int
	)

	idTemplateDataSlice := make([]*id_template_data_creator.IDTemplateData, 0, 8)
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
		idTemplateDataSlice = []*id_template_data_creator.IDTemplateData{}

		if len(commandDescription.GetRequiredFlags()) > 0 {
			for j = range commandDescription.GetRequiredFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetRequiredFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementRequiredFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*id_template_data_creator.IDTemplateData{}
		}

		if len(commandDescription.GetOptionalFlags()) > 0 {
			for j = range commandDescription.GetOptionalFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetOptionalFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementOptionalFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*id_template_data_creator.IDTemplateData{}
		}

		builder.WriteString(commandDescriptionSliceElementPostfix)
	}

	builder.WriteString(commandDescriptionSlicePostfix)

	return CommandDescriptionsSection(builder.String())
}

func joinCallNames(nameAndIDSlice []*id_template_data_creator.IDTemplateData) (res string) {
	for i := range nameAndIDSlice {
		res += fmt.Sprintf("\n\t\t\t\t\t%s: true,", nameAndIDSlice[i].GetNameID())
	}
	return res
}

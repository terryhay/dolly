package generate

import (
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
)

// Generate creates dolly.go file text body
func Generate(
	config *config_yaml.Config,
	flagDescriptionMap map[string]*config_yaml.FlagDescription,
) string {

	creator := id_template_data_creator.NewIDTemplateCreator()
	commandsIDTemplateData, namelessCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		config.GetCommandDescriptions(),
		config.GetHelpCommandDescription(),
		config.GetNamelessCommandDescription(),
		flagDescriptionMap)

	return GenArgParserFileBody(
		GenCommandIDListSection(commandsIDTemplateData, namelessCommandIDTemplateData),
		GenCommandListSection(commandsIDTemplateData),
		GenFlagIDConstantsSection(flagsIDTemplateData),
		GenAppDescriptionSection(config.GetAppHelpDescription()),
		GenFlagDescriptionsSection(config.GetFlagDescriptions(), flagsIDTemplateData),
		GenCommandDescriptionsSection(config.GetCommandDescriptions(), commandsIDTemplateData, flagsIDTemplateData),
		GenHelpCommandDescriptionSection(config.GetHelpCommandDescription(), commandsIDTemplateData),
		GenNamelessCommandComponent(config.GetNamelessCommandDescription(), namelessCommandIDTemplateData, flagsIDTemplateData),
		commandsIDTemplateData[config.GetHelpCommandDescription().GetCommand()].GetID())
}

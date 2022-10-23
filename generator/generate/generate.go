package generate

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/generator/id_template_data_creator"
)

// Generate creates dolly.go file text body
func Generate(
	config *confYML.ArgParserConfig,
	helpOutConfig *confYML.HelpOutConfig,
	flagDescriptionMap map[string]*confYML.FlagDescription,
) string {

	creator := id_template_data_creator.NewIDTemplateCreator()
	commandsIDTemplateData, namelessCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		config.GetCommandDescriptions(),
		config.GetHelpCommandDescription(),
		config.GetNamelessCommandDescription(),
		flagDescriptionMap)

	return GenArgParserFileBody(
		genImportSection(helpOutConfig.GetUsingTool()),
		genCommandIDListSection(commandsIDTemplateData, namelessCommandIDTemplateData),
		genCommandListSection(commandsIDTemplateData),
		genFlagIDConstantsSection(flagsIDTemplateData),
		genAppDescriptionSection(config.GetAppHelpDescription()),
		genFlagDescriptionsSection(config.GetFlagDescriptions(), flagsIDTemplateData),
		genCommandDescriptionsSection(config.GetCommandDescriptions(), commandsIDTemplateData, flagsIDTemplateData),
		genHelpCommandDescriptionSection(config.GetHelpCommandDescription(), commandsIDTemplateData),
		genNamelessCommandComponent(config.GetNamelessCommandDescription(), namelessCommandIDTemplateData, flagsIDTemplateData),
		commandsIDTemplateData[config.GetHelpCommandDescription().GetCommand()].GetID(),
		genHelpOutSection(helpOutConfig.GetUsingTool()),
	)
}

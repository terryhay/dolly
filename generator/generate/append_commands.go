package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendCommands creates a paste section with command descriptions
func appendCommands(builder *strings.Builder, marginLeft size.Width, configEntity ce.ConfigEntity) *strings.Builder {
	if builder == nil || len(configEntity.GetConfig().GetArgParserConfig().GetCommandsSorted()) == 0 {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "Commands: []*apConf.CommandOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	marginArray := marginLeft + size.WidthTab
	for _, command := range configEntity.GetConfig().GetArgParserConfig().GetCommandsSorted() {
		genDataCommand := configEntity.GenCompCommandByName(command.GetMainName())

		builder = sbt.NewRow(builder, marginArray, "{")
		{
			marginBody := marginArray + size.WidthTab
			builder = sbt.NewRow(builder, marginBody, "NameMain: ", genDataCommand.GetNameID(), ",")
			if len(command.GetAdditionalNames()) > 0 {
				builder = sbt.NewRow(builder, marginBody, "NamesAdditional: map[coty.NameCommand]struct{}{")
				{
					marginBodySecond := marginBody + size.WidthTab
					for _, name := range command.GetAdditionalNames() {
						genDataCommand = configEntity.GenCompCommandByName(name)
						builder = sbt.NewRow(builder, marginBodySecond, genDataCommand.GetNameID(), ",")
					}
				}
				builder = sbt.NewRow(builder, marginBody, "},")
			}
			builder = sbt.NewRow(builder, marginBody, `HelpInfo: "`, command.GetChapterDescriptionInfo().String(), `",`)
			builder = appendPlaceholders(builder, marginBody, configEntity, command.GetUsingPlaceholdersSorted())
		}
		builder = sbt.NewRow(builder, marginArray, "},")
	}

	return builder
}

package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendHelpOutCommand creates a paste section with help command description
func appendHelpOutCommand(builder *strings.Builder, marginLeft size.Width, configEntity ce.ConfigEntity) *strings.Builder {
	if builder == nil {
		return builder
	}

	var commandHelp *confYML.HelpCommand
	if commandHelp = configEntity.GetConfig().GetArgParserConfig().GetHelpCommand(); commandHelp == nil {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "CommandHelpOut: &apConf.HelpOutCommandOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	marginBody := marginLeft + size.WidthTab
	genDataCommandHelp := configEntity.GenCompCommandByName(commandHelp.GetMainName())
	builder = sbt.NewRow(builder, marginBody, "NameMain: ", genDataCommandHelp.GetNameID(), ",")

	if len(commandHelp.GetAdditionalNamesSorted()) > 0 {
		builder = sbt.NewRow(builder, marginBody, "NamesAdditional: map[coty.NameCommand]struct{}{")
		{
			marginBodySecond := marginBody + size.WidthTab
			builder = sbt.NewRow(builder, marginBodySecond, genDataCommandHelp.GetNameID(), ": {},")
			for _, name := range commandHelp.GetAdditionalNamesSorted() {
				genDataCommandHelp = configEntity.GenCompCommandByName(name)
				builder = sbt.NewRow(builder, marginBodySecond, genDataCommandHelp.GetNameID(), ": {},")
			}
		}
		builder = sbt.NewRow(builder, marginBody, "},")
	}

	return builder
}

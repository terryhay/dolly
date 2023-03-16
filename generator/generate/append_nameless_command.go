package generate

import (
	"strings"

	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendNamelessCommand creates a paste section with nameless command description
func appendNamelessCommand(builder *strings.Builder, marginLeft size.Width, configEntity ce.ConfigEntity) *strings.Builder {
	if builder == nil {
		return builder
	}

	var commandNameless *confYML.NamelessCommand
	if commandNameless = configEntity.GetConfig().GetArgParserConfig().GetNamelessCommand(); commandNameless == nil {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "CommandNameless: &apConf.NamelessCommandOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	builder = sbt.NewRow(builder, marginLeft+size.WidthTab, `HelpInfo: "`, commandNameless.GetChapterDescriptionInfo().String(), `",`)
	builder = appendPlaceholders(builder, marginLeft+size.WidthTab, configEntity, configEntity.GetConfig().GetArgParserConfig().GetNamelessCommand().GetUsingPlaceholders())

	return builder
}

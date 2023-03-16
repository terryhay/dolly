package generate

import (
	"strings"

	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// appendApplication creates a paste section with application description
func appendApplication(builder *strings.Builder, marginLeft size.Width, appHelp *confYML.AppHelp) *strings.Builder {
	if builder == nil || appHelp == nil {
		return builder
	}

	builder = sbt.NewRow(builder, marginLeft, "App: apConf.ApplicationOpt{")
	defer func() { builder = sbt.NewRow(builder, marginLeft, "},") }()

	builder = sbt.NewRow(builder, marginLeft+size.WidthTab, `AppName:         NameApp,`)
	builder = sbt.NewRow(builder, marginLeft+size.WidthTab, `InfoChapterNAME: "`, appHelp.GetHelpInfoChapterName().String(), `",`)

	return builder
}

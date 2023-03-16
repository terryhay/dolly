package generate

import (
	"strings"

	confYML "github.com/terryhay/dolly/generator/config_yaml"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

const (
	helpOutEntityPlainText = `
		helpOut.PrintHelpInfo(fmtd.NewFmtDecorator(), appArgConfig)
		return nil, nil`

	helpOutEntityManStyle = `
		pageView, err := pgv.NewPageView(tbd.NewTermBoxDecorator(), NameApp, hp.MakeBody(appArgConfig))
		if err != nil {
			return nil, err
		}

		if err = pageView.Run(); err != nil {
			return nil, err
		}

		return nil, nil`
)

// appendHelpOutEntity
func appendHelpOutEntity(builder *strings.Builder, helpOutTool confYML.HelpOutTool) *strings.Builder {
	if builder == nil || helpOutTool == confYML.HelpOutToolUndefined {
		return builder
	}

	switch {
	case helpOutTool == confYML.HelpOutToolPlainText:
		builder = sbt.Append(builder, helpOutEntityPlainText)

	case helpOutTool == confYML.HelpOutToolManStyle:
		builder = sbt.Append(builder, helpOutEntityManStyle)
	}

	return builder
}

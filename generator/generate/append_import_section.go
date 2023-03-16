package generate

import (
	"strings"

	confYML "github.com/terryhay/dolly/generator/config_yaml"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

const (
	importsManStyle = `
import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page"
	"github.com/terryhay/dolly/argparser/parsed"
	"github.com/terryhay/dolly/argparser/parser"
	pgv "github.com/terryhay/dolly/man_style_help/page_view"
	tbd "github.com/terryhay/dolly/man_style_help/termbox_decorator"
	coty "github.com/terryhay/dolly/tools/common_types"
)
`

	importsPlainText = `
import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
	"github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
	coty "github.com/terryhay/dolly/tools/common_types"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
)
`
)

func appendImports(builder *strings.Builder, helpOutTool confYML.HelpOutTool) *strings.Builder {
	if builder == nil || helpOutTool == confYML.HelpOutToolUndefined {
		return builder
	}

	switch {
	case helpOutTool == confYML.HelpOutToolPlainText:
		builder = sbt.Append(builder, importsPlainText)

	case helpOutTool == confYML.HelpOutToolManStyle:
		builder = sbt.Append(builder, importsManStyle)
	}

	return builder
}

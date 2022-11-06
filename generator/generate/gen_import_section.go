package generate

import confYML "github.com/terryhay/dolly/generator/config_yaml"

const (
	importSectionForManStyle = `import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	parsed "github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/argparser/parser"
	"github.com/terryhay/dolly/man_style_help/page"
	pgv "github.com/terryhay/dolly/man_style_help/page_view"
	tbd "github.com/terryhay/dolly/man_style_help/termbox_decorator"
)`
	importSectionForPlain = `import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	parsed "github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
)`
)

type importSection string

func genImportSection(helpOutTool confYML.HelpOutTool) importSection {
	if helpOutTool == confYML.HelpOutToolManStyle {
		return importSectionForManStyle
	}
	return importSectionForPlain
}

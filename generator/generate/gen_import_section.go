package generate

import confYML "github.com/terryhay/dolly/generator/config_yaml"

const (
	importSectionForManStyle = `import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/utils/dollyerr"
	"github.com/terryhay/dolly/argparser/helpdisplay/page"
	tbd "github.com/terryhay/dolly/argparser/helpdisplay/termbox_decorator"
	pgv "github.com/terryhay/dolly/argparser/helpdisplay/page_view"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/argparser/parser"
)`
	importSectionForPlain = `import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
	"github.com/terryhay/dolly/utils/dollyerr"
)`
)

type importSection string

func genImportSection(helpOutTool confYML.HelpOutTool) importSection {
	if helpOutTool == confYML.HelpOutToolManStyle {
		return importSectionForManStyle
	}
	return importSectionForPlain
}

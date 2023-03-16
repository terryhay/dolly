package plain_help_out

import (
	"errors"
	"strings"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
	"github.com/terryhay/dolly/tools/index"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

// ErrPrintHelpInfo - no chapters for print out
var ErrPrintHelpInfo = errors.New(`plain_help_out.PrintHelpInfo: no chapters for print out`)

// PrintHelpInfo prints help information by argParserConfig object
func PrintHelpInfo(decFmt fmtd.FmtDecorator, config apConf.ArgParserConfig) error {
	helpPage := hp.MakeBody(config)
	builder := &strings.Builder{}

	if helpPage.RowCount() == 0 {
		return ErrPrintHelpInfo
	}

	for i := index.Zero; i < helpPage.RowCount(); i++ {
		row := helpPage.Row(i)
		sbt.NewRow(builder, row.GetMarginLeft(), row.GetTextStyled())
	}

	if out := builder.String(); len(out) > 0 {
		decFmt.Println(out)
	}

	return nil
}

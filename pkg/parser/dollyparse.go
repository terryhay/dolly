package parser

import (
	"github.com/terryhay/dolly/internal/arg_parser_impl"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
)

func Parse(config dollyconf.ArgParserConfig, args []string) (*parsed_data.ParsedData, *dollyerr.Error) {
	return arg_parser_impl.NewCmdArgParserImpl(config).Parse(args)
}

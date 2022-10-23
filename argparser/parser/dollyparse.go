package parser

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/arg_parser_impl"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
)

func Parse(config apConf.ArgParserConfig, args []string) (*parsed_data.ParsedData, *dollyerr.Error) {
	return arg_parser_impl.NewCmdArgParserImpl(config).Parse(args)
}

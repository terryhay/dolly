package parser

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	impl "github.com/terryhay/dolly/argparser/arg_parser_impl"
	cmdArg "github.com/terryhay/dolly/argparser/cmd_arg"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
)

func Parse(config apConf.ArgParserConfig, args []string) (*parsed_data.ParsedData, *dollyerr.Error) {
	return impl.NewCmdArgParserImpl(config).Parse(cmdArg.MakeCmdArgIterator(args))
}

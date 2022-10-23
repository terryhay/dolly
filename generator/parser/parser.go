package parser

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/arg_parser_impl"
	"github.com/terryhay/dolly/argparser/parsed_data"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
	"github.com/terryhay/dolly/utils/dollyerr"
)

const (
	// CommandIDNullCommand - checks arguments types
	CommandIDNullCommand apConf.CommandID = iota + 1
	// CommandIDHelp - print help info
	CommandIDHelp
)

const (
	// CommandH - print help info
	CommandH apConf.Command = "-h"
	// CommandHelp - print help info
	CommandHelp = "help"
)

const (
	// FlagC - yaml file config path
	FlagC apConf.Flag = "-c"
	// FlagO - generate package path
	FlagO = "-o"
)

// Parse - processes command line arguments
func Parse(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
	appArgConfig := apConf.NewArgParserConfig(
		// appDescription
		apConf.ApplicationDescription{
			AppName:      "gen_dolly",
			NameHelpInfo: "code generator",
			DescriptionHelpInfo: []string{
				"generate parser package which contains a command line page parser",
			},
		},
		// flagDescriptions
		map[apConf.Flag]*apConf.FlagDescription{
			FlagC: {
				DescriptionHelpInfo: "yaml file config path",
				ArgDescription: &apConf.ArgumentsDescription{
					AmountType:              apConf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "file",
				},
			},
			FlagO: {
				DescriptionHelpInfo: "generate package path",
				ArgDescription: &apConf.ArgumentsDescription{
					AmountType:              apConf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "dir",
				},
			},
			CommandHelp: {},
		},
		// commandDescriptions
		nil,
		// helpCommandDescription
		apConf.NewHelpCommandDescription(
			CommandIDHelp,
			map[apConf.Command]bool{
				CommandHelp: true,
				CommandH:    true,
			},
		),
		// namelessCommandDescription
		apConf.NewNamelessCommandDescription(
			CommandIDNullCommand,
			"generate parser package",
			nil,
			map[apConf.Flag]bool{
				FlagC: true,
				FlagO: true,
			},
			nil,
		),
	)

	res, err = arg_parser_impl.NewCmdArgParserImpl(appArgConfig).Parse(args)
	if err != nil {
		return nil, err
	}

	if res.GetCommandID() == CommandIDHelp {
		helpOut.PrintHelpInfo(appArgConfig)
		return res, nil
	}

	return res, nil
}

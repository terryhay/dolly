package parser

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	impl "github.com/terryhay/dolly/argparser/arg_parser_impl"
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
	appArgConfig := apConf.ArgParserConfigSrc{
		AppDescription: apConf.ApplicationDescriptionSrc{
			AppName:      "gen_dolly",
			NameHelpInfo: "code generator",
			DescriptionHelpInfo: []string{
				"generate parser package which contains a command line page parser",
			},
		}.ToConst(),
		FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
			FlagC: apConf.FlagDescriptionSrc{
				DescriptionHelpInfo: "yaml file config path",
				ArgDescription: apConf.ArgumentsDescriptionSrc{
					AmountType:              apConf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "file",
				}.ToConstPtr(),
			}.ToConstPtr(),
			FlagO: apConf.FlagDescriptionSrc{
				DescriptionHelpInfo: "generate package path",
				ArgDescription: apConf.ArgumentsDescriptionSrc{
					AmountType:              apConf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "dir",
				}.ToConstPtr(),
			}.ToConstPtr(),
			CommandHelp: {},
		},
		HelpCommandDescription: apConf.NewHelpCommandDescription(
			CommandIDHelp,
			map[apConf.Command]bool{
				CommandHelp: true,
				CommandH:    true,
			},
		),
		NamelessCommandDescription: apConf.NewNamelessCommandDescription(
			CommandIDNullCommand,
			"generate parser package",
			nil,
			map[apConf.Flag]bool{
				FlagC: true,
				FlagO: true,
			},
			nil,
		),
	}.ToConst()

	res, err = impl.NewCmdArgParserImpl(appArgConfig).Parse(args)
	if err != nil {
		return nil, err
	}

	if res.GetCommandID() == CommandIDHelp {
		helpOut.PrintHelpInfo(appArgConfig)
		return res, nil
	}

	return res, nil
}

package parser

import (
	"github.com/terryhay/dolly/internal/arg_parser_impl"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpprinter"
	"github.com/terryhay/dolly/pkg/parsed_data"
)

const (
	// CommandIDNullCommand - checks arguments types
	CommandIDNullCommand dollyconf.CommandID = iota + 1
	// CommandIDHelp - print help info
	CommandIDHelp
)

const (
	// CommandH - print help info
	CommandH dollyconf.Command = "-h"
	// CommandHelp - print help info
	CommandHelp = "help"
)

const (
	// FlagC - yaml file config path
	FlagC dollyconf.Flag = "-c"
	// FlagO - generate package path
	FlagO = "-o"
)

// Parse - processes command line arguments
func Parse(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
	appArgConfig := dollyconf.NewArgParserConfig(
		// appDescription
		dollyconf.ApplicationDescription{
			AppName:      "gen_dolly",
			NameHelpInfo: "code generator",
			DescriptionHelpInfo: []string{
				"generate parser package which contains a command line data parser",
			},
		},
		// flagDescriptions
		map[dollyconf.Flag]*dollyconf.FlagDescription{
			FlagC: {
				DescriptionHelpInfo: "yaml file config path",
				ArgDescription: &dollyconf.ArgumentsDescription{
					AmountType:              dollyconf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "file",
				},
			},
			FlagO: {
				DescriptionHelpInfo: "generate package path",
				ArgDescription: &dollyconf.ArgumentsDescription{
					AmountType:              dollyconf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "dir",
				},
			},
			CommandHelp: {},
		},
		// commandDescriptions
		nil,
		// helpCommandDescription
		dollyconf.NewHelpCommandDescription(
			CommandIDHelp,
			map[dollyconf.Command]bool{
				CommandHelp: true,
				CommandH:    true,
			},
		),
		// namelessCommandDescription
		dollyconf.NewNamelessCommandDescription(
			CommandIDNullCommand,
			"generate parser package",
			nil,
			map[dollyconf.Flag]bool{
				FlagC: true,
				FlagO: true,
			},
			nil,
		),
	)

	if res, err = arg_parser_impl.NewCmdArgParserImpl(appArgConfig).Parse(args); err != nil {
		return nil, err
	}

	if res.GetCommandID() == CommandIDHelp {
		helpprinter.PrintHelpInfo(appArgConfig)
		return res, nil
	}

	return res, nil
}

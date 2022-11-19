// This code was generated by dolly.generator. DO NOT EDIT

package dolly

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	parsed "github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
)

const (
	// CommandIDNamelessCommand - checks arguments types
	CommandIDNamelessCommand apConf.CommandID = iota + 1
	// CommandIDPrintHelpInfo - print help info
	CommandIDPrintHelpInfo
)

const (
	// CommandHLw - print help info
	CommandHLw apConf.Command = "-h"
	// CommandHelp - print help info
	CommandHelp = "help"
)

const (
	// FlagCheck - check command arguments types
	FlagCheck apConf.Flag = "-check"
)

// Parse - processes command line arguments
func Parse(args []string) (*parsed.ParsedData, error) {
	appArgConfig := apConf.ArgParserConfigSrc{
		AppDescription: apConf.ApplicationDescriptionSrc{
			AppName:      "example2",
			NameHelpInfo: "shows how parser generator works without commands",
			DescriptionHelpInfo: []string{
				"you can write more detailed description here",
			},
		}.ToConst(),
		FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
			FlagCheck: apConf.FlagDescriptionSrc{
				DescriptionHelpInfo: "check command arguments types",
				ArgDescription: apConf.ArgumentsDescriptionSrc{
					AmountType:              apConf.ArgAmountTypeSingle,
					SynopsisHelpDescription: "str",
				}.ToConstPtr(),
			}.ToConstPtr(),
		},
		CommandDescriptions: nil,
		HelpCommandDescription: apConf.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[apConf.Command]bool{
				CommandHLw:  true,
				CommandHelp: true,
			},
		),
		NamelessCommandDescription: apConf.NewNamelessCommandDescription(
			CommandIDNamelessCommand,
			"checks arguments types",
			apConf.ArgumentsDescriptionSrc{
				AmountType:              apConf.ArgAmountTypeList,
				SynopsisHelpDescription: "str list",
			}.ToConstPtr(),
			map[apConf.Flag]bool{
				FlagCheck: true,
			},
			nil,
		)}.ToConst()

	res, err := parser.Parse(appArgConfig, args)
	if err != nil {
		return nil, err.Error()
	}

	if res.GetCommandID() == CommandIDPrintHelpInfo {
		helpOut.PrintHelpInfo(appArgConfig)
		return nil, nil
	}

	return res, nil
}

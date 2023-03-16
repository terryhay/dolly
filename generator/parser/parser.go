package parser

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
	impl "github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
	coty "github.com/terryhay/dolly/tools/common_types"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
)

const (
	// ArgGroupConfigPath - yaml file config path argument group
	ArgGroupConfigPath coty.IDPlaceholder = iota + 1

	// ArgGroupGeneratePackagePath - generate package path argument group
	ArgGroupGeneratePackagePath
)

const (
	// CommandH - print help info
	CommandH coty.NameCommand = "-h"

	// CommandHelp - print help info
	CommandHelp coty.NameCommand = "help"
)

const (
	// FlagC - yaml file config path
	FlagC coty.NameFlag = "-c"

	// FlagO - generate package path
	FlagO coty.NameFlag = "-o"
)

// Parse - processes command line arguments
func Parse(args []string) (res *parsed.Result, err error) {
	appArgConfig := apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{

		CommandNameless: &apConf.NamelessCommandOpt{
			Placeholders: []*apConf.PlaceholderOpt{
				{
					ID: ArgGroupConfigPath,
					FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
						FlagC: {
							NameMain: FlagC,
							HelpInfo: "yaml file config path",
						},
					},
					Argument: &apConf.ArgumentOpt{
						DescSynopsisHelp: "file",
					},
				},
				{
					ID: ArgGroupGeneratePackagePath,
					FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
						FlagO: {
							NameMain: FlagO,
							HelpInfo: "generate package path",
						},
					},
					Argument: &apConf.ArgumentOpt{
						DescSynopsisHelp: "dir",
					},
				},
			},
		},

		CommandHelpOut: &apConf.HelpOutCommandOpt{
			NameMain: CommandHelp,
			NamesAdditional: map[coty.NameCommand]struct{}{
				CommandH: {},
			},
		},

		App: apConf.ApplicationOpt{
			AppName:         "gen_dolly",
			InfoChapterNAME: "code generator",
		},

		HelpInfoChapterDESCRIPTION: []string{
			"generate parser package which contains a command line page parser",
		},
	})

	res, err = impl.Parse(appArgConfig, args)
	if err != nil {
		return nil, err
	}

	if res.GetCommandMainName() == CommandHelp {
		err = helpOut.PrintHelpInfo(fmtd.New(), appArgConfig)
		return res, err
	}

	return res, nil
}

package generate

import (
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	nameCommand := "cmd"
	nameCommandAdditional := "cmd-add"
	commandDescriptionHelpInfo := "nameCommand description help info"

	nameCommandHelp := "--help"
	nameCommandHelpAdditional := "-h"

	namePlaceholder1 := "placeholder1"
	namePlaceholder2 := "placeholder2"
	namePlaceholder3 := "placeholder3"
	namePlaceholder4 := "placeholder4"

	nameRequiredFlag := "-rf"
	nameRequiredFlag2 := "--req-flag"

	optionalFlag1 := "-f1"
	optionalFlag2 := "--opt-flag"

	configEntity, err := ce.MakeConfigEntity(
		confYML.NewConfig(&confYML.ConfigOpt{
			Version: "1.0.0",
			ArgParserConfig: &confYML.ArgParserConfigOpt{
				AppHelp: &confYML.AppHelpOpt{
					AppName:         "app",
					ChapterNameInfo: "chapter name info",
				},
				Placeholders: []*confYML.PlaceholderOpt{
					{
						Name: namePlaceholder1,
						Flags: []*confYML.FlagOpt{
							{
								MainName:               nameRequiredFlag,
								ChapterDescriptionInfo: "required flag",
							},
						},
						Argument: &confYML.ArgumentOpt{
							DefaultValues: []string{
								"cmdDefValue",
							},
							AllowedValues: []string{
								"cmdDefValue",
								"cmdAllValue",
							},
							HelpName: "arg",
						},
					},
					{
						Name: namePlaceholder2,
						Flags: []*confYML.FlagOpt{
							{
								MainName:               nameRequiredFlag2,
								ChapterDescriptionInfo: "required second flag",
							},
						},
						Argument: &confYML.ArgumentOpt{
							IsList: true,
							DefaultValues: []string{
								"f2DefValue",
							},
							AllowedValues: []string{
								"f2DefValue",
								"f2AllValue",
							},
							HelpName: "arg",
						},
					},
					{
						Name: namePlaceholder3,
						Flags: []*confYML.FlagOpt{
							{
								MainName:               optionalFlag1,
								ChapterDescriptionInfo: "first optional flag",
								IsOptional:             true,
							},
						},
					},
					{
						Name: namePlaceholder4,
						Flags: []*confYML.FlagOpt{
							{
								MainName:               optionalFlag2,
								ChapterDescriptionInfo: "second optional flag",
								IsOptional:             true,
							},
						},
					},
				},
				NamelessCommand: &confYML.NamelessCommandOpt{
					ChapterDescriptionInfo: "this is command without name",
					UsingPlaceholders: []string{
						namePlaceholder1,
						namePlaceholder2,
					},
				},
				Commands: []*confYML.CommandOpt{
					{
						MainName: nameCommand,
						AdditionalNames: []string{
							nameCommandAdditional,
						},
						ChapterDescriptionInfo: commandDescriptionHelpInfo,
						UsingPlaceholders: []string{
							namePlaceholder3,
							namePlaceholder4,
						},
					},
				},
				HelpCommand: &confYML.HelpCommandOpt{
					MainName: nameCommandHelp,
					AdditionalNames: []string{
						nameCommandHelpAdditional,
					},
				},
			},
		}),
	)
	require.NoError(t, err)

	argParserFileText := Generate(configEntity)

	require.Equal(t, `// This code was generated by dolly.generator. DO NOT EDIT

package dolly

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
	"github.com/terryhay/dolly/argparser/parser"
	helpOut "github.com/terryhay/dolly/argparser/plain_help_out"
	coty "github.com/terryhay/dolly/tools/common_types"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
)

const (
    // NameApp - name of the application
    NameApp coty.NameApp = "app"
)

const (
    // NameCommandCmd - nameCommand description help info
    NameCommandCmd coty.NameCommand = "cmd"

    // NameCommandCmdAdd - nameCommand description help info
    NameCommandCmdAdd coty.NameCommand = "cmd-add"

    // NameCommandHLw - print help info
    NameCommandHLw coty.NameCommand = "-h"

    // NameCommandHelp - print help info
    NameCommandHelp coty.NameCommand = "--help"

    // NameCommandNameless - this is command without name
    NameCommandNameless coty.NameCommand = "Nameless"
)

const (
    // IDPlaceholderPlaceholder1 - placeholder1
    IDPlaceholderPlaceholder1 coty.IDPlaceholder = iota + 1

    // IDPlaceholderPlaceholder2 - placeholder2
    IDPlaceholderPlaceholder2

    // IDPlaceholderPlaceholder3 - placeholder3
    IDPlaceholderPlaceholder3

    // IDPlaceholderPlaceholder4 - placeholder4
    IDPlaceholderPlaceholder4
)

const (
    // NameFlagF1 - first optional flag
    NameFlagF1 coty.NameFlag = "-f1"

    // NameFlagOptFlag - second optional flag
    NameFlagOptFlag coty.NameFlag = "--opt-flag"

    // NameFlagReqFlag - required second flag
    NameFlagReqFlag coty.NameFlag = "--req-flag"

    // NameFlagRf - required flag
    NameFlagRf coty.NameFlag = "-rf"
)

// Parse processes command line arguments
func Parse(args []string) (*parsed.Result, error) {
    appArgConfig := apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
        App: apConf.ApplicationOpt{
            AppName:         NameApp,
            InfoChapterNAME: "chapter name info",
        },
        CommandNameless: &apConf.NamelessCommandOpt{
            HelpInfo: "this is command without name",
            Placeholders: []*apConf.PlaceholderOpt{
                {
                    ID: IDPlaceholderPlaceholder1,
                    FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                        NameFlagRf: {
                            NameMain: NameFlagRf,
                            HelpInfo: "required flag",
                        },
                    },
                    Argument: &apConf.ArgumentOpt{
                        DefaultValues: []string{
                            "cmdDefValue",
                        },
                        AllowedValues: []string{
                            "cmdAllValue",
                            "cmdDefValue",
                        },
                    },
                },
                {
                    ID: IDPlaceholderPlaceholder2,
                    FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                        NameFlagReqFlag: {
                            NameMain: NameFlagReqFlag,
                            HelpInfo: "required second flag",
                        },
                    },
                    Argument: &apConf.ArgumentOpt{
                        IsList: coty.ArgAmountTypeList,
                        DefaultValues: []string{
                            "f2DefValue",
                        },
                        AllowedValues: []string{
                            "f2AllValue",
                            "f2DefValue",
                        },
                    },
                },
            },
        },
        Commands: []*apConf.CommandOpt{
            {
                NameMain: NameCommandCmd,
                NamesAdditional: map[coty.NameCommand]struct{}{
                    NameCommandCmdAdd,
                },
                HelpInfo: "nameCommand description help info",
                Placeholders: []*apConf.PlaceholderOpt{
                    {
                        ID: IDPlaceholderPlaceholder3,
                        FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                            NameFlagF1: {
                                NameMain: NameFlagF1,
                                HelpInfo: "first optional flag",
                                IsOptional:          true
                            },
                        },
                    },
                    {
                        ID: IDPlaceholderPlaceholder4,
                        FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                            NameFlagOptFlag: {
                                NameMain: NameFlagOptFlag,
                                HelpInfo: "second optional flag",
                                IsOptional:          true
                            },
                        },
                    },
                },
            },
        },
        CommandHelpOut: &apConf.HelpOutCommandOpt{
            NameMain: NameCommandHelp,
            NamesAdditional: map[coty.NameCommand]struct{}{
                NameCommandHelp: {},
                NameCommandHLw: {},
            },
        },
    })

    res, errParse := parser.Parse(appArgConfig, args)
    if errParse != nil {
        return nil, errParse
    }

    if res.GetCommandMainName() == NameCommandHelp {
		helpOut.PrintHelpInfo(fmtd.NewFmtDecorator(), appArgConfig)
		return nil, nil
    }

    return res, nil
}`, argParserFileText)
}

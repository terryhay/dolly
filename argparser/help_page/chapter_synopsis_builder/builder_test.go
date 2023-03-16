package chapter_synopsis_builder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func TestAppendRows(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		appName  coty.NameApp
		commands []*apConf.Command

		expStyled   string
		expUnstyled string
	}{
		{
			caseName: "no_data",
		},
		{
			caseName: "only_app_name",
			appName:  coty.RandNameApp(),
		},
		{
			caseName: "only_nameless_command",
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "",
					HelpInfo: "nameless command description",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-f": {
									NameMain: "-f",
									HelpInfo: "flag without arguments",
								},
							},
						},
					},
				}),
			},
		},
		{
			caseName: "command_without_placeholder",
			appName:  "app",
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "cmd",
					HelpInfo: "command description",
				}),
			},
			expStyled: `
[1mSYNOPSIS[0m
    [1mapp cmd [0m`,
			expUnstyled: `
SYNOPSIS
    app cmd `,
		},
		{
			caseName: "common",
			appName:  "app",
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "",
					HelpInfo: "nameless command description",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							Argument: &apConf.ArgumentOpt{
								IsOptional:       true,
								DescSynopsisHelp: "name",
							},
						},
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-a": {
									NameMain: "-a",
									HelpInfo: "a flag description",
								},
								"-b": {
									NameMain: "-b",
									HelpInfo: "b flag description",
								},
								"-B": {
									NameMain: "-B",
									HelpInfo: "B flag description",
								},
								"-c": {
									NameMain: "-c",
									HelpInfo: "c flag description",
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsList:           true,
								IsOptional:       true,
								DefaultValues:    []string{"0", "1"},
								DescSynopsisHelp: "str",
							},
						},
					},
				}),
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "cmd", // short command
					HelpInfo: "short command description",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"--long-flag": {
									NameMain: "--long-flag",
									NamesAdditional: map[coty.NameFlag]struct{}{
										"-fl": {},
									},
									HelpInfo: "long flag with optional argument",
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsOptional:       true,
								IsList:           true,
								DefaultValues:    []string{"0", "1"},
								DescSynopsisHelp: "num",
							},
						},
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"--long-second": {
									NameMain: "--long-second",
									HelpInfo: "long flag with optional argument",
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsList:           true,
								DefaultValues:    []string{"0", "1"},
								DescSynopsisHelp: "num",
							},
						},
					},
				}),
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "long_command",
					HelpInfo: "long command description",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-f": {
									NameMain: "-f",
									HelpInfo: "short flag description",
								},
								"-F": {
									NameMain: "-F",
									HelpInfo: "another one flag description",
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsOptional:       true,
								DefaultValues:    []string{"bubu"},
								DescSynopsisHelp: "str",
							},
						},
						{
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-na": {
									NameMain: "-na",
									HelpInfo: "flag without arguments",
								},
							},
						},
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-t": {
									NameMain: "-t",
									HelpInfo: "another one fucking flag with argument",
								},
							},
							Argument: &apConf.ArgumentOpt{
								DescSynopsisHelp: "num",
							},
						},
					},
				}),
			},
			expStyled: `
[1mSYNOPSIS[0m
    [1mapp [0m[[4mname[0m] [1m-aBbc[0m[[4mstr[0m [4m...[0m]
    [1mapp cmd [0m[1m--long-flag[0m[=[4mnum[0m [4m...[0m] [1m--long-second[0m=[4mnum[0m [4m...[0m
    [1mapp long_command [0m[[1m-F[0m | [1m-f[0m[[4mstr[0m]] [[1m-na[0m] [1m-t[0m [4mnum[0m`,
			expUnstyled: `
SYNOPSIS
    app [name] -aBbc[str ...]
    app cmd --long-flag[=num ...] --long-second=num ...
    app long_command [-F | -f[str]] [-na] -t num`,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			rows := AppendRows(make([]hp.Row, 0), tc.appName, tc.commands)

			builder := &strings.Builder{}
			for i := range rows {
				if i > 0 {
					builder = sbt.BreakRow(builder)
				}

				tab := hp.MakeRowChunkSpaces(rows[i].GetMarginLeft())
				builder = sbt.Append(builder, tab.GetText(), rows[i].GetTextStyled())
			}

			str := builder.String()
			require.Equal(t, tc.expStyled, str)
			require.Equal(t, tc.expUnstyled, hp.RemoveStyleTextMarkers(str))
		})
	}
}

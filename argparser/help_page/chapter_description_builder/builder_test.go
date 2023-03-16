package chapter_description_builder

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

		textsIntroduction []coty.InfoChapterDESCRIPTION
		commands          []*apConf.Command

		expStyled   string
		expUnstyled string
	}{
		{
			caseName: "no_data",
		},
		{
			caseName: "only_nameless_command",
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "",
					HelpInfo: "nameless command description",
				}),
			},
		},
		{
			caseName: "nameless_command_with_flag",
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
			expStyled: `
[1mDESCRIPTION[0m
The flags are as follows:
    [1m-f[0m      flag without arguments
`,
			expUnstyled: `
DESCRIPTION
The flags are as follows:
    -f      flag without arguments
`,
		},
		{
			caseName: "command_without_placeholder",
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "cmd",
					HelpInfo: "command description",
				}),
			},
			expStyled: `
[1mDESCRIPTION[0m
The commands are as follows:
    [1mcmd[0m     command description
`,
			expUnstyled: `
DESCRIPTION
The commands are as follows:
    cmd     command description
`,
		},
		{
			caseName: "common",
			textsIntroduction: []coty.InfoChapterDESCRIPTION{
				"First introduction line.",
				"Second introduction line.",
			},
			commands: []*apConf.Command{
				apConf.NewCommand(apConf.CommandOpt{
					NameMain: "",
					HelpInfo: "nameless command description",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"--opt-list": {
									NameMain: "--opt-list",
									NamesAdditional: map[coty.NameFlag]struct{}{
										"-ol": {},
										"-O":  {},
									},
									HelpInfo: "long flag description",
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
									HelpInfo: "long flag description",
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
					},
				}),
			},
			expStyled: `
[1mDESCRIPTION[0m
    First introduction line.

    Second introduction line.

The commands are as follows:
    [1m<empty>[0m nameless command description

    [1mcmd[0m     short command description

    [1mlong_command[0m
            long command description

The flags are as follows:
    [1m--long-flag[0m=[4mnum[0m ..., [1m-fl[0m [4mnum[0m ...
            long flag description

    [1m--opt-list[0m[=[4mstr[0m] ..., [1m-ol[0m[[4mstr[0m] ..., [1m-O[0m[[4mstr[0m] ...
            long flag description

    [1m-f[0m[[4mstr[0m] short flag description

    [1m-na[0m     flag without arguments
`,
			expUnstyled: `
DESCRIPTION
    First introduction line.

    Second introduction line.

The commands are as follows:
    <empty> nameless command description

    cmd     short command description

    long_command
            long command description

The flags are as follows:
    --long-flag=num ..., -fl num ...
            long flag description

    --opt-list[=str] ..., -ol[str] ..., -O[str] ...
            long flag description

    -f[str] short flag description

    -na     flag without arguments
`,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			rows := AppendRows(make([]hp.Row, 0), tc.textsIntroduction, tc.commands)

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

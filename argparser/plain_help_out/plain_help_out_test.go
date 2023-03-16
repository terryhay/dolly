package plain_help_out

import (
	"testing"

	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	fmtd "github.com/terryhay/dolly/tools/fmt_decorator"
)

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	t.Run("empty_config", func(t *testing.T) {
		t.Parallel()

		fmtCatcher := fmtd.NewCatcher()
		err := PrintHelpInfo(fmtCatcher, apConf.ArgParserConfig{})

		require.ErrorIs(t, err, ErrPrintHelpInfo)
		require.Empty(t, fmtCatcher.GetPrintln())
	})

	t.Run("common", func(t *testing.T) {
		t.Parallel()

		conf := apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
			App: apConf.ApplicationOpt{
				AppName:         "appname",
				InfoChapterNAME: "name help info",
			},
			CommandNameless: &apConf.NamelessCommandOpt{
				HelpInfo: "nameless command description",
			},
			Commands: []*apConf.CommandOpt{
				{
					NameMain: "cmd",
					HelpInfo: "command description help info",
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.IDPlaceholder(1),
							Argument: &apConf.ArgumentOpt{
								DescSynopsisHelp: "str",
							},
						},
						{
							ID: coty.IDPlaceholder(2),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-rf": {
									NameMain: "-rf",
									HelpInfo: "flag -rf description",
								},
							},
							Argument: &apConf.ArgumentOpt{
								DescSynopsisHelp: "str",
							},
						},
						{
							ID:             coty.IDPlaceholder(3),
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								"-of": {
									NameMain: "-of",
									HelpInfo: "flag -of description",
								},
							},
						},
					},
				},
				{
					NameMain: "longcommand",
					HelpInfo: "longcommand description help info",
				},
			},
		})

		fmtCatcher := fmtd.NewCatcher()
		require.NoError(t, PrintHelpInfo(fmtCatcher, conf))

		out := fmtCatcher.GetPrintln()

		require.Equal(t, `
[1mNAME[0m
    [1mappname[0m – name help info

[1mSYNOPSIS[0m
    [1mappname [0m
    [1mappname cmd [0m [4mstr[0m [1m-rf[0m [4mstr[0m [[1m-of[0m]
    [1mappname longcommand [0m

[1mDESCRIPTION[0m
The commands are as follows:
    [1m<empty>[0m nameless command description

    [1mcmd[0m     command description help info

    [1mlongcommand[0m
            longcommand description help info

The flags are as follows:
    [1m-of[0m     flag -of description

    [1m-rf[0m [4mstr[0m flag -rf description
`, out)

		require.Equal(t, `
NAME
    appname – name help info

SYNOPSIS
    appname 
    appname cmd  str -rf str [-of]
    appname longcommand 

DESCRIPTION
The commands are as follows:
    <empty> nameless command description

    cmd     command description help info

    longcommand
            longcommand description help info

The flags are as follows:
    -of     flag -of description

    -rf str flag -rf description
`, hp.RemoveStyleTextMarkers(out))
	})
}

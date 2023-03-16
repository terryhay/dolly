package help_page

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	hp "github.com/terryhay/dolly/argparser/help_page/page"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/index"
	sbt "github.com/terryhay/dolly/tools/string_builder_tools"
)

func TestBuild(t *testing.T) {
	t.Parallel()

	body := MakeBody(apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
		App: apConf.ApplicationOpt{
			AppName:         "app",
			InfoChapterNAME: "application description",
		},
		CommandNameless: &apConf.NamelessCommandOpt{
			HelpInfo: "nameless command description",
			Placeholders: []*apConf.PlaceholderOpt{
				{
					ID: coty.IDPlaceholder(1),
					FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
						"-f": {
							NameMain: "-f",
							HelpInfo: "nameless command flag description",
						},
					},
				},
			},
		},
		Commands: []*apConf.CommandOpt{
			{
				NameMain: "do",
				HelpInfo: "execute something fucking doing",
			},
		},
		CommandHelpOut: &apConf.HelpOutCommandOpt{
			NameMain: "-h",
		},
	}))

	builder := &strings.Builder{}
	for i := index.Zero; i < body.RowCount(); i++ {
		if i > 0 {
			builder = sbt.BreakRow(builder)
		}

		row := body.Row(i)
		tab := hp.MakeRowChunkSpaces(row.GetMarginLeft())
		builder = sbt.Append(builder, tab.GetText(), row.GetTextStyled())
	}

	str := builder.String()
	require.Equal(t, `[1mNAME[0m
    [1mapp[0m – application description

[1mSYNOPSIS[0m
    [1mapp [0m[1m-f[0m
    [1mapp do [0m

[1mDESCRIPTION[0m
The commands are as follows:
    [1m<empty>[0m nameless command description

    [1mdo[0m      execute something fucking doing

The flags are as follows:
    [1m-f[0m      nameless command flag description
`, str)

	require.Equal(t, `NAME
    app – application description

SYNOPSIS
    app -f
    app do 

DESCRIPTION
The commands are as follows:
    <empty> nameless command description

    do      execute something fucking doing

The flags are as follows:
    -f      nameless command flag description
`, hp.RemoveStyleTextMarkers(str))
}

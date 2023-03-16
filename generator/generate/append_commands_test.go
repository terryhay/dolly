package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendCommands(t *testing.T) {
	t.Parallel()

	configEntity, err := ce.MakeConfigEntity(confYML.NewConfig(&confYML.ConfigOpt{
		Version: "1.0.0",
		ArgParserConfig: &confYML.ArgParserConfigOpt{
			AppHelp: &confYML.AppHelpOpt{
				AppName:                "app name",
				ChapterNameInfo:        "app info",
				ChapterDescriptionInfo: []string{"app description"},
			},
			Placeholders: []*confYML.PlaceholderOpt{
				{
					Name: "placeholder",
					Flags: []*confYML.FlagOpt{
						{
							MainName:               "-f",
							ChapterDescriptionInfo: "some flag",
						},
					},
				},
			},
			Commands: []*confYML.CommandOpt{
				{
					MainName: "dosomething",
					AdditionalNames: []string{
						"command",
						"apply",
					},
					ChapterDescriptionInfo: "command without placeholders",
				},
				{
					MainName:               "doplacehold",
					ChapterDescriptionInfo: "command without placeholders",
					UsingPlaceholders: []string{
						"placeholder",
					},
				},
			},
			HelpCommand: &confYML.HelpCommandOpt{
				MainName: "--help",
			},
		},
	}))
	require.NoError(t, err)

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendCommands(nil, size.WidthTab, configEntity))
		require.Equal(t, 0, len(appendCommands(&strings.Builder{}, size.WidthTab, ce.ConfigEntity{}).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        Commands: []*apConf.CommandOpt{
            {
                NameMain: NameCommandDoplacehold,
                HelpInfo: "command without placeholders",
                Placeholders: []*apConf.PlaceholderOpt{
                    {
                        ID: IDPlaceholderPlaceholder,
                        FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                            NameFlagFLw: {
                                NameMain: NameFlagFLw,
                                HelpInfo: "some flag",
                            },
                        },
                    },
                },
            },
            {
                NameMain: NameCommandDosomething,
                NamesAdditional: map[coty.NameCommand]struct{}{
                    NameCommandApply,
                    NameCommandCommand,
                },
                HelpInfo: "command without placeholders",
            },
        },`, appendCommands(&strings.Builder{}, size.WidthTab+size.WidthTab, configEntity).String())
	})
}

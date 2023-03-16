package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendPlaceholders(t *testing.T) {
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
					Name: "placeholder1",
					Flags: []*confYML.FlagOpt{
						{
							MainName: "-f",
							AdditionalNames: []string{
								"--flag",
							},
							ChapterDescriptionInfo: "some flag",
							IsOptional:             true,
						},
					},
				},
				{
					Name: "placeholder2",
					Argument: &confYML.ArgumentOpt{
						HelpName:   "arg",
						IsList:     true,
						IsOptional: true,
						AllowedValues: []string{
							"v1",
							"v2",
						},
						DefaultValues: []string{
							"v1",
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
					UsingPlaceholders: []string{
						"placeholder1",
						"placeholder2",
					},
				},
			},
			HelpCommand: &confYML.HelpCommandOpt{
				MainName: "--help",
			},
		},
	}))
	require.NoError(t, err)

	namePlaceholders := []coty.NamePlaceholder{
		"placeholder1",
		"placeholder2",
	}

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendPlaceholders(nil, size.WidthTab, configEntity, namePlaceholders))
		require.Equal(t, 0, len(appendPlaceholders(&strings.Builder{}, size.WidthTab, configEntity, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        Placeholders: []*apConf.PlaceholderOpt{
            {
                ID: IDPlaceholderPlaceholder1,
                FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
                    NameFlagFLw: {
                        NameMain: NameFlagFLw,
                        NamesAdditional:      map[coty.NameFlag]struct{}{
                            NameFlagFlag,
                        },
                        HelpInfo: "some flag",
                        IsOptional:          true
                    },
                },
            },
            {
                ID: IDPlaceholderPlaceholder2,
                Argument: &apConf.ArgumentOpt{
                    IsList: coty.ArgAmountTypeList,
                    DefaultValues: []string{
                        "v1",
                    },
                    AllowedValues: []string{
                        "v1",
                        "v2",
                    },
                    IsOptional:          true,
                },
            },
        },`, appendPlaceholders(&strings.Builder{}, size.WidthTab+size.WidthTab, configEntity, namePlaceholders).String())
	})
}

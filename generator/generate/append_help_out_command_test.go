package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendHelpOutCommand(t *testing.T) {
	t.Parallel()

	configEntity, err := ce.MakeConfigEntity(confYML.NewConfig(&confYML.ConfigOpt{
		Version: "1.0.0",
		ArgParserConfig: &confYML.ArgParserConfigOpt{
			AppHelp: &confYML.AppHelpOpt{
				AppName:                "app name",
				ChapterNameInfo:        "app info",
				ChapterDescriptionInfo: []string{"app description"},
			},
			NamelessCommand: &confYML.NamelessCommandOpt{
				ChapterDescriptionInfo: "do something shit",
			},
			HelpCommand: &confYML.HelpCommandOpt{
				MainName: "--help",
				AdditionalNames: []string{
					"-h",
				},
			},
		},
	}))
	require.NoError(t, err)

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendHelpOutCommand(nil, size.WidthTab, configEntity))
		require.Equal(t, 0, len(appendHelpOutCommand(&strings.Builder{}, size.WidthTab, ce.ConfigEntity{}).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        CommandHelpOut: &apConf.HelpOutCommandOpt{
            NameMain: NameCommandHelp,
            NamesAdditional: map[coty.NameCommand]struct{}{
                NameCommandHelp: {},
                NameCommandHLw: {},
            },
        },`, appendHelpOutCommand(&strings.Builder{}, size.WidthTab+size.WidthTab, configEntity).String())
	})
}

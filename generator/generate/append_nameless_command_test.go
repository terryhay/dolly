package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	ce "github.com/terryhay/dolly/generator/config_entity"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendNamelessCommand(t *testing.T) {
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
			},
		},
	}))
	require.NoError(t, err)

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendNamelessCommand(nil, size.WidthTab, configEntity))
		require.Equal(t, 0, len(appendNamelessCommand(&strings.Builder{}, size.WidthTab, ce.ConfigEntity{}).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        CommandNameless: &apConf.NamelessCommandOpt{
            HelpInfo: "do something shit",
        },`, appendNamelessCommand(&strings.Builder{}, size.WidthTab+size.WidthTab, configEntity).String())
	})
}

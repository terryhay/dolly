package generate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/tools/size"
)

func TestAppendApplication(t *testing.T) {
	t.Parallel()

	appHelp := confYML.NewAppHelp(&confYML.AppHelpOpt{
		AppName:         "appname",
		ChapterNameInfo: "chapter name info",
	})

	t.Run("nil", func(t *testing.T) {
		require.Nil(t, appendApplication(nil, size.WidthTab, appHelp))
		require.Equal(t, 0, len(appendApplication(&strings.Builder{}, size.WidthTab, nil).String()))
	})

	t.Run("common", func(t *testing.T) {
		require.Equal(t, `
        App: apConf.ApplicationOpt{
            AppName:         NameApp,
            InfoChapterNAME: "chapter name info",
        },`, appendApplication(&strings.Builder{}, size.WidthTab+size.WidthTab, appHelp).String())
	})
}

package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	obj := ArgParserConfigSrc{
		AppDescription: ApplicationDescription{},
		FlagDescriptions: map[Flag]*FlagDescription{
			Flag(gofakeit.Name()): {},
		},
		CommandDescriptions: []*CommandDescription{
			{},
		},
	}.Cast()

	require.Equal(t, obj.appDescription, obj.GetAppDescription())
	require.Equal(t, obj.flagDescriptions, obj.GetFlagDescriptions())
	require.Equal(t, obj.commandDescriptions, obj.GetCommandDescriptions())
	require.Equal(t, obj.helpCommandDescription, obj.GetHelpCommandDescription())
	require.Equal(t, obj.namelessCommandDescription, obj.GetNamelessCommandDescription())
}

package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *ArgParserConfig

		require.Equal(t, pointer.GetAppDescription(), ApplicationDescription{})
		require.Nil(t, pointer.GetCommandDescriptions())
		require.Nil(t, pointer.GetFlagDescriptionSlice())
		require.Nil(t, pointer.ExtractFlagDescriptionMap())
		require.Nil(t, pointer.GetHelpCommandDescription())
		require.Nil(t, pointer.GetNamelessCommandDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		flag := Flag(gofakeit.Name())
		flagDesc := FlagDescriptionSrc{
			Flags: []Flag{
				flag,
			},
		}.ToConstPtr()

		src := ArgParserConfigSrc{
			AppDescription: ApplicationDescription{},
			FlagDescriptionSlice: []*FlagDescription{
				flagDesc,
			},
			CommandDescriptions: []*CommandDescription{
				{},
			},
		}
		pointer := src.ToConst()

		require.Equal(t, src.AppDescription, pointer.GetAppDescription())
		require.Equal(t, src.FlagDescriptionSlice, pointer.GetFlagDescriptionSlice())
		require.Equal(t, src.CommandDescriptions, pointer.GetCommandDescriptions())
		require.Equal(t, map[Flag]*FlagDescription{flag: flagDesc}, pointer.ExtractFlagDescriptionMap())
		require.Equal(t, src.HelpCommandDescription, pointer.GetHelpCommandDescription())
		require.Equal(t, src.NamelessCommandDescription, pointer.GetNamelessCommandDescription())
	})
}

package dollyconf

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	obj := NewArgParserConfig(
		ApplicationDescription{},
		map[Flag]*FlagDescription{
			Flag(gofakeit.Name()): {},
		},
		[]*CommandDescription{
			{},
		},
		nil,
		nil,
	)

	require.Equal(t, obj.AppDescription, obj.GetAppDescription())
	require.Equal(t, obj.FlagDescriptions, obj.GetFlagDescriptions())
	require.Equal(t, obj.CommandDescriptions, obj.GetCommandDescriptions())
	require.Equal(t, obj.HelpCommandDescription, obj.GetHelpCommandDescription())
	require.Equal(t, obj.NamelessCommandDescription, obj.GetNamelessCommandDescription())
}

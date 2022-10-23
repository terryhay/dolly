package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	helpCommandDescription := NewHelpCommandDescription(
		CommandID(gofakeit.Uint32()),
		map[Command]bool{
			Command(gofakeit.Color()): true,
		},
	)

	commandDescription := helpCommandDescription.(*CommandDescription)

	require.Equal(t, commandDescription.ID, helpCommandDescription.GetID())
	require.Equal(t, commandDescription.Commands, helpCommandDescription.GetCommands())
}

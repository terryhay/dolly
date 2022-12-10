package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	id := CommandID(gofakeit.Uint32())
	commands := map[Command]bool{
		Command(gofakeit.Color()): true,
	}

	helpCommandDescription := NewHelpCommandDescription(id, commands)

	require.Equal(t, id, helpCommandDescription.GetID())
	require.Equal(t, commands, helpCommandDescription.GetCommands())
}

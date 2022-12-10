package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *CommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, CommandIDUndefined, pointer.GetID())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetCommands())
		require.Nil(t, pointer.GetArgDescription())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := CommandDescriptionSrc{
			ID:                  CommandID(gofakeit.Uint32()),
			DescriptionHelpInfo: gofakeit.Name(),
			Commands: map[Command]bool{
				Command(gofakeit.Name()): true,
			},
			ArgDescription: &ArgumentsDescription{},
			RequiredFlags: map[Flag]bool{
				Flag(gofakeit.Name()): true,
			},
			OptionalFlags: map[Flag]bool{
				Flag(gofakeit.Name()): true,
			},
		}

		pointer = src.CastPtr()

		require.Equal(t, src.ID, pointer.GetID())
		require.Equal(t, src.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, src.Commands, pointer.GetCommands())
		require.Equal(t, src.ArgDescription, pointer.GetArgDescription())
		require.Equal(t, src.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, src.OptionalFlags, pointer.GetOptionalFlags())
	})
}

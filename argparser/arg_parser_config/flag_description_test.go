package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *FlagDescription

		require.Equal(t, FlagIDUndefined, pointer.GetID())
		require.Nil(t, pointer.GetFlags())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetArgDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := FlagDescriptionSrc{
			ID:                  FlagID(gofakeit.Uint32()),
			Flags:               []Flag{Flag(gofakeit.Name())},
			DescriptionHelpInfo: gofakeit.Name(),
			ArgDescription:      &ArgumentsDescription{},
		}
		pointer := src.ToConstPtr()

		require.Equal(t, src.ID, pointer.GetID())
		require.Equal(t, src.Flags, pointer.GetFlags())
		require.Equal(t, src.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, src.ArgDescription, pointer.GetArgDescription())
	})
}

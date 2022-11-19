package arg_parser_config

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *FlagDescription

	t.Run("nil_pointer", func(t *testing.T) {

		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetArgDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := FlagDescriptionSrc{
			DescriptionHelpInfo: gofakeit.Name(),
			ArgDescription:      &ArgumentsDescription{},
		}
		pointer = src.ToConstPtr()

		require.Equal(t, src.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, src.ArgDescription, pointer.GetArgDescription())
	})
}

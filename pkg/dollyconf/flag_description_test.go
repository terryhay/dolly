package dollyconf

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
		pointer = &FlagDescription{
			DescriptionHelpInfo: gofakeit.Name(),
			ArgDescription:      &ArgumentsDescription{},
		}

		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.ArgDescription, pointer.GetArgDescription())
	})
}

package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestParsedFlagDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedFlagData

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, dollyconf.Flag(""), pointer.GetFlag())
		require.Nil(t, pointer.GetArgData())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = NewParsedFlagData(
			dollyconf.Flag(gofakeit.Name()),
			NewParsedArgData(nil),
		)

		require.Equal(t, pointer.Flag, pointer.GetFlag())
		require.Equal(t, pointer.ArgData, pointer.GetArgData())
	})
}

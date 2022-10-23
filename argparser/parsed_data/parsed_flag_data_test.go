package parsed_data

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"testing"
)

func TestParsedFlagDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedFlagData

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, apConf.Flag(""), pointer.GetFlag())
		require.Nil(t, pointer.GetArgData())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = NewParsedFlagData(
			apConf.Flag(gofakeit.Name()),
			NewParsedArgData(nil),
		)

		require.Equal(t, pointer.Flag, pointer.GetFlag())
		require.Equal(t, pointer.ArgData, pointer.GetArgData())
	})
}

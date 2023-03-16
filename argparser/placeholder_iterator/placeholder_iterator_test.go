package placeholder_iterator

import (
	"testing"

	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestPlaceholderIterator(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Iterator

		require.True(t, pointer.IsEnded())
		require.Nil(t, pointer.Get())
		require.Nil(t, pointer.Next())
	})

	t.Run("initialized", func(t *testing.T) {
		t.Parallel()

		placeholders := []*apConf.Placeholder{
			apConf.NewPlaceholder(apConf.PlaceholderOpt{
				ID: coty.RandIDPlaceholder(),
			}),
			apConf.NewPlaceholder(apConf.PlaceholderOpt{
				ID: coty.RandIDPlaceholderSecond(),
			}),
		}

		it := Make(placeholders)
		i := -1
		for it.Next(); !it.IsEnded(); it.Next() {
			i++

			require.NotNil(t, it.Get())

			require.Equal(t, placeholders[i], it.Get())
		}
	})
}

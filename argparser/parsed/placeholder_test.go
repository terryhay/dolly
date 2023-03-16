package parsed

import (
	"testing"

	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestParsedFlagDataGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var opt *PlaceholderOpt

		var pointer = NewPlaceholder(opt)
		require.Nil(t, pointer)

		require.Equal(t, coty.ArgPlaceholderIDUndefined, pointer.GetID())
		require.Equal(t, coty.NameFlagUndefined, pointer.GetNameFlag())
		require.Nil(t, pointer.GetArgData())

		require.False(t, pointer.HasArg())
		require.False(t, pointer.HasFlag())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := &PlaceholderOpt{
			ID:   coty.RandIDPlaceholderSecond(),
			Flag: coty.RandNameFlagShort(),
			Argument: &ArgumentOpt{
				ArgValues: []ArgValue{
					RandArgValue(),
				},
			},
		}
		pointer := NewPlaceholder(opt)

		require.Equal(t, opt.ID, pointer.GetID())
		require.Equal(t, opt.Flag, pointer.GetNameFlag())
		require.Equal(t, MakeArgument(opt.Argument), pointer.GetArgData())

		require.True(t, pointer.HasArg())
		require.True(t, pointer.HasFlag())
	})
}

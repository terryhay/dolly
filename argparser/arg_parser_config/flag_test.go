package arg_parser_config

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Flag

		require.Equal(t, coty.NameFlagUndefined, pointer.GetNameMain())
		require.Nil(t, pointer.GetNamesAdditional())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := FlagOpt{
			NameMain: coty.RandNameFlagShort(),
			NamesAdditional: map[coty.NameFlag]struct{}{
				coty.RandNameFlagShortSecond(): {},
			},
			HelpInfo: gofakeit.Name(),
		}
		pointer := MakeFlag(opt)

		require.Equal(t, opt.NameMain, pointer.GetNameMain())
		require.Equal(t, opt.NamesAdditional, pointer.GetNamesAdditional())
		require.Equal(t, opt.HelpInfo, pointer.GetDescriptionHelpInfo())
	})
}

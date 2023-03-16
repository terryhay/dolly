package arg_parser_config

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestCommandDescription(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Command

		require.Equal(t, coty.NameCommandUndefined, pointer.GetNameMain())

		require.Nil(t, pointer.GetNamesAdditional())
		require.Nil(t, pointer.GetPlaceholders())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Equal(t, "", pointer.CreateStringWithCommandNames())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := CommandOpt{
			NameMain: coty.RandNameCommand(),
			NamesAdditional: map[coty.NameCommand]struct{}{
				coty.RandNameCommandSecond(): {},
			},
			Placeholders: []*PlaceholderOpt{
				{},
			},
			HelpInfo: gofakeit.Name(),
		}
		pointer := NewCommand(opt)

		require.Equal(t, opt.NameMain, pointer.GetNameMain())
		require.Equal(t, opt.NamesAdditional, pointer.GetNamesAdditional())
		require.Equal(t, createPlaceholders(opt.Placeholders), pointer.GetPlaceholders())
		require.Equal(t, opt.HelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, fmt.Sprintf("%s, %s", coty.RandNameCommand(), coty.RandNameCommandSecond()), pointer.CreateStringWithCommandNames())
	})
}

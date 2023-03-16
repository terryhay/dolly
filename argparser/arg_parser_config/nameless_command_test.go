package arg_parser_config

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestNamelessCommandGetters(t *testing.T) {
	t.Parallel()

	opt := NamelessCommandOpt{
		Placeholders: []*PlaceholderOpt{
			{},
		},
		HelpInfo: gofakeit.Name(),
	}
	pointer := NewNamelessCommand(&opt)

	require.Equal(t, coty.NameCommandUndefined, pointer.GetNameMain())
	require.Equal(t, 0, len(pointer.GetNamesAdditional()))
	require.Equal(t, opt.HelpInfo, pointer.GetDescriptionHelpInfo())
	require.Equal(t, 0, len(pointer.CreateStringWithCommandNames()))
}

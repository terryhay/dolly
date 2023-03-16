package arg_parser_config

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestArgGroupDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Placeholder

		require.Equal(t, coty.ArgPlaceholderIDUndefined, pointer.GetID())
		require.False(t, pointer.GetIsFlagOptional())
		require.Nil(t, pointer.GetDescriptionFlags())
		require.Nil(t, pointer.GetArgument())

		require.False(t, pointer.HasArg())
		require.False(t, pointer.HasFlags())
		require.False(t, pointer.HasFlagName(coty.RandNameFlagShort()))
		require.False(t, pointer.IsFlagRequired())
		require.False(t, pointer.IsArgRequired())
		require.Nil(t, pointer.FlagByName(coty.RandNameFlagShort()))
		require.True(t, len(pointer.CreateStringWithFlagNames()) == 0)
	})

	t.Run("initialized_pointer_with_flag", func(t *testing.T) {
		t.Parallel()

		opt := PlaceholderOpt{
			ID:             coty.RandIDPlaceholder(),
			IsFlagOptional: gofakeit.Bool(),
			FlagsByNames: map[coty.NameFlag]*FlagOpt{
				coty.RandNameFlagShort(): {},
			},
		}
		pointer := NewPlaceholder(opt)

		require.Equal(t, opt.ID, pointer.GetID())
		require.Equal(t, opt.IsFlagOptional, pointer.GetIsFlagOptional())
		require.Equal(t, createFlags(opt.FlagsByNames), pointer.GetDescriptionFlags())
		require.Nil(t, pointer.GetArgument())

		require.False(t, pointer.HasArg())
		require.True(t, pointer.HasFlags())
		require.True(t, pointer.HasFlagName(coty.RandNameFlagShort()))
		require.Equal(t, !opt.IsFlagOptional, pointer.IsFlagRequired())
		require.False(t, pointer.IsArgRequired())
		require.NotNil(t, pointer.FlagByName(coty.RandNameFlagShort()))
		require.Equal(t, coty.RandNameFlagShort().String(), pointer.CreateStringWithFlagNames())
	})

	t.Run("initialized_pointer_with_arg", func(t *testing.T) {
		t.Parallel()

		opt := PlaceholderOpt{
			ID:             coty.RandIDPlaceholder(),
			IsFlagOptional: gofakeit.Bool(),
			Argument:       &ArgumentOpt{},
			FlagsByNames: map[coty.NameFlag]*FlagOpt{
				coty.RandNameFlagShort(): {},
			},
		}
		pointer := NewPlaceholder(opt)

		require.Equal(t, opt.ID, pointer.GetID())
		require.Equal(t, opt.IsFlagOptional, pointer.GetIsFlagOptional())
		require.Equal(t, 1, len(pointer.GetDescriptionFlags()))
		require.Equal(t, MakeArgument(opt.Argument), pointer.GetArgument())

		require.Equal(t, !opt.IsFlagOptional, pointer.IsFlagRequired())
		require.True(t, pointer.HasFlags())
		require.True(t, pointer.HasArg())
		require.NotNil(t, pointer.FlagByName(coty.RandNameFlagShort()))
		require.True(t, len(pointer.CreateStringWithFlagNames()) > 0)
	})
}

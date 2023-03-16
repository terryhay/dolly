package arg_parser_config

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
)

func TestArgumentGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *Argument

		require.True(t, len(pointer.GetSynopsisHelpDescription()) == 0)
		require.Nil(t, pointer.GetDefaultValues())
		require.Nil(t, pointer.GetAllowedValues())
		require.False(t, pointer.GetIsOptional())
		require.False(t, pointer.GetIsList())

		require.NoError(t, pointer.IsArgAllowed(clArg.RandCmdArg()))
		require.False(t, pointer.IsRequired())
		require.Equal(t, 0, len(pointer.CreateStringWithArgInfo()))
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := ArgumentOpt{
			IsList:        true,
			DefaultValues: []string{clArg.RandCmdArg().String()},
			AllowedValues: map[string]struct{}{
				clArg.RandCmdArg().String(): {},
			},
			DescSynopsisHelp: gofakeit.Name(),
		}
		pointer := MakeArgument(&opt)

		require.Equal(t, opt.DescSynopsisHelp, pointer.GetSynopsisHelpDescription())
		require.Equal(t, opt.DefaultValues, pointer.GetDefaultValues())
		require.Equal(t, opt.AllowedValues, pointer.GetAllowedValues())
		require.Equal(t, opt.IsOptional, pointer.GetIsOptional())
		require.Equal(t, opt.IsList, pointer.GetIsList())

		require.NoError(t, pointer.IsArgAllowed(clArg.RandCmdArg()))
		require.ErrorIs(t, pointer.IsArgAllowed(clArg.RandCmdArgSecond()), ErrIsArgAllowed)
		require.Equal(t, !opt.IsOptional, pointer.IsRequired())
		require.Equal(t, fmt.Sprintf("list of arguments allowed [%s] default [%[1]s]", clArg.RandCmdArg()), pointer.CreateStringWithArgInfo())
	})
}

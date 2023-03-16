package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/tools/common_types"
)

func TestPlaceholderGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *PlaceholderOpt
		pointer := NewPlaceholder(opt)
		require.Nil(t, pointer)

		require.ErrorIs(t, pointer.IsValid(), ErrPlaceholderName)
		require.Equal(t, 0, len(pointer.GetName()))
		require.False(t, pointer.GetIsFlagOptional())
		require.Equal(t, 0, len(pointer.GetFlags()))
		require.Nil(t, pointer.GetArgument())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := PlaceholderOpt{
			Name: common_types.RandNamePlaceholder().String(),

			// one or two fields must be set
			Flags: []*FlagOpt{
				{},
			},
			Argument: &ArgumentOpt{},

			// optional
			IsFlagOptional: gofakeit.Bool(),
		}
		pointer := NewPlaceholder(&opt)
		require.NotNil(t, pointer)

		require.Equal(t, opt.Name, pointer.GetName().String())
		require.Equal(t, opt.IsFlagOptional, pointer.GetIsFlagOptional())
		require.Equal(t, toFlagSlice(opt.Flags), pointer.GetFlags())
		require.Equal(t, NewArgument(opt.Argument), pointer.GetArgument())
	})
}

func TestPlaceholderIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/placeholder_cases/err_invalid_placeholder_name.yaml",
			expError: ErrPlaceholderName,
		},
		{
			caseFile: "./test_cases/placeholder_cases/err_no_flags_and_arg.yaml",
			expError: ErrPlaceholderNoFlagsNoArg,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NotNil(t, config)
			require.NoError(t, err)

			err = config.IsValid()
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

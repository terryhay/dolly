package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/stretchr/testify/require"
)

func TestArgumentGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *ArgumentOpt
		pointer := NewArgument(opt)
		require.Nil(t, pointer)

		require.False(t, pointer.GetIsList())
		require.Equal(t, 0, len(pointer.GetHelpName()))
		require.Nil(t, pointer.GetDefaultValues())
		require.Nil(t, pointer.GetAllowedValues())
		require.False(t, pointer.GetIsOptional())
		require.NoError(t, pointer.IsValid())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := ArgumentOpt{
			IsList:        gofakeit.Bool(),
			HelpName:      gofakeit.Name(),
			DefaultValues: []string{gofakeit.Name()},
			AllowedValues: []string{gofakeit.Name()},
			IsOptional:    gofakeit.Bool(),
		}
		pointer := NewArgument(&opt)
		require.NotNil(t, pointer)

		require.Equal(t, opt.IsList, pointer.GetIsList())
		require.Equal(t, opt.HelpName, pointer.GetHelpName().String())
		require.Equal(t, coty.ToSliceTypesSorted[coty.ArgValue](opt.DefaultValues), pointer.GetDefaultValues())
		require.Equal(t, coty.ToSliceTypesSorted[coty.ArgValue](opt.AllowedValues), pointer.GetAllowedValues())
		require.Equal(t, opt.IsOptional, pointer.GetIsOptional())
	})
}

func TestArgumentsDescriptionIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/argument_cases/err_no_help_name.yaml",
			expError: ErrArgumentHelpName,
		},
		{
			caseFile: "./test_cases/argument_cases/err_stop_char_in_allowed_values.yaml",
			expError: ErrArgumentStopCharacter,
		},
		{
			caseFile: "./test_cases/argument_cases/err_stop_char_in_default_values.yaml",
			expError: ErrArgumentStopCharacter,
		},
		{
			caseFile: "./test_cases/argument_cases/err_default_value_is_not_allowed.yaml",
			expError: ErrArgumentDefaultValueIsNotAllowed,
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

func TestArgumentsDescriptionIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/argument_cases/no_amount_type.yaml",
		},
		{
			caseFile: "./test_cases/argument_cases/no_amount_type.yaml",
		},
		{
			caseFile: "./test_cases/argument_cases/no_default_values.yaml",
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NotNil(t, config)
			require.NoError(t, err)

			require.NoError(t, config.IsValid())
		})
	}
}

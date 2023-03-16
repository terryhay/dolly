package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/stretchr/testify/require"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *HelpCommandOpt
		pointer := NewHelpCommand(opt)
		require.Nil(t, pointer)

		require.Equal(t, 0, len(pointer.GetMainName()))
		require.Nil(t, pointer.GetAdditionalNamesSorted())
		require.Equal(t, coty.InfoChapterDESCRIPTIONUndefined, pointer.GetChapterDescriptionInfo())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := &HelpCommandOpt{
			MainName:        gofakeit.Name(),
			AdditionalNames: []string{gofakeit.Name()},
		}
		pointer := NewHelpCommand(opt)
		require.NotNil(t, pointer)

		require.Equal(t, opt.MainName, pointer.GetMainName().String())
		require.Equal(t, coty.ToSliceTypesSorted[coty.NameCommand](opt.AdditionalNames), pointer.GetAdditionalNamesSorted())
		require.Equal(t, "print help info", pointer.GetChapterDescriptionInfo().String())
	})
}

func TestHelpCommandDescriptionIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/help_command_cases/err_no_help_command.yaml",
			expError: ErrHelpCommandNilPointer,
		},
		{
			caseFile: "./test_cases/help_command_cases/err_no_main_name.yaml",
			expError: ErrHelpCommandMainName,
		},
		{
			caseFile: "./test_cases/help_command_cases/err_additional_name.yaml",
			expError: ErrHelpCommandAdditionalNames,
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

func TestHelpCommandDescriptionIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/help_command_cases/no_additional_names.yaml",
		},
		{
			caseFile: "./test_cases/help_command_cases/additional_names.yaml",
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NoError(t, err)
			require.NoError(t, config.IsValid())
		})
	}
}

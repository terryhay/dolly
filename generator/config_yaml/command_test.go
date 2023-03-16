package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/stretchr/testify/require"
)

func TestCommandGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *CommandOpt
		pointer := NewCommand(opt)
		require.Nil(t, pointer)

		require.Equal(t, coty.NameCommandUndefined, pointer.GetMainName())
		require.Nil(t, pointer.GetAdditionalNames())
		require.Nil(t, pointer.GetUsingPlaceholdersSorted())
		require.Equal(t, coty.InfoChapterDESCRIPTIONUndefined, pointer.GetChapterDescriptionInfo())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := CommandOpt{
			MainName:        coty.RandNameCommand().String(),
			AdditionalNames: []string{coty.RandNameCommandSecond().String()},
			UsingPlaceholders: []string{
				coty.RandNamePlaceholder().String(),
			},
			ChapterDescriptionInfo: gofakeit.Name(),
		}
		pointer := NewCommand(&opt)
		require.NotNil(t, opt)

		require.Equal(t, opt.MainName, pointer.GetMainName().String())
		require.Equal(t, coty.ToSliceTypesSorted[coty.NameCommand](opt.AdditionalNames), pointer.GetAdditionalNames())
		require.Equal(t, opt.ChapterDescriptionInfo, pointer.GetChapterDescriptionInfo().String())
		require.Equal(t, coty.ToSliceTypesSorted[coty.NamePlaceholder](opt.UsingPlaceholders), pointer.GetUsingPlaceholdersSorted())
	})
}

func TestCommandIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/command_cases/err_no_main_name.yaml",
			expError: ErrCommandMainName,
		},
		{
			caseFile: "./test_cases/command_cases/err_no_chapter_description_info.yaml",
			expError: ErrCommandNoChapterDescriptionInfo,
		},
		{
			caseFile: "./test_cases/command_cases/err_invalid_additional_name.yaml",
			expError: ErrCommandAdditionalNames,
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

func TestCommandIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/command_cases/no_additional_names.yaml",
		},
		{
			caseFile: "./test_cases/command_cases/no_argument.yaml",
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

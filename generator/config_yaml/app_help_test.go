package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/stretchr/testify/require"
)

func TestAppHelpGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *AppHelpOpt
		pointer := NewAppHelp(opt)
		require.Nil(t, pointer)

		require.ErrorIs(t, pointer.IsValid(), ErrAppHelpNilPointer)
		require.Equal(t, 0, len(pointer.GetApplicationName()))
		require.Equal(t, 0, len(pointer.GetHelpInfoChapterName()))
		require.Equal(t, 0, len(pointer.GetHelpInfoChapterDescription()))

	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := AppHelpOpt{
			AppName:         gofakeit.Name(),
			ChapterNameInfo: gofakeit.Name(),
			ChapterDescriptionInfo: []string{
				gofakeit.Name(),
			},
		}
		pointer := NewAppHelp(&opt)

		require.Equal(t, opt.AppName, pointer.GetApplicationName().String())
		require.Equal(t, opt.ChapterNameInfo, pointer.GetHelpInfoChapterName().String())
		require.Equal(t, coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.ChapterDescriptionInfo), pointer.GetHelpInfoChapterDescription())
	})
}

func TestAppHelpIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/app_help_cases/err_no_app_help.yaml",
			expError: ErrAppHelpNilPointer,
		},
		{
			caseFile: "./test_cases/app_help_cases/err_no_app_name.yaml",
			expError: ErrAppHelpNoAppName,
		},
		{
			caseFile: "./test_cases/app_help_cases/err_no_chapter_name_info.yaml",
			expError: ErrAppHelpNoChapterNameInfo,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NoError(t, err)

			err = config.IsValid()
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

func TestAppHelpIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/app_help_cases/no_chapter_description_info.yaml",
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

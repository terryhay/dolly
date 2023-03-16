package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	coty "github.com/terryhay/dolly/tools/common_types"

	"github.com/stretchr/testify/require"
)

func TestFlagGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *FlagOpt
		pointer := NewFlag(opt)
		require.Nil(t, pointer)

		require.Equal(t, 0, len(pointer.GetMainName()))
		require.Equal(t, 0, len(pointer.GetDescriptionHelpInfo()))
		require.False(t, pointer.GetIsOptional())
		require.Equal(t, 0, len(pointer.GetAdditionalNames()))
		require.Equal(t, 0, len(pointer.NamesSorted()))
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := FlagOpt{
			MainName:               coty.RandNameFlagShort().String(),
			ChapterDescriptionInfo: gofakeit.Name(),
			IsOptional:             gofakeit.Bool(),
			AdditionalNames:        []string{coty.RandNameFlagShortSecond().String()},
		}
		pointer := NewFlag(&opt)
		require.NotNil(t, pointer)

		require.Equal(t, opt.MainName, pointer.GetMainName().String())
		require.Equal(t, opt.ChapterDescriptionInfo, pointer.GetDescriptionHelpInfo().String())
		require.Equal(t, opt.IsOptional, pointer.GetIsOptional())

		require.Equal(t, coty.ToSliceTypesSorted[coty.NameFlag](opt.AdditionalNames), pointer.GetAdditionalNames())
		require.Equal(t, []coty.NameFlag{coty.RandNameFlagShort(), coty.RandNameFlagShortSecond()}, pointer.NamesSorted())
	})
}

func TestFlagIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/flag_cases/err_no_main_name.yaml",
			expError: ErrFlagMainName,
		},
		{
			caseFile: "./test_cases/flag_cases/err_no_chapter_description_info.yaml",
			expError: ErrFlagNoChapterDescriptionInfo,
		},
		{
			caseFile: "./test_cases/flag_cases/err_invalid_additional_name.yaml",
			expError: ErrFlagAdditionalName,
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

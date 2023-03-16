package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestNamelessCommandGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *NamelessCommandOpt
		pointer := NewNamelessCommand(opt)
		require.Nil(t, pointer)

		require.Equal(t, coty.InfoChapterDESCRIPTIONUndefined, pointer.GetChapterDescriptionInfo())
		require.Nil(t, pointer.GetUsingPlaceholders())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := NamelessCommandOpt{
			ChapterDescriptionInfo: gofakeit.Name(),
			UsingPlaceholders: []string{
				coty.RandNamePlaceholder().String(),
			},
		}
		pointer := NewNamelessCommand(&opt)
		require.NotNil(t, pointer)

		require.Equal(t, pointer.chapterDescriptionInfo, pointer.GetChapterDescriptionInfo())
		require.Equal(t, coty.ToSliceTypesSorted[coty.NamePlaceholder](opt.UsingPlaceholders), pointer.GetUsingPlaceholders())
	})
}

func TestNamelessCommandIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/nameless_command_cases/err_no_chapter_description_info.yaml",
			expError: ErrNamelessCommandNoChapterDescriptionInfo,
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

func TestNamelessCommandDescriptionIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/nameless_command_cases/no_placeholders.yaml",
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NotNil(t, config)
			require.NoError(t, err)
		})
	}
}

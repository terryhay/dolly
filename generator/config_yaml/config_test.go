package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *ConfigOpt
		pointer := NewConfig(opt)
		require.Nil(t, pointer)

		require.ErrorIs(t, pointer.IsValid(), ErrConfigNilPointer)
		require.Equal(t, "", pointer.GetVersion())
		require.Nil(t, pointer.GetArgParserConfig())
		require.Nil(t, pointer.GetHelpOutConfig())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := &ConfigOpt{
			Version:         gofakeit.Name(),
			ArgParserConfig: &ArgParserConfigOpt{},
			HelpOutConfig:   &HelpOutConfigOpt{},
		}
		pointer := NewConfig(opt)
		require.NotNil(t, pointer)

		require.Equal(t, opt.Version, pointer.GetVersion())
		require.Equal(t, NewArgParserConfig(opt.ArgParserConfig), pointer.GetArgParserConfig())
		require.Equal(t, NewHelpOutConfig(opt.HelpOutConfig), pointer.GetHelpOutConfig())
	})
}

func TestConfigIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/config_cases/err_no_version.yaml",
			expError: ErrConfigNoVersion,
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

func TestConfigIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/config_cases/valid.yaml",
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

package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/stretchr/testify/require"
)

func TestHelpOutConfig(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *HelpOutConfigOpt
		pointer := NewHelpOutConfig(opt)
		require.Nil(t, pointer)

		require.Equal(t, HelpOutToolPlainText, pointer.GetUsingTool())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := HelpOutConfigOpt{
			UsingTool: usingToolManStyle,
		}
		pointer := NewHelpOutConfig(&opt)
		require.NotNil(t, pointer)

		require.Equal(t, makeHelpOutTool(opt.UsingTool), pointer.GetUsingTool())
	})
}

func TestHelpOutConfigIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/help_out_config_cases/err_invalid_using_tool_value.yaml",
			expError: ErrHelpOutUnexpectedTool,
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

func TestHelpOutConfigIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/help_out_config_cases/plain.yaml",
		},
		{
			caseFile: "./test_cases/help_out_config_cases/man_style.yaml",
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

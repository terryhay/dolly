package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName       string
		configYamlPath string
		expErr         error
	}{
		{
			caseName:       "empty_string_path",
			configYamlPath: "",
			expErr:         ErrLoadReadConfigFile,
		},
		{
			caseName:       "not_existed_path",
			configYamlPath: "./non-exist/path",
			expErr:         ErrLoadReadConfigFile,
		},
		{
			caseName:       "unmarshal_config_file_error",
			configYamlPath: "./test_cases/config_cases/err_unmarshal.yaml",
			expErr:         ErrLoadUnmarshal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseName, func(t *testing.T) {
			config, err := Load(osd.New(), tc.configYamlPath)
			require.Nil(t, config)
			require.ErrorIs(t, err, tc.expErr)
		})
	}

}

func TestLoadConfigSuccess(t *testing.T) {
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

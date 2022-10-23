package config_yaml

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestGetConfigErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName          string
		configYamlPath    string
		expectedErrorCode dollyerr.Code
	}{
		{
			caseName:          "empty_string_path",
			configYamlPath:    "",
			expectedErrorCode: dollyerr.CodeGetConfigReadFileError,
		},
		{
			caseName:          "not_existed_path",
			configYamlPath:    "./non-exist/path",
			expectedErrorCode: dollyerr.CodeGetConfigReadFileError,
		},
		{
			caseName:          "unmarshal_config_file_error",
			configYamlPath:    "./test_cases/config_cases/no_version.yaml",
			expectedErrorCode: dollyerr.CodeGetConfigUnmarshalError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			config, err := GetConfig(td.configYamlPath)
			require.Nil(t, config)

			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}

}

func TestGetConfig(t *testing.T) {
	t.Parallel()

	config, err := GetConfig("./test_cases/config_cases/no_flag_descriptions.yaml")
	require.NotNil(t, config)
	require.Nil(t, err)
}

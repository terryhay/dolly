package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigGetters(t *testing.T) {
	t.Parallel()

	var pointer *Config

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetVersion())
		require.Nil(t, pointer.GetArgParserConfig())
		require.Nil(t, pointer.GetHelpOutConfig())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &Config{
			Version:         gofakeit.Name(),
			ArgParserConfig: &ArgParserConfig{},
			HelpOutConfig:   &HelpOutConfig{},
		}

		require.Equal(t, pointer.Version, pointer.GetVersion())
		require.Equal(t, pointer.ArgParserConfig, pointer.GetArgParserConfig())
		require.Equal(t, pointer.HelpOutConfig, pointer.GetHelpOutConfig())
	})
}

func TestConfigUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_version.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: config unmarshal error: no required field \"version\"",
		},
		{
			yamlFileName:      "no_arg_parser_config.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: config unmarshal error: no required field \"arg_parser_config\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/config_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &Config{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

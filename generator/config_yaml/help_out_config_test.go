package config_yaml

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestHelpOutConfig(t *testing.T) {
	t.Parallel()

	var pointer *HelpOutConfig

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, HelpOutToolPlainText, pointer.GetUsingTool())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &HelpOutConfig{
			UsingTool: HelpOutToolManStyle,
		}

		require.Equal(t, pointer.UsingTool, pointer.GetUsingTool())
	})

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer = &HelpOutConfig{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestHelpOutConfigUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "invalid_using_tool_value.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: help_out_config.using_tool unmarshal error: unexpected value invalid_value (expected: \"plain\", \"man_style\")",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/help_out_config_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}
}

func TestUnmarshalHelpOutConfig(t *testing.T) {
	t.Parallel()

	config, err := GetConfig("./test_cases/help_out_config_cases/plain.yaml")
	require.Nil(t, err)
	require.NotNil(t, config)

	config, err = GetConfig("./test_cases/help_out_config_cases/man_style.yaml")
	require.Nil(t, err)
	require.NotNil(t, config)
}

package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigGetters(t *testing.T) {
	t.Parallel()

	var pointer *Config

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetVersion())
		require.Nil(t, pointer.GetAppHelpDescription())
		require.Nil(t, pointer.GetHelpCommandDescription())
		require.Nil(t, pointer.GetNamelessCommandDescription())
		require.Nil(t, pointer.GetCommandDescriptions())
		require.Nil(t, pointer.GetFlagDescriptions())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &Config{
			Version:                    gofakeit.Name(),
			AppHelpDescription:         &AppHelpDescription{},
			HelpCommandDescription:     &HelpCommandDescription{},
			NamelessCommandDescription: &NamelessCommandDescription{},
			CommandDescriptions:        []*CommandDescription{{}},
			FlagDescriptions:           []*FlagDescription{{}},
		}

		require.Equal(t, pointer.Version, pointer.GetVersion())
		require.Equal(t, pointer.AppHelpDescription, pointer.GetAppHelpDescription())
		require.Equal(t, pointer.HelpCommandDescription, pointer.GetHelpCommandDescription())
		require.Equal(t, pointer.NamelessCommandDescription, pointer.GetNamelessCommandDescription())
		require.Equal(t, pointer.CommandDescriptions, pointer.GetCommandDescriptions())
		require.Equal(t, pointer.FlagDescriptions, pointer.GetFlagDescriptions())
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
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: config unmarshal error: no required field \"version\"",
		},
		{
			yamlFileName:      "no_app_help_description.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: config unmarshal error: no required field \"app_help_description\"",
		},
		{
			yamlFileName:      "no_help_description.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: config unmarshal error: no required field \"help_command_description\"",
		},
		{
			yamlFileName:      "no_help_command_description.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: config unmarshal error: no required field \"help_command_description\"",
		},
		{
			yamlFileName:      "no_command_description_and_nameless_command.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: config unmarshal error: one or more of fields \"nameless_command_description\" or \"command_descriptions\" must be set",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/config_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
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

func TestConfigUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_flag_descriptions.yaml",
		},
		{
			yamlFileName: "no_command_descriptions_but_has_nameless_command_description.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/config_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

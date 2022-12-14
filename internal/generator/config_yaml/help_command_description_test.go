package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *HelpCommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetCommand())
		require.Nil(t, pointer.GetAdditionalCommands())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &HelpCommandDescription{
			Command:            gofakeit.Name(),
			AdditionalCommands: []string{gofakeit.Name()},
		}

		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.AdditionalCommands, pointer.GetAdditionalCommands())
	})
}

func TestHelpCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_command.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: helpCommandDescription unmarshal error: no required field \"command\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/help_command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &HelpCommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestHelpCommandDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_additional_commands.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/help_command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *HelpCommandDescription

		require.Equal(t, "", pointer.GetCommand())
		require.Nil(t, pointer.GetAdditionalCommands())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := &HelpCommandDescriptionSrc{
			Command:            gofakeit.Name(),
			AdditionalCommands: []string{gofakeit.Name()},
		}
		pointer := src.ToConstPtr()

		require.Equal(t, pointer.command, pointer.GetCommand())
		require.Equal(t, pointer.additionalCommands, pointer.GetAdditionalCommands())
	})
}

func TestHelpCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_command.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: helpCommandDescription unmarshal error: no required field \"command\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/help_command_description_cases/%s", tc.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, tc.expectedErrorText, err.Error().Error())
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

	testCases := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_additional_commands.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/help_command_description_cases/%s", tc.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

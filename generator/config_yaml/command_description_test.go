package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *CommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetCommand())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
		require.Nil(t, pointer.GetAdditionalCommands())
		require.Nil(t, pointer.GetArgumentsDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &CommandDescription{
			command:              gofakeit.Name(),
			descriptionHelpInfo:  gofakeit.Name(),
			requiredFlags:        []string{gofakeit.Name()},
			optionalFlags:        []string{gofakeit.Name()},
			additionalCommands:   []string{gofakeit.Name()},
			argumentsDescription: &ArgumentsDescription{},
		}

		require.Equal(t, pointer.command, pointer.GetCommand())
		require.Equal(t, pointer.descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.requiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.optionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, pointer.additionalCommands, pointer.GetAdditionalCommands())
		require.Equal(t, pointer.argumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_command.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"command\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &CommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestCommandDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_additional_names.yaml",
		},
		{
			yamlFileName: "no_arguments_description.yaml",
		},
		{
			yamlFileName: "no_required_flags.yaml",
		},
		{
			yamlFileName: "no_optional_flags.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

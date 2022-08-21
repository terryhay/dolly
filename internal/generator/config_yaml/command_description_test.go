package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/pkg/dollyerr"
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
			Command:              gofakeit.Name(),
			DescriptionHelpInfo:  gofakeit.Name(),
			RequiredFlags:        []string{gofakeit.Name()},
			OptionalFlags:        []string{gofakeit.Name()},
			AdditionalCommands:   []string{gofakeit.Name()},
			ArgumentsDescription: &ArgumentsDescription{},
		}

		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.OptionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, pointer.AdditionalCommands, pointer.GetAdditionalCommands())
		require.Equal(t, pointer.ArgumentsDescription, pointer.GetArgumentsDescription())
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
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"command\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "config_yaml.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
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

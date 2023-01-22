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

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *CommandDescription

		require.Equal(t, "", pointer.GetCommand())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
		require.Nil(t, pointer.GetAdditionalCommands())
		require.Nil(t, pointer.GetArgumentsDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := CommandDescriptionSrc{
			Command:              gofakeit.Name(),
			DescriptionHelpInfo:  gofakeit.Name(),
			RequiredFlags:        []string{gofakeit.Name()},
			OptionalFlags:        []string{gofakeit.Name()},
			AdditionalCommands:   []string{gofakeit.Name()},
			ArgumentsDescription: &ArgumentsDescription{},
		}
		pointer := src.ToConstPtr()

		require.Equal(t, src.Command, pointer.GetCommand())
		require.Equal(t, src.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, src.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, src.OptionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, src.AdditionalCommands, pointer.GetAdditionalCommands())
		require.Equal(t, src.ArgumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/command_description_cases/%s", tc.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, tc.expectedErrorText, err.Error().Error())
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

	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/command_description_cases/%s", tc.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

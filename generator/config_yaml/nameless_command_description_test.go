package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestNamelessCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *NamelessCommandDescription

		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
		require.Nil(t, pointer.GetArgumentsDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := NamelessCommandDescriptionSrc{
			DescriptionHelpInfo:  gofakeit.Name(),
			RequiredFlags:        []string{gofakeit.Name()},
			OptionalFlags:        []string{gofakeit.Name()},
			ArgumentsDescription: &ArgumentsDescription{},
		}
		pointer := src.ToConstPtr()

		require.Equal(t, pointer.descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.requiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.optionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, pointer.argumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestNamelessCommandDescriptionErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/nameless_command_description_cases/%s", tc.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, tc.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &NamelessCommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestNamelessCommandDescriptionNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_arguments_description.yaml",
		},
		{
			yamlFileName: "no_optional_flags.yaml",
		},
		{
			yamlFileName: "no_required_flags.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/nameless_command_description_cases/%s", tc.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

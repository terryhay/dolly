package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgumentsDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *ArgumentsDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, apConf.ArgAmountTypeNoArgs, pointer.GetAmountType())
		require.Equal(t, "", pointer.GetSynopsisHelpDescription())
		require.Nil(t, pointer.GetDefaultValues())
		require.Nil(t, pointer.GetAllowedValues())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &ArgumentsDescription{
			AmountType:              apConf.ArgAmountTypeSingle,
			SynopsisHelpDescription: gofakeit.Name(),
			DefaultValues:           []string{gofakeit.Name()},
			AllowedValues:           []string{gofakeit.Name()},
		}

		require.Equal(t, pointer.AmountType, pointer.GetAmountType())
		require.Equal(t, pointer.SynopsisHelpDescription, pointer.GetSynopsisHelpDescription())
		require.Equal(t, pointer.DefaultValues, pointer.GetDefaultValues())
		require.Equal(t, pointer.AllowedValues, pointer.GetAllowedValues())
	})
}

func TestArgumentsDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_amount_type.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: no required field \"amount_type\"",
		},
		{
			yamlFileName:      "no_synopsis_description.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: no required field \"synopsis_description\"",
		},
		{
			yamlFileName:      "unexpected_amount_type.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: can't convert string value \"amount_type\": unexpected \"amount_type\" value: trash\\nallowed values: \"single\", \"array\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/arguments_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &ArgumentsDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestArgumentsDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_default_values.yaml",
		},
		{
			yamlFileName: "no_allowed_values.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/arguments_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

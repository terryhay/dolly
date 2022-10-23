package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *FlagDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetFlag())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Equal(t, "", pointer.GetSynopsisDescription())
		require.Nil(t, pointer.GetArgumentsDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &FlagDescription{
			Flag:                 gofakeit.Name(),
			DescriptionHelpInfo:  gofakeit.Name(),
			SynopsisDescription:  gofakeit.Name(),
			ArgumentsDescription: &ArgumentsDescription{},
		}

		require.Equal(t, pointer.Flag, pointer.GetFlag())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.SynopsisDescription, pointer.GetSynopsisDescription())
		require.Equal(t, pointer.ArgumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestFlagDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_flag.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: flagDescription unmarshal error: no required field \"flag\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: flagDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/flag_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &FlagDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestFlagDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_arguments_description.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/flag_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

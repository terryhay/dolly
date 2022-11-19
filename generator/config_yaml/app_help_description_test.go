package config_yaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppHelpDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *AppHelpDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, "", pointer.GetApplicationName())
		require.Equal(t, "", pointer.GetNameHelpInfo())
		require.Nil(t, pointer.GetDescriptionHelpInfo())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &AppHelpDescription{
			applicationName: gofakeit.Name(),
			nameHelpInfo:    gofakeit.Name(),
			descriptionHelpInfo: []string{
				gofakeit.Name(),
			},
		}

		require.Equal(t, pointer.applicationName, pointer.GetApplicationName())
		require.Equal(t, pointer.nameHelpInfo, pointer.GetNameHelpInfo())
		require.Equal(t, pointer.descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
	})
}

func TestAppHelpDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_app_name.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"app_name\"",
		},
		{
			yamlFileName:      "no_name_help_info.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"name_help_info\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "confYML.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/app_help_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error().Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &AppHelpDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

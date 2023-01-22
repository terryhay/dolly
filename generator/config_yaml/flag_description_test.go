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

	t.Run("nil_pointer", func(t *testing.T) {
		var pointer *FlagDescription

		require.Equal(t, "", pointer.GetFlag())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Equal(t, "", pointer.GetSynopsisDescription())
		require.Nil(t, pointer.GetArgumentsDescription())
		require.Nil(t, nil, pointer.GetAdditionalFlags())
		require.Nil(t, nil, pointer.ExtractSortedFlags())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		src := FlagDescriptionSrc{
			Flag:                 "flag1",
			DescriptionHelpInfo:  gofakeit.Name(),
			SynopsisDescription:  gofakeit.Name(),
			ArgumentsDescription: &ArgumentsDescription{},
			AdditionalFlags:      []string{"flag2"},
		}
		pointer := src.ToConstPtr()

		require.Equal(t, src.Flag, pointer.GetFlag())
		require.Equal(t, src.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, src.SynopsisDescription, pointer.GetSynopsisDescription())
		require.Equal(t, src.ArgumentsDescription, pointer.GetArgumentsDescription())
		require.Equal(t, src.AdditionalFlags, pointer.GetAdditionalFlags())
		require.Equal(t, []string{"flag1", "flag2"}, pointer.ExtractSortedFlags())
	})
}

func TestFlagDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/flag_description_cases/%s", tc.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, dollyerr.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, tc.expectedErrorText, err.Error().Error())
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

	testCases := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_arguments_description.yaml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./test_cases/flag_description_cases/%s", tc.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}

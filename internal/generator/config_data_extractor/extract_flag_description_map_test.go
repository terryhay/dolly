package config_data_extractor

import (
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractFlagDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName          string
		flagDescriptions  []*config_yaml.FlagDescription
		expectedErrorCode dollyerr.Code
	}{
		{
			caseName: "single_empty_flag_description",
			flagDescriptions: []*config_yaml.FlagDescription{
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_front",
			flagDescriptions: []*config_yaml.FlagDescription{
				nil,
				{},
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_back",
			flagDescriptions: []*config_yaml.FlagDescription{
				{},
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions",
			flagDescriptions: []*config_yaml.FlagDescription{
				{
					Flag: "flag",
				},
				{
					Flag: "flag",
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractFlagDescriptionMap(td.flagDescriptions)
			require.Nil(t, flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}

func TestExtractFlagDescriptionMap(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName         string
		flagDescriptions []*config_yaml.FlagDescription
		expectedMap      map[string]*config_yaml.FlagDescription
	}{
		{
			caseName:         "no_flag_description",
			flagDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			flagDescriptions: []*config_yaml.FlagDescription{
				{
					Flag: "flag",
				},
			},
			expectedMap: map[string]*config_yaml.FlagDescription{
				"flag": {
					Flag: "flag",
				},
			},
		},
		{
			caseName: "two_flag_descriptions",
			flagDescriptions: []*config_yaml.FlagDescription{
				{
					Flag: "flag1",
				},
				{
					Flag: "flag2",
				},
			},
			expectedMap: map[string]*config_yaml.FlagDescription{
				"flag1": {
					Flag: "flag1",
				},
				"flag2": {
					Flag: "flag2",
				},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractFlagDescriptionMap(td.flagDescriptions)
			require.Nil(t, err)
			require.Equal(t, len(td.expectedMap), len(flagDescriptionMap))

			for flag, expectedFlagDescription := range td.expectedMap {
				flagDescription, contain := flagDescriptionMap[flag]
				require.True(t, contain)
				require.Equal(t, expectedFlagDescription.Flag, flagDescription.Flag)
			}
		})
	}
}

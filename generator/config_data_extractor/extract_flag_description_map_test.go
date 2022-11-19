package config_data_extractor

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractFlagDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName          string
		flagDescriptions  []*confYML.FlagDescription
		expectedErrorCode dollyerr.Code
	}{
		{
			caseName: "single_empty_flag_description",
			flagDescriptions: []*confYML.FlagDescription{
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_front",
			flagDescriptions: []*confYML.FlagDescription{
				nil,
				{},
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_back",
			flagDescriptions: []*confYML.FlagDescription{
				{},
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions",
			flagDescriptions: []*confYML.FlagDescription{
				confYML.FlagDescriptionSrc{Flag: "flag"}.ToConstPtr(),
				confYML.FlagDescriptionSrc{Flag: "flag"}.ToConstPtr(),
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
		flagDescriptions []*confYML.FlagDescription
		expectedMap      map[string]*confYML.FlagDescription
	}{
		{
			caseName:         "no_flag_description",
			flagDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			flagDescriptions: []*confYML.FlagDescription{
				confYML.FlagDescriptionSrc{Flag: "flag"}.ToConstPtr(),
			},
			expectedMap: map[string]*confYML.FlagDescription{
				"flag": confYML.FlagDescriptionSrc{Flag: "flag"}.ToConstPtr(),
			},
		},
		{
			caseName: "two_flag_descriptions",
			flagDescriptions: []*confYML.FlagDescription{
				confYML.FlagDescriptionSrc{Flag: "flag1"}.ToConstPtr(),
				confYML.FlagDescriptionSrc{Flag: "flag2"}.ToConstPtr(),
			},
			expectedMap: map[string]*confYML.FlagDescription{
				"flag1": confYML.FlagDescriptionSrc{Flag: "flag1"}.ToConstPtr(),
				"flag2": confYML.FlagDescriptionSrc{Flag: "flag2"}.ToConstPtr(),
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
				require.Equal(t, expectedFlagDescription.GetFlag(), flagDescription.GetFlag())
			}
		})
	}
}

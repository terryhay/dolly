package config_data_extractor

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCommandDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName            string
		commandDescriptions []*confYML.CommandDescription
		expectedErrorCode   dollyerr.Code
	}{
		{
			caseName: "single_empty_command_description",
			commandDescriptions: []*confYML.CommandDescription{
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_front",
			commandDescriptions: []*confYML.CommandDescription{
				nil,
				{},
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_back",
			commandDescriptions: []*confYML.CommandDescription{
				{},
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_1",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
					AdditionalCommands: []string{
						"command",
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_10",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
					AdditionalCommands: []string{
						"command",
						"command1",
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_01",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
					AdditionalCommands: []string{
						"command1",
						"command",
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
				}.ToConstPtr(),
				confYML.CommandDescriptionSrc{
					Command: "command",
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractCommandDescriptionMap(td.commandDescriptions)
			require.Nil(t, flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}

func TestExtractCommandDescriptionMap(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName            string
		commandDescriptions []*confYML.CommandDescription
		expectedMap         map[string]*confYML.CommandDescription
	}{
		{
			caseName:            "no_flag_description",
			commandDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
				}.ToConstPtr(),
			},
			expectedMap: map[string]*confYML.CommandDescription{
				"command": confYML.CommandDescriptionSrc{
					Command: "command",
				}.ToConstPtr(),
			},
		},
		{
			caseName: "single_flag_description_with_additional_command",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command",
					AdditionalCommands: []string{
						"command1",
					},
				}.ToConstPtr(),
			},
			expectedMap: map[string]*confYML.CommandDescription{
				"command": confYML.CommandDescriptionSrc{
					Command: "command",
					AdditionalCommands: []string{
						"command1",
					},
				}.ToConstPtr(),
			},
		},

		{
			caseName: "two_flag_descriptions",
			commandDescriptions: []*confYML.CommandDescription{
				confYML.CommandDescriptionSrc{
					Command: "command1",
				}.ToConstPtr(),
				confYML.CommandDescriptionSrc{
					Command: "command2",
				}.ToConstPtr(),
			},
			expectedMap: map[string]*confYML.CommandDescription{
				"command1": confYML.CommandDescriptionSrc{
					Command: "command1",
				}.ToConstPtr(),
				"command2": confYML.CommandDescriptionSrc{
					Command: "command2",
				}.ToConstPtr(),
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractCommandDescriptionMap(td.commandDescriptions)
			require.Nil(t, err)

			require.Equal(t, len(td.expectedMap), len(flagDescriptionMap))

			for command, expectedCommandDescription := range td.expectedMap {
				flagDescription, contain := flagDescriptionMap[command]
				require.True(t, contain)
				require.Equal(t, expectedCommandDescription.GetCommand(), flagDescription.GetCommand())
			}
		})
	}
}

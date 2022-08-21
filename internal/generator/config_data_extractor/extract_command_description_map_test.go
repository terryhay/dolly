package config_data_extractor

import (
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCommandDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName            string
		commandDescriptions []*config_yaml.CommandDescription
		expectedErrorCode   dollyerr.Code
	}{
		{
			caseName: "single_empty_command_description",
			commandDescriptions: []*config_yaml.CommandDescription{
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_front",
			commandDescriptions: []*config_yaml.CommandDescription{
				nil,
				{},
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_back",
			commandDescriptions: []*config_yaml.CommandDescription{
				{},
				nil,
			},
			expectedErrorCode: dollyerr.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_1",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []string{
						"command",
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_10",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []string{
						"command",
						"command1",
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_01",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []string{
						"command1",
						"command",
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
				},
				{
					Command: "command",
				},
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
		commandDescriptions []*config_yaml.CommandDescription
		expectedMap         map[string]*config_yaml.CommandDescription
	}{
		{
			caseName:            "no_flag_description",
			commandDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
				},
			},
			expectedMap: map[string]*config_yaml.CommandDescription{
				"command": {
					Command: "command",
				},
			},
		},
		{
			caseName: "single_flag_description_with_additional_command",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []string{
						"command1",
					},
				},
			},
			expectedMap: map[string]*config_yaml.CommandDescription{
				"command": {
					Command: "command",
					AdditionalCommands: []string{
						"command1",
					},
				},
			},
		},

		{
			caseName: "two_flag_descriptions",
			commandDescriptions: []*config_yaml.CommandDescription{
				{
					Command: "command1",
				},
				{
					Command: "command2",
				},
			},
			expectedMap: map[string]*config_yaml.CommandDescription{
				"command1": {
					Command: "command1",
				},
				"command2": {
					Command: "command2",
				},
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
				require.Equal(t, expectedCommandDescription.Command, flagDescription.Command)
			}
		})
	}
}

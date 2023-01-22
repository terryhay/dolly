package id_template_data_creator

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"testing"
)

func TestIDTemplateDataCreator(t *testing.T) {
	t.Parallel()

	command := gofakeit.Color()
	additionalCommand := gofakeit.Color()
	commandDescriptionHelpInfo := gofakeit.Name()

	helpCommand := gofakeit.Color()
	additionalHelpCommand := gofakeit.Color()

	flag1 := gofakeit.Color()
	flag2 := "F"

	creator := NewIDTemplateCreator()
	commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		[]*confYML.CommandDescription{
			confYML.CommandDescriptionSrc{
				Command: command,
				AdditionalCommands: []string{
					additionalCommand,
				},
				DescriptionHelpInfo: commandDescriptionHelpInfo,
			}.ToConstPtr(),
			{
				// fake empty command
			},
		},
		confYML.HelpCommandDescriptionSrc{
			Command: helpCommand,
			AdditionalCommands: []string{
				additionalHelpCommand,
			},
		}.ToConstPtr(),
		&confYML.NamelessCommandDescription{},
		map[string]*confYML.FlagDescription{
			flag1: confYML.FlagDescriptionSrc{
				Flag: flag1,
			}.ToConstPtr(),
			flag2: confYML.FlagDescriptionSrc{
				Flag: flag2,
			}.ToConstPtr(),
		})

	expectedCommandID := creator.CreateID(PrefixCommandID, command)
	expectedHelpCommandID := "CommandIDPrintHelpInfo"
	expectedCommandsIDTemplateData := map[string]*IDTemplateData{
		command: {
			id:       expectedCommandID,
			nameID:   creator.CreateID(PrefixCommandStringID, command),
			callName: command,
			comment:  commandDescriptionHelpInfo,
		},
		additionalCommand: {
			id:       expectedCommandID,
			nameID:   creator.CreateID(PrefixCommandStringID, additionalCommand),
			callName: additionalCommand,
			comment:  commandDescriptionHelpInfo,
		},
		helpCommand: {
			id:       expectedHelpCommandID,
			nameID:   creator.CreateID(PrefixCommandStringID, helpCommand),
			callName: helpCommand,
			comment:  helpCommandComment,
		},
		additionalHelpCommand: {
			id:       expectedHelpCommandID,
			nameID:   creator.CreateID(PrefixCommandStringID, additionalHelpCommand),
			callName: additionalHelpCommand,
			comment:  helpCommandComment,
		},
		"": {},
	}

	require.Equal(t, len(expectedCommandsIDTemplateData), len(commandsIDTemplateData))
	for expectedCommand, expectedIDTemplateData := range expectedCommandsIDTemplateData {
		idTemplateData, ok := commandsIDTemplateData[expectedCommand]
		require.True(t, ok)

		require.Equal(t, expectedIDTemplateData, idTemplateData)
	}

	require.Equal(t, &IDTemplateData{id: "CommandIDNamelessCommand"}, nullCommandIDTemplateData)

	{
		flagIDTemplateData, ok := flagsIDTemplateData[flag1]
		require.True(t, ok)
		require.Equal(t, &IDTemplateData{
			id:       "",
			nameID:   creator.CreateID(PrefixFlagStringID, flag1),
			callName: flag1,
		}, flagIDTemplateData)
	}
	{
		flagIDTemplateData, ok := flagsIDTemplateData[flag2]
		require.True(t, ok)
		require.Equal(t, &IDTemplateData{
			id:       "",
			nameID:   creator.CreateID(PrefixFlagStringID, flag2),
			callName: flag2,
		}, flagIDTemplateData)
	}
}

package plain_help_out

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	tools "github.com/terryhay/dolly/utils/test_tools"
	"testing"
)

func TestCreateDescriptionChapter(t *testing.T) {
	t.Parallel()

	randDescriptionHelpInfo := gofakeit.Name()
	randCommand := apConf.Command(gofakeit.Color())
	randFlag := apConf.Flag(gofakeit.Color())
	randFlagDescriptionHelpInfo := gofakeit.Name()

	testData := []struct {
		caseName string

		descriptionHelpInfo        []string
		namelessCommandDescription apConf.NamelessCommandDescription
		commandDescriptions        []*apConf.CommandDescription
		flagDescriptions           map[apConf.Flag]*apConf.FlagDescription

		expected string
	}{
		{
			caseName:            "empty",
			descriptionHelpInfo: nil,
			flagDescriptions:    nil,

			expected: descriptionChapterTitle,
		},
		{
			caseName:            "two_flags",
			descriptionHelpInfo: []string{randDescriptionHelpInfo},
			flagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
	%s

The flags are as follows:
	[1m%s[0m
		%s
`,
				randDescriptionHelpInfo,
				randFlag,
				randFlagDescriptionHelpInfo,
			),
		},
		{
			caseName:            "command_and_flag_descriptions",
			descriptionHelpInfo: []string{randDescriptionHelpInfo},
			commandDescriptions: []*apConf.CommandDescription{
				{
					Commands: map[apConf.Command]bool{randCommand: true},
				},
			},
			flagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
	%s

The commands are as follows:
	[1m%s[0m
		
The flags are as follows:
	[1m%s[0m
		%s
`, randDescriptionHelpInfo, randCommand, randFlag, randFlagDescriptionHelpInfo),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			actual := createDescriptionChapter(
				td.descriptionHelpInfo,
				td.namelessCommandDescription,
				td.commandDescriptions,
				td.flagDescriptions)

			ok, msg := tools.CheckSpaces(actual)
			require.True(t, ok, msg)

			require.Equal(t, td.expected, actual)
		})
	}
}

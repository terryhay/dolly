package helpprinter

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/internal/test_tools"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestCreateDescriptionChapter(t *testing.T) {
	t.Parallel()

	randDescriptionHelpInfo := gofakeit.Name()
	randCommand := dollyconf.Command(gofakeit.Color())
	randFlag := dollyconf.Flag(gofakeit.Color())
	randFlagDescriptionHelpInfo := gofakeit.Name()

	testData := []struct {
		caseName string

		descriptionHelpInfo        []string
		namelessCommandDescription dollyconf.NamelessCommandDescription
		commandDescriptions        []*dollyconf.CommandDescription
		flagDescriptions           map[dollyconf.Flag]*dollyconf.FlagDescription

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
			flagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
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
			commandDescriptions: []*dollyconf.CommandDescription{
				{
					Commands: map[dollyconf.Command]bool{randCommand: true},
				},
			},
			flagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
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
			actual := CreateDescriptionChapter(
				td.descriptionHelpInfo,
				td.namelessCommandDescription,
				td.commandDescriptions,
				td.flagDescriptions)

			ok, msg := test_tools.CheckSpaces(actual)
			require.True(t, ok, msg)

			require.Equal(t, td.expected, actual)
		})
	}
}

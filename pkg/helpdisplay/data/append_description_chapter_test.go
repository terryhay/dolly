package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
	"testing"
)

func TestAppendDescriptionChapter(t *testing.T) {
	t.Parallel()

	randDescriptionHelpInfo := gofakeit.Name()
	randDescriptionHelpInfoSecond := gofakeit.Name()
	randCommand := dollyconf.Command(gofakeit.Color())
	randCommandSecond := dollyconf.Command(gofakeit.Color())
	randNamelessCommandDescription := gofakeit.Name()
	randFlag := dollyconf.Flag(gofakeit.Color())
	randFlagDescriptionHelpInfo := gofakeit.Name()
	randFlagSecond := dollyconf.Flag("-s")

	testData := []struct {
		caseName string

		appDescription             dollyconf.ApplicationDescription
		namelessCommandDescription dollyconf.NamelessCommandDescription
		commandDescriptions        []*dollyconf.CommandDescription
		flagDescriptions           map[dollyconf.Flag]*dollyconf.FlagDescription

		expected string
	}{
		{
			caseName:         "empty",
			flagDescriptions: nil,
		},
		{
			caseName: "two_flags",
			appDescription: dollyconf.ApplicationDescription{
				DescriptionHelpInfo: []string{
					randDescriptionHelpInfo,
					randDescriptionHelpInfoSecond,
				},
			},
			flagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
    %s

    %s

The flags are as follows:
    [1m%s[0m
        %s`,
				randDescriptionHelpInfo,
				randDescriptionHelpInfoSecond,
				randFlag,
				randFlagDescriptionHelpInfo,
			),
		},
		{
			caseName: "command_and_flag_descriptions",

			appDescription: dollyconf.ApplicationDescription{
				DescriptionHelpInfo: []string{randDescriptionHelpInfo},
			},
			namelessCommandDescription: dollyconf.NewNamelessCommandDescription(0, randNamelessCommandDescription, nil, nil, nil),
			commandDescriptions: []*dollyconf.CommandDescription{
				{
					Commands: map[dollyconf.Command]bool{randCommand: true},
				},
				{
					Commands: map[dollyconf.Command]bool{randCommandSecond: true},
				},
			},
			flagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
				randFlagSecond: {},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
    %s

The commands are as follows:
    [1m<empty>[0m %s

    [1m%s[0m
        

    [1m%s[0m 

The flags are as follows:
    [1m-s[0m 

    [1m%s[0m
        %s`, randDescriptionHelpInfo, randNamelessCommandDescription, randCommand, randCommandSecond, randFlag, randFlagDescriptionHelpInfo),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {

			paragraphs := appendDescriptionChapter(make([]*Paragraph, 0),
				td.appDescription,
				td.namelessCommandDescription,
				td.commandDescriptions,
				td.flagDescriptions)

			paragraphTexts := make([]string, 0, len(paragraphs))
			for i := range paragraphs {
				paragraphTexts = append(paragraphTexts, paragraphs[i].String())
			}
			text := strings.Join(paragraphTexts, "\n")

			//ok, msg := test_tools.CheckSpaces(text)
			//require.True(t, ok, msg)

			require.Equal(t, td.expected, text)
		})
	}
}

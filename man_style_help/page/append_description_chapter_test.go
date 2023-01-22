package page

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
	"testing"
)

func TestAppendDescriptionChapter(t *testing.T) {
	t.Parallel()

	randDescriptionHelpInfo := gofakeit.Name()
	randDescriptionHelpInfoSecond := gofakeit.Name()
	randCommand := apConf.Command(gofakeit.Color())
	randCommandSecond := apConf.Command(gofakeit.Color())
	randNamelessCommandDescription := gofakeit.Name()
	randFlag := apConf.Flag(gofakeit.Color())
	randFlagDescriptionHelpInfo := gofakeit.Name()
	randFlagSecond := apConf.Flag("-s")

	testCases := []struct {
		caseName string

		appDescription             apConf.ApplicationDescription
		namelessCommandDescription apConf.NamelessCommandDescription
		commandDescriptions        []*apConf.CommandDescription
		flagDescriptions           map[apConf.Flag]*apConf.FlagDescription

		expected string
	}{
		{
			caseName:         "empty",
			flagDescriptions: nil,
		},
		{
			caseName: "two_flags",
			appDescription: apConf.ApplicationDescriptionSrc{
				DescriptionHelpInfo: []string{
					randDescriptionHelpInfo,
					randDescriptionHelpInfoSecond,
				},
			}.ToConst(),
			flagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
				randFlag: apConf.FlagDescriptionSrc{
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				}.ToConstPtr(),
			},

			expected: fmt.Sprintf(`
[1mDESCRIPTION[0m
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

			appDescription: apConf.ApplicationDescriptionSrc{
				DescriptionHelpInfo: []string{randDescriptionHelpInfo},
			}.ToConst(),
			namelessCommandDescription: apConf.NewNamelessCommandDescription(0, randNamelessCommandDescription, nil, nil, nil),
			commandDescriptions: []*apConf.CommandDescription{
				apConf.CommandDescriptionSrc{
					Commands: map[apConf.Command]bool{randCommand: true},
				}.ToConstPtr(),
				apConf.CommandDescriptionSrc{
					Commands: map[apConf.Command]bool{randCommandSecond: true},
				}.ToConstPtr(),
			},
			flagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
				randFlag: apConf.FlagDescriptionSrc{
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				}.ToConstPtr(),
				randFlagSecond: {},
			},

			expected: fmt.Sprintf(`
[1mDESCRIPTION[0m
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

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {

			paragraphs := appendDescriptionChapter(make([]Paragraph, 0),
				tc.appDescription,
				tc.namelessCommandDescription,
				tc.commandDescriptions,
				tc.flagDescriptions)

			paragraphTexts := make([]string, 0, len(paragraphs))
			for i := range paragraphs {
				paragraphTexts = append(paragraphTexts, paragraphs[i].String())
			}
			text := strings.Join(paragraphTexts, "\n")

			//ok, msg := test_tools.CheckSpaces(text)
			//require.True(t, ok, msg)

			require.Equal(t, tc.expected, text)
		})
	}
}

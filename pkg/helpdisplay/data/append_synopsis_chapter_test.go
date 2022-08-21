package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"strings"
	"testing"
)

func TestCreateSynopsisChapter(t *testing.T) {
	t.Parallel()

	appDescription := dollyconf.ApplicationDescription{
		AppName:      gofakeit.Color(),
		NameHelpInfo: gofakeit.Name(),
		DescriptionHelpInfo: []string{
			gofakeit.Name(),
			gofakeit.Name(),
		},
	}

	namelessCmdRequiredFlag := dollyconf.Flag("-rf")
	namelessCmdOptionalFlag := dollyconf.Flag("-of")
	command := dollyconf.Command(gofakeit.Color())
	commandFlagWithSingleArgument := dollyconf.Flag("-sa")
	commandFlagWithListArgument := dollyconf.Flag("-la")
	commandFlagDescriptionWithSingleArgument := &dollyconf.FlagDescription{
		ArgDescription: &dollyconf.ArgumentsDescription{
			AmountType:              dollyconf.ArgAmountTypeSingle,
			SynopsisHelpDescription: "arg",
		},
	}
	commandFlagDescriptionWithListArgumentDefaultValue := "val1"
	commandFlagDescriptionWithListArgumentAllowedValue := "val2"
	commandFlagDescriptionWithListArgument := &dollyconf.FlagDescription{
		ArgDescription: &dollyconf.ArgumentsDescription{
			AmountType:              dollyconf.ArgAmountTypeList,
			SynopsisHelpDescription: "str",
			DefaultValues: []string{
				commandFlagDescriptionWithListArgumentDefaultValue,
			},
			AllowedValues: map[string]bool{
				commandFlagDescriptionWithListArgumentAllowedValue: true,
			},
		},
	}

	testData := []struct {
		caseName string

		appDescription             dollyconf.ApplicationDescription
		namelessCommandDescription dollyconf.NamelessCommandDescription
		commandDescriptions        []*dollyconf.CommandDescription
		flagDescriptions           map[dollyconf.Flag]*dollyconf.FlagDescription

		expected string
	}{
		{
			caseName: "full_data",

			appDescription: appDescription,
			namelessCommandDescription: dollyconf.NewNamelessCommandDescription(
				0,
				"nameless command description",
				&dollyconf.ArgumentsDescription{
					SynopsisHelpDescription: "args",
				},
				map[dollyconf.Flag]bool{
					namelessCmdRequiredFlag: true,
				},
				map[dollyconf.Flag]bool{
					namelessCmdOptionalFlag: true,
				},
			),
			commandDescriptions: []*dollyconf.CommandDescription{
				{
					Commands: map[dollyconf.Command]bool{
						command: true,
					},
					RequiredFlags: map[dollyconf.Flag]bool{
						commandFlagWithSingleArgument: true,
					},
					OptionalFlags: map[dollyconf.Flag]bool{
						commandFlagWithListArgument: true,
					},
				},
			},
			flagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
				namelessCmdRequiredFlag: {
					ArgDescription: &dollyconf.ArgumentsDescription{
						SynopsisHelpDescription: "arg1",
					},
				},
				commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
				commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
			},

			expected: fmt.Sprintf(`[1mSYNOPSIS[0m
    [1m%s[0m [1m-rf[0m [[1m-of[0m]
    [1m%s %s[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]`, appDescription.GetAppName(), appDescription.GetAppName(), command),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			paragraphs := appendSynopsisChapterParagraphs(make([]*Paragraph, 0),
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

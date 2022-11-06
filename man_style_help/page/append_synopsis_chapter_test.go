package page

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"strings"
	"testing"
)

func TestCreateSynopsisChapter(t *testing.T) {
	t.Parallel()

	appDescription := apConf.ApplicationDescription{
		AppName:      gofakeit.Color(),
		NameHelpInfo: gofakeit.Name(),
		DescriptionHelpInfo: []string{
			gofakeit.Name(),
			gofakeit.Name(),
		},
	}

	namelessCmdRequiredFlag := apConf.Flag("-rf")
	namelessCmdOptionalFlag := apConf.Flag("-of")
	command := apConf.Command(gofakeit.Color())
	commandFlagWithSingleArgument := apConf.Flag("-sa")
	commandFlagWithListArgument := apConf.Flag("-la")
	commandFlagDescriptionWithSingleArgument := &apConf.FlagDescription{
		ArgDescription: &apConf.ArgumentsDescription{
			AmountType:              apConf.ArgAmountTypeSingle,
			SynopsisHelpDescription: "arg",
		},
	}
	commandFlagDescriptionWithListArgumentDefaultValue := "val1"
	commandFlagDescriptionWithListArgumentAllowedValue := "val2"
	commandFlagDescriptionWithListArgument := &apConf.FlagDescription{
		ArgDescription: &apConf.ArgumentsDescription{
			AmountType:              apConf.ArgAmountTypeList,
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

		appDescription             apConf.ApplicationDescription
		namelessCommandDescription apConf.NamelessCommandDescription
		commandDescriptions        []*apConf.CommandDescription
		flagDescriptions           map[apConf.Flag]*apConf.FlagDescription

		expected string
	}{
		{
			caseName: "full_data",

			appDescription: appDescription,
			namelessCommandDescription: apConf.NewNamelessCommandDescription(
				0,
				"nameless command description",
				&apConf.ArgumentsDescription{
					SynopsisHelpDescription: "args",
				},
				map[apConf.Flag]bool{
					namelessCmdRequiredFlag: true,
				},
				map[apConf.Flag]bool{
					namelessCmdOptionalFlag: true,
				},
			),
			commandDescriptions: []*apConf.CommandDescription{
				{
					Commands: map[apConf.Command]bool{
						command: true,
					},
					RequiredFlags: map[apConf.Flag]bool{
						commandFlagWithSingleArgument: true,
					},
					OptionalFlags: map[apConf.Flag]bool{
						commandFlagWithListArgument: true,
					},
				},
			},
			flagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
				namelessCmdRequiredFlag: {
					ArgDescription: &apConf.ArgumentsDescription{
						SynopsisHelpDescription: "arg1",
					},
				},
				commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
				commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
			},

			expected: fmt.Sprintf(`
[1mSYNOPSIS[0m
    [1m%s[0m [1m-rf[0m [[1m-of[0m]
    [1m%s %s[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]`, appDescription.GetAppName(), appDescription.GetAppName(), command),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			paragraphs := appendSynopsisChapterParagraphs(make([]Paragraph, 0),
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

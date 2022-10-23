package plain_help_out

import (
	"fmt"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/utils/test_tools"
	"testing"
)

func TestCreateSynopsisChapter(t *testing.T) {
	t.Parallel()

	var (
		appName = "appname"
		chapter string

		expectedChapter string
	)

	t.Run("full_data", func(t *testing.T) {

		nullCommandRequiredFlag := apConf.Flag("-rf")
		nullCommandOptionalFlag := apConf.Flag("-of")

		commandFlagWithSingleArgument := apConf.Flag("-sa")
		commandFlagDescriptionWithSingleArgument := &apConf.FlagDescription{
			ArgDescription: &apConf.ArgumentsDescription{
				AmountType:              apConf.ArgAmountTypeSingle,
				SynopsisHelpDescription: "arg",
			},
		}

		commandFlagWithListArgument := apConf.Flag("-la")
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

		command := apConf.Command("command")

		namelessCommandDescription := apConf.NewNamelessCommandDescription(
			0,
			"nameless command description",
			&apConf.ArgumentsDescription{
				SynopsisHelpDescription: "args",
			},
			map[apConf.Flag]bool{
				nullCommandRequiredFlag: true,
			},
			map[apConf.Flag]bool{
				nullCommandOptionalFlag: true,
			},
		)
		commandDescriptions := []*apConf.CommandDescription{
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
		}
		flagDescriptions := map[apConf.Flag]*apConf.FlagDescription{
			nullCommandRequiredFlag: {
				ArgDescription: &apConf.ArgumentsDescription{
					SynopsisHelpDescription: "arg1",
				},
			},
			commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
			commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
		}

		chapter = createSynopsisChapter(appName, namelessCommandDescription, commandDescriptions, flagDescriptions)

		ok, msg := test_tools.CheckSpaces(chapter)
		require.True(t, ok, msg)

		require.Equal(t,
			`[1mSYNOPSIS[0m
	[1mappname[0m [1m-rf[0m [[1m-of[0m]
	[1mappname command[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]

`,
			chapter)
	})

	t.Run("no_commands", func(t *testing.T) {
		commandDescriptions := []*apConf.CommandDescription{
			{},
		}

		expectedChapter = fmt.Sprintf(`[1mSYNOPSIS[0m
	[1m%s[0m

`, appName)

		chapter = createSynopsisChapter(appName, nil, commandDescriptions, nil)
		ok, msg := test_tools.CheckSpaces(chapter)
		require.True(t, ok, msg)

		require.Equal(t, expectedChapter, chapter)
	})
}
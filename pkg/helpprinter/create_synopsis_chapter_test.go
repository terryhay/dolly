package helpprinter

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/internal/test_tools"
	"github.com/terryhay/dolly/pkg/dollyconf"
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

		nullCommandRequiredFlag := dollyconf.Flag("-rf")
		nullCommandOptionalFlag := dollyconf.Flag("-of")

		commandFlagWithSingleArgument := dollyconf.Flag("-sa")
		commandFlagDescriptionWithSingleArgument := &dollyconf.FlagDescription{
			ArgDescription: &dollyconf.ArgumentsDescription{
				AmountType:              dollyconf.ArgAmountTypeSingle,
				SynopsisHelpDescription: "arg",
			},
		}

		commandFlagWithListArgument := dollyconf.Flag("-la")
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

		command := dollyconf.Command("command")

		namelessCommandDescription := dollyconf.NewNamelessCommandDescription(
			0,
			"nameless command description",
			&dollyconf.ArgumentsDescription{
				SynopsisHelpDescription: "args",
			},
			map[dollyconf.Flag]bool{
				nullCommandRequiredFlag: true,
			},
			map[dollyconf.Flag]bool{
				nullCommandOptionalFlag: true,
			},
		)
		commandDescriptions := []*dollyconf.CommandDescription{
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
		}
		flagDescriptions := map[dollyconf.Flag]*dollyconf.FlagDescription{
			nullCommandRequiredFlag: {
				ArgDescription: &dollyconf.ArgumentsDescription{
					SynopsisHelpDescription: "arg1",
				},
			},
			commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
			commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
		}

		chapter = CreateSynopsisChapter(appName, namelessCommandDescription, commandDescriptions, flagDescriptions)

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
		commandDescriptions := []*dollyconf.CommandDescription{
			{},
		}

		expectedChapter = fmt.Sprintf(`[1mSYNOPSIS[0m
	[1m%s[0m

`, appName)

		chapter = CreateSynopsisChapter(appName, nil, commandDescriptions, nil)
		ok, msg := test_tools.CheckSpaces(chapter)
		require.True(t, ok, msg)

		require.Equal(t, expectedChapter, chapter)
	})
}

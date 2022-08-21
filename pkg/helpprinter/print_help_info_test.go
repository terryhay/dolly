package helpprinter

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/internal/test_tools"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"testing"
)

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	t.Run("empty_config", func(t *testing.T) {
		out := test_tools.CatchStdOut(func() {
			PrintHelpInfo(dollyconf.ArgParserConfig{})
		})

		ok, msg := test_tools.CheckSpaces(out)
		require.True(t, ok, msg)

		require.Equal(t, `[1mNAME[0m
	[1m[0m – 

[1mSYNOPSIS[0m

[1mDESCRIPTION[0m

`, out)
	})

	t.Run("simple_case", func(t *testing.T) {
		out := test_tools.CatchStdOut(func() {
			PrintHelpInfo(dollyconf.NewArgParserConfig(
				dollyconf.ApplicationDescription{
					AppName:      "appname",
					NameHelpInfo: "name help info",
				},
				nil,
				[]*dollyconf.CommandDescription{
					{
						ID:                  1,
						DescriptionHelpInfo: "command id 1 description help info",
						Commands: map[dollyconf.Command]bool{
							"command": true,
						},
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType:              dollyconf.ArgAmountTypeSingle,
							SynopsisHelpDescription: "str",
						},
						RequiredFlags: map[dollyconf.Flag]bool{
							"-rf1": true,
						},
						OptionalFlags: map[dollyconf.Flag]bool{
							"-of1": true,
						},
					},
					{
						ID:                  2,
						DescriptionHelpInfo: "command id 2 description help info",
						Commands: map[dollyconf.Command]bool{
							"longcommand": true,
						},
					},
				},
				nil,
				dollyconf.NewNamelessCommandDescription(
					0,
					"nameless command description",
					nil,
					nil,
					nil,
				),
			))
		})

		ok, msg := test_tools.CheckSpaces(out)
		require.True(t, ok, msg)

		require.Equal(t, `[1mNAME[0m
	[1mappname[0m – name help info

[1mSYNOPSIS[0m
	[1mappname[0m
	[1mappname command[0m [4mstr[0m [1m-rf1[0m [[1m-of1[0m]
	[1mappname longcommand[0m

[1mDESCRIPTION[0m

The commands are as follows:
	[1m<empty>[0m	nameless command description

	[1mcommand[0m	command id 1 description help info

	[1mlongcommand[0m
		command id 2 description help info

`, out)
	})
}

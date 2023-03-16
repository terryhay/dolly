package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestCheckParseResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		command *apConf.Command
		data    *parsed.Result

		expErr error
	}{
		{
			caseName: "nil_arguments",
		},
		{
			caseName: "skip_empty_placeholder",
			command: apConf.NewCommand(apConf.CommandOpt{
				NameMain: coty.RandNameCommand(),
				Placeholders: []*apConf.PlaceholderOpt{
					{},
				},
			}),
		},
		{
			caseName: "required_flag_is_not_set",
			command: apConf.NewCommand(apConf.CommandOpt{
				NameMain: coty.RandNameCommand(),
				Placeholders: []*apConf.PlaceholderOpt{
					{
						FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
							coty.RandNameFlagShort(): {
								NameMain: coty.RandNameFlagShort(),
							},
						},
					},
				},
			}),
			expErr: ErrCheckParseResultRequiredFlagIsNotSet,
		},
		{
			caseName: "command_arg_is_not_set",
			command: apConf.NewCommand(apConf.CommandOpt{
				Placeholders: []*apConf.PlaceholderOpt{
					{
						Argument: &apConf.ArgumentOpt{},
					},
				},
			}),
			data: parsed.MakeResult(&parsed.ResultOpt{
				CommandMainName: coty.RandNameCommand(),
			}),
			expErr: ErrCheckParseResultRequiredArgIsNotSet,
		},
		{
			caseName: "empty_parsed_placeholder",
			command: apConf.NewCommand(apConf.CommandOpt{
				Placeholders: []*apConf.PlaceholderOpt{
					{
						ID:             coty.RandIDPlaceholder(),
						IsFlagOptional: true,
					},
					{
						ID: coty.RandIDPlaceholderSecond(),
						Argument: &apConf.ArgumentOpt{
							IsOptional: true,
						},
					},
				},
			}),
			data: parsed.MakeResult(&parsed.ResultOpt{
				CommandMainName: coty.RandNameCommand(),
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholderSecond(): {
						ID: coty.RandIDPlaceholderSecond(),
					},
				},
			}),
			expErr: ErrCheckParsedResultEmptyPlaceholder,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			err := checkParseResult(tc.command, tc.data)

			if tc.expErr == nil {
				require.NoError(t, err)
				return
			}

			require.ErrorIs(t, err, tc.expErr)
		})
	}
}

package parser

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string

		config apConf.ArgParserConfig
		args   []string

		expectedResult *parsed.Result
		expErr         error
	}{
		{
			caseName: "no_config_and_arguments",
			expErr:   ErrUsingCommandDescriptionNoCommands,
		},
		{
			caseName: "no_config_and_one_argument",
			args: []string{
				parsed.RandArgValue().String(),
			},
			expErr: ErrUsingCommandDescriptionNoCommands,
		},
		{
			caseName: "arg_is_not_command_name",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				Commands: []*apConf.CommandOpt{
					{NameMain: coty.RandNameCommand()},
				},
			}),

			expErr: ErrUsingCommandDescriptionNoCommands,
		},
		{
			caseName: "no_args_for_nameless_command_without_args",

			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{}),
		},
		{
			caseName: "no_args_for_nameless_command_with_required_flag",

			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrCheckParseResultRequiredFlagIsNotSet,
		},
		{
			caseName: "no_args_for_nameless_command_with_required_argument",

			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							Argument: &apConf.ArgumentOpt{},
						},
					},
				},
			}),

			expErr: ErrCheckParseResultRequiredArgIsNotSet,
		},
		{
			caseName: "waste_arg_for_nameless_command",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{},
			}),

			expErr: ErrProcessProcessFindingPlaceholder,
		},
		{
			caseName: "arg_for_nameless_command_with_incorrect_arg_group_description",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{},
					},
				},
			}),

			expErr: ErrFindingPlaceholderNotFound,
		},
		{
			caseName: "wrong_arg_for_nameless_command_with_required_flag",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrProcessProcessFindingPlaceholder,
		},
		{
			caseName: "arg_for_nameless_command_with_required_arg",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:       coty.RandIDPlaceholder(),
							Argument: &apConf.ArgumentOpt{},
						},
					},
				},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID: coty.RandIDPlaceholder(),
						Argument: &parsed.ArgumentOpt{
							ArgValues: []parsed.ArgValue{
								parsed.RandArgValue(),
							},
						},
					},
				},
			}),
		},
		{
			caseName: "try_to_set_flag_name_instead_arg_for_nameless_command",

			args: []string{
				coty.RandNameFlagShort().String(),
				coty.RandNameFlagLong().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
							Argument: &apConf.ArgumentOpt{},
						},
					},
				},
			}),

			expErr: ErrNotSetArgValueCaseNoRequiredArg,
		},
		{
			caseName: "arg_for_second_arg_list_group_of_nameless_command_with_required_flag",

			args: []string{
				parsed.RandArgValue().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:             coty.RandIDPlaceholder(),
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							Argument: &apConf.ArgumentOpt{
								IsList: true,
							},
						},
					},
				},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholderSecond(): {
						ID: coty.RandIDPlaceholderSecond(),
						Argument: &parsed.ArgumentOpt{
							ArgValues: []parsed.ArgValue{
								parsed.RandArgValue(),
							},
						},
					},
				},
			}),
		},
		{
			caseName: "duplicate_flags",

			args: []string{
				coty.RandNameFlagShort().String(),
				coty.RandNameFlagShort().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrProcessProcessFindingPlaceholder,
		},
		{
			caseName: "arg_with_dash_in_front_in_single_argument_case",

			args: []string{
				coty.RandNameFlagShort().String(),
				coty.RandNameFlagLong().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:             coty.RandIDPlaceholder(),
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							Argument: &apConf.ArgumentOpt{
								IsOptional: true,
							},
						},
					},
				},
			}),

			expErr: ErrProcessProcessFindingPlaceholder,
		},
		{
			caseName: "arg_with_dash_in_front_in_list_argument_case",

			args: []string{
				coty.RandNameFlagShort().String(),
				coty.RandNameFlagLong().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:             coty.RandIDPlaceholder(),
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsOptional: true,
								IsList:     true,
							},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
						},
					},
				},
			}),

			expErr: ErrFindingPlaceholderNotFound,
		},
		{
			caseName: "unexpected_flag_in_list_argument_case",

			args: []string{
				coty.RandNameFlagShort().String(),
				parsed.RandArgValue().String(),
				coty.RandNameFlagLong().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsList: true,
							},
						},
					},
				},
				CommandHelpOut: &apConf.HelpOutCommandOpt{
					NameMain: coty.RandNameCommand(),
					NamesAdditional: map[coty.NameCommand]struct{}{
						coty.RandNameCommand(): {},
					},
				},
			}),

			expErr: ErrProcessProcessReadingArgumentList,
		},
		{
			caseName: "duplicated_flag_in_list_argument_case",

			args: []string{
				coty.RandNameFlagShort().String(),
				parsed.RandArgValue().String(),
				coty.RandNameFlagShort().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsList: true,
							},
						},
					},
				},
			}),

			expErr: ErrProcessProcessReadingArgumentList,
		},
		{
			caseName: "correct_list_argument_case",

			args: []string{
				coty.RandNameFlagShort().String(),
				parsed.RandArgValue().String(),
				coty.RandNameFlagLong().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
							Argument: &apConf.ArgumentOpt{
								IsList: true,
							},
						},
						{
							ID:             coty.RandIDPlaceholderSecond(),
							IsFlagOptional: true,
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagLong(): {
									NameMain: coty.RandNameFlagLong(),
								},
							},
						},
					},
				},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID:   coty.RandIDPlaceholder(),
						Flag: coty.RandNameFlagShort(),
						Argument: &parsed.ArgumentOpt{
							ArgValues: []parsed.ArgValue{
								parsed.RandArgValue(),
							},
						},
					},

					coty.RandIDPlaceholderSecond(): {
						ID:   coty.RandIDPlaceholderSecond(),
						Flag: coty.RandNameFlagLong(),
					},
				},
			}),
		},
		{
			caseName: "failed_final_parsed_data_checking",

			args: []string{parsed.RandArgValue().String()},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:       coty.RandIDPlaceholder(),
							Argument: &apConf.ArgumentOpt{},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrCheckParseResultRequiredFlagIsNotSet,
		},
		{
			caseName: "arg_and_no_default_value",

			args: []string{coty.RandNameFlagShort().String()},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID:       coty.RandIDPlaceholder(),
							Argument: &apConf.ArgumentOpt{},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrNotSetArgValueCaseNoRequiredArg,
		},
		{
			caseName: "success_using_default_value",

			args: []string{
				coty.RandNameFlagShort().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							Argument: &apConf.ArgumentOpt{

								DefaultValues: []string{
									parsed.RandArgValue().String(),
								},
								AllowedValues: map[string]struct{}{
									parsed.RandArgValue().String():       {},
									parsed.RandArgValueSecond().String(): {},
								},
							},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID: coty.RandIDPlaceholder(),
						Argument: &parsed.ArgumentOpt{
							ArgValues: []parsed.ArgValue{
								parsed.RandArgValue(),
							},
						},
					},
					coty.RandIDPlaceholderSecond(): {
						ID:   coty.RandIDPlaceholderSecond(),
						Flag: coty.RandNameFlagShort(),
					},
				},
			}),
		},
		{
			caseName: "not_allowed_value",

			args: []string{
				gofakeit.Color(), coty.RandNameFlagShort().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							Argument: &apConf.ArgumentOpt{

								AllowedValues: map[string]struct{}{
									gofakeit.Color(): {},
								},
							},
						},
						{
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),

			expErr: ErrProcessProcessFindingPlaceholder,
		},
		{
			caseName: "success_allowed_value_checking",

			args: []string{
				parsed.RandArgValue().String(),
				coty.RandNameFlagShort().String(),
			},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				CommandNameless: &apConf.NamelessCommandOpt{
					Placeholders: []*apConf.PlaceholderOpt{
						{
							ID: coty.RandIDPlaceholder(),
							Argument: &apConf.ArgumentOpt{

								AllowedValues: map[string]struct{}{
									parsed.RandArgValue().String(): {},
								},
							},
						},
						{
							ID: coty.RandIDPlaceholderSecond(),
							FlagsByNames: map[coty.NameFlag]*apConf.FlagOpt{
								coty.RandNameFlagShort(): {
									NameMain: coty.RandNameFlagShort(),
								},
							},
						},
					},
				},
			}),
			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				PlaceholdersByID: map[coty.IDPlaceholder]*parsed.PlaceholderOpt{
					coty.RandIDPlaceholder(): {
						ID: coty.RandIDPlaceholder(),
						Argument: &parsed.ArgumentOpt{
							ArgValues: []parsed.ArgValue{
								parsed.RandArgValue(),
							},
						},
					},

					coty.RandIDPlaceholderSecond(): {
						ID:   coty.RandIDPlaceholderSecond(),
						Flag: coty.RandNameFlagShort(),
					},
				},
			}),
		},
		{
			caseName: "command_without_flags_and_arguments",

			args: []string{coty.RandNameCommand().String()},
			config: apConf.MakeArgParserConfig(apConf.ArgParserConfigOpt{
				Commands: []*apConf.CommandOpt{
					{NameMain: coty.RandNameCommand()},
				},
			}),

			expectedResult: parsed.MakeResult(&parsed.ResultOpt{
				CommandMainName: coty.RandNameCommand(),
			}),
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			data, err := Parse(tc.config, tc.args)

			if tc.expErr != nil {
				require.Nil(t, data)
				require.ErrorIs(t, err, tc.expErr)

				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.expectedResult, data)
		})
	}
}

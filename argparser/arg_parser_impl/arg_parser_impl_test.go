package arg_parser_impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	cmdArg "github.com/terryhay/dolly/argparser/cmd_arg"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var (
		command   = apConf.Command(gofakeit.Name())
		commandID = apConf.CommandID(gofakeit.Uint32())

		namelessCommandID = apConf.CommandID(gofakeit.Uint32())
		requiredFlag      = apConf.Flag("-" + gofakeit.Color())
		optionalFlag      = apConf.Flag("-" + gofakeit.Color())
		arg               = gofakeit.Color()
	)

	testCases := []struct {
		caseName string

		config apConf.ArgParserConfig
		args   []string

		expectedParsedData *parsed_data.ParsedData
		expectedErr        *dollyerr.Error
	}{
		{
			caseName:    "empty_config",
			expectedErr: fakeError(dollyerr.CodeArgParserNamelessCommandUndefined),
		},
		{
			caseName:    "no_command_descriptions",
			args:        []string{arg},
			expectedErr: fakeError(dollyerr.CodeArgParserIsNotInitialized),
		},
		{
			caseName: "unexpected_arg_error",
			args: []string{
				gofakeit.Color(),
			},
			config: apConf.ArgParserConfigSrc{
				CommandDescriptions: []*apConf.CommandDescription{
					apConf.CommandDescriptionSrc{
						Commands: map[apConf.Command]bool{
							apConf.Command(gofakeit.Color()): true,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeCantFindFlagNameInGroupSpec),
		},
		{
			caseName: "no_args_for_nameless_command",
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					nil,
					nil),
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
			},
		},
		{
			caseName: "no_args_for_nameless_command_with_required_flag",
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "no_args_for_nameless_command_with_required_argument",
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
					}.ToConstPtr(),
					nil,
					nil),
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserCommandDoesNotContainArgs),
		},
		{
			caseName: "waste_arg_for_nameless_command",
			args:     []string{arg},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					nil,
					nil),
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserUnexpectedArg),
		},
		{
			caseName: "arg_for_nameless_command_with_required_arg",
			args:     []string{arg},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
					}.ToConstPtr(),
					nil,
					nil),
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
				ArgData: &parsed_data.ParsedArgData{
					ArgValues: []parsed_data.ArgValue{parsed_data.ArgValue(arg)},
				},
			},
		},
		{
			caseName: "duplicate_flags",
			args: []string{
				string(requiredFlag),
				string(requiredFlag),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "arg_with_dash_in_front_in_single_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
						ArgDescription: apConf.ArgumentsDescriptionSrc{
							AmountType: apConf.ArgAmountTypeSingle,
						}.ToConstPtr(),
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "arg_with_dash_in_front_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
						ArgDescription: apConf.ArgumentsDescriptionSrc{
							AmountType: apConf.ArgAmountTypeList,
						}.ToConstPtr(),
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "unexpected_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				gofakeit.Color(),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
						ArgDescription: apConf.ArgumentsDescriptionSrc{
							AmountType: apConf.ArgAmountTypeList,
						}.ToConstPtr(),
					}.ToConstPtr(),
				},
				HelpCommandDescription: apConf.NewHelpCommandDescription(
					apConf.CommandID(gofakeit.Uint32()),
					map[apConf.Command]bool{
						apConf.Command(gofakeit.Color()): true,
					}),
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserUnexpectedFlag),
		},
		{
			caseName: "duplicated_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(requiredFlag),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
						ArgDescription: apConf.ArgumentsDescriptionSrc{
							AmountType: apConf.ArgAmountTypeList,
						}.ToConstPtr(),
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "correct_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(optionalFlag),
			},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					map[apConf.Flag]bool{
						optionalFlag: true,
					}),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
						ArgDescription: apConf.ArgumentsDescriptionSrc{
							AmountType: apConf.ArgAmountTypeList,
						}.ToConstPtr(),
					}.ToConstPtr(),
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							optionalFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
				FlagDataMap: map[apConf.Flag]*parsed_data.ParsedFlagData{
					requiredFlag: {
						Flag: requiredFlag,
						ArgData: &parsed_data.ParsedArgData{
							ArgValues: []parsed_data.ArgValue{
								parsed_data.ArgValue(arg),
							},
						},
					},
					optionalFlag: {
						Flag: optionalFlag,
					},
				},
			},
		},
		{
			caseName: "failed_final_parsed_data_checking",
			args:     []string{arg},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
					}.ToConstPtr(),
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "arg_and_no_default_value",
			args:     []string{string(requiredFlag)},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
					}.ToConstPtr(),
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserFlagMustHaveArg),
		},
		{
			caseName: "success_using_default_value",
			args:     []string{string(requiredFlag)},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
						DefaultValues: []string{
							arg,
						},
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					}.ToConstPtr(),
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
				ArgData: &parsed_data.ParsedArgData{
					ArgValues: []parsed_data.ArgValue{
						parsed_data.ArgValue(arg),
					},
				},
				FlagDataMap: map[apConf.Flag]*parsed_data.ParsedFlagData{
					requiredFlag: {
						Flag: requiredFlag,
					},
				},
			},
		},
		{
			caseName: "not_allowed_value",
			args:     []string{gofakeit.Color(), string(requiredFlag)},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					}.ToConstPtr(),
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedErr: fakeError(dollyerr.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "success_allowed_value_checking",
			args:     []string{arg, string(requiredFlag)},
			config: apConf.ArgParserConfigSrc{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					apConf.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							arg: true,
						},
					}.ToConstPtr(),
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptionSlice: []*apConf.FlagDescription{
					apConf.FlagDescriptionSrc{
						Flags: []apConf.Flag{
							requiredFlag,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
				ArgData: &parsed_data.ParsedArgData{
					ArgValues: []parsed_data.ArgValue{
						parsed_data.ArgValue(arg),
					},
				},
				FlagDataMap: map[apConf.Flag]*parsed_data.ParsedFlagData{
					requiredFlag: {
						Flag: requiredFlag,
					},
				},
			},
		},

		{
			caseName: "command_without_flags_and_arguments",
			args:     []string{string(command)},
			config: apConf.ArgParserConfigSrc{
				CommandDescriptions: []*apConf.CommandDescription{
					apConf.CommandDescriptionSrc{
						ID: commandID,
						Commands: map[apConf.Command]bool{
							command: true,
						},
					}.ToConstPtr(),
				},
			}.ToConst(),
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: commandID,
				Command:   command,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			impl := NewCmdArgParserImpl(tc.config)
			data, err := impl.Parse(cmdArg.MakeCmdArgIterator(tc.args))

			if tc.expectedErr != nil {
				require.Nil(t, data)
				require.NotNil(t, err)

				require.Equal(t, tc.expectedErr.Code(), err.Code())
				return
			}

			require.Nil(t, err)
			require.Equal(t, tc.expectedParsedData, data)
		})
	}
}

func fakeError(code dollyerr.Code) *dollyerr.Error {
	return dollyerr.NewError(code, fmt.Errorf(""))
}

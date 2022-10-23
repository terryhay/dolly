package arg_parser_impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var (
		namelessCommandID = apConf.CommandID(gofakeit.Uint32())
		requiredFlag      = apConf.Flag("-" + gofakeit.Color())
		optionalFlag      = apConf.Flag("-" + gofakeit.Color())
		arg               = gofakeit.Color()
	)

	testData := []struct {
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
			config: apConf.ArgParserConfig{
				CommandDescriptions: []*apConf.CommandDescription{
					{
						Commands: map[apConf.Command]bool{
							apConf.Command(gofakeit.Color()): true,
						},
					},
				},
			},
			expectedErr: fakeError(dollyerr.CodeCantFindFlagNameInGroupSpec),
		},
		{
			caseName: "no_args_for_nameless_command",
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					nil,
					nil),
			},
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
			},
		},
		{
			caseName: "no_args_for_nameless_command_with_required_flag",
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "no_args_for_nameless_command_with_required_argument",
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
					},
					nil,
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserCommandDoesNotContainArgs),
		},
		{
			caseName: "waste_arg_for_nameless_command",
			args:     []string{arg},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					nil,
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserUnexpectedArg),
		},
		{
			caseName: "arg_for_nameless_command_with_required_arg",
			args:     []string{arg},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
					},
					nil,
					nil),
			},
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
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "arg_with_dash_in_front_in_single_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {
						ArgDescription: &apConf.ArgumentsDescription{
							AmountType: apConf.ArgAmountTypeSingle,
						},
					},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "arg_with_dash_in_front_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {
						ArgDescription: &apConf.ArgumentsDescription{
							AmountType: apConf.ArgAmountTypeList,
						},
					},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "unexpected_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				gofakeit.Color(),
				"-" + gofakeit.Color(),
			},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						apConf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {
						ArgDescription: &apConf.ArgumentsDescription{
							AmountType: apConf.ArgAmountTypeList,
						},
					},
				},
				HelpCommandDescription: apConf.NewHelpCommandDescription(
					apConf.CommandID(gofakeit.Uint32()),
					map[apConf.Command]bool{
						apConf.Command(gofakeit.Color()): true,
					}),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserUnexpectedFlag),
		},
		{
			caseName: "duplicated_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(requiredFlag),
			},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {
						ArgDescription: &apConf.ArgumentsDescription{
							AmountType: apConf.ArgAmountTypeList,
						},
					},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "correct_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(optionalFlag),
			},
			config: apConf.ArgParserConfig{
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
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {
						ArgDescription: &apConf.ArgumentsDescription{
							AmountType: apConf.ArgAmountTypeList,
						},
					},
					optionalFlag: {},
				},
			},
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
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
					},
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "arg_and_no_default_value",
			args:     []string{string(requiredFlag)},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
					},
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserFlagMustHaveArg),
		},
		{
			caseName: "success_using_default_value",
			args:     []string{string(requiredFlag)},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
						DefaultValues: []string{
							arg,
						},
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					},
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {},
				},
			},
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
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					},
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "success_allowed_value_checking",
			args:     []string{arg, string(requiredFlag)},
			config: apConf.ArgParserConfig{
				NamelessCommandDescription: apConf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&apConf.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							arg: true,
						},
					},
					map[apConf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[apConf.Flag]*apConf.FlagDescription{
					requiredFlag: {},
				},
			},
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
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			impl := NewCmdArgParserImpl(td.config)
			data, err := impl.Parse(td.args)

			if td.expectedErr != nil {
				require.Nil(t, data)
				require.NotNil(t, err)

				require.Equal(t, td.expectedErr.Code(), err.Code())
				return
			}

			require.Nil(t, err)
			require.Equal(t, td.expectedParsedData, data)
		})
	}
}

func fakeError(code dollyerr.Code) *dollyerr.Error {
	return dollyerr.NewError(code, fmt.Errorf(""))
}
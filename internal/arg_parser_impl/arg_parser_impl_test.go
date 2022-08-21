package arg_parser_impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var (
		namelessCommandID = dollyconf.CommandID(gofakeit.Uint32())
		requiredFlag      = dollyconf.Flag("-" + gofakeit.Color())
		optionalFlag      = dollyconf.Flag("-" + gofakeit.Color())
		arg               = gofakeit.Color()
	)

	testData := []struct {
		caseName string

		config dollyconf.ArgParserConfig
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
			config: dollyconf.ArgParserConfig{
				CommandDescriptions: []*dollyconf.CommandDescription{
					{
						Commands: map[dollyconf.Command]bool{
							dollyconf.Command(gofakeit.Color()): true,
						},
					},
				},
			},
			expectedErr: fakeError(dollyerr.CodeCantFindFlagNameInGroupSpec),
		},
		{
			caseName: "no_args_for_nameless_command",
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "no_args_for_nameless_command_with_required_argument",
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
					},
					nil,
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserCommandDoesNotContainArgs),
		},
		{
			caseName: "waste_arg_for_nameless_command",
			args:     []string{arg},
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						dollyconf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						dollyconf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType: dollyconf.ArgAmountTypeSingle,
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						dollyconf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType: dollyconf.ArgAmountTypeList,
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						dollyconf.Flag(arg): true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType: dollyconf.ArgAmountTypeList,
						},
					},
				},
				HelpCommandDescription: dollyconf.NewHelpCommandDescription(
					dollyconf.CommandID(gofakeit.Uint32()),
					map[dollyconf.Command]bool{
						dollyconf.Command(gofakeit.Color()): true,
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType: dollyconf.ArgAmountTypeList,
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					nil,
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					map[dollyconf.Flag]bool{
						optionalFlag: true,
					}),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {
						ArgDescription: &dollyconf.ArgumentsDescription{
							AmountType: dollyconf.ArgAmountTypeList,
						},
					},
					optionalFlag: {},
				},
			},
			expectedParsedData: &parsed_data.ParsedData{
				CommandID: namelessCommandID,
				FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
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
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
					},
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
			},
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "arg_and_no_default_value",
			args:     []string{string(requiredFlag)},
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
					},
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserFlagMustHaveArg),
		},
		{
			caseName: "success_using_default_value",
			args:     []string{string(requiredFlag)},
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
						DefaultValues: []string{
							arg,
						},
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					},
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
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
				FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
					requiredFlag: {
						Flag: requiredFlag,
					},
				},
			},
		},
		{
			caseName: "not_allowed_value",
			args:     []string{gofakeit.Color(), string(requiredFlag)},
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							gofakeit.Color(): true,
						},
					},
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "success_allowed_value_checking",
			args:     []string{arg, string(requiredFlag)},
			config: dollyconf.ArgParserConfig{
				NamelessCommandDescription: dollyconf.NewNamelessCommandDescription(
					namelessCommandID,
					"",
					&dollyconf.ArgumentsDescription{
						AmountType: dollyconf.ArgAmountTypeSingle,
						AllowedValues: map[string]bool{
							arg: true,
						},
					},
					map[dollyconf.Flag]bool{
						requiredFlag: true,
					},
					nil),
				FlagDescriptions: map[dollyconf.Flag]*dollyconf.FlagDescription{
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
				FlagDataMap: map[dollyconf.Flag]*parsed_data.ParsedFlagData{
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

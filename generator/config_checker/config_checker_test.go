package config_checker

import (
	"github.com/brianvoe/gofakeit"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigCheckerCorrectResponse(t *testing.T) {
	t.Parallel()

	require.Nil(t, Check(nil, nil, nil))
}

func TestConfigCheckerErrors(t *testing.T) {
	t.Parallel()

	value := gofakeit.Color()

	flag := "-" + gofakeit.Color()
	if len(flag) >= maxFlagLen {
		flag = flag[:maxFlagLen]
	}

	testData := []struct {
		caseName                   string
		namelessCommandDescription *confYML.NamelessCommandDescription
		commandDescriptionMap      map[string]*confYML.CommandDescription
		flagDescriptionMap         map[string]*confYML.FlagDescription
		expectedErrorCode          dollyerr.Code
	}{
		{
			caseName: "default_value_with_no_args_amount_type_in_nameless_command",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				ArgumentsDescription: &confYML.ArgumentsDescription{
					DefaultValues: []string{value},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "default_value_with_no_args_amount_type_in_command",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					ArgumentsDescription: &confYML.ArgumentsDescription{
						DefaultValues: []string{value},
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "twp_default_values_with_no_args_amount_type",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				ArgumentsDescription: &confYML.ArgumentsDescription{
					DefaultValues: []string{
						value,
						gofakeit.Color(),
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "default_value_is_not_allowed",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				ArgumentsDescription: &confYML.ArgumentsDescription{
					AmountType: apConf.ArgAmountTypeList,
					DefaultValues: []string{
						value,
						gofakeit.Color(),
					},
					AllowedValues: []string{
						value,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigDefaultValueIsNotAllowed,
		},
		{
			caseName: "flag_with_check_arg_error",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				RequiredFlags: []string{
					flag,
				},
			},
			flagDescriptionMap: map[string]*confYML.FlagDescription{
				flag: {
					Flag: flag,
					ArgumentsDescription: &confYML.ArgumentsDescription{
						DefaultValues: []string{value},
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "duplicate_flag_in_required_list",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					ArgumentsDescription: &confYML.ArgumentsDescription{
						AmountType:    apConf.ArgAmountTypeSingle,
						DefaultValues: []string{gofakeit.Color()},
					},
					RequiredFlags: []string{
						flag,
						flag,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_optional_list",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					ArgumentsDescription: &confYML.ArgumentsDescription{
						AmountType: apConf.ArgAmountTypeSingle,
						DefaultValues: []string{
							value,
						},
						AllowedValues: []string{
							value,
						},
					},
					OptionalFlags: []string{
						flag,
						flag,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_required_and_optional_lists",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					RequiredFlags: []string{
						flag,
					},
					OptionalFlags: []string{
						flag,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},

		{
			caseName: "unused_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {},
			},
			flagDescriptionMap: map[string]*confYML.FlagDescription{
				flag: {
					Flag: flag,
				},
			},
			expectedErrorCode: dollyerr.CodeConfigFlagIsNotUsedInCommands,
		},

		{
			caseName: "undefined_required_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					RequiredFlags: []string{
						flag,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUndefinedFlag,
		},
		{
			caseName: "undefined_optional_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {
					OptionalFlags: []string{
						flag,
					},
				},
			},
			expectedErrorCode: dollyerr.CodeConfigUndefinedFlag,
		},

		{
			caseName: "nameless_command_description_with_duplicate_required_flags",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				RequiredFlags: []string{
					flag,
					flag,
				},
			},

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_optional_flags",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				OptionalFlags: []string{
					flag,
					flag,
				},
			},

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_required_and_optional_flags",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				RequiredFlags: []string{
					flag,
				},
				OptionalFlags: []string{
					flag,
				},
			},

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_required_flag_does_not_have_dash_in_front",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				RequiredFlags: []string{
					flag[1:],
				},
				OptionalFlags: []string{
					flag,
				},
			},

			expectedErrorCode: dollyerr.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "nameless_command_optional_flag_has_russian_char",
			namelessCommandDescription: &confYML.NamelessCommandDescription{
				RequiredFlags: []string{
					flag,
				},
				OptionalFlags: []string{
					"-йцукен",
				},
			},

			expectedErrorCode: dollyerr.CodeConfigIncorrectCharacterInFlagName,
		},
		{
			caseName: "command_with_too_long_required_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Color(): {
					RequiredFlags: []string{
						flag + "d",
					},
				},
			},

			expectedErrorCode: dollyerr.CodeConfigIncorrectFlagLen,
		},
		{
			caseName: "command_with_empty_optional_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Color(): {
					OptionalFlags: []string{
						"",
					},
				},
			},

			expectedErrorCode: dollyerr.CodeConfigIncorrectCharacterInFlagName,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Check(td.namelessCommandDescription, td.commandDescriptionMap, td.flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}

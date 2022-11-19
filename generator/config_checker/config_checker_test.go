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
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
					DefaultValues: []string{value},
				}.ToConstPtr(),
			}.ToConstPtr(),
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "default_value_with_no_args_amount_type_in_command",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
						DefaultValues: []string{value},
					}.ToConstPtr(),
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "twp_default_values_with_no_args_amount_type",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
					DefaultValues: []string{
						value,
						gofakeit.Color(),
					},
				}.ToConstPtr(),
			}.ToConstPtr(),
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "default_value_is_not_allowed",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
					AmountType: apConf.ArgAmountTypeList,
					DefaultValues: []string{
						value,
						gofakeit.Color(),
					},
					AllowedValues: []string{
						value,
					},
				}.ToConstPtr(),
			}.ToConstPtr(),
			expectedErrorCode: dollyerr.CodeConfigDefaultValueIsNotAllowed,
		},
		{
			caseName: "flag_with_check_arg_error",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				RequiredFlags: []string{
					flag,
				},
			}.ToConstPtr(),
			flagDescriptionMap: map[string]*confYML.FlagDescription{
				flag: confYML.FlagDescriptionSrc{
					Flag: flag,
					ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
						DefaultValues: []string{value},
					}.ToConstPtr(),
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigUnexpectedDefaultValue,
		},
		{
			caseName: "duplicate_flag_in_required_list",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
						AmountType:    apConf.ArgAmountTypeSingle,
						DefaultValues: []string{gofakeit.Color()},
					}.ToConstPtr(),
					RequiredFlags: []string{
						flag,
						flag,
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_optional_list",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					ArgumentsDescription: confYML.ArgumentsDescriptionSrc{
						AmountType: apConf.ArgAmountTypeSingle,
						DefaultValues: []string{
							value,
						},
						AllowedValues: []string{
							value,
						},
					}.ToConstPtr(),
					OptionalFlags: []string{
						flag,
						flag,
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_required_and_optional_lists",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					RequiredFlags: []string{
						flag,
					},
					OptionalFlags: []string{
						flag,
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},

		{
			caseName: "unused_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): {},
			},
			flagDescriptionMap: map[string]*confYML.FlagDescription{
				flag: confYML.FlagDescriptionSrc{
					Flag: flag,
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigFlagIsNotUsedInCommands,
		},

		{
			caseName: "undefined_required_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					RequiredFlags: []string{
						flag,
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigUndefinedFlag,
		},
		{
			caseName: "undefined_optional_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Name(): confYML.CommandDescriptionSrc{
					OptionalFlags: []string{
						flag,
					},
				}.ToConstPtr(),
			},
			expectedErrorCode: dollyerr.CodeConfigUndefinedFlag,
		},

		{
			caseName: "nameless_command_description_with_duplicate_required_flags",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				RequiredFlags: []string{
					flag,
					flag,
				},
			}.ToConstPtr(),

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_optional_flags",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				OptionalFlags: []string{
					flag,
					flag,
				},
			}.ToConstPtr(),

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_required_and_optional_flags",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				RequiredFlags: []string{
					flag,
				},
				OptionalFlags: []string{
					flag,
				},
			}.ToConstPtr(),

			expectedErrorCode: dollyerr.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_required_flag_does_not_have_dash_in_front",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				RequiredFlags: []string{
					flag[1:],
				},
				OptionalFlags: []string{
					flag,
				},
			}.ToConstPtr(),

			expectedErrorCode: dollyerr.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "nameless_command_optional_flag_has_russian_char",
			namelessCommandDescription: confYML.NamelessCommandDescriptionSrc{
				RequiredFlags: []string{
					flag,
				},
				OptionalFlags: []string{
					"-йцукен",
				},
			}.ToConstPtr(),

			expectedErrorCode: dollyerr.CodeConfigIncorrectCharacterInFlagName,
		},
		{
			caseName: "command_with_too_long_required_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Color(): confYML.CommandDescriptionSrc{
					RequiredFlags: []string{
						flag + "d",
					},
				}.ToConstPtr(),
			},

			expectedErrorCode: dollyerr.CodeConfigIncorrectFlagLen,
		},
		{
			caseName: "command_with_empty_optional_flag",
			commandDescriptionMap: map[string]*confYML.CommandDescription{
				gofakeit.Color(): confYML.CommandDescriptionSrc{
					OptionalFlags: []string{
						"",
					},
				}.ToConstPtr(),
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

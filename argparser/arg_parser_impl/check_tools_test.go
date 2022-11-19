package arg_parser_impl

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
	"testing"
)

func TestCheckNoDashInFront(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		arg           string
		expectedValue bool
	}{
		{
			caseName:      "empty",
			expectedValue: true,
		},
		{
			caseName:      "dash_in_front",
			arg:           "-" + gofakeit.Color(),
			expectedValue: false,
		},
		{
			caseName:      "no_dash_in_front",
			arg:           gofakeit.Color(),
			expectedValue: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			require.Equal(t, td.expectedValue, checkNoDashInFront(td.arg))
		})
	}
}

func TestCheckParsedData(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		commandDescription *apConf.CommandDescription
		data               *parsed_data.ParsedData

		expectedErr *dollyerr.Error
	}{
		{
			caseName: "nil_arguments",
		},
		{
			caseName: "required flag is not set",
			commandDescription: apConf.CommandDescriptionSrc{
				RequiredFlags: map[apConf.Flag]bool{
					apConf.Flag(gofakeit.Color()): true,
				},
			}.ToConstPtr(),
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "command arg is not set",
			commandDescription: apConf.CommandDescriptionSrc{
				ArgDescription: apConf.ArgumentsDescriptionSrc{
					AmountType: apConf.ArgAmountTypeSingle,
				}.ToConstPtr(),
			}.ToConstPtr(),
			expectedErr: fakeError(dollyerr.CodeArgParserCommandDoesNotContainArgs),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := checkParsedData(td.commandDescription, td.data)

			if td.expectedErr == nil {
				require.Nil(t, err)
				return
			}

			require.Equal(t, td.expectedErr.Code(), err.Code())
		})
	}
}

func TestIsValueAllowed(t *testing.T) {
	t.Parallel()

	value := gofakeit.Color()

	testData := []struct {
		caseName string

		argDescription *apConf.ArgumentsDescription
		value          string

		expectedErr *dollyerr.Error
	}{
		{
			caseName:    "nil_arguments",
			expectedErr: fakeError(dollyerr.CodeArgParserCheckValueAllowabilityError),
		},
		{
			caseName: "no_allowed_values",
			argDescription: apConf.ArgumentsDescriptionSrc{
				AmountType: apConf.ArgAmountTypeNoArgs,
			}.ToConstPtr(),
		},
		{
			caseName: "value_is_not_allowed",
			argDescription: apConf.ArgumentsDescriptionSrc{
				AmountType: apConf.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			}.ToConstPtr(),
			value:       gofakeit.Color(),
			expectedErr: fakeError(dollyerr.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "value_is_allowed",
			argDescription: apConf.ArgumentsDescriptionSrc{
				AmountType: apConf.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			}.ToConstPtr(),
			value: value,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := isValueAllowed(td.argDescription, td.value)

			if td.expectedErr == nil {
				require.Nil(t, err)
				return
			}

			require.Equal(t, td.expectedErr.Code(), err.Code())
		})
	}
}

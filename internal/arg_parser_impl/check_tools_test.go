package arg_parser_impl

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
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

		commandDescription *dollyconf.CommandDescription
		data               *parsed_data.ParsedData

		expectedErr *dollyerr.Error
	}{
		{
			caseName: "nil_arguments",
		},
		{
			caseName: "required flag is not set",
			commandDescription: &dollyconf.CommandDescription{
				RequiredFlags: map[dollyconf.Flag]bool{
					dollyconf.Flag(gofakeit.Color()): true,
				},
			},
			expectedErr: fakeError(dollyerr.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "command arg is not set",
			commandDescription: &dollyconf.CommandDescription{
				ArgDescription: &dollyconf.ArgumentsDescription{
					AmountType: dollyconf.ArgAmountTypeSingle,
				},
			},
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

		argDescription *dollyconf.ArgumentsDescription
		value          string

		expectedErr *dollyerr.Error
	}{
		{
			caseName:    "nil_arguments",
			expectedErr: fakeError(dollyerr.CodeArgParserCheckValueAllowabilityError),
		},
		{
			caseName: "no_allowed_values",
			argDescription: &dollyconf.ArgumentsDescription{
				AmountType: dollyconf.ArgAmountTypeNoArgs,
			},
		},
		{
			caseName: "value_is_not_allowed",
			argDescription: &dollyconf.ArgumentsDescription{
				AmountType: dollyconf.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			},
			value:       gofakeit.Color(),
			expectedErr: fakeError(dollyerr.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "value_is_allowed",
			argDescription: &dollyconf.ArgumentsDescription{
				AmountType: dollyconf.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			},
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

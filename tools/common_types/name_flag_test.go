package common_types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameFlagIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string
		nameFlag NameFlag
		expErr   error
	}{
		{
			caseName: "empty_name",
			expErr:   ErrNameFlagIsValidEmpty,
		},
		{
			caseName: "name_without_dash_in_front",
			nameFlag: "f",
			expErr:   ErrNameFlagIsValidNoDashInFront,
		},
		{
			caseName: "dash",
			nameFlag: "-",
			expErr:   ErrNameFlagIsValidTooShort,
		},
		{
			caseName: "name_with_unexpected_char",
			nameFlag: "-$",
			expErr:   ErrNameFlagIsValidUnexpectedChar,
		},
		{
			caseName: "double_dash",
			nameFlag: "--",
			expErr:   ErrNameFlagIsValidTooShort,
		},
		{
			caseName: "triple_dash",
			nameFlag: "---",
			expErr:   ErrNameFlagIsValidTooManyDashesInFront,
		},
		{
			caseName: "too_long_name_with_two_dashes_in_front",
			nameFlag: "--too-long-name",
			expErr:   ErrNameFlagIsValidTooLong,
		},
		{
			caseName: "too_long_name_with_one_dash_in_front",
			nameFlag: "-flag",
			expErr:   ErrNameFlagIsValidTooLong,
		},
		{
			caseName: "valid_short_flag_name",
			nameFlag: "-f",
		},
		{
			caseName: "valid_long_flag_name",
			nameFlag: "--flag",
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			err := tc.nameFlag.IsValid()
			if tc.expErr == nil {
				require.NoError(t, err)
				return
			}

			require.ErrorIs(t, err, tc.expErr)
		})
	}
}

func TestIsDoubleDash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName string
		nameFlag NameFlag
		exp      bool
	}{
		{
			caseName: "empty",
		},
		{
			caseName: "no_dash",
			nameFlag: "f",
		},
		{
			caseName: "dash",
			nameFlag: "-f",
		},
		{
			caseName: "two_dash",
			nameFlag: "--f",
			exp:      true,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, tc.nameFlag.IsDoubleDash())
		})
	}
}

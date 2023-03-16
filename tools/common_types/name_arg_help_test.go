package common_types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameArgHelpIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName       string
		nameArgHelpSrc string
		expErr         error
	}{
		{
			caseName:       "invalid_tooShort",
			nameArgHelpSrc: "b",
			expErr:         ErrNameArgHelpTooShort,
		},
		{
			caseName:       "valid",
			nameArgHelpSrc: "bubu",
		},
		{
			caseName:       "invalid_tooShort",
			nameArgHelpSrc: "bububu",
			expErr:         ErrNameArgHelpTooLong,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			nameArgHelp := NameArgHelp(tc.nameArgHelpSrc)
			require.Equal(t, tc.nameArgHelpSrc, nameArgHelp.String())

			require.ErrorIs(t, nameArgHelp.IsValid(), tc.expErr)
		})
	}
}

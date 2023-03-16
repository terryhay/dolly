package common_types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameCommandIsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseName          string
		nameCommand       NameCommand
		thisIsHelpCommand bool

		expErr error
	}{
		{
			caseName: "empty_name",
			expErr:   ErrNameCommandEmpty,
		},
		{
			caseName:    "too_short_name",
			nameCommand: RandNameCommandTooShort(),
			expErr:      ErrNameCommandTooShort,
		},
		{
			caseName:          "too_short_help_command",
			nameCommand:       RandNameCommandTooShort(),
			thisIsHelpCommand: true,
			expErr:            ErrNameCommandHelpCommandTooShort,
		},
		{
			caseName:    "too_long_name",
			nameCommand: RandNameCommandTooLong() + "qwerty",
			expErr:      ErrNameCommandTooLong,
		},
		{
			caseName:    "unexpected_char",
			nameCommand: "$" + RandNameCommandShort(),
			expErr:      ErrNameCommandUnexpectedChar,
		},
		{
			caseName:    "dash_in_front",
			nameCommand: "-" + RandNameCommandShort(),
			expErr:      ErrNameCommandDashInFront,
		},
		{
			caseName:    "valid_command_name_short",
			nameCommand: RandNameCommandShort(),
		},
		{
			caseName:    "valid_command_name_long",
			nameCommand: RandNameCommandLong(),
		},
		{
			caseName:          "exception_help_command__h",
			nameCommand:       NameCommandHelpHException,
			thisIsHelpCommand: true,
		},
		{
			caseName:          "exception_help_command_dash_h",
			nameCommand:       NameCommandHelpDashHException,
			thisIsHelpCommand: true,
		},
		{
			caseName:          "exception_help_command_dash_dash_help",
			nameCommand:       NameCommandHelpDashHelpException,
			thisIsHelpCommand: true,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			err := tc.nameCommand.IsValid(tc.thisIsHelpCommand)
			require.ErrorIs(t, err, tc.expErr)
		})
	}
}

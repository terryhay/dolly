package common_types

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/terryhay/dolly/tools/size"
)

// NameCommand - name type of command
type NameCommand string

const (
	// NameCommandUndefined - default NameCommand value, nameless command has this name
	NameCommandUndefined NameCommand = ""

	// NameCommandHelpHException - exception name for 'h' command name
	NameCommandHelpHException NameCommand = "h"

	// NameCommandHelpDashHException - exception name for '-h' command name
	NameCommandHelpDashHException NameCommand = "-h"

	// NameCommandHelpDashHelpException - exception name for '--help' command name
	NameCommandHelpDashHelpException NameCommand = "--help"
)

const (
	// LenNameCommandMin - min len of NameCommand
	LenNameCommandMin size.Width = 3

	// LenNameCommandMax - max len of NameCommand
	LenNameCommandMax size.Width = 12
)

// String implements Stringer interface
func (n NameCommand) String() string {
	return string(n)
}

var (
	// ErrNameCommandDashInFront - unexpected dash in front
	ErrNameCommandDashInFront = fmt.Errorf("NameCommand: must not have dash in front; exceptions are '%s' and '%s' for help command", NameCommandHelpDashHException, NameCommandHelpDashHelpException)

	// ErrNameCommandEmpty - empty string
	ErrNameCommandEmpty = errors.New("NameCommand: is empty")

	// ErrNameCommandHelpCommandTooShort - len is less than LenNameCommandMin
	ErrNameCommandHelpCommandTooShort = errors.New("NameCommand: help command is too short; exception '-h'")

	// ErrNameCommandTooLong - len is more than LenNameCommandMax
	ErrNameCommandTooLong = errors.New("NameCommand: is too long")

	// ErrNameCommandTooShort - len is less than LenNameCommandMin
	ErrNameCommandTooShort = errors.New("NameCommand: is too short")

	// ErrNameCommandUnexpectedChar - NameCommand has unexpected char
	ErrNameCommandUnexpectedChar = errors.New("NameCommand: unexpected character")
)

var patternBlackListCommandChar = regexp.MustCompile(`[^a-zA-Z-_]`)

// IsValid checks if NameCommand doesn't have dash in front, isn't too short/long
func (n NameCommand) IsValid(thisIsHelpCommand bool) error {
	const dash = "-"

	switch {
	case len(n) == 0:
		return ErrNameCommandEmpty

	case len(n) < LenNameCommandMin.Int():
		if thisIsHelpCommand {
			if n != NameCommandHelpDashHException && n != NameCommandHelpHException {
				return fmt.Errorf("%w: '%s' has len='%d'; permissible len inteval is [%d; %d]",
					ErrNameCommandHelpCommandTooShort, n, len(n), LenNameCommandMin, LenNameCommandMax)
			}
			return nil
		}

		return fmt.Errorf("%w: '%s' has len='%d'; permissible len inteval is [%d; %d]",
			ErrNameCommandTooShort, n, len(n), LenNameCommandMin, LenNameCommandMax)

	case len(n) > LenNameCommandMax.Int():
		return fmt.Errorf("%w: '%s' has len='%d'; permissible len inteval is [%d; %d]",
			ErrNameCommandTooLong, n, len(n), LenNameCommandMin, LenNameCommandMax)

	case len(patternBlackListCommandChar.FindString(n.String())) > 0:
		return fmt.Errorf("%w: '%s' has unexpected character '%s'",
			ErrNameCommandUnexpectedChar, n, patternBlackListFlagChar.FindString(n.String()))

	case string(n[0]) == dash:
		if thisIsHelpCommand && (n == NameCommandHelpDashHException || n == NameCommandHelpDashHelpException) {
			return nil
		}

		return fmt.Errorf("%w: '%s'", ErrNameCommandDashInFront, n)

	default:
		return nil
	}
}

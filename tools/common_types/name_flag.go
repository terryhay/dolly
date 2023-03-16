package common_types

import (
	"errors"
	"fmt"
	"regexp"
)

// NameFlag - name type of command flag
type NameFlag string

const (
	// NameFlagUndefined - default NameFlag value, don't use it
	NameFlagUndefined NameFlag = ""

	// LenNameFlagWithOneDashMax - max amount characters in NameFlag which is like "-f"
	LenNameFlagWithOneDashMax = 3

	// LenNameFlagWithTwoDashesMax - max amount characters in NameFlag which is like "--flag"
	LenNameFlagWithTwoDashesMax = 12
)

var (
	// ErrNameFlagIsValidEmpty - NameFlag is empty string
	ErrNameFlagIsValidEmpty = errors.New(`common_types: NameFlag.IsValid: is empty`)

	// ErrNameFlagIsValidNoDashInFront - NameFlag must have dash in front
	ErrNameFlagIsValidNoDashInFront = errors.New(`common_types: NameFlag.IsValid: flag must have dash in front`)

	// ErrNameFlagIsValidTooLong - NameFlag len is more than LenNameFlagWithTwoDashesMax
	ErrNameFlagIsValidTooLong = errors.New(`common_types: NameFlag.IsValid: too long`)

	// ErrNameFlagIsValidTooManyDashesInFront - NameFlag has three or more dashes in front
	ErrNameFlagIsValidTooManyDashesInFront = errors.New(`common_types: NameFlag.IsValid: flag must have only one ore two dashes in front`)

	// ErrNameFlagIsValidTooShort - NameFlag len is too short
	ErrNameFlagIsValidTooShort = errors.New(`common_types: NameFlag.IsValid: too short`)

	// ErrNameFlagIsValidUnexpectedChar - NameFlag has unexpected char
	ErrNameFlagIsValidUnexpectedChar = errors.New(`common_types: NameFlag.IsValid: unexpected character`)
)

var patternBlackListFlagChar = regexp.MustCompile(`[^a-zA-Z-0-9]`)

// String implements Stringer interface
func (n NameFlag) String() string {
	return string(n)
}

// IsValid checks if NameFlag has dash in front and is not too long
func (n NameFlag) IsValid() error {
	const dash = "-"

	switch {
	case len(n) == 0:
		return ErrNameFlagIsValidEmpty

	case n[:1] != dash:
		return fmt.Errorf(`%w: flag "%s"; example "-n"`, ErrNameFlagIsValidNoDashInFront, n)

	case len(n) == 1:
		// flag is "-"
		return fmt.Errorf(`%w: flag "%s"; example "-n"`, ErrNameFlagIsValidTooShort, n)

	case len(patternBlackListFlagChar.FindString(n.String())) > 0:
		return fmt.Errorf(`%w: flag "%s"; char "%s"`,
			ErrNameFlagIsValidUnexpectedChar, n, patternBlackListFlagChar.FindString(n.String()))

	case string(n[1]) == dash:
		// flag is like --flag
		if len(n) == 2 {
			return fmt.Errorf(`%w: flag "%s"; examples are "-n", "--flag"`,
				ErrNameFlagIsValidTooShort, n)
		}
		if string(n[2]) == dash {
			return fmt.Errorf(`%w: flag "%s"; examples are "-n", "--flag"`,
				ErrNameFlagIsValidTooManyDashesInFront, n)
		}
		if len(n) > LenNameFlagWithTwoDashesMax {
			return fmt.Errorf(`%w: flag "%s"; len "%d"; max len "%d"`,
				ErrNameFlagIsValidTooLong, n, len(n), LenNameFlagWithTwoDashesMax)
		}
		return nil

	default:
		// flag is like -n
		if len(n) > LenNameFlagWithOneDashMax {
			return fmt.Errorf(`%w: flag "%s"; len "%d"; max len "%d"; use two dash in front for long names like "--flag"`,
				ErrNameFlagIsValidTooLong, n, len(n), LenNameFlagWithOneDashMax)
		}
		return nil
	}
}

// IsDoubleDash returns if NameFlag has double dash prefix
func (n NameFlag) IsDoubleDash() bool {
	if len(n) < 2 {
		return false
	}

	const dash = '-'
	return n[0] == dash && n[1] == dash
}

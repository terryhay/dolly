package common_types

import (
	"errors"
	"fmt"
)

// NameArgHelp - name of argument for create help page
type NameArgHelp string

// String implements Stringer interface
func (n NameArgHelp) String() string {
	return string(n)
}

const (
	// LenNameArgHelpMin - min NameArgHelp length
	LenNameArgHelpMin = 2

	// LenNameArgHelpMax - max NameArgHelp length
	LenNameArgHelpMax = 5
)

var (
	// ErrNameArgHelpTooShort - NameArgHelp is too short
	ErrNameArgHelpTooShort = errors.New(`NameArgHelp: too short`)

	// ErrNameArgHelpTooLong - NameArgHelp is too long
	ErrNameArgHelpTooLong = errors.New(`NameArgHelp: too long`)
)

// IsValid checks if NameArgHelp is valid
func (n NameArgHelp) IsValid() error {
	switch {
	case len(n) < LenNameArgHelpMin:
		return fmt.Errorf(`%w: NameArgHelp '%s' has len=%d; permissible len inteval is [%d; %d]`,
			ErrNameArgHelpTooShort, n, len(n), LenNameArgHelpMin, LenNameArgHelpMax)

	case len(n) > LenNameArgHelpMax:
		return fmt.Errorf(`%w: NameArgHelp '%s' has len=%d; permissible len inteval is [%d; %d]`,
			ErrNameArgHelpTooLong, n, len(n), LenNameArgHelpMin, LenNameArgHelpMax)

	default:
		return nil
	}
}

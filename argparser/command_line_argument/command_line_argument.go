package command_line_argument

import coty "github.com/terryhay/dolly/tools/common_types"

// Argument is a command line argument
type Argument string

// ArgumentEmpty - empty value of Argument
const ArgumentEmpty Argument = ""

// String implements Stringer interface
func (a Argument) String() string {
	return string(a)
}

// ToNameCommand converts Argument to CommandName type
func (a Argument) ToNameCommand() coty.NameCommand {
	return coty.NameCommand(a)
}

// ToNameFlag converts Argument to FlagName type
func (a Argument) ToNameFlag() coty.NameFlag {
	return coty.NameFlag(a)
}

// IsValid returns if Argument is not empty string
func (a Argument) IsValid() bool {
	return len(a) > 0
}

// IsFlag checks if Argument has a dash in front (like a flag)
func (a Argument) IsFlag() bool {
	if len(a) == 0 {
		return false
	}
	const dash = "-"
	return a[:1] == dash
}

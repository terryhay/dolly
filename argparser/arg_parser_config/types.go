package arg_parser_config

type (
	// CommandID - id of command which contains some amount of flags and arguments (0 value is undefined value)
	CommandID uint

	// Command is name type of command
	Command string

	// FlagID - id of flag (0 value is undefined value)
	FlagID uint

	// Flag is name type of command flag
	Flag string

	// ArgAmountType defines an amount of arguments
	ArgAmountType uint8
)

const (
	// CommandIDUndefined - a default return CommandID value, don't use it
	CommandIDUndefined CommandID = 0

	// CommandUndefined - a default Command value, don't use it
	CommandUndefined Command = ""
)

const (
	// FlagIDUndefined - default FlagID value, don't use it
	FlagIDUndefined FlagID = 0

	// FlagUndefined - a default Flag value, don't use it
	FlagUndefined Flag = ""
)

const (
	// ArgAmountTypeNoArgs - default value, no arguments
	ArgAmountTypeNoArgs ArgAmountType = iota

	// ArgAmountTypeSingle - single value is expected
	ArgAmountTypeSingle

	// ArgAmountTypeList - value array is expected
	ArgAmountTypeList
)

func (c Command) ToString() string {
	return string(c)
}

func (f Flag) ToString() string {
	return string(f)
}

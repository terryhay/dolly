package dollyconf

// CommandID is id of command which contains some amount of flags and arguments (0 value is reserved)
type CommandID uint

// CommandIDUndefined - a default return CommandID value, don't use it
const CommandIDUndefined CommandID = 0

// Command is name type of command
type Command string

// CommandUndefined - a default Command value, don't use it
const CommandUndefined Command = ""

// Flag is name type of command flag
type Flag string

// FlagUndefined - a default Flag value, don't use it
const FlagUndefined Flag = ""

// ArgAmountType defines an amount of arguments
type ArgAmountType uint8

const (
	// ArgAmountTypeNoArgs - default value, no arguments
	ArgAmountTypeNoArgs ArgAmountType = iota

	// ArgAmountTypeSingle - single value is expected
	ArgAmountTypeSingle

	// ArgAmountTypeList - value array is expected
	ArgAmountTypeList
)

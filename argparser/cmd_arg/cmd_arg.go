package cmd_arg

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	parsedData "github.com/terryhay/dolly/argparser/parsed_data"
)

// CmdArg is a command line argument
type CmdArg string

// CmdArgEmpty is empty value of CmdArg
const CmdArgEmpty CmdArg = ""

// ToString converts CmdArg to string
func (a CmdArg) ToString() string {
	return string(a)
}

// ToCommand converts CmdArg to Command type
func (a CmdArg) ToCommand() apConf.Command {
	return apConf.Command(a)
}

// ToFlag converts CmdArg to Flag type
func (a CmdArg) ToFlag() apConf.Flag {
	return apConf.Flag(a)
}

// ToArgValue converts CmdArg to ArgValue
func (a CmdArg) ToArgValue() parsedData.ArgValue {
	return parsedData.ArgValue(a)
}

// IsEmpty returns if CmdArg object is empty string
func (a CmdArg) IsEmpty() bool {
	return len(a) == 0
}

// IsValid returns if CmdArg is not empty string
func (a CmdArg) IsValid() bool {
	return len(a) > 0
}

func (a CmdArg) HasDashInFront() bool {
	if len(a) == 0 {
		return true
	}
	const dash = "-"
	return a[:1] == dash
}

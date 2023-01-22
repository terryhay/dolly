package cmd_arg

import "github.com/terryhay/dolly/utils/index"

// CmdArgIterator provides an iterating by command line arguments
type CmdArgIterator struct {
	args  []string
	index index.Index
}

// MakeCmdArgIterator constructs CmdArgIterator object in a stack
func MakeCmdArgIterator(args []string) CmdArgIterator {
	return CmdArgIterator{
		args: args,
	}
}

// IsEnded returns if CmdArgIterator doesn't have any more arguments
func (it *CmdArgIterator) IsEnded() bool {
	if it == nil {
		return true
	}
	return it.index.ToInt() >= len(it.args)
}

// Next increments an index and return next command line argument
func (it *CmdArgIterator) Next() CmdArg {
	if it == nil {
		return CmdArgEmpty
	}

	it.index = index.Append(it.index, 1)
	if it.index.ToInt() >= len(it.args) {
		return CmdArgEmpty
	}
	return CmdArg(it.args[it.index.ToInt()])
}

// GetCmdArg returns a current command line argument
func (it *CmdArgIterator) GetCmdArg() CmdArg {
	if it == nil || it.index.ToInt() >= len(it.args) {
		return CmdArgEmpty
	}
	return CmdArg(it.args[it.index.ToInt()])
}

package command_line_argument

import (
	"github.com/terryhay/dolly/tools/index"
)

// Iterator provides an iterating by command line arguments
type Iterator struct {
	args  []string
	index index.Index
}

// MakeIterator constructs Iterator object in a stack
func MakeIterator(args []string) Iterator {
	return Iterator{
		args: args,
	}
}

// IsEnded returns if Iterator doesn't have any more arguments
func (it Iterator) IsEnded() bool {
	return it.index.Int() >= len(it.args)
}

// Next increments an index and return next command line argument
func (it Iterator) Next() Iterator {
	it.index = index.Inc(it.index)
	return it
}

// GetArg returns a current command line argument
func (it Iterator) GetArg() Argument {
	if it.index.Int() >= len(it.args) {
		return ArgumentEmpty
	}
	return Argument(it.args[it.index.Int()])
}

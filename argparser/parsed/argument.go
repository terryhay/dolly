package parsed

import "unsafe"

// Argument contains parsed argument values of a nameMainCommand or a flagName
type Argument struct {
	argValues []ArgValue
}

// ArgumentOpt contains source data for cast to Argument
type ArgumentOpt struct {
	ArgValues []ArgValue
}

// MakeArgument converts opt to Result pointer
func MakeArgument(opt *ArgumentOpt) *Argument {
	if opt == nil {
		return nil
	}

	return (*Argument)(unsafe.Pointer(opt))
}

// GetArgValues gets ArgValues field
func (i *Argument) GetArgValues() []ArgValue {
	if i == nil {
		return nil
	}
	return i.argValues
}

package arg_parser_config

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"

	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
)

// Argument contains specification of flag arguments
type Argument struct {
	isList           bool
	defaultValues    []string
	allowedValues    map[string]struct{}
	isOptional       bool
	descSynopsisHelp string
}

// ArgumentOpt contains source data for cast to Argument
type ArgumentOpt struct {
	IsList           bool
	DefaultValues    []string
	AllowedValues    map[string]struct{}
	IsOptional       bool
	DescSynopsisHelp string
}

// MakeArgument converts opt to Argument pointer
func MakeArgument(opt *ArgumentOpt) *Argument {
	if opt == nil {
		return nil
	}
	return (*Argument)(unsafe.Pointer(opt))
}

// GetIsList gets IsList field
func (a *Argument) GetIsList() bool {
	if a == nil {
		return false
	}
	return a.isList
}

// GetDefaultValues gets DefaultValues field
func (a *Argument) GetDefaultValues() []string {
	if a == nil {
		return nil
	}
	return a.defaultValues
}

// GetAllowedValues gets AllowedValues field
func (a *Argument) GetAllowedValues() map[string]struct{} {
	if a == nil {
		return nil
	}
	return a.allowedValues
}

// GetIsOptional gets isOptional field
func (a *Argument) GetIsOptional() bool {
	if a == nil {
		return false
	}
	return a.isOptional
}

// GetSynopsisHelpDescription gets DescSynopsisHelp field
func (a *Argument) GetSynopsisHelpDescription() string {
	if a == nil {
		return ""
	}
	return a.descSynopsisHelp
}

// ErrIsArgAllowed -  argument is not allowed
var ErrIsArgAllowed = errors.New(`Argument.IsArgAllowed: argument is not allowed`)

// IsArgAllowed checks if allowedValues field contains the argument
func (a *Argument) IsArgAllowed(arg clArg.Argument) error {
	if a == nil {
		return nil
	}

	if len(a.allowedValues) > 0 {
		if _, contain := a.allowedValues[arg.String()]; !contain {
			return fmt.Errorf(`%w: arg "%s"; %s`,
				ErrIsArgAllowed, arg.String(), a.CreateStringWithArgInfo())
		}
	}
	return nil
}

// IsRequired returns if Argument is not optional
func (a *Argument) IsRequired() bool {
	if a == nil {
		return false
	}
	return !a.isOptional
}

// CreateStringWithArgInfo  creates string info about default and allowed arguments
func (a *Argument) CreateStringWithArgInfo() string {
	if a == nil {
		return ""
	}

	const (
		patternSingleArgument  = "single argument"
		patternListOfArguments = "list of arguments"
	)

	res := patternSingleArgument
	if a.isList {
		res = patternListOfArguments
	}

	if len(a.allowedValues) > 0 {
		values := make([]string, 0, len(a.allowedValues))
		for v := range a.allowedValues {
			values = append(values, v)
		}

		const patternAllowedValues = " allowed [%s]"
		res += fmt.Sprintf(patternAllowedValues, strings.Join(values, ", "))
	}
	if len(a.defaultValues) > 0 {
		values := make([]string, 0, len(a.defaultValues))
		for v := range a.allowedValues {
			values = append(values, v)
		}

		const patternDefaultValues = " default [%s]"
		res += fmt.Sprintf(patternDefaultValues, strings.Join(values, ", "))
	}

	return res
}

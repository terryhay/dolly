package config_yaml

import (
	"errors"
	"fmt"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Placeholder contains description of flag/argument placeholder
type Placeholder struct {
	name coty.NamePlaceholder

	// one or two fields must be set
	flags    []*Flag
	argument *Argument

	// optional
	isFlagOptional bool
}

// GetName gets name field
func (p *Placeholder) GetName() coty.NamePlaceholder {
	if p == nil {
		return coty.NamePlaceholderUndefined
	}
	return p.name
}

// GetIsFlagOptional gets isOptional isFlagOptional
func (p *Placeholder) GetIsFlagOptional() bool {
	if p == nil {
		return false
	}
	return p.isFlagOptional
}

// GetFlags gets flags field
func (p *Placeholder) GetFlags() []*Flag {
	if p == nil {
		return nil
	}
	return p.flags
}

// GetArgument gets argument field
func (p *Placeholder) GetArgument() *Argument {
	if p == nil {
		return nil
	}
	return p.argument
}

var (
	// ErrPlaceholderName - no 'name' filed in source yaml file
	ErrPlaceholderName = errors.New(`Placeholder.IsValid: 'arg_parser.placeholders' element must have 'name'`)

	// ErrPlaceholderNoFlagsNoArg - no 'flags' and 'argument' in source yaml file (at least one of them must be set)
	ErrPlaceholderNoFlagsNoArg = errors.New(`Placeholder.IsValid: 'arg_parser.placeholders' element must have one of fields 'flags' or 'argument'`)

	// ErrPlaceholderIsValid - some inside component of Placeholder is not valid
	ErrPlaceholderIsValid = errors.New(`Placeholder.IsValid`)
)

// IsValid checks if Placeholder is valid
func (p *Placeholder) IsValid() error {
	if len(p.GetName()) == 0 {
		return ErrPlaceholderName
	}

	if len(p.GetFlags()) == 0 && p.GetArgument() == nil {
		return fmt.Errorf(`%w: placeholder name '%s'`, ErrPlaceholderNoFlagsNoArg, p.GetName())
	}
	for _, flag := range p.GetFlags() {
		if err := flag.IsValid(); err != nil {
			return errors.Join(fmt.Errorf(`%w: placeholder name '%s'`, ErrPlaceholderIsValid, p.GetName()), err)
		}
	}

	if err := p.GetArgument().IsValid(); err != nil {
		return errors.Join(fmt.Errorf(`%w: placeholder name '%s'`, ErrPlaceholderIsValid, p.GetName()), err)
	}

	return nil
}

// PlaceholderOpt contains source data with flag and argument descriptions for cast to Placeholder
type PlaceholderOpt struct {
	Name string `yaml:"name"`

	// one or two fields must be set
	Flags    []*FlagOpt   `yaml:"flags"`
	Argument *ArgumentOpt `yaml:"argument"`

	IsFlagOptional bool `yaml:"is_flag_optional"`
}

// NewPlaceholder converts opt to Placeholder pointer
func NewPlaceholder(opt *PlaceholderOpt) *Placeholder {
	if opt == nil {
		return nil
	}

	return &Placeholder{
		name: coty.NamePlaceholder(opt.Name),

		flags:    toFlagSlice(opt.Flags),
		argument: NewArgument(opt.Argument),

		isFlagOptional: opt.IsFlagOptional,
	}
}

func toFlagSlice(from []*FlagOpt) []*Flag {
	if len(from) == 0 {
		return nil
	}

	to := make([]*Flag, 0, len(from))
	for _, opt := range from {
		to = append(to, NewFlag(opt))
	}

	return to
}

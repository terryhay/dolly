package arg_parser_config

import (
	"sort"
	"strings"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Placeholder contains specification of argument placeholder
// which contains namesAdditional and arguments
type Placeholder struct {
	id             coty.IDPlaceholder
	isFlagOptional bool
	flagsByNames   map[coty.NameFlag]*Flag
	argument       *Argument
}

// PlaceholderOpt contains source data with flag and argument descriptions for cast to Placeholder
type PlaceholderOpt struct {
	ID             coty.IDPlaceholder
	IsFlagOptional bool
	FlagsByNames   map[coty.NameFlag]*FlagOpt
	Argument       *ArgumentOpt
}

// NewPlaceholder converts opt to Placeholder pointer
func NewPlaceholder(opt PlaceholderOpt) *Placeholder {
	return &Placeholder{
		id:             opt.ID,
		isFlagOptional: opt.IsFlagOptional,
		flagsByNames:   createFlags(opt.FlagsByNames),
		argument:       MakeArgument(opt.Argument),
	}
}

// GetID gets id field
func (p *Placeholder) GetID() coty.IDPlaceholder {
	if p == nil {
		return coty.ArgPlaceholderIDUndefined
	}
	return p.id
}

// GetIsFlagOptional gets isOptional isFlagOptional
func (p *Placeholder) GetIsFlagOptional() bool {
	if p == nil {
		return false
	}
	return p.isFlagOptional
}

// GetDescriptionFlags gets flagsByNames field
func (p *Placeholder) GetDescriptionFlags() map[coty.NameFlag]*Flag {
	if p == nil {
		return nil
	}
	return p.flagsByNames
}

// GetArgument gets argument field
func (p *Placeholder) GetArgument() *Argument {
	if p == nil {
		return nil
	}
	return p.argument
}

// IsFlagRequired returns if Placeholder is required
func (p *Placeholder) IsFlagRequired() bool {
	if p == nil {
		return false
	}
	return len(p.flagsByNames) > 0 && !p.isFlagOptional
}

// IsArgRequired returns if Argument is required
func (p *Placeholder) IsArgRequired() bool {
	if p == nil {
		return false
	}
	return p.argument.IsRequired()
}

// HasArg returns if Placeholder has not empty argument field data
func (p *Placeholder) HasArg() bool {
	if p == nil {
		return false
	}
	return p.argument != nil
}

// HasFlags returns if Placeholder has not empty flagsByNames field data
func (p *Placeholder) HasFlags() bool {
	if p == nil {
		return false
	}
	return len(p.flagsByNames) > 0
}

// FlagByName returns Flag by FlagName
func (p *Placeholder) FlagByName(name coty.NameFlag) *Flag {
	if p == nil {
		return nil
	}
	return p.flagsByNames[name]
}

// HasFlagName returns if Placeholder contains set NameFlag
func (p *Placeholder) HasFlagName(name coty.NameFlag) bool {
	if p == nil {
		return false
	}
	_, contain := p.flagsByNames[name]
	return contain
}

// CreateStringWithFlagNames creates sequence of flag names as a string
func (p *Placeholder) CreateStringWithFlagNames() string {
	if p == nil || len(p.flagsByNames) == 0 {
		return ""
	}

	flags := make([]string, 0, len(p.flagsByNames))
	for flag := range p.flagsByNames {
		flags = append(flags, flag.String())
	}
	sort.Strings(flags)

	return strings.Join(flags, ", ")
}

func createFlags(opt map[coty.NameFlag]*FlagOpt) map[coty.NameFlag]*Flag {
	if len(opt) == 0 {
		return nil
	}

	flagByNames := make(map[coty.NameFlag]*Flag, len(opt))
	for name, optFlag := range opt {
		flagByNames[name] = MakeFlag(*optFlag)
	}

	return flagByNames
}

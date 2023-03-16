package parsed

import (
	coty "github.com/terryhay/dolly/tools/common_types"
)

// Placeholder contains parsed flagName and arguments array
type Placeholder struct {
	id       coty.IDPlaceholder
	nameFlag coty.NameFlag
	argument *Argument
}

// PlaceholderOpt contains source data for cast to Placeholder
type PlaceholderOpt struct {
	ID       coty.IDPlaceholder
	Flag     coty.NameFlag
	Argument *ArgumentOpt
}

// NewPlaceholder converts opt to Placeholder pointer
func NewPlaceholder(opt *PlaceholderOpt) *Placeholder {
	if opt == nil {
		return nil
	}

	return &Placeholder{
		id:       opt.ID,
		nameFlag: opt.Flag,
		argument: MakeArgument(opt.Argument),
	}
}

// GetID gets id field
func (p *Placeholder) GetID() coty.IDPlaceholder {
	if p == nil {
		return coty.ArgPlaceholderIDUndefined
	}
	return p.id
}

// GetNameFlag gets flagName field
func (p *Placeholder) GetNameFlag() coty.NameFlag {
	if p == nil {
		return coty.NameFlagUndefined
	}
	return p.nameFlag
}

// GetArgData gets Argument field
func (p *Placeholder) GetArgData() *Argument {
	if p == nil {
		return nil
	}
	return p.argument
}

// HasArg returns if Placeholder has Argument option
func (p *Placeholder) HasArg() bool {
	if p == nil {
		return false
	}
	return p.argument != nil
}

// HasFlag returns if Placeholder has Flag option
func (p *Placeholder) HasFlag() bool {
	if p == nil {
		return false
	}
	return len(p.nameFlag) > 0
}

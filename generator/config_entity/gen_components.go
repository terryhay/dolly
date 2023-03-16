package config_entity

import (
	"fmt"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Prefix - type of prefix for constants generating
type Prefix string

const (
	// PrefixNameCommand - for command constants naming
	PrefixNameCommand Prefix = "NameCommand"

	// PrefixPlaceholderID - for placeholder IDs
	PrefixPlaceholderID Prefix = "IDPlaceholder"

	// PrefixFlagName - for flag constants naming
	PrefixFlagName Prefix = "NameFlag"

	// NamelessCommandIDPostfix - for nameless command naming
	NamelessCommandIDPostfix = coty.NameCommand("Nameless")
)

// String implements Stringer interface
func (p Prefix) String() string {
	return string(p)
}

// Name - abstract name
type Name string

// String implements Stringer interface
func (n Name) String() string {
	return string(n)
}

// NameCommand converts Name to NameCommand
func (n Name) NameCommand() coty.NameCommand {
	return coty.NameCommand(n)
}

// NamePlaceholder converts Name to NamePlaceholder
func (n Name) NamePlaceholder() coty.NamePlaceholder {
	return coty.NamePlaceholder(n)
}

// NameFlag converts Name to NameFlag
func (n Name) NameFlag() coty.NameFlag {
	return coty.NameFlag(n)
}

// GenComponentsOpt contain options for construct GenComponents object
type GenComponentsOpt struct {
	PrefixID Prefix
	Name     fmt.Stringer
	Comment  fmt.Stringer
}

// NewGenComponents construct GenComponents object
func NewGenComponents(opt GenComponentsOpt) *GenComponents {
	return &GenComponents{
		nameID:  createID(opt.PrefixID, opt.Name),
		name:    Name(opt.Name.String()),
		comment: opt.Comment.String(),
	}
}

// GenComponents contains data for generate help information
type GenComponents struct {
	nameID  string
	name    Name
	comment string
}

// GetNameID gets nameID field
func (gc *GenComponents) GetNameID() string {
	if gc == nil {
		return ""
	}
	return gc.nameID
}

// GetName gets name field
func (gc *GenComponents) GetName() Name {
	if gc == nil {
		return ""
	}
	return gc.name
}

// GetComment gets comment field
func (gc *GenComponents) GetComment() string {
	if gc == nil {
		return ""
	}
	return gc.comment
}

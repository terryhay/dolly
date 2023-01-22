package arg_parser_config

import "unsafe"

// FlagDescription contains specification of flag with flag's arguments
type FlagDescription struct {
	id                  FlagID
	flags               []Flag
	descriptionHelpInfo string
	argDescription      *ArgumentsDescription
}

// FlagDescriptionSrc contains source data for cast to FlagDescription
type FlagDescriptionSrc struct {
	ID                  FlagID
	Flags               []Flag
	DescriptionHelpInfo string
	ArgDescription      *ArgumentsDescription
}

// ToConstPtr converts src to FlagDescription pointer
func (src FlagDescriptionSrc) ToConstPtr() *FlagDescription {
	return (*FlagDescription)(unsafe.Pointer(&src))
}

// GetID gets DescriptionHelpInfo field
func (fd *FlagDescription) GetID() FlagID {
	if fd == nil {
		return FlagIDUndefined
	}
	return fd.id
}

// GetFlags gets Flags field
func (fd *FlagDescription) GetFlags() []Flag {
	if fd == nil {
		return nil
	}
	return fd.flags
}

// GetDescriptionHelpInfo gets DescriptionHelpInfo field
func (fd *FlagDescription) GetDescriptionHelpInfo() string {
	if fd == nil {
		return ""
	}
	return fd.descriptionHelpInfo
}

// GetArgDescription gets ArgDescription field
func (fd *FlagDescription) GetArgDescription() *ArgumentsDescription {
	if fd == nil {
		return nil
	}
	return fd.argDescription
}

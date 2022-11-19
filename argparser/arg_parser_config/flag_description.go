package arg_parser_config

import "unsafe"

// FlagDescription contains specification of flag with flag's arguments
type FlagDescription struct {
	descriptionHelpInfo string
	argDescription      *ArgumentsDescription
}

// FlagDescriptionSrc contains source data for cast to FlagDescription
type FlagDescriptionSrc struct {
	DescriptionHelpInfo string
	ArgDescription      *ArgumentsDescription
}

// ToConstPtr converts src to FlagDescription pointer
func (src FlagDescriptionSrc) ToConstPtr() *FlagDescription {
	return (*FlagDescription)(unsafe.Pointer(&src))
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *FlagDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.descriptionHelpInfo
}

// GetArgDescription - ArgDescription field getter
func (i *FlagDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.argDescription
}

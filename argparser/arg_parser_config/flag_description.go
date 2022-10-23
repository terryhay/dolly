package arg_parser_config

// FlagDescription contains specification of flag with flag's arguments
type FlagDescription struct {
	DescriptionHelpInfo string
	ArgDescription      *ArgumentsDescription
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *FlagDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetArgDescription - ArgDescription field getter
func (i *FlagDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgDescription
}

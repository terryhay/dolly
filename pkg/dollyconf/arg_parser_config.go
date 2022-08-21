package dollyconf

// ArgParserConfig contains specifications of flags, arguments and command groups of application
type ArgParserConfig struct {
	AppDescription             ApplicationDescription
	FlagDescriptions           map[Flag]*FlagDescription
	CommandDescriptions        []*CommandDescription
	HelpCommandDescription     HelpCommandDescription
	NamelessCommandDescription NamelessCommandDescription
}

// GetAppDescription - AppDescription field getter
func (i ArgParserConfig) GetAppDescription() ApplicationDescription {
	return i.AppDescription
}

// GetCommandDescriptions - CommandDescriptions field getter
func (i ArgParserConfig) GetCommandDescriptions() []*CommandDescription {
	return i.CommandDescriptions
}

// GetFlagDescriptions - FlagDescriptions field getter
func (i ArgParserConfig) GetFlagDescriptions() map[Flag]*FlagDescription {
	return i.FlagDescriptions
}

// GetHelpCommandDescription i.HelpCommandDescription
func (i ArgParserConfig) GetHelpCommandDescription() HelpCommandDescription {
	return i.HelpCommandDescription
}

// GetNamelessCommandDescription - NamelessCommandDescription field getter
func (i ArgParserConfig) GetNamelessCommandDescription() NamelessCommandDescription {
	return i.NamelessCommandDescription
}

package arg_parser_config

// ArgParserConfig contains specifications of flags, arguments and command groups of application
type ArgParserConfig struct {
	AppDescription             ApplicationDescription
	FlagDescriptions           map[Flag]*FlagDescription
	CommandDescriptions        []*CommandDescription
	HelpCommandDescription     HelpCommandDescription
	NamelessCommandDescription NamelessCommandDescription
}

// NewArgParserConfig - ArgParserConfig constructor
func NewArgParserConfig(
	appDescription ApplicationDescription,
	flagDescriptions map[Flag]*FlagDescription,
	commandDescriptions []*CommandDescription,
	helpCommandDescription HelpCommandDescription,
	namelessCommandDescription NamelessCommandDescription,
) ArgParserConfig {

	return ArgParserConfig{
		AppDescription:             appDescription,
		FlagDescriptions:           flagDescriptions,
		CommandDescriptions:        commandDescriptions,
		HelpCommandDescription:     helpCommandDescription,
		NamelessCommandDescription: namelessCommandDescription,
	}
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

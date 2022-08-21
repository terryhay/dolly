package dollyconf

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

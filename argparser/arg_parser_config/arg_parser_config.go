package arg_parser_config

import "unsafe"

// ArgParserConfig contains specifications of flags, arguments and command groups of application
type ArgParserConfig struct {
	appDescription             ApplicationDescription
	flagDescriptions           map[Flag]*FlagDescription
	commandDescriptions        []*CommandDescription
	helpCommandDescription     HelpCommandDescription
	namelessCommandDescription NamelessCommandDescription
}

// ArgParserConfigSrc contains source data for cast to ArgParserConfig
type ArgParserConfigSrc struct {
	AppDescription             ApplicationDescription
	FlagDescriptions           map[Flag]*FlagDescription
	CommandDescriptions        []*CommandDescription
	HelpCommandDescription     HelpCommandDescription
	NamelessCommandDescription NamelessCommandDescription
}

// Cast converts src to ArgParserConfig object
func (src ArgParserConfigSrc) Cast() ArgParserConfig {
	return *(*ArgParserConfig)(unsafe.Pointer(&src))
}

// GetAppDescription - appDescription field getter
func (i ArgParserConfig) GetAppDescription() ApplicationDescription {
	return i.appDescription
}

// GetCommandDescriptions - commandDescriptions field getter
func (i ArgParserConfig) GetCommandDescriptions() []*CommandDescription {
	return i.commandDescriptions
}

// GetFlagDescriptions - flagDescriptions field getter
func (i ArgParserConfig) GetFlagDescriptions() map[Flag]*FlagDescription {
	return i.flagDescriptions
}

// GetHelpCommandDescription i.helpCommandDescription
func (i ArgParserConfig) GetHelpCommandDescription() HelpCommandDescription {
	return i.helpCommandDescription
}

// GetNamelessCommandDescription - NamelessCommandDescription field getter
func (i ArgParserConfig) GetNamelessCommandDescription() NamelessCommandDescription {
	return i.namelessCommandDescription
}

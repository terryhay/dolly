package arg_parser_config

import "unsafe"

// ArgParserConfig contains specifications of flags, arguments and command groups of application
type ArgParserConfig struct {
	appDescription             ApplicationDescription
	flagDescriptionSlice       []*FlagDescription
	commandDescriptions        []*CommandDescription
	helpCommandDescription     HelpCommandDescription
	namelessCommandDescription NamelessCommandDescription
}

// ArgParserConfigSrc contains source data for cast to ArgParserConfig
type ArgParserConfigSrc struct {
	AppDescription             ApplicationDescription
	FlagDescriptionSlice       []*FlagDescription
	CommandDescriptions        []*CommandDescription
	HelpCommandDescription     HelpCommandDescription
	NamelessCommandDescription NamelessCommandDescription
}

// ToConst converts src to ArgParserConfig object
func (src ArgParserConfigSrc) ToConst() ArgParserConfig {
	return *(*ArgParserConfig)(unsafe.Pointer(&src))
}

// GetAppDescription gets appDescription field
func (apc *ArgParserConfig) GetAppDescription() ApplicationDescription {
	if apc == nil {
		return ApplicationDescription{}
	}
	return apc.appDescription
}

// GetCommandDescriptions gets commandDescriptions field
func (apc *ArgParserConfig) GetCommandDescriptions() []*CommandDescription {
	if apc == nil {
		return nil
	}
	return apc.commandDescriptions
}

// GetFlagDescriptionSlice gets flagDescriptionSlice field
func (apc *ArgParserConfig) GetFlagDescriptionSlice() []*FlagDescription {
	if apc == nil {
		return nil
	}
	return apc.flagDescriptionSlice
}

// ExtractFlagDescriptionMap creates and returns flagDescriptionSlice field
func (apc *ArgParserConfig) ExtractFlagDescriptionMap() map[Flag]*FlagDescription {
	if apc == nil || len(apc.flagDescriptionSlice) == 0 {
		return nil
	}

	flagDescriptionMap := make(map[Flag]*FlagDescription, len(apc.flagDescriptionSlice)*2)
	for _, description := range apc.flagDescriptionSlice {
		for _, flag := range description.GetFlags() {
			flagDescriptionMap[flag] = description
		}
	}

	return flagDescriptionMap
}

// GetHelpCommandDescription gets helpCommandDescription field
func (apc *ArgParserConfig) GetHelpCommandDescription() HelpCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.helpCommandDescription
}

// GetNamelessCommandDescription gets namelessCommandDescription field
func (apc *ArgParserConfig) GetNamelessCommandDescription() NamelessCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.namelessCommandDescription
}

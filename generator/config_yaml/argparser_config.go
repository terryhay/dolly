package config_yaml

import "fmt"

// ArgParserConfig contains arg parser configuration
type ArgParserConfig struct {
	AppHelpDescription     *AppHelpDescription
	HelpCommandDescription *HelpCommandDescription

	// one or more of these field must be set
	NamelessCommandDescription *NamelessCommandDescription
	CommandDescriptions        []*CommandDescription

	// optional
	FlagDescriptions []*FlagDescription
}

// GetAppHelpDescription - AppHelpDescription field getter
func (apc *ArgParserConfig) GetAppHelpDescription() *AppHelpDescription {
	if apc == nil {
		return nil
	}
	return apc.AppHelpDescription
}

// GetHelpCommandDescription - HelpCommandDescription field getter
func (apc *ArgParserConfig) GetHelpCommandDescription() *HelpCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.HelpCommandDescription
}

// GetNamelessCommandDescription - NamelessCommandDescription field getter
func (apc *ArgParserConfig) GetNamelessCommandDescription() *NamelessCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.NamelessCommandDescription
}

// GetCommandDescriptions - CommandDescriptions field getter
func (apc *ArgParserConfig) GetCommandDescriptions() []*CommandDescription {
	if apc == nil {
		return nil
	}
	return apc.CommandDescriptions
}

// GetFlagDescriptions - FlagDescriptions field getter
func (apc *ArgParserConfig) GetFlagDescriptions() []*FlagDescription {
	if apc == nil {
		return nil
	}
	return apc.FlagDescriptions
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (apc *ArgParserConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = apc

	proxy := struct {
		AppHelpDescription     *AppHelpDescription     `yaml:"app_help_description"`
		HelpCommandDescription *HelpCommandDescription `yaml:"help_command_description"`

		// one or more of these field must be set
		NamelessCommandDescription *NamelessCommandDescription `yaml:"nameless_command_description"`
		CommandDescriptions        []*CommandDescription       `yaml:"command_descriptions"`

		// optional
		FlagDescriptions []*FlagDescription `yaml:"flag_descriptions"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if proxy.AppHelpDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "app_help_description"`)
	}
	apc.AppHelpDescription = proxy.AppHelpDescription

	if proxy.HelpCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "help_command_description"`)
	}
	apc.HelpCommandDescription = proxy.HelpCommandDescription

	if len(proxy.CommandDescriptions) == 0 && proxy.NamelessCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: one or more of fields "nameless_command_description" or "command_descriptions" must be set`)
	}

	// don't check optional fields
	apc.CommandDescriptions = proxy.CommandDescriptions
	apc.FlagDescriptions = proxy.FlagDescriptions
	apc.NamelessCommandDescription = proxy.NamelessCommandDescription

	return nil
}

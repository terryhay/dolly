package config_yaml

import (
	"fmt"
	"unsafe"
)

// ArgParserConfig - arg parser configuration
type ArgParserConfig struct {
	appHelpDescription     *AppHelpDescription
	helpCommandDescription *HelpCommandDescription

	// one or more of these field must be set
	namelessCommandDescription *NamelessCommandDescription
	commandDescriptions        []*CommandDescription

	// optional
	flagDescriptions []*FlagDescription
}

// GetAppHelpDescription - AppHelpDescription field getter
func (apc *ArgParserConfig) GetAppHelpDescription() *AppHelpDescription {
	if apc == nil {
		return nil
	}
	return apc.appHelpDescription
}

// GetHelpCommandDescription - HelpCommandDescription field getter
func (apc *ArgParserConfig) GetHelpCommandDescription() *HelpCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.helpCommandDescription
}

// GetNamelessCommandDescription - NamelessCommandDescription field getter
func (apc *ArgParserConfig) GetNamelessCommandDescription() *NamelessCommandDescription {
	if apc == nil {
		return nil
	}
	return apc.namelessCommandDescription
}

// GetCommandDescriptions - commandDescriptions field getter
func (apc *ArgParserConfig) GetCommandDescriptions() []*CommandDescription {
	if apc == nil {
		return nil
	}
	return apc.commandDescriptions
}

// GetFlagDescriptions - flagDescriptions field getter
func (apc *ArgParserConfig) GetFlagDescriptions() []*FlagDescription {
	if apc == nil {
		return nil
	}
	return apc.flagDescriptions
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (apc *ArgParserConfig) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = apc

	src := ArgParserConfigSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if src.AppHelpDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "app_help_description"`)
	}
	if src.HelpCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "help_command_description"`)
	}
	if len(src.CommandDescriptions) == 0 && src.NamelessCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: one or more of fields "nameless_command_description" or "command_descriptions" must be set`)
	}

	*apc = *src.ToConstPtr()

	return nil
}

// ArgParserConfigSrc - source for construct an arg parser configuration
type ArgParserConfigSrc struct {
	AppHelpDescription     *AppHelpDescription     `yaml:"app_help_description"`
	HelpCommandDescription *HelpCommandDescription `yaml:"help_command_description"`

	// one or more of these field must be set
	NamelessCommandDescription *NamelessCommandDescription `yaml:"nameless_command_description"`
	CommandDescriptions        []*CommandDescription       `yaml:"command_descriptions"`

	// optional
	FlagDescriptions []*FlagDescription `yaml:"flag_descriptions"`
}

// ToConstPtr converts src to AppHelpDescription pointer
func (m ArgParserConfigSrc) ToConstPtr() *ArgParserConfig {
	return (*ArgParserConfig)(unsafe.Pointer(&m))
}

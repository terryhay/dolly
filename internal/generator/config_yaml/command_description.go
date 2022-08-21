package config_yaml

import "fmt"

// CommandDescription - description of a command line command (which can contain flags and arguments)
type CommandDescription struct {
	Command             string
	DescriptionHelpInfo string

	// optional
	RequiredFlags        []string
	OptionalFlags        []string
	AdditionalCommands   []string
	ArgumentsDescription *ArgumentsDescription
}

// GetCommand - Command field getter
func (i *CommandDescription) GetCommand() string {
	if i == nil {
		return ""
	}
	return i.Command
}

// GetAdditionalCommands - AdditionalCommands field getter
func (i *CommandDescription) GetAdditionalCommands() []string {
	if i == nil {
		return nil
	}
	return i.AdditionalCommands
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *CommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (i *CommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgumentsDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (i *CommandDescription) GetRequiredFlags() []string {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *CommandDescription) GetOptionalFlags() []string {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}

type commandDescriptionSource struct {
	Command             string `yaml:"command"`
	DescriptionHelpInfo string `yaml:"description_help_info"`

	// optional
	RequiredFlags        []string              `yaml:"required_flags"`
	OptionalFlags        []string              `yaml:"optional_flags"`
	AdditionalCommands   []string              `yaml:"additional_names"`
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *CommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(commandDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.Command) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "command"`)
	}
	i.Command = source.Command

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}
	i.DescriptionHelpInfo = source.DescriptionHelpInfo

	// don't check optional fields
	i.RequiredFlags = source.RequiredFlags
	i.OptionalFlags = source.OptionalFlags
	i.AdditionalCommands = source.AdditionalCommands
	i.ArgumentsDescription = source.ArgumentsDescription

	return nil
}

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
func (cd *CommandDescription) GetCommand() string {
	if cd == nil {
		return ""
	}
	return cd.Command
}

// GetAdditionalCommands - AdditionalCommands field getter
func (cd *CommandDescription) GetAdditionalCommands() []string {
	if cd == nil {
		return nil
	}
	return cd.AdditionalCommands
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (cd *CommandDescription) GetDescriptionHelpInfo() string {
	if cd == nil {
		return ""
	}
	return cd.DescriptionHelpInfo
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (cd *CommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if cd == nil {
		return nil
	}
	return cd.ArgumentsDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (cd *CommandDescription) GetRequiredFlags() []string {
	if cd == nil {
		return nil
	}
	return cd.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (cd *CommandDescription) GetOptionalFlags() []string {
	if cd == nil {
		return nil
	}
	return cd.OptionalFlags
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (cd *CommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = cd

	source := struct {
		Command             string `yaml:"command"`
		DescriptionHelpInfo string `yaml:"description_help_info"`

		// optional
		RequiredFlags        []string              `yaml:"required_flags"`
		OptionalFlags        []string              `yaml:"optional_flags"`
		AdditionalCommands   []string              `yaml:"additional_names"`
		ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
	}{}
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.Command) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "command"`)
	}
	cd.Command = source.Command

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}
	cd.DescriptionHelpInfo = source.DescriptionHelpInfo

	// don't check optional fields
	cd.RequiredFlags = source.RequiredFlags
	cd.OptionalFlags = source.OptionalFlags
	cd.AdditionalCommands = source.AdditionalCommands
	cd.ArgumentsDescription = source.ArgumentsDescription

	return nil
}

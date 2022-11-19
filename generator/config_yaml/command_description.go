package config_yaml

import (
	"fmt"
	"unsafe"
)

// CommandDescription - description of a command line command (which can contain flags and arguments)
type CommandDescription struct {
	command             string
	descriptionHelpInfo string

	// optional
	requiredFlags        []string
	optionalFlags        []string
	additionalCommands   []string
	argumentsDescription *ArgumentsDescription
}

// GetCommand - command field getter
func (cd *CommandDescription) GetCommand() string {
	if cd == nil {
		return ""
	}
	return cd.command
}

// GetAdditionalCommands - additionalCommands field getter
func (cd *CommandDescription) GetAdditionalCommands() []string {
	if cd == nil {
		return nil
	}
	return cd.additionalCommands
}

// GetDescriptionHelpInfo - descriptionHelpInfo field getter
func (cd *CommandDescription) GetDescriptionHelpInfo() string {
	if cd == nil {
		return ""
	}
	return cd.descriptionHelpInfo
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (cd *CommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if cd == nil {
		return nil
	}
	return cd.argumentsDescription
}

// GetRequiredFlags - requiredFlags field getter
func (cd *CommandDescription) GetRequiredFlags() []string {
	if cd == nil {
		return nil
	}
	return cd.requiredFlags
}

// GetOptionalFlags - optionalFlags field getter
func (cd *CommandDescription) GetOptionalFlags() []string {
	if cd == nil {
		return nil
	}
	return cd.optionalFlags
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (cd *CommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = cd

	src := CommandDescriptionSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if len(src.Command) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "command"`)
	}

	if len(src.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}

	*cd = *src.ToConstPtr()

	return nil
}

// CommandDescriptionSrc - source for construct a description of a command line command
// (which can contain flags and arguments)
type CommandDescriptionSrc struct {
	Command             string `yaml:"command"`
	DescriptionHelpInfo string `yaml:"description_help_info"`

	// optional
	RequiredFlags        []string              `yaml:"required_flags"`
	OptionalFlags        []string              `yaml:"optional_flags"`
	AdditionalCommands   []string              `yaml:"additional_names"`
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// ToConstPtr converts src to CommandDescription pointer
func (m CommandDescriptionSrc) ToConstPtr() *CommandDescription {
	return (*CommandDescription)(unsafe.Pointer(&m))
}

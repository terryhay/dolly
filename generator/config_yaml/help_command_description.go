package config_yaml

import (
	"fmt"
	"unsafe"
)

// HelpCommandDescription - special description of a help command
type HelpCommandDescription struct {
	command string

	// optional
	additionalCommands []string
}

// GetCommand - command field getter
func (hcd *HelpCommandDescription) GetCommand() string {
	if hcd == nil {
		return ""
	}
	return hcd.command
}

// GetAdditionalCommands - additionalCommands field getter
func (hcd *HelpCommandDescription) GetAdditionalCommands() []string {
	if hcd == nil {
		return nil
	}
	return hcd.additionalCommands
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (hcd *HelpCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = hcd

	src := HelpCommandDescriptionSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if len(src.Command) == 0 {
		return fmt.Errorf(`helpCommandDescription unmarshal error: no required field "command"`)
	}

	*hcd = *src.ToConstPtr()

	return nil
}

// HelpCommandDescriptionSrc -source for construct a special description of a help command
type HelpCommandDescriptionSrc struct {
	Command string `yaml:"command"`

	// optional
	AdditionalCommands []string `yaml:"additional_commands"`
}

// ToConstPtr converts src to HelpCommandDescription pointer
func (m HelpCommandDescriptionSrc) ToConstPtr() *HelpCommandDescription {
	return (*HelpCommandDescription)(unsafe.Pointer(&m))
}

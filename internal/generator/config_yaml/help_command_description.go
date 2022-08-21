package config_yaml

import "fmt"

// HelpCommandDescription - special description of a help command
type HelpCommandDescription struct {
	Command string

	// optional
	AdditionalCommands []string
}

// GetCommand - Command field getter
func (i *HelpCommandDescription) GetCommand() string {
	if i == nil {
		return ""
	}
	return i.Command
}

// GetAdditionalCommands - AdditionalCommands field getter
func (i *HelpCommandDescription) GetAdditionalCommands() []string {
	if i == nil {
		return nil
	}
	return i.AdditionalCommands
}

type helpCommandDescriptionSource struct {
	Command            string   `yaml:"command"`
	AdditionalCommands []string `yaml:"additional_commands"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *HelpCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(helpCommandDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.Command) == 0 {
		return fmt.Errorf(`helpCommandDescription unmarshal error: no required field "command"`)
	}
	i.Command = source.Command

	// don't check optional fields
	i.AdditionalCommands = source.AdditionalCommands

	return nil
}

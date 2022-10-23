package config_yaml

import "fmt"

// HelpCommandDescription - special description of a help command
type HelpCommandDescription struct {
	Command string

	// optional
	AdditionalCommands []string
}

// GetCommand - Command field getter
func (hcd *HelpCommandDescription) GetCommand() string {
	if hcd == nil {
		return ""
	}
	return hcd.Command
}

// GetAdditionalCommands - AdditionalCommands field getter
func (hcd *HelpCommandDescription) GetAdditionalCommands() []string {
	if hcd == nil {
		return nil
	}
	return hcd.AdditionalCommands
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (hcd *HelpCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = hcd

	proxy := struct {
		Command            string   `yaml:"command"`
		AdditionalCommands []string `yaml:"additional_commands"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if len(proxy.Command) == 0 {
		return fmt.Errorf(`helpCommandDescription unmarshal error: no required field "command"`)
	}
	hcd.Command = proxy.Command

	// don't check optional fields
	hcd.AdditionalCommands = proxy.AdditionalCommands

	return nil
}

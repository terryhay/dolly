package config_yaml

import "fmt"

// FlagDescription - description of a command line flag
type FlagDescription struct {
	Flag                string
	DescriptionHelpInfo string
	SynopsisDescription string

	// optional
	ArgumentsDescription *ArgumentsDescription
}

// GetFlag - Flag getter
func (fd *FlagDescription) GetFlag() string {
	if fd == nil {
		return ""
	}
	return fd.Flag
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (fd *FlagDescription) GetDescriptionHelpInfo() string {
	if fd == nil {
		return ""
	}
	return fd.DescriptionHelpInfo
}

// GetSynopsisDescription - SynopsisDescription field getter
func (fd *FlagDescription) GetSynopsisDescription() string {
	if fd == nil {
		return ""
	}
	return fd.SynopsisDescription
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (fd *FlagDescription) GetArgumentsDescription() *ArgumentsDescription {
	if fd == nil {
		return nil
	}
	return fd.ArgumentsDescription
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (fd *FlagDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	proxy := struct {
		Flag                string `yaml:"flag"`
		DescriptionHelpInfo string `yaml:"description_help_info"`

		ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if len(proxy.Flag) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "flag"`)
	}
	fd.Flag = proxy.Flag

	if len(proxy.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "description_help_info"`)
	}
	fd.DescriptionHelpInfo = proxy.DescriptionHelpInfo

	// don't check optional field
	fd.ArgumentsDescription = proxy.ArgumentsDescription

	return nil
}

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
func (i *FlagDescription) GetFlag() string {
	if i == nil {
		return ""
	}
	return i.Flag
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *FlagDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetSynopsisDescription - SynopsisDescription field getter
func (i *FlagDescription) GetSynopsisDescription() string {
	if i == nil {
		return ""
	}
	return i.SynopsisDescription
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (i *FlagDescription) GetArgumentsDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgumentsDescription
}

type flagDescriptionSource struct {
	Flag                string `yaml:"flag"`
	DescriptionHelpInfo string `yaml:"description_help_info"`

	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *FlagDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(flagDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.Flag) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "flag"`)
	}
	i.Flag = source.Flag

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "description_help_info"`)
	}
	i.DescriptionHelpInfo = source.DescriptionHelpInfo

	// don't check optional field
	i.ArgumentsDescription = source.ArgumentsDescription

	return nil
}

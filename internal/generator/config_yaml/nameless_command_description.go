package config_yaml

import "fmt"

// NamelessCommandDescription - pseudo command which doesn't have call name
type NamelessCommandDescription struct {
	DescriptionHelpInfo string

	// optional
	RequiredFlags        []string
	OptionalFlags        []string
	ArgumentsDescription *ArgumentsDescription
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *NamelessCommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetRequiredFlags - RequiredFlags field getter
func (i *NamelessCommandDescription) GetRequiredFlags() []string {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *NamelessCommandDescription) GetOptionalFlags() []string {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (i *NamelessCommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgumentsDescription
}

type nullCommandDescriptionSource struct {
	DescriptionHelpInfo string `yaml:"description_help_info"`

	// optional
	RequiredFlags        []string              `yaml:"required_flags"`
	OptionalFlags        []string              `yaml:"optional_flags"`
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *NamelessCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(nullCommandDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}
	i.DescriptionHelpInfo = source.DescriptionHelpInfo

	// don't check optional fields
	i.RequiredFlags = source.RequiredFlags
	i.OptionalFlags = source.OptionalFlags
	i.ArgumentsDescription = source.ArgumentsDescription

	return nil
}

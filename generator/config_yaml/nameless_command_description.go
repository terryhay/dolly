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
func (ncd *NamelessCommandDescription) GetDescriptionHelpInfo() string {
	if ncd == nil {
		return ""
	}
	return ncd.DescriptionHelpInfo
}

// GetRequiredFlags - RequiredFlags field getter
func (ncd *NamelessCommandDescription) GetRequiredFlags() []string {
	if ncd == nil {
		return nil
	}
	return ncd.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (ncd *NamelessCommandDescription) GetOptionalFlags() []string {
	if ncd == nil {
		return nil
	}
	return ncd.OptionalFlags
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (ncd *NamelessCommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if ncd == nil {
		return nil
	}
	return ncd.ArgumentsDescription
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (ncd *NamelessCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = ncd

	proxy := struct {
		DescriptionHelpInfo string `yaml:"description_help_info"`

		// optional
		RequiredFlags        []string              `yaml:"required_flags"`
		OptionalFlags        []string              `yaml:"optional_flags"`
		ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if len(proxy.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}
	ncd.DescriptionHelpInfo = proxy.DescriptionHelpInfo

	// don't check optional fields
	ncd.RequiredFlags = proxy.RequiredFlags
	ncd.OptionalFlags = proxy.OptionalFlags
	ncd.ArgumentsDescription = proxy.ArgumentsDescription

	return nil
}

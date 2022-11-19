package config_yaml

import (
	"fmt"
	"unsafe"
)

// NamelessCommandDescription - pseudo command which doesn't have call name
type NamelessCommandDescription struct {
	descriptionHelpInfo string

	// optional
	requiredFlags        []string
	optionalFlags        []string
	argumentsDescription *ArgumentsDescription
}

// GetDescriptionHelpInfo - descriptionHelpInfo field getter
func (ncd *NamelessCommandDescription) GetDescriptionHelpInfo() string {
	if ncd == nil {
		return ""
	}
	return ncd.descriptionHelpInfo
}

// GetRequiredFlags - requiredFlags field getter
func (ncd *NamelessCommandDescription) GetRequiredFlags() []string {
	if ncd == nil {
		return nil
	}
	return ncd.requiredFlags
}

// GetOptionalFlags - optionalFlags field getter
func (ncd *NamelessCommandDescription) GetOptionalFlags() []string {
	if ncd == nil {
		return nil
	}
	return ncd.optionalFlags
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (ncd *NamelessCommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if ncd == nil {
		return nil
	}
	return ncd.argumentsDescription
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (ncd *NamelessCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = ncd

	src := NamelessCommandDescriptionSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if len(src.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}

	*ncd = *src.ToConstPtr()

	return nil
}

// NamelessCommandDescriptionSrc - cource for construct a pseudo command which doesn't have call name
type NamelessCommandDescriptionSrc struct {
	DescriptionHelpInfo string `yaml:"description_help_info"`

	// optional
	RequiredFlags        []string              `yaml:"required_flags"`
	OptionalFlags        []string              `yaml:"optional_flags"`
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// ToConstPtr converts src to NamelessCommandDescription pointer
func (m NamelessCommandDescriptionSrc) ToConstPtr() *NamelessCommandDescription {
	return (*NamelessCommandDescription)(unsafe.Pointer(&m))
}

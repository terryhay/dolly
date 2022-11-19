package config_yaml

import (
	"fmt"
	"unsafe"
)

// FlagDescription - description of a command line flag
type FlagDescription struct {
	flag                string
	descriptionHelpInfo string
	synopsisDescription string

	// optional
	argumentsDescription *ArgumentsDescription
}

// GetFlag - Flag getter
func (fd *FlagDescription) GetFlag() string {
	if fd == nil {
		return ""
	}
	return fd.flag
}

// GetDescriptionHelpInfo - descriptionHelpInfo field getter
func (fd *FlagDescription) GetDescriptionHelpInfo() string {
	if fd == nil {
		return ""
	}
	return fd.descriptionHelpInfo
}

// GetSynopsisDescription - SynopsisDescription field getter
func (fd *FlagDescription) GetSynopsisDescription() string {
	if fd == nil {
		return ""
	}
	return fd.synopsisDescription
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (fd *FlagDescription) GetArgumentsDescription() *ArgumentsDescription {
	if fd == nil {
		return nil
	}
	return fd.argumentsDescription
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (fd *FlagDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	src := FlagDescriptionSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if len(src.Flag) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "flag"`)
	}

	if len(src.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`flagDescription unmarshal error: no required field "description_help_info"`)
	}

	*fd = *src.ToConstPtr()

	return nil
}

// FlagDescriptionSrc - source description of a command line flag
type FlagDescriptionSrc struct {
	Flag                string `yaml:"flag"`
	DescriptionHelpInfo string `yaml:"description_help_info"`
	SynopsisDescription string `yaml:"synopsis_description"`

	// optional
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// ToConstPtr converts src to FlagDescription pointer
func (m FlagDescriptionSrc) ToConstPtr() *FlagDescription {
	return (*FlagDescription)(unsafe.Pointer(&m))
}

package config_yaml

import (
	"fmt"
	"unsafe"
)

// AppHelpDescription - application help info page
type AppHelpDescription struct {
	applicationName     string
	nameHelpInfo        string
	descriptionHelpInfo []string
}

// GetApplicationName - applicationName field getter
func (ahd *AppHelpDescription) GetApplicationName() string {
	if ahd == nil {
		return ""
	}
	return ahd.applicationName
}

// GetNameHelpInfo - nameHelpInfo field getter
func (ahd *AppHelpDescription) GetNameHelpInfo() string {
	if ahd == nil {
		return ""
	}
	return ahd.nameHelpInfo
}

// GetDescriptionHelpInfo - descriptionHelpInfo field getter
func (ahd *AppHelpDescription) GetDescriptionHelpInfo() []string {
	if ahd == nil {
		return nil
	}
	return ahd.descriptionHelpInfo
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (ahd *AppHelpDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = ahd

	src := AppHelpDescriptionSrc{}
	if err = unmarshal(&src); err != nil {
		return err
	}

	if len(src.ApplicationName) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "app_name"`)
	}

	if len(src.NameHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "name_help_info"`)
	}

	if len(src.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "description_help_info"`)
	}

	*ahd = *src.ToConstPtr()

	return nil
}

// AppHelpDescriptionSrc - source for construct an application help info page
type AppHelpDescriptionSrc struct {
	ApplicationName     string   `yaml:"app_name"`
	NameHelpInfo        string   `yaml:"name_help_info"`
	DescriptionHelpInfo []string `yaml:"description_help_info"`
}

// ToConstPtr converts src to AppHelpDescription pointer
func (m AppHelpDescriptionSrc) ToConstPtr() *AppHelpDescription {
	return (*AppHelpDescription)(unsafe.Pointer(&m))
}

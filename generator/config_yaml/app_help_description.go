package config_yaml

import "fmt"

// AppHelpDescription - application help info page
type AppHelpDescription struct {
	ApplicationName     string
	NameHelpInfo        string
	DescriptionHelpInfo []string
}

// GetApplicationName - ApplicationName field getter
func (ahd *AppHelpDescription) GetApplicationName() string {
	if ahd == nil {
		return ""
	}
	return ahd.ApplicationName
}

// GetNameHelpInfo - NameHelpInfo field getter
func (ahd *AppHelpDescription) GetNameHelpInfo() string {
	if ahd == nil {
		return ""
	}
	return ahd.NameHelpInfo
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (ahd *AppHelpDescription) GetDescriptionHelpInfo() []string {
	if ahd == nil {
		return nil
	}
	return ahd.DescriptionHelpInfo
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (ahd *AppHelpDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = ahd

	proxy := struct {
		ApplicationName     string   `yaml:"app_name"`
		NameHelpInfo        string   `yaml:"name_help_info"`
		DescriptionHelpInfo []string `yaml:"description_help_info"`
	}{}
	if err = unmarshal(&proxy); err != nil {
		return err
	}

	if len(proxy.ApplicationName) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "app_name"`)
	}
	ahd.ApplicationName = proxy.ApplicationName

	if len(proxy.NameHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "name_help_info"`)
	}
	ahd.NameHelpInfo = proxy.NameHelpInfo

	if len(proxy.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "description_help_info"`)
	}
	ahd.DescriptionHelpInfo = proxy.DescriptionHelpInfo

	return nil
}

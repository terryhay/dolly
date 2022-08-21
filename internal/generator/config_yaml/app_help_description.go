package config_yaml

import "fmt"

// AppHelpDescription - application help info data
type AppHelpDescription struct {
	ApplicationName     string
	NameHelpInfo        string
	DescriptionHelpInfo []string
}

// GetApplicationName - ApplicationName field getter
func (i *AppHelpDescription) GetApplicationName() string {
	if i == nil {
		return ""
	}
	return i.ApplicationName
}

// GetNameHelpInfo - NameHelpInfo field getter
func (i *AppHelpDescription) GetNameHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.NameHelpInfo
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *AppHelpDescription) GetDescriptionHelpInfo() []string {
	if i == nil {
		return nil
	}
	return i.DescriptionHelpInfo
}

type appHelpDescriptionSource struct {
	ApplicationName     string   `yaml:"app_name"`
	NameHelpInfo        string   `yaml:"name_help_info"`
	DescriptionHelpInfo []string `yaml:"description_help_info"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *AppHelpDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(appHelpDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.ApplicationName) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "app_name"`)
	}
	i.ApplicationName = source.ApplicationName

	if len(source.NameHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "name_help_info"`)
	}
	i.NameHelpInfo = source.NameHelpInfo

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`appHelpDescription unmarshal error: no required field "description_help_info"`)
	}
	i.DescriptionHelpInfo = source.DescriptionHelpInfo

	return nil
}

package config_yaml

import (
	"fmt"
	"unsafe"
)

// Config - code struct of a config yaml file
type Config struct {
	version         string
	argParserConfig *ArgParserConfig
	helpOutConfig   *HelpOutConfig
}

// GetVersion - version field getter
func (c *Config) GetVersion() string {
	if c == nil {
		return ""
	}
	return c.version
}

func (c *Config) GetArgParserConfig() *ArgParserConfig {
	if c == nil {
		return nil
	}
	return c.argParserConfig
}

func (c *Config) GetHelpOutConfig() *HelpOutConfig {
	if c == nil {
		return nil
	}
	return c.helpOutConfig
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = c

	src := ConfigSrc{}
	err := unmarshal(&src)
	if err != nil {
		return err
	}

	if len(src.Version) == 0 {
		return fmt.Errorf(`config unmarshal error: no required field "version"`)
	}
	if src.ArgParserConfig == nil {
		return fmt.Errorf(`config unmarshal error: no required field "arg_parser_config"`)
	}

	*c = *src.ToConstPtr()

	return nil
}

// ConfigSrc - source for construct a struct of a config file
type ConfigSrc struct {
	Version         string           `yaml:"version"`
	ArgParserConfig *ArgParserConfig `yaml:"arg_parser_config"`
	HelpOutConfig   *HelpOutConfig   `yaml:"help_out_config"`
}

// ToConstPtr converts src to Config pointer
func (m ConfigSrc) ToConstPtr() *Config {
	return (*Config)(unsafe.Pointer(&m))
}

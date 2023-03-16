package config_yaml

import (
	"errors"
	"fmt"
)

// Config - code struct of a config yaml file
type Config struct {
	version         string
	argParserConfig *ArgParserConfig
	helpOutConfig   *HelpOutConfig
}

// GetVersion gets version field
func (c *Config) GetVersion() string {
	if c == nil {
		return ""
	}

	return c.version
}

// GetArgParserConfig gets ArgParserConfig
func (c *Config) GetArgParserConfig() *ArgParserConfig {
	if c == nil {
		return nil
	}

	return c.argParserConfig
}

// GetHelpOutConfig gets HelpOutConfig
func (c *Config) GetHelpOutConfig() *HelpOutConfig {
	if c == nil {
		return nil
	}

	return c.helpOutConfig
}

var (
	// ErrConfigNilPointer - Config pointer is nil
	ErrConfigNilPointer = errors.New(`Config.IsValid: Config pointer is nil`)

	// ErrConfigNoVersion - no 'version' filed in source yaml file
	ErrConfigNoVersion = errors.New(`Config.IsValid: no required field 'version' in source yaml file`)
)

// IsValid cascade checks if Config is valid
func (c *Config) IsValid() error {
	if c == nil {
		return ErrConfigNilPointer
	}

	if len(c.version) == 0 {
		return ErrConfigNoVersion
	}

	if err := c.argParserConfig.IsValid(); err != nil {
		return fmt.Errorf(`Config.IsValid: %w`, err)
	}

	if err := c.helpOutConfig.IsValid(); err != nil {
		return fmt.Errorf(`Config.IsValid: %w`, err)
	}

	return nil
}

// ConfigOpt - source for construct a struct of a config file
type ConfigOpt struct {
	Version         string              `yaml:"version"`
	ArgParserConfig *ArgParserConfigOpt `yaml:"arg_parser"`
	HelpOutConfig   *HelpOutConfigOpt   `yaml:"help_out_config"`
}

// NewConfig converts opt to Config pointer
func NewConfig(opt *ConfigOpt) *Config {
	if opt == nil {
		return nil
	}

	return &Config{
		version:         opt.Version,
		argParserConfig: NewArgParserConfig(opt.ArgParserConfig),
		helpOutConfig:   NewHelpOutConfig(opt.HelpOutConfig),
	}
}

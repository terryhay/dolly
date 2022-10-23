package config_yaml

import "fmt"

// Config - code struct of a config yaml file
type Config struct {
	Version         string
	ArgParserConfig *ArgParserConfig
	HelpOutConfig   *HelpOutConfig
}

// GetVersion - Version field getter
func (c *Config) GetVersion() string {
	if c == nil {
		return ""
	}
	return c.Version
}

func (c *Config) GetArgParserConfig() *ArgParserConfig {
	if c == nil {
		return nil
	}
	return c.ArgParserConfig
}

func (c *Config) GetHelpOutConfig() *HelpOutConfig {
	if c == nil {
		return nil
	}
	return c.HelpOutConfig
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	_ = c

	proxy := struct {
		Version         string           `yaml:"version"`
		ArgParserConfig *ArgParserConfig `yaml:"arg_parser_config"`
		HelpOutConfig   *HelpOutConfig   `yaml:"help_out_config"`
	}{}
	err = unmarshal(&proxy)
	if err != nil {
		return err
	}

	if len(proxy.Version) == 0 {
		return fmt.Errorf(`config unmarshal error: no required field "version"`)
	}
	if proxy.ArgParserConfig == nil {
		return fmt.Errorf(`config unmarshal error: no required field "arg_parser_config"`)
	}

	*c = Config(proxy)

	return nil
}

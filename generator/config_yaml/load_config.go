package config_yaml

import (
	"errors"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"gopkg.in/yaml.v3"
)

var (
	// ErrLoadReadConfigFile - read config file error
	ErrLoadReadConfigFile = errors.New(`config_yaml.Load: read config file error`)

	// ErrLoadUnmarshal - unmarshal error
	ErrLoadUnmarshal = errors.New(`config_yaml.Load: unmarshal error`)
)

// Load loads config yaml file and unmarshal it into Config object
func Load(decOS osd.Proxy, configPath string) (*Config, error) {
	configYAML, err := decOS.ReadFile(configPath)
	if err != nil {
		return nil, errors.Join(ErrLoadReadConfigFile, err)
	}

	opt := &ConfigOpt{}
	if err = yaml.Unmarshal(configYAML, opt); err != nil {
		return nil, errors.Join(ErrLoadUnmarshal, err)
	}

	return NewConfig(opt), nil
}

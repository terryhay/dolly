package config_yaml

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// GetConfig - loads a config yaml file and unmarshal it into Config object
func GetConfig(configPath string) (*Config, *dollyerr.Error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, dollyerr.NewError(
			dollyerr.CodeGetConfigReadFileError,
			fmt.Errorf("config_yaml.GetConfig: read config file error: %v", err))
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, dollyerr.NewError(
			dollyerr.CodeGetConfigUnmarshalError,
			fmt.Errorf("config_yaml.GetConfig: unmarshal error: %v", err))
	}

	return config, nil
}

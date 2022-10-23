package config_yaml

import (
	"fmt"
	"github.com/terryhay/dolly/utils/dollyerr"
	"os"

	"gopkg.in/yaml.v3"
)

// GetConfig - loads a config yaml file and unmarshal it into Config object
func GetConfig(configPath string) (*Config, *dollyerr.Error) {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, dollyerr.NewError(
			dollyerr.CodeGetConfigReadFileError,
			fmt.Errorf("confYML.GetConfig: read config file error: %s", err.Error()))
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, dollyerr.NewError(
			dollyerr.CodeGetConfigUnmarshalError,
			fmt.Errorf("confYML.GetConfig: unmarshal error: %s", err.Error()))
	}

	return config, nil
}

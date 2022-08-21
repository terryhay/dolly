package config_data_extractor

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/pkg/dollyerr"
)

// ExtractFlagDescriptionMap extracts flag descriptions by flags from config object
func ExtractFlagDescriptionMap(flagDescriptions []*config_yaml.FlagDescription) (flagDescriptionMap map[string]*config_yaml.FlagDescription, error *dollyerr.Error) {
	descriptionCount := len(flagDescriptions)
	if descriptionCount == 0 {
		return nil, nil
	}
	flagDescriptionMap = make(map[string]*config_yaml.FlagDescription, descriptionCount)

	var contain bool
	for _, flagDescription := range flagDescriptions {
		if flagDescription == nil {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeUndefinedError,
					fmt.Errorf(`ExtractFlagDescriptionMap: config object contains zero flag description pointer`))
		}

		if _, contain = flagDescriptionMap[flagDescription.GetFlag()]; contain {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeConfigContainsDuplicateFlags,
					fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, flagDescription.GetFlag()))
		}

		flagDescriptionMap[flagDescription.GetFlag()] = flagDescription
	}

	return flagDescriptionMap, nil
}

// ExtractCommandDescriptionMap extracts command descriptions by commands from config object
func ExtractCommandDescriptionMap(commandDescriptions []*config_yaml.CommandDescription) (commandDescriptionMap map[string]*config_yaml.CommandDescription, error *dollyerr.Error) {
	descriptionCount := len(commandDescriptions)
	if descriptionCount == 0 {
		return nil, nil
	}
	commandDescriptionMap = make(map[string]*config_yaml.CommandDescription, descriptionCount)
	checkDuplicationsMap := make(map[string]bool, descriptionCount)

	var contain bool
	for _, commandDescription := range commandDescriptions {
		if commandDescription == nil {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeUndefinedError,
					fmt.Errorf(`ExtractFlagDescriptionMap: config object contains zero command description pointer`))
		}

		if _, contain = checkDuplicationsMap[commandDescription.GetCommand()]; contain {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeConfigContainsDuplicateCommands,
					fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, commandDescription.GetCommand()))
		}
		checkDuplicationsMap[commandDescription.GetCommand()] = true
		for _, command := range commandDescription.GetAdditionalCommands() {
			if _, contain = checkDuplicationsMap[command]; contain {
				return nil,
					dollyerr.NewError(
						dollyerr.CodeConfigContainsDuplicateCommands,
						fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, command))
			}
			checkDuplicationsMap[command] = true
		}

		commandDescriptionMap[commandDescription.GetCommand()] = commandDescription
	}

	return commandDescriptionMap, nil
}

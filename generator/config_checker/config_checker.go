package config_checker

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"github.com/terryhay/dolly/utils/dollyerr"
	"regexp"
	"strings"
)

const (
	maxFlagLen = 12
)

// Check checks command and flag descriptions for duplicates
func Check(
	namelessCommandDescription *confYML.NamelessCommandDescription,
	commandDescriptions map[string]*confYML.CommandDescription,
	flagDescriptions map[string]*confYML.FlagDescription,
) *dollyerr.Error {

	var (
		contain bool
		err     *dollyerr.Error
	)
	if err = checkArgumentDescription(namelessCommandDescription.GetArgumentsDescription()); err != nil {
		return err
	}

	var allUsingFlags map[string]string
	allUsingFlags, err = getAllFlagsFromCommandDescriptions(namelessCommandDescription, commandDescriptions)
	if err != nil {
		return err
	}

	for flag, flagDescription := range flagDescriptions {
		if _, contain = allUsingFlags[flag]; !contain {
			return dollyerr.NewError(dollyerr.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(`config_checker.Check: flag "%s" is not found in command descriptions`, flag))
		}
		if err = checkArgumentDescription(flagDescription.GetArgumentsDescription()); err != nil {
			return err
		}
	}

	for flag, command := range allUsingFlags {
		if _, contain = flagDescriptions[flag]; !contain {
			return dollyerr.NewError(dollyerr.CodeConfigUndefinedFlag, fmt.Errorf(`config_checker.Check: command "%s" conains undefined flag "%s"`, command, flag))
		}
	}

	return nil
}

// CheckFlag checks if flag has dash in front and is not too long
func CheckFlag(checkFlagCharsFunc func(s string) bool, flag string) *dollyerr.Error {
	if !checkFlagCharsFunc(flag) {
		return dollyerr.NewError(
			dollyerr.CodeConfigIncorrectCharacterInFlagName,
			fmt.Errorf("config_checker.CheckFlag: flag \"%s\" must contain a dash in front and latin chars", flag))
	}

	flagLen := len(flag)
	if flagLen > maxFlagLen {
		return dollyerr.NewError(
			dollyerr.CodeConfigIncorrectFlagLen,
			fmt.Errorf("config_checker.CheckFlag: flag \"%s\" has len=%d, max len=%d", flag, flagLen, maxFlagLen))
	}

	if flag[:1] != "-" {
		return dollyerr.NewError(
			dollyerr.CodeConfigFlagMustHaveDashInFront,
			fmt.Errorf("config_checker.CheckFlag: flag \"%s\" must have a dash in front", flag))
	}

	return nil
}

func checkArgumentDescription(argDescription *confYML.ArgumentsDescription) *dollyerr.Error {
	defaultValuesCount := len(argDescription.GetDefaultValues())
	if defaultValuesCount == 0 {
		return nil
	}

	if defaultValuesCount == 1 {
		if argDescription.GetAmountType() != apConf.ArgAmountTypeSingle {
			return dollyerr.NewError(
				dollyerr.CodeConfigUnexpectedDefaultValue,
				fmt.Errorf(`config_checker.checkArgumentDescription: you need to set amount_type "single" if you want to use default_values logic`))
		}
	} else {
		if argDescription.GetAmountType() != apConf.ArgAmountTypeList {
			return dollyerr.NewError(
				dollyerr.CodeConfigUnexpectedDefaultValue,
				fmt.Errorf(`config_checker.checkArgumentDescription: you need to set amount_type "list" if you want to use default_values logic`))
		}
	}

	allowedValuesCount := len(argDescription.GetAllowedValues())
	if allowedValuesCount == 0 {
		return nil
	}

	var allowed bool
	for i := 0; i < defaultValuesCount; i++ {
		allowed = false
		for j := 0; j < allowedValuesCount; j++ {
			if argDescription.GetDefaultValues()[i] == argDescription.GetAllowedValues()[j] {
				allowed = true
			}
		}

		if !allowed {
			return dollyerr.NewError(
				dollyerr.CodeConfigDefaultValueIsNotAllowed,
				fmt.Errorf(`config_checker.checkArgumentDescription: default value "%s" is not found in allowed values list: [%s]`,
					argDescription.GetDefaultValues()[i], strings.Join(argDescription.GetAllowedValues(), ", ")),
			)
		}
	}

	return nil
}

func getAllFlagsFromCommandDescriptions(
	namelessCommandDescription *confYML.NamelessCommandDescription,
	commandDescriptionMap map[string]*confYML.CommandDescription,
) (allUsingFlagMap map[string]string, err *dollyerr.Error) {

	checkFlagCharsFunc := regexp.MustCompile(`^[a-zA-Z-]+$`).MatchString
	allUsingFlagMap = make(map[string]string, 2*len(commandDescriptionMap))
	checkDuplicateFlagMap := make(map[string]bool, 2*len(commandDescriptionMap))

	var (
		contain bool
		flag    string
	)

	// checking for nameless command
	const namelessCommand string = "NamelessCommand"
	for _, flag = range namelessCommandDescription.GetRequiredFlags() {
		if err = CheckFlag(checkFlagCharsFunc, flag); err != nil {
			return nil, err
		}

		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, dollyerr.NewError(dollyerr.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, namelessCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = namelessCommand
	}

	for _, flag = range namelessCommandDescription.GetOptionalFlags() {
		if err = CheckFlag(checkFlagCharsFunc, flag); err != nil {
			return nil, err
		}
		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, dollyerr.NewError(dollyerr.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, namelessCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = namelessCommand
	}

	// checking for commands
	for _, commandDescription := range commandDescriptionMap {
		checkDuplicateFlagMap = map[string]bool{}

		if err = checkArgumentDescription(commandDescription.GetArgumentsDescription()); err != nil {
			return nil, err
		}

		for _, flag = range commandDescription.GetRequiredFlags() {
			err = CheckFlag(checkFlagCharsFunc, flag)
			if err != nil {
				return nil, err
			}
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, dollyerr.NewError(dollyerr.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}

		for _, flag = range commandDescription.GetOptionalFlags() {
			err = CheckFlag(checkFlagCharsFunc, flag)
			if err != nil {
				return nil, err
			}
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, dollyerr.NewError(dollyerr.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}
	}

	return allUsingFlagMap, nil
}

package arg_parser_impl

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
	"strings"
)

func checkNoDashInFront(arg string) bool {
	if len(arg) == 0 {
		return true
	}
	return arg[:1] != "-"
}

func checkParsedData(
	usingCommandDescription *apConf.CommandDescription,
	data *parsed_data.ParsedData,
) *dollyerr.Error {
	var (
		argDescription *apConf.ArgumentsDescription
		contain        bool
		flag           apConf.Flag
		parsedFlagData = data.GetFlagDataMap()
	)

	// check if all required flags is set
	for flag = range usingCommandDescription.GetRequiredFlags() {
		if _, contain = parsedFlagData[flag]; !contain {
			return dollyerr.NewError(
				dollyerr.CodeArgParserRequiredFlagIsNotSet,
				fmt.Errorf("CmdArgParser.checkParsedData: required flag is not set: %s", flag))
		}
	}

	// check command arguments
	argDescription = usingCommandDescription.GetArgDescription()
	if argDescription.GetAmountType() != apConf.ArgAmountTypeNoArgs {
		if data.GetAgrData() == nil {
			return dollyerr.NewError(
				dollyerr.CodeArgParserCommandDoesNotContainArgs,
				fmt.Errorf("CmdArgParser.checkParsedData: command arg is not set: %s", flag))
		}
	}

	return nil
}

func isValueAllowed(argDescription *apConf.ArgumentsDescription, value string) *dollyerr.Error {
	if argDescription == nil {
		return dollyerr.NewError(
			dollyerr.CodeArgParserCheckValueAllowabilityError,
			fmt.Errorf("isValueAllowed: try to check a value \"%s\" allowability by nil pointer", value))
	}

	if len(argDescription.GetAllowedValues()) == 0 {
		return nil
	}

	if _, allow := argDescription.GetAllowedValues()[value]; !allow {
		allowedValuesSlice := make([]string, 0, len(argDescription.GetAllowedValues()))
		for allowedValue := range argDescription.GetAllowedValues() {
			allowedValuesSlice = append(allowedValuesSlice, allowedValue)
		}

		return dollyerr.NewError(
			dollyerr.CodeArgParserArgValueIsNotAllowed,
			fmt.Errorf(`isValueAllowed: value "%s" is not found in list allowed values: [%s]`,
				value, strings.Join(allowedValuesSlice, ", ")))
	}

	return nil
}

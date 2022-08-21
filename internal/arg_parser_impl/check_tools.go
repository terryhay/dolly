package arg_parser_impl

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
	"strings"
)

func checkNoDashInFront(arg string) bool {
	if len(arg) == 0 {
		return true
	}
	return arg[:1] != "-"
}

func checkParsedData(
	usingCommandDescription *dollyconf.CommandDescription,
	data *parsed_data.ParsedData,
) *dollyerr.Error {
	var (
		argDescription *dollyconf.ArgumentsDescription
		contain        bool
		flag           dollyconf.Flag
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
	if argDescription.GetAmountType() != dollyconf.ArgAmountTypeNoArgs {
		if data.GetAgrData() == nil {
			return dollyerr.NewError(
				dollyerr.CodeArgParserCommandDoesNotContainArgs,
				fmt.Errorf("CmdArgParser.checkParsedData: command arg is not set: %s", flag))
		}
	}

	return nil
}

func isValueAllowed(argDescription *dollyconf.ArgumentsDescription, value string) *dollyerr.Error {
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

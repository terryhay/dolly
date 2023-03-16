package parser

import (
	"errors"
	"fmt"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	"github.com/terryhay/dolly/argparser/parsed"
)

var (
	// ErrCheckParseResultRequiredFlagIsNotSet - required flag is not set
	ErrCheckParseResultRequiredFlagIsNotSet = errors.New(`checkParseResult: required flag is not set`)

	// ErrCheckParseResultRequiredArgIsNotSet - required argument is not set
	ErrCheckParseResultRequiredArgIsNotSet = errors.New(`checkParseResult: required argument is not set`)

	// ErrCheckParsedResultEmptyPlaceholder - parsed placeholder doesn't contain neater flag not argument
	ErrCheckParsedResultEmptyPlaceholder = errors.New(`checkParseResult: parsed placeholder doesn't contain neater flag not argument`)
)

func checkParseResult(command *apConf.Command, parseResult *parsed.Result) error {

	// check if all required groups is set
	for _, placeholder := range command.GetPlaceholders() {
		if placeholder.GetIsFlagOptional() {
			continue
		}

		placeholderParsed := parseResult.PlaceholderByID(placeholder.GetID())

		if placeholder.IsFlagRequired() && !placeholderParsed.HasFlag() {
			return fmt.Errorf(`%w: placeholder id '%d'; flag '%s'`,
				ErrCheckParseResultRequiredFlagIsNotSet, placeholder.GetID(), placeholder.CreateStringWithFlagNames())
		}

		if placeholder.IsArgRequired() && !placeholderParsed.HasArg() {
			return fmt.Errorf(`%w: placeholder id '%d'; expected argument: %s`,
				ErrCheckParseResultRequiredArgIsNotSet, placeholder.GetID(), placeholder.GetArgument().CreateStringWithArgInfo())
		}

		if placeholderParsed != nil && !placeholderParsed.HasFlag() && !placeholderParsed.HasArg() {
			return fmt.Errorf(`%w: placeholder id '%d'`, ErrCheckParsedResultEmptyPlaceholder, placeholderParsed.GetID())
		}
	}

	return nil
}

package parser

import (
	"errors"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
	"github.com/terryhay/dolly/argparser/parsed"
)

type parseState uint8

const (
	parseStateFindingArgGroupDescription parseState = iota
	parseStateReadingSingleArg
	parseStateReadingArgList
)

var (
	// ErrParseUsingCommandDescription - usingCommandDescription returned error
	ErrParseUsingCommandDescription = errors.New(`parser.Parse: usingCommandDescription returned error`)

	// ErrParseProcess - process returned error
	ErrParseProcess = errors.New(`parser.Parse: process returned error`)

	// ErrParseCheckParsedData - checkParseResult returned error
	ErrParseCheckParsedData = errors.New(`parser.Parse: checkParseResult returned error`)
)

// Parse extracts command line arguments and returns them as parsed.Result object
func Parse(config apConf.ArgParserConfig, args []string) (*parsed.Result, error) {
	itCmdArg, descriptionCommand, err := usingCommandDescription(config, clArg.MakeIterator(args))
	if err != nil {
		return nil, errors.Join(ErrParseUsingCommandDescription, err)
	}

	var res *parsed.Result
	res, err = process(descriptionCommand, itCmdArg)
	if err != nil {
		return nil, errors.Join(ErrParseProcess, err)
	}

	if err = checkParseResult(descriptionCommand, res); err != nil {
		return nil, errors.Join(ErrParseCheckParsedData, err)
	}

	return res, nil
}

// ErrUsingCommandDescriptionNoCommands - arguments are not set, but nameless command is not defined in config object
var ErrUsingCommandDescriptionNoCommands = errors.New(`usingCommandDescription: arguments are not set, but nameless command is not defined in config object`)

func usingCommandDescription(config apConf.ArgParserConfig, it clArg.Iterator) (clArg.Iterator, *apConf.Command, error) {
	description := config.CommandByName(it.GetArg().ToNameCommand())
	if description == nil {
		description = config.GetCommandNameless()
		if description == nil {
			return clArg.Iterator{}, nil, ErrUsingCommandDescriptionNoCommands
		}

		return it, description, nil
	}

	return it.Next(), description, nil
}

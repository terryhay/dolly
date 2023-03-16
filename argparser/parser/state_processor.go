package parser

import (
	"errors"
	"fmt"

	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	clArg "github.com/terryhay/dolly/argparser/command_line_argument"
	"github.com/terryhay/dolly/argparser/parsed"
	plit "github.com/terryhay/dolly/argparser/placeholder_iterator"
)

type parseProcessData struct {
	State parseState

	ItPlaceholder plit.Iterator
	ArgConfig     *apConf.Argument

	Result parsed.ResultOpt
}

var (
	// ErrProcessProcessFindingPlaceholder - findingPlaceholder returned error
	ErrProcessProcessFindingPlaceholder = errors.New(`process: findingPlaceholder returned error`)

	// ErrProcessProcessReadingSingleArgument - setSingleArg returned error
	ErrProcessProcessReadingSingleArgument = errors.New(`process: setSingleArg returned error`)

	// ErrProcessProcessReadingArgumentList - processReadingArgumentList returned error
	ErrProcessProcessReadingArgumentList = errors.New(`process: processReadingArgumentList returned error`)
)

func process(command *apConf.Command, itCmdArg clArg.Iterator) (*parsed.Result, error) {
	processData := parseProcessData{
		State:         parseStateFindingArgGroupDescription,
		ItPlaceholder: plit.Make(command.GetPlaceholders()),
		Result: parsed.ResultOpt{
			CommandMainName: command.GetNameMain(),
		},
	}

	var err error
	for ; !itCmdArg.IsEnded(); itCmdArg = itCmdArg.Next() {
		switch {
		case processData.State == parseStateFindingArgGroupDescription:
			processData, err = findingPlaceholder(processData, itCmdArg.GetArg())
			if err != nil {
				return nil, errors.Join(ErrProcessProcessFindingPlaceholder, err)
			}

		case processData.State == parseStateReadingSingleArg:
			processData, err = setSingleArg(processData, itCmdArg.GetArg())
			if err != nil {
				return nil, errors.Join(ErrProcessProcessReadingSingleArgument, err)
			}

		case processData.State == parseStateReadingArgList:
			processData, err = processReadingArgumentList(processData, itCmdArg.GetArg())
			if err != nil {
				return nil, errors.Join(ErrProcessProcessReadingArgumentList, err)
			}
		}
	}

	return parsed.MakeResult(&processData.Result), nil
}

var (
	// ErrFindingPlaceholderRequiredFlagIsNotSet - placeholder can't be skipped
	ErrFindingPlaceholderRequiredFlagIsNotSet = errors.New(`findingPlaceholder: placeholder can't be skipped`)

	// ErrFindingPlaceholderNotFound - placeholder is not found
	ErrFindingPlaceholderNotFound = errors.New(`findingPlaceholder: placeholder is not found`)
)

func findingPlaceholder(processData parseProcessData, arg clArg.Argument) (
	parseProcessData, error,
) {
	for placeholder := processData.ItPlaceholder.Next(); placeholder != nil; placeholder = processData.ItPlaceholder.Next() {
		switch {
		case placeholder.HasFlags():
			if !placeholder.HasFlagName(arg.ToNameFlag()) {
				if placeholder.IsFlagRequired() {
					return processData,
						fmt.Errorf(`%w: placeholder id "%d; command line argConfig "%s"`,
							ErrFindingPlaceholderRequiredFlagIsNotSet, placeholder.GetID(), arg.String())
				}
				// Current placeholder is optional, so just skip it
				continue
			}
			processData.Result.SetFlagName(placeholder.GetID(), arg.ToNameFlag())

			if argConfig := placeholder.GetArgument(); argConfig != nil {
				state := parseStateReadingSingleArg
				if argConfig.GetIsList() {
					state = parseStateReadingArgList
				}

				return parseProcessData{
					State:         state,
					ItPlaceholder: processData.ItPlaceholder,
					ArgConfig:     argConfig,
					Result:        processData.Result,
				}, nil
			}

			return parseProcessData{
				State:         parseStateFindingArgGroupDescription,
				ItPlaceholder: processData.ItPlaceholder,
				Result:        processData.Result,
			}, nil

		case placeholder.HasArg():
			if placeholder.GetArgument().GetIsList() {
				return processReadingArgumentList(
					parseProcessData{
						State:         parseStateReadingArgList,
						ItPlaceholder: processData.ItPlaceholder,
						ArgConfig:     placeholder.GetArgument(),
						Result:        processData.Result,
					},
					arg)
			}

			return setSingleArg(
				parseProcessData{
					State:         parseStateReadingSingleArg,
					ItPlaceholder: processData.ItPlaceholder,
					ArgConfig:     placeholder.GetArgument(),
					Result:        processData.Result,
				},
				arg)
		}
	}

	return parseProcessData{}, fmt.Errorf(`%w: command line argument "%s"`,
		ErrFindingPlaceholderNotFound, arg.String())
}

var (
	// ErrSetSingleArg - command line argument is not allowed
	ErrSetSingleArg = errors.New(`setSingleArg: command line argument is not allowed`)
)

func setSingleArg(processData parseProcessData, arg clArg.Argument) (
	parseProcessData, error,
) {
	if arg.IsFlag() {
		return notSetArgValueCase(processData, arg)
	}

	if err := processData.ArgConfig.IsArgAllowed(arg); err != nil {
		return parseProcessData{},
			errors.Join(
				fmt.Errorf(`%w: placeholder id "%d"`,
					ErrSetSingleArg, processData.ItPlaceholder.Get().GetID()),
				err,
			)
	}

	processData.Result.SetArg(
		processData.ItPlaceholder.Get().GetID(),
		parsed.MakeArgValue(arg),
	)

	return parseProcessData{
		State:         parseStateFindingArgGroupDescription,
		ItPlaceholder: processData.ItPlaceholder,
		Result:        processData.Result,
	}, nil
}

func processReadingArgumentList(processData parseProcessData, arg clArg.Argument) (
	parseProcessData, error,
) {
	if arg.IsFlag() {
		// argument sequence is ended
		if processData.Result.PlaceholderDoesNotHaveArgs(processData.ItPlaceholder.Get().GetID()) {
			return notSetArgValueCase(processData, arg)
		}

		return findingPlaceholder(
			parseProcessData{
				State:         parseStateFindingArgGroupDescription,
				ItPlaceholder: processData.ItPlaceholder,
				Result:        processData.Result,
			}, arg,
		)
	}

	processData.Result.SetArg(processData.ItPlaceholder.Get().GetID(), parsed.MakeArgValue(arg))

	return parseProcessData{
		State:         parseStateReadingArgList,
		ItPlaceholder: processData.ItPlaceholder,
		ArgConfig:     processData.ArgConfig,
		Result:        processData.Result,
	}, nil
}

var (
	// ErrNotSetArgValueCaseNoRequiredArg - can't set command line argument as placeholder arg
	ErrNotSetArgValueCaseNoRequiredArg = errors.New(`notSetArgValueCase: can't set command line argument as placeholder arg`)
)

func notSetArgValueCase(processData parseProcessData, arg clArg.Argument) (parseProcessData, error) {
	// arg is a flag, ok. but we are expecting flag/command argument value now,
	// so let's try to get it from default values slice
	defaultValues := processData.ArgConfig.GetDefaultValues()
	if len(defaultValues) == 0 {
		if processData.ItPlaceholder.Get().IsArgRequired() {

			placeholderParsed := processData.Result.PlaceholderByID(processData.ItPlaceholder.Get().GetID())
			if placeholderParsed != nil && len(placeholderParsed.Flag) != 0 {
				return parseProcessData{},
					fmt.Errorf(`%w: placeholder id "%d"; flag "%s; arg "%s"; command`,
						ErrNotSetArgValueCaseNoRequiredArg, placeholderParsed.ID, placeholderParsed.Flag, arg.String())
			}

			return parseProcessData{},
				fmt.Errorf(`%w: placeholder id "%d"; arg "%s"; command`,
					ErrNotSetArgValueCaseNoRequiredArg, parsed.NewPlaceholder(placeholderParsed).GetID(), arg.String())
		}

		return findingPlaceholder(
			parseProcessData{
				State:         parseStateFindingArgGroupDescription,
				ItPlaceholder: processData.ItPlaceholder,
				Result:        processData.Result,
			}, arg,
		)
	}

	for _, value := range defaultValues {
		processData.Result.SetArg(processData.ItPlaceholder.Get().GetID(), parsed.MakeArgValue(clArg.Argument(value)))
	}

	// default value is set, good
	// time to switch logic to process a flag
	return findingPlaceholder(
		parseProcessData{
			State:         parseStateFindingArgGroupDescription,
			ItPlaceholder: processData.ItPlaceholder,
			Result:        processData.Result,
		}, arg,
	)
}

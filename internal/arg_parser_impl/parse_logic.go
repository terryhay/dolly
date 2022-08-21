package arg_parser_impl

import (
	"fmt"
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
)

// Parse processes command line arguments
func (i *ArgParserImpl) Parse(args []string) (*parsed_data.ParsedData, *dollyerr.Error) {
	_ = i // check if pointer is nil here and check no further

	var (
		argIndexStartValue      = 1
		commandArgData          *parsed_data.ParsedArgData
		res                     *parsed_data.ParsedData
		usingCommandDescription *dollyconf.CommandDescription
	)

	if len(args) == 0 {
		if i.namelessCommandDescription == nil {
			return nil, dollyerr.NewError(
				dollyerr.CodeArgParserNamelessCommandUndefined,
				fmt.Errorf(`ArgParserImpl.Parse: arguments are not set, but nameless command is not defined in config object`))
		}

		res = parsed_data.NewParsedData(i.namelessCommandDescription.GetID(), "", nil, nil)
		if i.mErr = checkParsedData(i.namelessCommandDescription, res); i.mErr != nil {
			return nil, i.mErr
		}
		return res, nil
	}
	if len(i.commandDescriptions) == 0 && i.namelessCommandDescription == nil {
		return nil,
			dollyerr.NewError(
				dollyerr.CodeArgParserIsNotInitialized,
				fmt.Errorf(`ArgParserImpl.Parse: parser is not initialized`))
	}

	// Determinate command
	command := dollyconf.Command(args[0])
	usingCommandDescription, i.mOK = i.commandDescriptions[command]
	if !i.mOK {
		if i.namelessCommandDescription == nil {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeCantFindFlagNameInGroupSpec,
					fmt.Errorf(`ArgParserImpl.Parse: unexpected command: "%s"`, command))
		}
		usingCommandDescription = i.namelessCommandDescription
		command = ""
		argIndexStartValue = 0
	}

	if i.mArgDescription = usingCommandDescription.GetArgDescription(); i.mArgDescription != nil {
		commandArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
		i.mParsedArgData = commandArgData
	}
	i.mParseState = getParseState(i.mArgDescription)

	i.rParsedFlagDataMap = make(
		map[dollyconf.Flag]*parsed_data.ParsedFlagData,
		len(usingCommandDescription.GetRequiredFlags())+len(usingCommandDescription.GetOptionalFlags()))

	for argIndex := argIndexStartValue; argIndex < len(args); argIndex++ {
		if i.mErr = i.stateProcessors[i.mParseState](args[argIndex]); i.mErr != nil {
			return nil, i.mErr
		}
	}

	res = parsed_data.NewParsedData(usingCommandDescription.GetID(), command, commandArgData, i.rParsedFlagDataMap)
	if i.mErr = checkParsedData(usingCommandDescription, res); i.mErr != nil {
		return nil, i.mErr
	}

	return res, nil
}

func (i *ArgParserImpl) processReadingFlag(arg string) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	i.mFlag = dollyconf.Flag(arg)
	if i.mFlagDescription, i.mOK = i.flagDescriptions[i.mFlag]; !i.mOK {
		return dollyerr.NewError(
			dollyerr.CodeArgParserUnexpectedArg,
			fmt.Errorf(`ArgParserImpl.Parse: unexpected argument: "%s"`, arg))
	}

	i.mParsedArgData = nil
	if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
		i.mParsedArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
	}

	if _, i.mOK = i.rParsedFlagDataMap[i.mFlag]; i.mOK {
		return dollyerr.NewError(
			dollyerr.CodeArgParserDuplicateFlags,
			fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, arg))
	}
	i.rParsedFlagDataMap[i.mFlag] = parsed_data.NewParsedFlagData(i.mFlag, i.mParsedArgData)

	i.mParseState = getParseState(i.mArgDescription)
	return nil
}

func (i *ArgParserImpl) processReadingSingleArgument(arg string) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	if !checkNoDashInFront(arg) {
		return i.notSetArgValueCase(arg)
	}

	if i.mErr = isValueAllowed(i.mArgDescription, arg); i.mErr != nil {
		return dollyerr.NewError(
			i.mErr.Code(),
			fmt.Errorf(`ArgParserImpl.Parse: flag "%s": %s`, i.mFlag, i.mErr.Error()))
	}

	i.mParsedArgData.ArgValues = []parsed_data.ArgValue{parsed_data.ArgValue(arg)}

	i.mParseState = parseStateReadingFlag
	return nil
}

func (i *ArgParserImpl) processReadingArgumentList(arg string) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	if !checkNoDashInFront(arg) {
		if len(i.mParsedArgData.ArgValues) == 0 {
			return i.notSetArgValueCase(arg)
		}

		if i.mFlagDescription, i.mOK = i.flagDescriptions[dollyconf.Flag(arg)]; !i.mOK {
			return dollyerr.NewError(
				dollyerr.CodeArgParserUnexpectedFlag,
				fmt.Errorf(`ArgParserImpl.Parse: unexpected flag: "%s"`, arg))
		}

		i.mParsedArgData = nil
		if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
			i.mParsedArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
		}

		i.mFlag = dollyconf.Flag(arg)
		if _, i.mOK = i.rParsedFlagDataMap[i.mFlag]; i.mOK {
			return dollyerr.NewError(
				dollyerr.CodeArgParserDuplicateFlags,
				fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, arg))
		}
		i.rParsedFlagDataMap[i.mFlag] = parsed_data.NewParsedFlagData(i.mFlag, i.mParsedArgData)

		i.mParseState = getParseState(i.mArgDescription)
		return nil
	}

	i.mParsedArgData.ArgValues = append(i.mParsedArgData.ArgValues, parsed_data.ArgValue(arg))

	return nil
}

func (i *ArgParserImpl) notSetArgValueCase(arg string) *dollyerr.Error {
	// current command line argument looks like a flag
	// let's check if it is a flag
	i.mTmpFlagDescription, i.mOK = i.flagDescriptions[dollyconf.Flag(arg)]
	if !i.mOK {
		return dollyerr.NewError(
			dollyerr.CodeArgParserDashInFrontOfArg,
			fmt.Errorf(`ArgParserImpl.Parse: argument "%s" contains a dash in front`, arg))
	}

	// arg is a flag, ok. but we are expecting flag/command argument value now,
	// so let's try to get it from default values slice
	if len(i.mArgDescription.GetDefaultValues()) == 0 {
		return dollyerr.NewError(
			dollyerr.CodeArgParserFlagMustHaveArg,
			fmt.Errorf(`ArgParserImpl.Parse: flag "%s" must have an arg`, arg))
	}
	i.mParsedArgData.ArgValues = copyDefaultValues2ArgValues(i.mArgDescription.GetDefaultValues(), i.mParsedArgData.ArgValues)

	// default value is set, good
	// time to switch logic to process a flag
	i.mParseState = parseStateReadingFlag
	return i.processReadingFlag(arg)
}

func getParseState(argumentsDescription *dollyconf.ArgumentsDescription) parseState {
	switch argumentsDescription.GetAmountType() {
	case dollyconf.ArgAmountTypeSingle:
		return parseStateReadingSingleArgument
	case dollyconf.ArgAmountTypeList:
		return parseStateReadingArgumentList
	default:
		return parseStateReadingFlag
	}
}

func copyDefaultValues2ArgValues(defaultValueSlice []string, argValueSlice []parsed_data.ArgValue) []parsed_data.ArgValue {
	for i := 0; i < len(defaultValueSlice); i++ {
		argValueSlice = append(argValueSlice, parsed_data.ArgValue(defaultValueSlice[i]))
	}

	return argValueSlice
}

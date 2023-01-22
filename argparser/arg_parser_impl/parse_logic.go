package arg_parser_impl

import (
	"fmt"
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	cmdArg "github.com/terryhay/dolly/argparser/cmd_arg"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
)

// Parse processes command line arguments
func (i *ArgParserImpl) Parse(it cmdArg.CmdArgIterator) (*parsed_data.ParsedData, *dollyerr.Error) {
	_ = i // check if pointer is nil here and check no further

	var (
		commandArgData          *parsed_data.ParsedArgData
		res                     *parsed_data.ParsedData
		usingCommandDescription *apConf.CommandDescription
	)

	if it.IsEnded() {
		if i.namelessCommandDescription == nil {
			return nil, dollyerr.NewError(
				dollyerr.CodeArgParserNamelessCommandUndefined,
				fmt.Errorf(`ArgParserImpl.Parse: arguments are not set, but nameless command is not defined in config object`))
		}

		res = parsed_data.NewParsedData(i.namelessCommandDescription.GetID(), "", nil, nil)
		err := checkParsedData(i.namelessCommandDescription, res)
		if err != nil {
			return nil, dollyerr.Append(err, fmt.Errorf("ArgParserImpl.Parse: checkParsedData error"))
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
	commandCandidate := it.GetCmdArg()
	usingCommandDescription, i.mOK = i.commandDescriptions[commandCandidate.ToCommand()]
	if !i.mOK {
		if i.namelessCommandDescription == nil {
			return nil,
				dollyerr.NewError(
					dollyerr.CodeCantFindFlagNameInGroupSpec,
					fmt.Errorf(`ArgParserImpl.Parse: unexpected command: "%s"`, commandCandidate),
				)
		}
		usingCommandDescription = i.namelessCommandDescription
		commandCandidate = cmdArg.CmdArgEmpty
	} else {
		it.Next()
	}

	if i.mArgDescription = usingCommandDescription.GetArgDescription(); i.mArgDescription != nil {
		commandArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
		i.mParsedArgData = commandArgData
	}
	i.mParseState = getParseState(i.mArgDescription)

	i.rParsedFlagDataMap = make(
		map[apConf.Flag]*parsed_data.ParsedFlagData,
		len(usingCommandDescription.GetRequiredFlags())+len(usingCommandDescription.GetOptionalFlags()))

	for arg := it.GetCmdArg(); arg.IsValid(); arg = it.Next() {
		err := i.stateProcessors[i.mParseState](arg)
		if err != nil {
			return nil, dollyerr.Append(err, fmt.Errorf("ArgParserImpl.Parse: stateProcessor error"))
		}
	}

	res = parsed_data.NewParsedData(usingCommandDescription.GetID(), commandCandidate.ToCommand(), commandArgData, i.rParsedFlagDataMap)
	err := checkParsedData(usingCommandDescription, res)
	if err != nil {
		return nil, dollyerr.Append(err, fmt.Errorf("ArgParserImpl.Parse: checkParsedData error"))
	}

	return res, nil
}

func (i *ArgParserImpl) processReadingFlag(flagCandidate cmdArg.CmdArg) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	i.mFlag = flagCandidate.ToFlag()
	if i.mFlagDescription, i.mOK = i.flagDescriptions[flagCandidate.ToFlag()]; !i.mOK {
		return dollyerr.NewError(
			dollyerr.CodeArgParserUnexpectedArg,
			fmt.Errorf(`ArgParserImpl.Parse: unexpected argument: "%s"`, flagCandidate))
	}

	i.mParsedArgData = nil
	if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
		i.mParsedArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
	}

	if _, i.mOK = i.rParsedFlagDataMap[flagCandidate.ToFlag()]; i.mOK {
		return dollyerr.NewError(
			dollyerr.CodeArgParserDuplicateFlags,
			fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, flagCandidate))
	}
	i.rParsedFlagDataMap[flagCandidate.ToFlag()] = parsed_data.NewParsedFlagData(flagCandidate.ToFlag(), i.mParsedArgData)

	i.mParseState = getParseState(i.mArgDescription)
	return nil
}

func (i *ArgParserImpl) processReadingSingleArgument(argCandidate cmdArg.CmdArg) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	if argCandidate.HasDashInFront() {
		return i.notSetArgValueCase(argCandidate)
	}

	err := isValueAllowed(i.mArgDescription, argCandidate.ToString())
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf(`ArgParserImpl.processReadingSingleArgument: flag "%s"`, i.mFlag))
	}

	i.mParsedArgData.ArgValues = []parsed_data.ArgValue{argCandidate.ToArgValue()}

	i.mParseState = parseStateReadingFlag
	return nil
}

func (i *ArgParserImpl) processReadingArgumentList(argListCandidate cmdArg.CmdArg) *dollyerr.Error {
	_ = i // check if pointer is nil here and check no further

	if argListCandidate.HasDashInFront() {
		flagCandidate := argListCandidate
		if len(i.mParsedArgData.GetArgValues()) == 0 {
			return i.notSetArgValueCase(flagCandidate)
		}

		if i.mFlagDescription, i.mOK = i.flagDescriptions[flagCandidate.ToFlag()]; !i.mOK {
			return dollyerr.NewError(
				dollyerr.CodeArgParserUnexpectedFlag,
				fmt.Errorf(`ArgParserImpl.Parse: unexpected flag: "%s"`, flagCandidate))
		}

		i.mParsedArgData = nil
		if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
			i.mParsedArgData = parsed_data.NewParsedArgData(make([]parsed_data.ArgValue, 0, 8))
		}

		i.mFlag = flagCandidate.ToFlag()
		if _, i.mOK = i.rParsedFlagDataMap[flagCandidate.ToFlag()]; i.mOK {
			return dollyerr.NewError(
				dollyerr.CodeArgParserDuplicateFlags,
				fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, flagCandidate))
		}
		i.rParsedFlagDataMap[flagCandidate.ToFlag()] = parsed_data.NewParsedFlagData(flagCandidate.ToFlag(), i.mParsedArgData)

		i.mParseState = getParseState(i.mArgDescription)
		return nil
	}

	i.mParsedArgData.ArgValues = append(i.mParsedArgData.ArgValues, argListCandidate.ToArgValue())

	return nil
}

func (i *ArgParserImpl) notSetArgValueCase(arg cmdArg.CmdArg) *dollyerr.Error {
	// current command line argument looks like a flag
	// let's check if it is a flag
	i.mTmpFlagDescription, i.mOK = i.flagDescriptions[apConf.Flag(arg)]
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

func getParseState(argumentsDescription *apConf.ArgumentsDescription) parseState {
	switch argumentsDescription.GetAmountType() {
	case apConf.ArgAmountTypeSingle:
		return parseStateReadingSingleArgument
	case apConf.ArgAmountTypeList:
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

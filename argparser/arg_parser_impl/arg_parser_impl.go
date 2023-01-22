package arg_parser_impl

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
	cmdArg "github.com/terryhay/dolly/argparser/cmd_arg"
	"github.com/terryhay/dolly/argparser/parsed_data"
	"github.com/terryhay/dolly/utils/dollyerr"
)

type parseState uint8

const (
	parseStateReadingFlag parseState = iota
	parseStateReadingSingleArgument
	parseStateReadingArgumentList
)

const countOfParseStates = int(parseStateReadingArgumentList) + 1

// ArgParserImpl implements command line argument parser
type ArgParserImpl struct {
	commandDescriptions        map[apConf.Command]*apConf.CommandDescription
	namelessCommandDescription *apConf.CommandDescription
	flagDescriptions           map[apConf.Flag]*apConf.FlagDescription
	stateProcessors            [countOfParseStates]func(arg cmdArg.CmdArg) *dollyerr.Error

	// mutable page
	mArgDescription     *apConf.ArgumentsDescription
	mFlag               apConf.Flag
	mFlagDescription    *apConf.FlagDescription
	mTmpFlagDescription *apConf.FlagDescription
	mOK                 bool
	mParsedArgData      *parsed_data.ParsedArgData
	mParseState         parseState

	// res page
	rParsedFlagDataMap map[apConf.Flag]*parsed_data.ParsedFlagData
}

// NewCmdArgParserImpl constructs ArgParserImpl object
func NewCmdArgParserImpl(config apConf.ArgParserConfig) *ArgParserImpl {
	res := &ArgParserImpl{
		commandDescriptions: createCommandDescriptionMap(
			config.GetCommandDescriptions(),
			config.GetHelpCommandDescription()),
		namelessCommandDescription: castNamelessCommandDescriptionToPointer(config.GetNamelessCommandDescription()),
		flagDescriptions:           createFlagDescriptionMap(config.GetFlagDescriptionSlice()),
	}
	res.stateProcessors = [countOfParseStates]func(arg cmdArg.CmdArg) *dollyerr.Error{
		parseStateReadingFlag:           res.processReadingFlag,
		parseStateReadingSingleArgument: res.processReadingSingleArgument,
		parseStateReadingArgumentList:   res.processReadingArgumentList,
	}

	return res
}

func createCommandDescriptionMap(
	commandDescriptions []*apConf.CommandDescription,
	helpCommandDescription apConf.HelpCommandDescription,
) map[apConf.Command]*apConf.CommandDescription {

	res := make(map[apConf.Command]*apConf.CommandDescription, len(commandDescriptions)*2)

	if desc := castHelpCommandDescriptionToPointer(helpCommandDescription); desc != nil {
		for command := range desc.GetCommands() {
			res[command] = desc
		}
	}

	for _, desc := range commandDescriptions {
		for command := range desc.GetCommands() {
			res[command] = desc
		}
	}

	return res
}

func castHelpCommandDescriptionToPointer(helpCommandDescription apConf.HelpCommandDescription) *apConf.CommandDescription {
	if helpCommandDescription == nil {
		return nil
	}
	return helpCommandDescription.(*apConf.CommandDescription)
}

func castNamelessCommandDescriptionToPointer(namelessCommandDescription apConf.NamelessCommandDescription) *apConf.CommandDescription {
	if namelessCommandDescription == nil {
		return nil
	}
	return namelessCommandDescription.(*apConf.CommandDescription)
}

func createFlagDescriptionMap(flagDescriptions []*apConf.FlagDescription) map[apConf.Flag]*apConf.FlagDescription {
	res := make(map[apConf.Flag]*apConf.FlagDescription, len(flagDescriptions)*2)
	for _, desc := range flagDescriptions {
		for _, flag := range desc.GetFlags() {
			res[flag] = desc
		}
	}

	return res
}

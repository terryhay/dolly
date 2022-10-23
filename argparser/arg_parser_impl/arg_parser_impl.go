package arg_parser_impl

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
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

// ArgParserImpl - implementation of command line argument parser
type ArgParserImpl struct {
	commandDescriptions        map[apConf.Command]*apConf.CommandDescription
	namelessCommandDescription *apConf.CommandDescription
	flagDescriptions           map[apConf.Flag]*apConf.FlagDescription
	stateProcessors            [countOfParseStates]func(arg string) *dollyerr.Error

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

// NewCmdArgParserImpl - ArgParserImpl object constructor
func NewCmdArgParserImpl(config apConf.ArgParserConfig) *ArgParserImpl {
	res := &ArgParserImpl{
		commandDescriptions: createCommandDescriptionMap(
			config.GetCommandDescriptions(),
			config.GetHelpCommandDescription()),
		namelessCommandDescription: castNamelessCommandDescriptionToPointer(config.GetNamelessCommandDescription()),
		flagDescriptions:           config.GetFlagDescriptions(),
	}
	res.stateProcessors = [countOfParseStates]func(arg string) *dollyerr.Error{
		parseStateReadingFlag:           res.processReadingFlag,
		parseStateReadingSingleArgument: res.processReadingSingleArgument,
		parseStateReadingArgumentList:   res.processReadingArgumentList,
	}

	return res
}

func createCommandDescriptionMap(
	commandsDescriptionSlice []*apConf.CommandDescription,
	helpCommandDescription apConf.HelpCommandDescription,
) map[apConf.Command]*apConf.CommandDescription {

	commandsCount := 0
	commandDescription := castHelpCommandDescriptionToPointer(helpCommandDescription)
	if commandDescription != nil {
		commandsCount++
	}
	commandsCount += len(commandsDescriptionSlice)

	res := make(map[apConf.Command]*apConf.CommandDescription, commandsCount)

	var command apConf.Command
	for command = range commandDescription.GetCommands() {
		res[command] = commandDescription
	}

	for i := range commandsDescriptionSlice {
		for command = range commandsDescriptionSlice[i].GetCommands() {
			res[command] = commandDescription
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

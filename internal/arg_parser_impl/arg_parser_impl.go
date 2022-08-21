package arg_parser_impl

import (
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/parsed_data"
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
	commandDescriptions        map[dollyconf.Command]*dollyconf.CommandDescription
	namelessCommandDescription *dollyconf.CommandDescription
	flagDescriptions           map[dollyconf.Flag]*dollyconf.FlagDescription
	stateProcessors            [countOfParseStates]func(arg string) *dollyerr.Error

	// mutable data
	mArgDescription     *dollyconf.ArgumentsDescription
	mErr                *dollyerr.Error
	mFlag               dollyconf.Flag
	mFlagDescription    *dollyconf.FlagDescription
	mTmpFlagDescription *dollyconf.FlagDescription
	mOK                 bool
	mParsedArgData      *parsed_data.ParsedArgData
	mParseState         parseState

	// res data
	rParsedFlagDataMap map[dollyconf.Flag]*parsed_data.ParsedFlagData
}

// NewCmdArgParserImpl - ArgParserImpl object constructor
func NewCmdArgParserImpl(config dollyconf.ArgParserConfig) *ArgParserImpl {
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
	commandsDescriptionSlice []*dollyconf.CommandDescription,
	helpCommandDescription dollyconf.HelpCommandDescription,
) map[dollyconf.Command]*dollyconf.CommandDescription {

	commandsCount := 0
	commandDescription := castHelpCommandDescriptionToPointer(helpCommandDescription)
	if commandDescription != nil {
		commandsCount++
	}
	commandsCount += len(commandsDescriptionSlice)

	res := make(map[dollyconf.Command]*dollyconf.CommandDescription, commandsCount)

	var command dollyconf.Command
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

func castHelpCommandDescriptionToPointer(helpCommandDescription dollyconf.HelpCommandDescription) *dollyconf.CommandDescription {
	if helpCommandDescription == nil {
		return nil
	}
	return helpCommandDescription.(*dollyconf.CommandDescription)
}

func castNamelessCommandDescriptionToPointer(namelessCommandDescription dollyconf.NamelessCommandDescription) *dollyconf.CommandDescription {
	if namelessCommandDescription == nil {
		return nil
	}
	return namelessCommandDescription.(*dollyconf.CommandDescription)
}

package arg_parser_config

import coty "github.com/terryhay/dolly/tools/common_types"

// ArgParserConfig contains specifications of namesAdditional, arguments and command groups of application
type ArgParserConfig struct {
	commandNameless *Command
	commands        []*Command
	commandHelpOut  *Command
	app             Application

	helpInfoChapterDESCRIPTION []coty.InfoChapterDESCRIPTION
}

// ArgParserConfigOpt contains source data for cast to ArgParserConfig
type ArgParserConfigOpt struct {
	CommandNameless *NamelessCommandOpt
	Commands        []*CommandOpt
	CommandHelpOut  *HelpOutCommandOpt
	App             ApplicationOpt

	HelpInfoChapterDESCRIPTION []string
}

// MakeArgParserConfig converts opt to ArgParserConfig object
func MakeArgParserConfig(opt ArgParserConfigOpt) ArgParserConfig {
	return ArgParserConfig{
		commandNameless: NewNamelessCommand(opt.CommandNameless),
		commands:        toCommandSlice(opt.Commands),
		commandHelpOut:  NewHelpOutCommand(opt.CommandHelpOut),
		app:             MakeApplication(opt.App),

		helpInfoChapterDESCRIPTION: coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.HelpInfoChapterDESCRIPTION),
	}
}

// CommandByName returns Command by CommandName
func (apc *ArgParserConfig) CommandByName(name coty.NameCommand) *Command {
	for _, command := range apc.GetCommands() {
		if name == command.GetNameMain() {
			return command
		}

		if _, contain := command.GetNamesAdditional()[name]; contain {
			return command
		}
	}

	if apc.GetCommandHelpOut() != nil {
		if name == apc.GetCommandHelpOut().GetNameMain() {
			return apc.GetCommandHelpOut()
		}

		if _, contain := apc.GetCommandHelpOut().GetNamesAdditional()[name]; contain {
			return apc.commandHelpOut
		}
	}

	return nil
}

// GetCommandNameless gets commandNameless field
func (apc *ArgParserConfig) GetCommandNameless() *Command {
	if apc == nil {
		return nil
	}
	return apc.commandNameless
}

// GetCommands gets commands field
func (apc *ArgParserConfig) GetCommands() []*Command {
	if apc == nil {
		return nil
	}
	return apc.commands
}

// GetCommandHelpOut gets commandHelpOut field
func (apc *ArgParserConfig) GetCommandHelpOut() *Command {
	if apc == nil {
		return nil
	}
	return apc.commandHelpOut
}

// GetAppDescription gets app field
func (apc *ArgParserConfig) GetAppDescription() Application {
	if apc == nil {
		return Application{}
	}
	return apc.app
}

// GetHelpInfoChapterDESCRIPTION gets infoChapterNAME field
func (apc *ArgParserConfig) GetHelpInfoChapterDESCRIPTION() []coty.InfoChapterDESCRIPTION {
	if apc == nil {
		return nil
	}
	return apc.helpInfoChapterDESCRIPTION
}

func toCommandSlice(opts []*CommandOpt) []*Command {
	if len(opts) == 0 {
		return nil
	}

	commands := make([]*Command, 0, len(opts))
	for _, opt := range opts {
		commands = append(commands, NewCommand(*opt))
	}
	return commands
}

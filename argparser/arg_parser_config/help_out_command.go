package arg_parser_config

import coty "github.com/terryhay/dolly/tools/common_types"

// HelpOutCommandOpt contains source data for cast to HelpCommandDescription
type HelpOutCommandOpt struct {
	NameMain        coty.NameCommand
	NamesAdditional map[coty.NameCommand]struct{}
}

// NewHelpOutCommand converts opt to HelpCommandDescription object
func NewHelpOutCommand(opt *HelpOutCommandOpt) *Command {
	if opt == nil {
		return nil
	}
	return &Command{
		nameMain:        opt.NameMain,
		namesAdditional: opt.NamesAdditional,
	}
}

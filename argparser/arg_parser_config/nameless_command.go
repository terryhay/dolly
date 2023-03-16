package arg_parser_config

// NamelessCommandOpt contains source data for cast to NamelessCommandDescription
type NamelessCommandOpt struct {
	Placeholders []*PlaceholderOpt
	HelpInfo     string
}

// NewNamelessCommand converts opt to NamelessCommandDescription pointer
func NewNamelessCommand(opt *NamelessCommandOpt) *Command {
	if opt == nil {
		return nil
	}

	return NewCommand(CommandOpt{
		Placeholders: opt.Placeholders,
		HelpInfo:     opt.HelpInfo,
	})
}

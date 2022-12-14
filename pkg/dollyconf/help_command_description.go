package dollyconf

// HelpCommandDescription - interface for specification of print help info command
type HelpCommandDescription interface {
	GetID() CommandID
	GetCommands() map[Command]bool
}

func NewHelpCommandDescription(id CommandID, commands map[Command]bool) HelpCommandDescription {
	return &CommandDescription{
		ID:       id,
		Commands: commands,
	}
}

package dollyconf

// CommandDescription contains specification of command group which contains flags and arguments
type CommandDescription struct {
	ID                  CommandID
	DescriptionHelpInfo string
	Commands            map[Command]bool
	ArgDescription      *ArgumentsDescription
	RequiredFlags       map[Flag]bool
	OptionalFlags       map[Flag]bool
}

// GetID - ID field getter
func (i *CommandDescription) GetID() CommandID {
	if i == nil {
		return CommandIDUndefined
	}
	return i.ID
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *CommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetCommands - Commands field getter
func (i *CommandDescription) GetCommands() map[Command]bool {
	if i == nil {
		return nil
	}
	return i.Commands
}

// GetArgDescription - ArgDescription field getter
func (i *CommandDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (i *CommandDescription) GetRequiredFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *CommandDescription) GetOptionalFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}

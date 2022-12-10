package arg_parser_config

import "unsafe"

// CommandDescription contains specification of command group which contains flags and arguments
type CommandDescription struct {
	id                  CommandID
	descriptionHelpInfo string
	commands            map[Command]bool
	argDescription      *ArgumentsDescription
	requiredFlags       map[Flag]bool
	optionalFlags       map[Flag]bool
}

// CommandDescriptionSrc contains source data for cast to CommandDescription
type CommandDescriptionSrc struct {
	ID                  CommandID
	DescriptionHelpInfo string
	Commands            map[Command]bool
	ArgDescription      *ArgumentsDescription
	RequiredFlags       map[Flag]bool
	OptionalFlags       map[Flag]bool
}

// CastPtr converts src to CommandDescription pointer
func (src CommandDescriptionSrc) CastPtr() *CommandDescription {
	return (*CommandDescription)(unsafe.Pointer(&src))
}

// GetID - ID field getter
func (i *CommandDescription) GetID() CommandID {
	if i == nil {
		return CommandIDUndefined
	}
	return i.id
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *CommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.descriptionHelpInfo
}

// GetCommands - Commands field getter
func (i *CommandDescription) GetCommands() map[Command]bool {
	if i == nil {
		return nil
	}
	return i.commands
}

// GetArgDescription - ArgDescription field getter
func (i *CommandDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.argDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (i *CommandDescription) GetRequiredFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.requiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *CommandDescription) GetOptionalFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.optionalFlags
}

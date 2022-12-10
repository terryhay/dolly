package arg_parser_config

// NamelessCommandDescription - interface for specification of a command without call name
type NamelessCommandDescription interface {
	GetID() CommandID
	GetDescriptionHelpInfo() string
	GetArgDescription() *ArgumentsDescription
	GetRequiredFlags() map[Flag]bool
	GetOptionalFlags() map[Flag]bool
}

func NewNamelessCommandDescription(
	id CommandID,
	descriptionHelpInfo string,
	argDescription *ArgumentsDescription,
	requiredFlags map[Flag]bool,
	optionalFlags map[Flag]bool,
) NamelessCommandDescription {
	return CommandDescriptionSrc{
		ID:                  id,
		DescriptionHelpInfo: descriptionHelpInfo,
		ArgDescription:      argDescription,
		RequiredFlags:       requiredFlags,
		OptionalFlags:       optionalFlags,
	}.CastPtr()
}

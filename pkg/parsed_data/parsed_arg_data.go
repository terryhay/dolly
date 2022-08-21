package parsed_data

// ParsedArgData - parsed argument values of a command or a flag
type ParsedArgData struct {
	ArgValues []ArgValue
}

func NewParsedArgData(argValues []ArgValue) *ParsedArgData {
	return &ParsedArgData{
		ArgValues: argValues,
	}
}

// GetArgValues - ArgValues field getter
func (i *ParsedArgData) GetArgValues() []ArgValue {
	if i == nil {
		return nil
	}
	return i.ArgValues
}

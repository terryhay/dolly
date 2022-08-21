package parsed_data

import "github.com/terryhay/dolly/pkg/dollyconf"

// ParsedFlagData - parsed flag and arguments array
type ParsedFlagData struct {
	Flag    dollyconf.Flag
	ArgData *ParsedArgData
}

// NewParsedFlagData - ParsedFlagData object constructor
func NewParsedFlagData(flag dollyconf.Flag, argData *ParsedArgData) *ParsedFlagData {
	return &ParsedFlagData{
		Flag:    flag,
		ArgData: argData,
	}
}

// GetFlag - flag field getter
func (i *ParsedFlagData) GetFlag() dollyconf.Flag {
	if i == nil {
		return dollyconf.FlagUndefined
	}
	return i.Flag
}

// GetArgData - ArgData field getter
func (i *ParsedFlagData) GetArgData() *ParsedArgData {
	if i == nil {
		return nil
	}
	return i.ArgData
}

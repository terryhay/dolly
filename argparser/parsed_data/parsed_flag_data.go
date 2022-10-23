package parsed_data

import apConf "github.com/terryhay/dolly/argparser/arg_parser_config"

// ParsedFlagData - parsed flag and arguments array
type ParsedFlagData struct {
	Flag    apConf.Flag
	ArgData *ParsedArgData
}

// NewParsedFlagData - ParsedFlagData object constructor
func NewParsedFlagData(flag apConf.Flag, argData *ParsedArgData) *ParsedFlagData {
	return &ParsedFlagData{
		Flag:    flag,
		ArgData: argData,
	}
}

// GetFlag - flag field getter
func (i *ParsedFlagData) GetFlag() apConf.Flag {
	if i == nil {
		return apConf.FlagUndefined
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

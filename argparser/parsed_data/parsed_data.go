package parsed_data

import (
	apConf "github.com/terryhay/dolly/argparser/arg_parser_config"
)

// ParsedData - all parsed Command line page
type ParsedData struct {
	CommandID   apConf.CommandID
	Command     apConf.Command
	ArgData     *ParsedArgData
	FlagDataMap map[apConf.Flag]*ParsedFlagData
}

// NewParsedData - ParsedData object constructor
func NewParsedData(
	commandID apConf.CommandID,
	command apConf.Command,
	argData *ParsedArgData,
	flagDataMap map[apConf.Flag]*ParsedFlagData,
) *ParsedData {
	if len(flagDataMap) == 0 {
		flagDataMap = nil
	}
	return &ParsedData{
		CommandID:   commandID,
		Command:     command,
		ArgData:     argData,
		FlagDataMap: flagDataMap,
	}
}

// GetCommandID - CommandID field getter
func (i *ParsedData) GetCommandID() apConf.CommandID {
	if i == nil {
		return apConf.CommandIDUndefined
	}
	return i.CommandID
}

// GetCommand - Command field getter
func (i *ParsedData) GetCommand() apConf.Command {
	if i == nil {
		return apConf.CommandUndefined
	}
	return i.Command
}

// GetAgrData - AgrData field getter
func (i *ParsedData) GetAgrData() *ParsedArgData {
	if i == nil {
		return nil
	}
	return i.ArgData
}

// GetFlagDataMap - FlagDataMap field getter
func (i *ParsedData) GetFlagDataMap() map[apConf.Flag]*ParsedFlagData {
	if i == nil {
		return nil
	}
	return i.FlagDataMap
}

// GetFlagArgValue - extract flag argument value
func (i *ParsedData) GetFlagArgValue(flag apConf.Flag) (value ArgValue, ok bool) {
	var values []ArgValue
	values, ok = i.GetFlagArgValues(flag)
	if !ok || len(values) == 0 {
		return value, false
	}

	return values[0], true
}

// GetFlagArgValues - extract flag argument value slice
func (i *ParsedData) GetFlagArgValues(flag apConf.Flag) (values []ArgValue, ok bool) {
	if i == nil {
		return nil, false
	}
	var parsedFlagData *ParsedFlagData
	parsedFlagData, ok = i.GetFlagDataMap()[flag]
	if !ok {
		return nil, false
	}

	return parsedFlagData.GetArgData().GetArgValues(), true
}

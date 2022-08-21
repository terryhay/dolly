package parsed_data

import (
	"github.com/terryhay/dolly/pkg/dollyconf"
)

// ParsedData - all parsed Command line data
type ParsedData struct {
	CommandID   dollyconf.CommandID
	Command     dollyconf.Command
	ArgData     *ParsedArgData
	FlagDataMap map[dollyconf.Flag]*ParsedFlagData
}

// NewParsedData - ParsedData object constructor
func NewParsedData(
	commandID dollyconf.CommandID,
	command dollyconf.Command,
	argData *ParsedArgData,
	flagDataMap map[dollyconf.Flag]*ParsedFlagData,
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
func (i *ParsedData) GetCommandID() dollyconf.CommandID {
	if i == nil {
		return dollyconf.CommandIDUndefined
	}
	return i.CommandID
}

// GetCommand - Command field getter
func (i *ParsedData) GetCommand() dollyconf.Command {
	if i == nil {
		return dollyconf.CommandUndefined
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
func (i *ParsedData) GetFlagDataMap() map[dollyconf.Flag]*ParsedFlagData {
	if i == nil {
		return nil
	}
	return i.FlagDataMap
}

// GetFlagArgValue - extract flag argument value
func (i *ParsedData) GetFlagArgValue(flag dollyconf.Flag) (value ArgValue, ok bool) {
	var values []ArgValue
	values, ok = i.GetFlagArgValues(flag)
	if !ok || len(values) == 0 {
		return value, false
	}

	return values[0], true
}

// GetFlagArgValues - extract flag argument value slice
func (i *ParsedData) GetFlagArgValues(flag dollyconf.Flag) (values []ArgValue, ok bool) {
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

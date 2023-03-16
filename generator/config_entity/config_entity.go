package config_entity

import (
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

// ConfigEntity provides generator config entity
type ConfigEntity struct {
	config *confYML.Config

	placeholdersByNames map[coty.NamePlaceholder]*confYML.Placeholder

	genCompCommandsByNames     map[coty.NameCommand]*GenComponents
	genCompPlaceholdersByNames map[coty.NamePlaceholder]*GenComponents
	genCompFlagsByNames        map[coty.NameFlag]*GenComponents

	genCompCommandsSorted    []*GenComponents
	genCompPlaceholderSorted []*GenComponents
	genCompFlagsSorted       []*GenComponents
}

// MakeConfigEntity constructs ConfigEntity object in stack
func MakeConfigEntity(config *confYML.Config) (ConfigEntity, error) {
	return makeConfigEntity(config)
}

// GetConfig gets config field
func (ce *ConfigEntity) GetConfig() *confYML.Config {
	if ce == nil {
		return nil
	}
	return ce.config
}

// GetGenCompCommandsSorted gets genCompCommandsSorted field
func (ce *ConfigEntity) GetGenCompCommandsSorted() []*GenComponents {
	if ce == nil {
		return nil
	}
	return ce.genCompCommandsSorted
}

// GetGenCompPlaceholdersSorted gets genCompPlaceholderSorted field
func (ce *ConfigEntity) GetGenCompPlaceholdersSorted() []*GenComponents {
	if ce == nil {
		return nil
	}
	return ce.genCompPlaceholderSorted
}

// GetGenCompFlagsSorted gets genCompFlagsSorted field
func (ce *ConfigEntity) GetGenCompFlagsSorted() []*GenComponents {
	if ce == nil {
		return nil
	}
	return ce.genCompFlagsSorted
}

// GenCompCommandByName returns command GenComponents by name
func (ce *ConfigEntity) GenCompCommandByName(name coty.NameCommand) *GenComponents {
	if ce == nil {
		return nil
	}

	return ce.genCompCommandsByNames[name]
}

// PlaceholderByName returns Placeholder by name
func (ce *ConfigEntity) PlaceholderByName(name coty.NamePlaceholder) *confYML.Placeholder {
	if ce == nil {
		return nil
	}

	return ce.placeholdersByNames[name]
}

// GenCompPlaceholderByName returns placeholder GenComponents by name
func (ce *ConfigEntity) GenCompPlaceholderByName(name coty.NamePlaceholder) *GenComponents {
	if ce == nil {
		return nil
	}

	return ce.genCompPlaceholdersByNames[name]
}

// GenCompFlagByName returns flag GenComponents by name
func (ce *ConfigEntity) GenCompFlagByName(name coty.NameFlag) *GenComponents {
	if ce == nil {
		return nil
	}

	return ce.genCompFlagsByNames[name]
}

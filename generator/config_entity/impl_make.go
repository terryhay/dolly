package config_entity

import (
	"errors"

	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

// ErrMakeConfigEntity - invalid config object
var ErrMakeConfigEntity = errors.New(`makeConfigEntity: invalid config object`)

func makeConfigEntity(config *confYML.Config) (ConfigEntity, error) {
	if err := config.IsValid(); err != nil {
		return ConfigEntity{}, errors.Join(ErrMakeConfigEntity, err)
	}

	placeholders := placeholdersByNames(config.GetArgParserConfig())
	genComps := createGenComponents(config.GetArgParserConfig())

	return ConfigEntity{
		config: config,

		placeholdersByNames: placeholders,

		genCompCommandsByNames:     genComps.genCompCommandsByNames,
		genCompPlaceholdersByNames: genComps.genCompPlaceholdersByNames,
		genCompFlagsByNames:        genComps.genCompFlagsByNames,

		genCompCommandsSorted:    genComps.genCompCommandsSorted,
		genCompPlaceholderSorted: genComps.genCompPlaceholdersSorted,
		genCompFlagsSorted:       genComps.genCompFlagsSorted,
	}, nil
}

func placeholdersByNames(conf *confYML.ArgParserConfig) map[coty.NamePlaceholder]*confYML.Placeholder {
	res := make(map[coty.NamePlaceholder]*confYML.Placeholder, len(conf.GetCommandsSorted())+len(conf.GetNamelessCommand().GetUsingPlaceholders()))
	for _, placeholder := range conf.GetPlaceholders() {
		if placeholder == nil {
			continue
		}
		res[placeholder.GetName()] = placeholder
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

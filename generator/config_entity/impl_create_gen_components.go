package config_entity

import (
	"fmt"
	"sort"
	"unicode"

	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

type createGenComponentsResult struct {
	genCompCommandsSorted  []*GenComponents
	genCompCommandsByNames map[coty.NameCommand]*GenComponents

	genCompPlaceholdersSorted  []*GenComponents
	genCompPlaceholdersByNames map[coty.NamePlaceholder]*GenComponents

	genCompFlagsSorted  []*GenComponents
	genCompFlagsByNames map[coty.NameFlag]*GenComponents
}

func createGenComponents(conf *confYML.ArgParserConfig) createGenComponentsResult {

	// 2: nameless and help commands
	// 2*len(commands): let's think every command has one additional name
	countCommandsExpected := 2 + 2*len(conf.GetCommandsSorted()) + len(conf.GetHelpCommand().GetAdditionalNamesSorted())

	genCompCommands := make([]*GenComponents, 0, countCommandsExpected)
	genCompCommandsByNames := make(map[coty.NameCommand]*GenComponents, countCommandsExpected)

	{
		if commandNameless := conf.GetNamelessCommand(); commandNameless != nil {
			genComponent := NewGenComponents(GenComponentsOpt{
				PrefixID: PrefixNameCommand,
				Name:     NamelessCommandIDPostfix,
				Comment:  commandNameless.GetChapterDescriptionInfo(),
			})

			genCompCommands = append(genCompCommands, genComponent)
			genCompCommandsByNames[genComponent.GetName().NameCommand()] = genComponent
		}

		for _, command := range conf.GetCommandsSorted() {
			genComponent := NewGenComponents(GenComponentsOpt{
				PrefixID: PrefixNameCommand,
				Name:     command.GetMainName(),
				Comment:  command.GetChapterDescriptionInfo(),
			})

			genCompCommands = append(genCompCommands, genComponent)
			genCompCommandsByNames[genComponent.GetName().NameCommand()] = genComponent

			for _, name := range command.GetAdditionalNames() {
				genComponent = NewGenComponents(GenComponentsOpt{
					PrefixID: PrefixNameCommand,
					Name:     name,
					Comment:  command.GetChapterDescriptionInfo(),
				})

				genCompCommands = append(genCompCommands, genComponent)
				genCompCommandsByNames[name] = genComponent
			}
		}

		// help command
		if commandHelp := conf.GetHelpCommand(); commandHelp != nil {
			genComponent := NewGenComponents(GenComponentsOpt{
				PrefixID: PrefixNameCommand,
				Name:     commandHelp.GetMainName(),
				Comment:  commandHelp.GetChapterDescriptionInfo(),
			})

			genCompCommands = append(genCompCommands, genComponent)
			genCompCommandsByNames[genComponent.GetName().NameCommand()] = genComponent

			for _, name := range commandHelp.GetAdditionalNamesSorted() {
				genComponent = NewGenComponents(GenComponentsOpt{
					PrefixID: PrefixNameCommand,
					Name:     name,
					Comment:  commandHelp.GetChapterDescriptionInfo(),
				})

				genCompCommands = append(genCompCommands, genComponent)
				genCompCommandsByNames[genComponent.GetName().NameCommand()] = genComponent
			}
		}

		if len(genCompCommandsByNames) == 0 {
			genCompCommands = nil
			genCompCommandsByNames = nil
		}

		sort.Slice(genCompCommands, func(i, j int) bool {
			return genCompCommands[i].GetNameID() < genCompCommands[j].GetNameID()
		})
	}

	genCompPlaceholders := make([]*GenComponents, 0, len(conf.GetPlaceholders()))
	genCompPlaceholdersByNames := make(map[coty.NamePlaceholder]*GenComponents, len(conf.GetPlaceholders()))

	// 2*len(conf.GetPlaceholders()): let's think every placeholder has two flags
	countFlagsExpected := 2 * len(conf.GetPlaceholders())

	genCompFlags := make([]*GenComponents, 0, countFlagsExpected)
	genCompFlagsByNames := make(map[coty.NameFlag]*GenComponents, countFlagsExpected)

	{
		for _, placeholder := range conf.GetPlaceholders() {
			genComponent := NewGenComponents(GenComponentsOpt{
				PrefixID: PrefixPlaceholderID,
				Name:     placeholder.GetName(),
				Comment:  coty.InfoChapterDESCRIPTION(placeholder.GetName().String()),
			})

			genCompPlaceholders = append(genCompPlaceholders, genComponent)
			genCompPlaceholdersByNames[genComponent.GetName().NamePlaceholder()] = genComponent

			for _, flag := range placeholder.GetFlags() {
				genComponent = NewGenComponents(GenComponentsOpt{
					PrefixID: PrefixFlagName,
					Name:     flag.GetMainName(),
					Comment:  flag.GetDescriptionHelpInfo(),
				})

				genCompFlags = append(genCompFlags, genComponent)
				genCompFlagsByNames[genComponent.GetName().NameFlag()] = genComponent

				for _, name := range flag.GetAdditionalNames() {
					genComponent = NewGenComponents(GenComponentsOpt{
						PrefixID: PrefixFlagName,
						Name:     name,
						Comment:  flag.GetDescriptionHelpInfo(),
					})

					genCompFlags = append(genCompFlags, genComponent)
					genCompFlagsByNames[genComponent.GetName().NameFlag()] = genComponent
				}
			}
		}

		if len(genCompPlaceholdersByNames) == 0 {
			genCompPlaceholders = nil
			genCompPlaceholdersByNames = nil
		}

		sort.Slice(genCompPlaceholders, func(i, j int) bool {
			return genCompPlaceholders[i].GetNameID() < genCompPlaceholders[j].GetNameID()
		})

		if len(genCompFlagsByNames) == 0 {
			genCompFlags = nil
			genCompFlagsByNames = nil
		}

		sort.Slice(genCompFlags, func(i, j int) bool {
			return genCompFlags[i].GetNameID() < genCompFlags[j].GetNameID()
		})
	}

	return createGenComponentsResult{
		genCompCommandsSorted:  genCompCommands,
		genCompCommandsByNames: genCompCommandsByNames,

		genCompPlaceholdersSorted:  genCompPlaceholders,
		genCompPlaceholdersByNames: genCompPlaceholdersByNames,

		genCompFlagsSorted:  genCompFlags,
		genCompFlagsByNames: genCompFlagsByNames,
	}
}

func createID(prefix, name fmt.Stringer) string {
	nameCamelCase := nameToCamelCaseRunes(name.String())

	switch {
	case len(nameCamelCase) == 0:
		return ""

	case len(nameCamelCase) == 1:
		const (
			postfixUp = "Up"
			postfixLw = "Lw"
		)
		if unicode.IsUpper([]rune(name.String())[1]) {
			return fmt.Sprintf("%s%s%s", prefix, string(unicode.ToUpper(nameCamelCase[0])), postfixUp)
		}
		return fmt.Sprintf("%s%s%s", prefix, string(unicode.ToUpper(nameCamelCase[0])), postfixLw)

	default:
		return fmt.Sprintf("%s%s%s", prefix, string(unicode.ToUpper(nameCamelCase[0])), string(nameCamelCase[1:]))
	}
}

func nameToCamelCaseRunes(name string) []rune {
	const (
		runeDash      = rune('-')
		runeDashLower = rune('_')
	)

	nameRunes := []rune(name)
	nameCamelCaseRunes := make([]rune, 0, len(nameRunes))

	stateFindingDashes := false
	for _, r := range nameRunes {
		switch {
		case stateFindingDashes:
			if r != runeDash && r != runeDashLower {
				nameCamelCaseRunes = append(nameCamelCaseRunes, r)
				continue
			}
			stateFindingDashes = false

		default:
			if r == runeDash || r == runeDashLower {
				continue
			}
			stateFindingDashes = true
			nameCamelCaseRunes = append(nameCamelCaseRunes, unicode.ToUpper(r))
		}
	}

	return nameCamelCaseRunes
}

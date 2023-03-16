package config_entity

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestCreateGenComponents(t *testing.T) {
	t.Parallel()

	t.Run("nil_config", func(t *testing.T) {
		require.Equal(t, createGenComponentsResult{}, createGenComponents(nil))
	})

	t.Run("commands", func(t *testing.T) {

		commandNameless := &confYML.NamelessCommandOpt{
			ChapterDescriptionInfo: "nameless command",
		}
		commandNamelessGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     NamelessCommandIDPostfix,
			Comment:  confYML.NewNamelessCommand(commandNameless).GetChapterDescriptionInfo(),
		})

		commandFirst := &confYML.CommandOpt{
			MainName:               coty.RandNameCommand().String(),
			ChapterDescriptionInfo: "first command description",
			AdditionalNames:        []string{coty.RandNameCommandSecond().String()},
		}
		commandFirstGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     confYML.NewCommand(commandFirst).GetMainName(),
			Comment:  confYML.NewCommand(commandFirst).GetChapterDescriptionInfo(),
		})
		commandFirstGenCompSecond := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     confYML.NewCommand(commandFirst).GetAdditionalNames()[0],
			Comment:  confYML.NewCommand(commandFirst).GetChapterDescriptionInfo(),
		})

		commandSecond := &confYML.CommandOpt{
			MainName:               coty.RandNameCommandThird().String(),
			ChapterDescriptionInfo: "second command description",
		}
		commandSecondGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     confYML.NewCommand(commandSecond).GetMainName(),
			Comment:  confYML.NewCommand(commandSecond).GetChapterDescriptionInfo(),
		})

		commandHelp := &confYML.HelpCommandOpt{
			MainName:        "--help",
			AdditionalNames: []string{"-h"},
		}
		commandHelpGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     confYML.NewHelpCommand(commandHelp).GetMainName(),
			Comment:  confYML.NewHelpCommand(commandHelp).GetChapterDescriptionInfo(),
		})
		commandHelpGenCompSecond := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixNameCommand,
			Name:     confYML.NewHelpCommand(commandHelp).GetAdditionalNamesSorted()[0],
			Comment:  confYML.NewHelpCommand(commandHelp).GetChapterDescriptionInfo(),
		})

		exp := createGenComponentsResult{
			genCompCommandsSorted: func() []*GenComponents {
				genCompCommands := []*GenComponents{
					commandNamelessGenComp,
					commandHelpGenComp,
					commandHelpGenCompSecond,
					commandFirstGenComp,
					commandFirstGenCompSecond,
					commandSecondGenComp,
				}
				sort.Slice(genCompCommands, func(i, j int) bool {
					return genCompCommands[i].GetNameID() < genCompCommands[j].GetNameID()
				})
				return genCompCommands
			}(),
			genCompCommandsByNames: map[coty.NameCommand]*GenComponents{
				commandNamelessGenComp.GetName().NameCommand():    commandNamelessGenComp,
				commandFirstGenComp.GetName().NameCommand():       commandFirstGenComp,
				commandFirstGenCompSecond.GetName().NameCommand(): commandFirstGenCompSecond,
				commandSecondGenComp.GetName().NameCommand():      commandSecondGenComp,
				commandHelpGenComp.GetName().NameCommand():        commandHelpGenComp,
				commandHelpGenCompSecond.GetName().NameCommand():  commandHelpGenCompSecond,
			},
		}
		res := createGenComponents(confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
			AppHelp: &confYML.AppHelpOpt{
				AppName:         "app",
				ChapterNameInfo: "some description for NAME chapter",
			},
			NamelessCommand: commandNameless,
			Commands: []*confYML.CommandOpt{
				commandFirst,
				commandSecond,
			},
			HelpCommand: commandHelp,
		}))

		require.Equal(t, exp, res)
	})

	t.Run("placeholders", func(t *testing.T) {

		placeholder := &confYML.PlaceholderOpt{
			Name: coty.NamePlaceholderUndefined.String(), // empty string for testing!,
		}
		placeholderGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixPlaceholderID,
			Name:     confYML.NewPlaceholder(placeholder).GetName(),
			Comment:  coty.InfoChapterDESCRIPTION(confYML.NewPlaceholder(placeholder).GetName().String()),
		})

		exp := createGenComponentsResult{
			genCompPlaceholdersSorted: func() []*GenComponents {
				genCompPlaceholders := []*GenComponents{
					placeholderGenComp,
				}
				sort.Slice(genCompPlaceholders, func(i, j int) bool {
					return genCompPlaceholders[i].GetNameID() < genCompPlaceholders[j].GetNameID()
				})
				return genCompPlaceholders
			}(),
			genCompPlaceholdersByNames: map[coty.NamePlaceholder]*GenComponents{
				placeholderGenComp.GetName().NamePlaceholder(): placeholderGenComp,
			},
		}
		res := createGenComponents(confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
			Placeholders: []*confYML.PlaceholderOpt{placeholder},
		}))

		require.Equal(t, exp, res)
	})

	t.Run("flags", func(t *testing.T) {

		flag := &confYML.FlagOpt{
			MainName: coty.RandNameFlagShort().String(),
			AdditionalNames: []string{
				coty.RandNameFlagLong().String(),
				coty.RandNameFlagOneLetter().String(),
			},
			ChapterDescriptionInfo: coty.RandInfoChapterDescription().String(),
		}

		flagShortGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixFlagName,
			Name:     confYML.NewFlag(flag).GetMainName(),
			Comment:  confYML.NewFlag(flag).GetDescriptionHelpInfo(),
		})
		flagLongGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixFlagName,
			Name:     confYML.NewFlag(flag).GetAdditionalNames()[0],
			Comment:  confYML.NewFlag(flag).GetDescriptionHelpInfo(),
		})
		flagOneLetterGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixFlagName,
			Name:     confYML.NewFlag(flag).GetAdditionalNames()[1],
			Comment:  confYML.NewFlag(flag).GetDescriptionHelpInfo(),
		})

		placeholder := &confYML.PlaceholderOpt{
			Name: coty.RandNamePlaceholder().String(),
			Flags: []*confYML.FlagOpt{
				flag,
			},
		}
		placeholderGenComp := NewGenComponents(GenComponentsOpt{
			PrefixID: PrefixPlaceholderID,
			Name:     confYML.NewPlaceholder(placeholder).GetName(),
			Comment:  coty.InfoChapterDESCRIPTION(confYML.NewPlaceholder(placeholder).GetName().String()),
		})

		exp := createGenComponentsResult{
			genCompPlaceholdersSorted: []*GenComponents{
				placeholderGenComp,
			},
			genCompPlaceholdersByNames: map[coty.NamePlaceholder]*GenComponents{
				placeholderGenComp.GetName().NamePlaceholder(): placeholderGenComp,
			},

			genCompFlagsSorted: func() []*GenComponents {
				genCompFlags := []*GenComponents{
					flagShortGenComp,
					flagLongGenComp,
					flagOneLetterGenComp,
				}
				sort.Slice(genCompFlags, func(i, j int) bool {
					return genCompFlags[i].GetNameID() < genCompFlags[j].GetNameID()
				})
				return genCompFlags
			}(),
			genCompFlagsByNames: map[coty.NameFlag]*GenComponents{
				flagShortGenComp.GetName().NameFlag():     flagShortGenComp,
				flagLongGenComp.GetName().NameFlag():      flagLongGenComp,
				flagOneLetterGenComp.GetName().NameFlag(): flagOneLetterGenComp,
			},
		}
		res := createGenComponents(confYML.NewArgParserConfig(&confYML.ArgParserConfigOpt{
			Placeholders: []*confYML.PlaceholderOpt{placeholder},
		}))

		require.Equal(t, exp, res)
	})
}

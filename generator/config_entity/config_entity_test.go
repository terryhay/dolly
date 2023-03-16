package config_entity

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestConfigEntityNilPointer(t *testing.T) {
	t.Parallel()

	_, err := MakeConfigEntity(nil)
	require.ErrorIs(t, err, ErrMakeConfigEntity)

	var pointer *ConfigEntity

	require.Nil(t, pointer.GetConfig())
	require.Nil(t, pointer.GetGenCompCommandsSorted())
	require.Nil(t, pointer.GenCompCommandByName(coty.RandNameCommand()))
	require.Nil(t, pointer.GenCompFlagByName(coty.RandNameFlagShort()))
	require.Nil(t, pointer.GetGenCompFlagsSorted())
	require.Nil(t, pointer.GetGenCompPlaceholdersSorted())
	require.Nil(t, pointer.PlaceholderByName(coty.RandNamePlaceholder()))
	require.Nil(t, pointer.GenCompPlaceholderByName(coty.RandNamePlaceholder()))
}

func TestConfigEntityMethods(t *testing.T) {
	t.Parallel()

	flagFirst := &confYML.FlagOpt{
		MainName:               coty.RandNameFlagShort().String(),
		ChapterDescriptionInfo: "flag description",
	}
	flagFirstGenComp := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixFlagName,
		Name:     confYML.NewFlag(flagFirst).GetMainName(),
		Comment:  confYML.NewFlag(flagFirst).GetDescriptionHelpInfo(),
	})
	flagSecond := &confYML.FlagOpt{
		MainName: coty.RandNameFlagShortSecond().String(),
		AdditionalNames: []string{
			coty.RandNameFlagShortThird().String(),
		},
		ChapterDescriptionInfo: "flag description",
	}
	flagSecondGenComp := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixFlagName,
		Name:     confYML.NewFlag(flagSecond).GetMainName(),
		Comment:  confYML.NewFlag(flagSecond).GetDescriptionHelpInfo(),
	})
	flagSecondGenCompSecond := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixFlagName,
		Name:     confYML.NewFlag(flagSecond).GetAdditionalNames()[0],
		Comment:  confYML.NewFlag(flagSecond).GetDescriptionHelpInfo(),
	})

	placeholderFirst := &confYML.PlaceholderOpt{
		Name: coty.RandNamePlaceholder().String(),
		Flags: []*confYML.FlagOpt{
			flagFirst,
		},
	}
	placeholderFirstGenComp := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixPlaceholderID,
		Name:     confYML.NewPlaceholder(placeholderFirst).GetName(),
		Comment:  coty.InfoChapterDESCRIPTION(confYML.NewPlaceholder(placeholderFirst).GetName().String()),
	})
	placeholderSecond := &confYML.PlaceholderOpt{
		Name: coty.RandNamePlaceholderSecond().String(),
		Flags: []*confYML.FlagOpt{
			flagSecond,
		},
	}
	placeholderSecondGenComp := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixPlaceholderID,
		Name:     confYML.NewPlaceholder(placeholderSecond).GetName(),
		Comment:  coty.InfoChapterDESCRIPTION(confYML.NewPlaceholder(placeholderSecond).GetName().String()),
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

	commandNameless := &confYML.NamelessCommandOpt{
		ChapterDescriptionInfo: "nameless command",
		UsingPlaceholders: []string{
			placeholderFirst.Name,
			placeholderSecond.Name,
		},
	}
	commandNamelessGenComp := NewGenComponents(GenComponentsOpt{
		PrefixID: PrefixNameCommand,
		Name:     NamelessCommandIDPostfix,
		Comment:  confYML.NewNamelessCommand(commandNameless).GetChapterDescriptionInfo(),
	})

	config := confYML.NewConfig(&confYML.ConfigOpt{
		Version: "1.0.0",
		ArgParserConfig: &confYML.ArgParserConfigOpt{
			AppHelp: &confYML.AppHelpOpt{
				AppName:         "app",
				ChapterNameInfo: "some description for NAME chapter",
			},
			Placeholders: []*confYML.PlaceholderOpt{
				placeholderFirst,
				placeholderSecond,
			},
			NamelessCommand: commandNameless,
			Commands: []*confYML.CommandOpt{
				commandFirst,
			},
			HelpCommand: commandHelp,
		},
	})

	t.Run("IsArgParserConfigValid", func(t *testing.T) {
		_, err := MakeConfigEntity(config)
		require.NoError(t, err)
	})

	t.Run("GetConfig", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		require.Equal(t, config, configEntity.GetConfig())
	})

	t.Run("GetGenCompCommandsSorted", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		genCompCommands := []*GenComponents{
			commandNamelessGenComp,
			commandHelpGenComp,
			commandHelpGenCompSecond,
			commandFirstGenComp,
			commandFirstGenCompSecond,
		}
		sort.Slice(genCompCommands, func(i, j int) bool {
			return genCompCommands[i].GetNameID() < genCompCommands[j].GetNameID()
		})

		require.Equal(t, genCompCommands, configEntity.GetGenCompCommandsSorted())
	})

	t.Run("GetGenCompPlaceholdersSorted", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		genCompCommands := []*GenComponents{
			placeholderFirstGenComp,
			placeholderSecondGenComp,
		}
		sort.Slice(genCompCommands, func(i, j int) bool {
			return genCompCommands[i].GetNameID() < genCompCommands[j].GetNameID()
		})

		require.Equal(t, genCompCommands, configEntity.GetGenCompPlaceholdersSorted())
	})

	t.Run("GetGenCompFlagsSorted", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		genCompCommands := []*GenComponents{
			flagFirstGenComp,
			flagSecondGenComp,
			flagSecondGenCompSecond,
		}
		sort.Slice(genCompCommands, func(i, j int) bool {
			return genCompCommands[i].GetNameID() < genCompCommands[j].GetNameID()
		})

		require.Equal(t, genCompCommands, configEntity.GetGenCompFlagsSorted())
	})

	t.Run("GenCompCommandByName", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		require.Equal(t, commandFirstGenComp, configEntity.GenCompCommandByName(confYML.NewCommand(commandFirst).GetMainName()))
		require.Equal(t, commandFirstGenCompSecond, configEntity.GenCompCommandByName(confYML.NewCommand(commandFirst).GetAdditionalNames()[0]))

		require.Equal(t, commandHelpGenComp, configEntity.GenCompCommandByName(confYML.NewHelpCommand(commandHelp).GetMainName()))
		require.Equal(t, commandHelpGenCompSecond, configEntity.GenCompCommandByName(confYML.NewHelpCommand(commandHelp).GetAdditionalNamesSorted()[0]))
	})

	t.Run("PlaceholderByName", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		require.Equal(t, confYML.NewPlaceholder(placeholderFirst), configEntity.PlaceholderByName(confYML.NewPlaceholder(placeholderFirst).GetName()))
		require.Equal(t, confYML.NewPlaceholder(placeholderSecond), configEntity.PlaceholderByName(confYML.NewPlaceholder(placeholderSecond).GetName()))
	})

	t.Run("GenCompPlaceholderByName", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		require.Equal(t, placeholderFirstGenComp, configEntity.GenCompPlaceholderByName(confYML.NewPlaceholder(placeholderFirst).GetName()))
		require.Equal(t, placeholderSecondGenComp, configEntity.GenCompPlaceholderByName(confYML.NewPlaceholder(placeholderSecond).GetName()))
	})

	t.Run("GenCompFlagByName", func(t *testing.T) {
		configEntity, err := MakeConfigEntity(config)
		require.NoError(t, err)

		require.Equal(t, flagFirstGenComp, configEntity.GenCompFlagByName(confYML.NewFlag(flagFirst).GetMainName()))

		require.Equal(t, flagSecondGenComp, configEntity.GenCompFlagByName(confYML.NewFlag(flagSecond).GetMainName()))
		require.Equal(t, flagSecondGenCompSecond, configEntity.GenCompFlagByName(confYML.NewFlag(flagSecond).GetAdditionalNames()[0]))
	})
}

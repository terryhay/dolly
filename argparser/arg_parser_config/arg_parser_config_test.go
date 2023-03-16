package arg_parser_config

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		t.Parallel()

		var pointer *ArgParserConfig

		require.Nil(t, pointer.GetCommandNameless())
		require.Nil(t, pointer.GetCommands())
		require.Nil(t, pointer.GetCommandHelpOut())
		require.Equal(t, pointer.GetAppDescription(), Application{})
		require.Nil(t, pointer.GetHelpInfoChapterDESCRIPTION())
		require.Nil(t, pointer.CommandByName(coty.RandNameCommand()))
	})

	t.Run("empty_opt", func(t *testing.T) {
		t.Parallel()

		opt := ArgParserConfigOpt{}
		argParserConfig := MakeArgParserConfig(opt)

		require.Nil(t, argParserConfig.GetCommandNameless())
		require.Nil(t, argParserConfig.GetCommands())
		require.Equal(t, NewHelpOutCommand(opt.CommandHelpOut), argParserConfig.GetCommandHelpOut())
		require.Equal(t, argParserConfig.GetAppDescription(), Application{})
		require.Nil(t, argParserConfig.GetHelpInfoChapterDESCRIPTION())
		require.Nil(t, argParserConfig.CommandByName(coty.RandNameCommand()))
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		t.Parallel()

		opt := ArgParserConfigOpt{
			CommandNameless: &NamelessCommandOpt{},
			Commands: []*CommandOpt{
				{
					NameMain: coty.RandNameCommand(),
					NamesAdditional: map[coty.NameCommand]struct{}{
						coty.RandNameCommandSecond(): {},
					},
				},
				{
					NameMain: coty.RandNameCommandThird(),
				},
			},
			CommandHelpOut: &HelpOutCommandOpt{
				NameMain: coty.RandNameCommandFourth(),
				NamesAdditional: map[coty.NameCommand]struct{}{
					coty.RandNameCommandFifth(): {},
				},
			},
			App: ApplicationOpt{},
			HelpInfoChapterDESCRIPTION: []string{
				gofakeit.Name(),
				gofakeit.Name(),
			},
		}
		pointer := MakeArgParserConfig(opt)

		require.Equal(t, NewNamelessCommand(opt.CommandNameless), pointer.GetCommandNameless())
		require.Equal(t, toCommandSlice(opt.Commands), pointer.GetCommands())
		require.Equal(t, NewHelpOutCommand(opt.CommandHelpOut), pointer.GetCommandHelpOut())
		require.Equal(t, MakeApplication(opt.App), pointer.GetAppDescription())
		require.Equal(t, coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.HelpInfoChapterDESCRIPTION), pointer.GetHelpInfoChapterDESCRIPTION())

		for _, command := range pointer.GetCommands() {
			require.Equal(t, command, pointer.CommandByName(command.GetNameMain()))

			for name := range command.GetNamesAdditional() {
				require.Equal(t, command, pointer.CommandByName(name))
			}
		}

		command := pointer.GetCommandHelpOut()
		require.Equal(t, command, pointer.CommandByName(command.GetNameMain()))

		for name := range command.GetNamesAdditional() {
			require.Equal(t, command, pointer.CommandByName(name))
		}
	})
}

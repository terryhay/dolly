package config_yaml

import (
	"testing"

	osd "github.com/terryhay/dolly/generator/proxyes/os_proxy"

	"github.com/stretchr/testify/require"
	coty "github.com/terryhay/dolly/tools/common_types"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer", func(t *testing.T) {
		var opt *ArgParserConfigOpt
		pointer := NewArgParserConfig(opt)
		require.Nil(t, pointer)

		require.ErrorIs(t, pointer.IsValid(), ErrArgParserConfigNilPointer)
		require.Nil(t, pointer.GetAppHelp())
		require.Nil(t, pointer.GetHelpCommand())
		require.Nil(t, pointer.GetNamelessCommand())
		require.Nil(t, pointer.GetCommandsSorted())
		require.Nil(t, pointer.GetPlaceholders())
		require.Nil(t, pointer.GetChapterDescriptionInfo())

	})

	t.Run("initialized_pointer", func(t *testing.T) {
		opt := ArgParserConfigOpt{
			AppHelp:                &AppHelpOpt{},
			HelpCommand:            &HelpCommandOpt{},
			NamelessCommand:        &NamelessCommandOpt{},
			Commands:               []*CommandOpt{{}},
			InfoChapterDESCRIPTION: []string{coty.RandInfoChapterDescription().String()},
		}
		pointer := NewArgParserConfig(&opt)

		require.Equal(t, NewAppHelp(opt.AppHelp), pointer.GetAppHelp())
		require.Equal(t, NewHelpCommand(opt.HelpCommand), pointer.GetHelpCommand())
		require.Equal(t, NewNamelessCommand(opt.NamelessCommand), pointer.GetNamelessCommand())
		require.Equal(t, toCommandSortedSlice(opt.Commands), pointer.GetCommandsSorted())
		require.Equal(t, toPlaceholderSlice(opt.Placeholders), pointer.GetPlaceholders())
		require.Equal(t, coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.InfoChapterDESCRIPTION), pointer.GetChapterDescriptionInfo())
	})
}

func TestArgParserConfigIsValidErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
		expError error
	}{
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_no_arg_parser.yaml",
			expError: ErrArgParserConfigNilPointer,
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_no_any_command.yaml",
			expError: ErrArgParserConfigNoAnyCommand,
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_duplicate_command_name.yaml",
			expError: ErrArgParserConfigDuplicateCommandName,
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_duplicate_command_additional_name.yaml",
			expError: ErrArgParserConfigDuplicateCommandName,
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_duplicate_flag_name.yaml",
			expError: ErrArgParserConfigDuplicateFlagName,
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/err_duplicate_flag_additional_name.yaml",
			expError: ErrArgParserConfigDuplicateFlagName,
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NotNil(t, config)
			require.NoError(t, err)

			err = config.IsValid()
			require.ErrorIs(t, err, tc.expError)
		})
	}
}

func TestArgParserConfigIsValidSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		caseFile string
	}{
		{
			caseFile: "./test_cases/arg_parser_config_cases/no_placeholders.yaml",
		},
		{
			caseFile: "./test_cases/arg_parser_config_cases/has_only_nameless_command.yaml",
		},
	}

	for _, tc := range tests {
		t.Run(tc.caseFile, func(t *testing.T) {
			config, err := Load(osd.New(), tc.caseFile)
			require.NotNil(t, config)
			require.NoError(t, err)

			require.NoError(t, config.IsValid())
		})
	}
}

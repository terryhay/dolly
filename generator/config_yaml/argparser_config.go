package config_yaml

import (
	"errors"
	"fmt"
	"sort"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// ArgParserConfig - argument parser configuration
type ArgParserConfig struct {
	appHelp     *AppHelp
	helpCommand *HelpCommand

	// one or more of these field must be set
	namelessCommand *NamelessCommand
	commandsSorted  []*Command

	// optional
	placeholders           []*Placeholder
	infoChapterDESCRIPTION []coty.InfoChapterDESCRIPTION
}

// GetAppHelp gets appHelp field
func (apc *ArgParserConfig) GetAppHelp() *AppHelp {
	if apc == nil {
		return nil
	}
	return apc.appHelp
}

// GetHelpCommand gets helpCommand field
func (apc *ArgParserConfig) GetHelpCommand() *HelpCommand {
	if apc == nil {
		return nil
	}
	return apc.helpCommand
}

// GetNamelessCommand gets namelessCommand field
func (apc *ArgParserConfig) GetNamelessCommand() *NamelessCommand {
	if apc == nil {
		return nil
	}
	return apc.namelessCommand
}

// GetCommandsSorted gets commandsSorted field
func (apc *ArgParserConfig) GetCommandsSorted() []*Command {
	if apc == nil {
		return nil
	}
	return apc.commandsSorted
}

// GetPlaceholders gets placeholders field
func (apc *ArgParserConfig) GetPlaceholders() []*Placeholder {
	if apc == nil {
		return nil
	}
	return apc.placeholders
}

// GetChapterDescriptionInfo gets infoChapterDESCRIPTION field
func (apc *ArgParserConfig) GetChapterDescriptionInfo() []coty.InfoChapterDESCRIPTION {
	if apc == nil {
		return nil
	}
	return apc.infoChapterDESCRIPTION
}

var (
	// ErrArgParserConfigNilPointer - ArgParserConfig pointer is nil
	ErrArgParserConfigNilPointer = errors.New(`'arg_parser' field is required in yaml source file`)

	// ErrArgParserConfigNoAnyCommand - no neither 'nameless_command' nor 'commands' fields in source yaml file
	ErrArgParserConfigNoAnyCommand = errors.New(`one or more of fields 'arg_parser.nameless_command' or 'arg_parser.commands' must be set`)

	// ErrArgParserConfigDuplicateCommandName - some command names are duplicates
	ErrArgParserConfigDuplicateCommandName = errors.New(`duplicate command name`)

	// ErrArgParserConfigDuplicateFlagName - some flag names are duplicates
	ErrArgParserConfigDuplicateFlagName = errors.New(`duplicate flag name`)
)

// IsValid cascade checks if ArgParserConfig is valid
func (apc *ArgParserConfig) IsValid() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ArgParserConfig.IsValid: %w", err)
		}
	}()

	if apc == nil {
		return ErrArgParserConfigNilPointer
	}

	if err = apc.appHelp.IsValid(); err != nil {
		return err
	}

	if err = apc.helpCommand.IsValid(); err != nil {
		return err
	}

	if apc.namelessCommand == nil && len(apc.commandsSorted) == 0 {
		return ErrArgParserConfigNoAnyCommand
	}
	if err = apc.namelessCommand.IsValid(); apc.namelessCommand != nil && err != nil {
		return err
	}

	usedCommandNames := make(map[coty.NameCommand]struct{}, 2*len(apc.commandsSorted))
	for _, command := range apc.commandsSorted {
		if err = command.IsValid(); err != nil {
			return err
		}

		if _, duplicate := usedCommandNames[command.GetMainName()]; duplicate {
			err = fmt.Errorf(`%w: '%s'`, ErrArgParserConfigDuplicateCommandName, command.GetMainName())
			return err
		}
		usedCommandNames[command.GetMainName()] = struct{}{}

		for _, name := range command.GetAdditionalNames() {
			if _, duplicate := usedCommandNames[name]; duplicate {
				err = fmt.Errorf(`%w: '%s'`, ErrArgParserConfigDuplicateCommandName, name)
				return err
			}
			usedCommandNames[name] = struct{}{}
		}
	}

	usedFlagNames := make(map[coty.NameFlag]struct{}, 2*len(apc.placeholders))
	for _, placeholder := range apc.placeholders {
		if err = placeholder.IsValid(); err != nil {
			return err
		}

		for _, flag := range placeholder.GetFlags() {
			if _, duplicate := usedFlagNames[flag.GetMainName()]; duplicate {
				err = fmt.Errorf(`%w: '%s'`, ErrArgParserConfigDuplicateFlagName, flag.GetMainName())
				return err
			}
			usedFlagNames[flag.GetMainName()] = struct{}{}

			for _, name := range flag.GetAdditionalNames() {
				if _, duplicate := usedFlagNames[name]; duplicate {
					err = fmt.Errorf(`%w: '%s'`, ErrArgParserConfigDuplicateFlagName, name)
					return err
				}
				usedFlagNames[name] = struct{}{}
			}
		}
	}

	return nil
}

// ArgParserConfigOpt - source for construct ArgParserConfig object
type ArgParserConfigOpt struct {
	AppHelp     *AppHelpOpt     `yaml:"app_help"`
	HelpCommand *HelpCommandOpt `yaml:"help_command"`

	// one or more of these field must be set
	NamelessCommand *NamelessCommandOpt `yaml:"nameless_command"`
	Commands        []*CommandOpt       `yaml:"commands"`

	// optional
	Placeholders           []*PlaceholderOpt `yaml:"placeholders"`
	InfoChapterDESCRIPTION []string          `yaml:"chapter_description_info"`
}

// NewArgParserConfig converts opt to AppHelp pointer
func NewArgParserConfig(opt *ArgParserConfigOpt) *ArgParserConfig {
	if opt == nil {
		return nil
	}

	return &ArgParserConfig{
		appHelp:     NewAppHelp(opt.AppHelp),
		helpCommand: NewHelpCommand(opt.HelpCommand),

		namelessCommand: NewNamelessCommand(opt.NamelessCommand),
		commandsSorted:  toCommandSortedSlice(opt.Commands),

		placeholders:           toPlaceholderSlice(opt.Placeholders),
		infoChapterDESCRIPTION: coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.InfoChapterDESCRIPTION),
	}
}

func toCommandSortedSlice(from []*CommandOpt) []*Command {
	if len(from) == 0 {
		return nil
	}

	to := make([]*Command, 0, len(from))
	for _, opt := range from {
		to = append(to, NewCommand(opt))
	}

	sort.Slice(to, func(l, r int) bool {
		return to[l].GetMainName() < to[r].GetMainName()
	})

	return to
}

func toPlaceholderSlice(from []*PlaceholderOpt) []*Placeholder {
	if len(from) == 0 {
		return nil
	}

	to := make([]*Placeholder, 0, len(from))
	for _, opt := range from {
		to = append(to, NewPlaceholder(opt))
	}

	return to
}

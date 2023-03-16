package config_yaml

import (
	"errors"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// HelpCommand - special description of a help mainName
type HelpCommand struct {
	mainName coty.NameCommand

	// optional
	additionalNamesSorted []coty.NameCommand
}

// GetMainName gets mainName field
func (hc *HelpCommand) GetMainName() coty.NameCommand {
	if hc == nil {
		return coty.NameCommandUndefined
	}
	return hc.mainName
}

// GetAdditionalNamesSorted gets additionalNamesSorted field
func (hc *HelpCommand) GetAdditionalNamesSorted() []coty.NameCommand {
	if hc == nil {
		return nil
	}
	return hc.additionalNamesSorted
}

// GetChapterDescriptionInfo returns infoChapterDESCRIPTION field
func (hc *HelpCommand) GetChapterDescriptionInfo() coty.InfoChapterDESCRIPTION {
	if hc == nil {
		return coty.InfoChapterDESCRIPTIONUndefined
	}
	const chapterDescriptionInfo = "print help info"
	return chapterDescriptionInfo
}

var (
	// ErrHelpCommandNilPointer - no required field 'arg_parser.help_command' in source yaml file
	ErrHelpCommandNilPointer = errors.New(`HelpCommand.IsValid: no required field 'arg_parser.help_command' in source yaml file`)

	// ErrHelpCommandMainName - invalid field 'arg_parser.help_command.main_name'
	ErrHelpCommandMainName = errors.New(`HelpCommand.IsValid: invalid field 'arg_parser.help_command.main_name'`)

	// ErrHelpCommandAdditionalNames - invalid element of field 'arg_parser.help_command.additional_names
	ErrHelpCommandAdditionalNames = errors.New(`HelpCommand.IsValid: invalid element of field 'arg_parser.help_command.additional_names'`)
)

// IsValid checks if HelpCommand is valid
func (hc *HelpCommand) IsValid() error {
	if hc == nil {
		return ErrHelpCommandNilPointer
	}

	if err := hc.GetMainName().IsValid(true); err != nil {
		return errors.Join(ErrHelpCommandMainName, err)
	}

	for _, name := range hc.GetAdditionalNamesSorted() {
		if err := name.IsValid(true); err != nil {
			return errors.Join(ErrHelpCommandAdditionalNames, err)
		}
	}

	return nil
}

// HelpCommandOpt -source for construct a special description of a help mainName
type HelpCommandOpt struct {
	MainName string `yaml:"main_name"`

	// optional
	AdditionalNames []string `yaml:"additional_names"`
}

// NewHelpCommand converts opt to HelpCommand pointer
func NewHelpCommand(opt *HelpCommandOpt) *HelpCommand {
	if opt == nil {
		return nil
	}

	return &HelpCommand{
		mainName:              coty.NameCommand(opt.MainName),
		additionalNamesSorted: coty.ToSliceTypesSorted[coty.NameCommand](opt.AdditionalNames),
	}
}

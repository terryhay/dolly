package config_yaml

import (
	"errors"
	"fmt"
	"sort"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Command - description of a mainName line mainName (which can contain flags and arguments)
type Command struct {
	mainName               coty.NameCommand
	chapterDescriptionInfo coty.InfoChapterDESCRIPTION

	// optional
	usingPlaceholdersSorted []coty.NamePlaceholder
	additionalNames         []coty.NameCommand
}

// GetMainName gets mainName field
func (c *Command) GetMainName() coty.NameCommand {
	if c == nil {
		return coty.NameCommandUndefined
	}
	return c.mainName
}

// GetAdditionalNames gets additionalNames field
func (c *Command) GetAdditionalNames() []coty.NameCommand {
	if c == nil {
		return nil
	}
	return c.additionalNames
}

// GetUsingPlaceholdersSorted gets usingPlaceholdersSorted field
func (c *Command) GetUsingPlaceholdersSorted() []coty.NamePlaceholder {
	if c == nil {
		return nil
	}
	return c.usingPlaceholdersSorted
}

// GetChapterDescriptionInfo gets infoChapterDESCRIPTION field
func (c *Command) GetChapterDescriptionInfo() coty.InfoChapterDESCRIPTION {
	if c == nil {
		return coty.InfoChapterDESCRIPTIONUndefined
	}
	return c.chapterDescriptionInfo
}

var (
	// ErrCommandMainName - 'arg_parser.commandsSorted' element has invalid 'main_name'
	ErrCommandMainName = errors.New(`'arg_parser.commandsSorted' element has invalid 'main_name' in source yaml file`)

	// ErrCommandNoChapterDescriptionInfo - 'arg_parser.commandsSorted' element must have 'chapter_description_info' in source yaml file
	ErrCommandNoChapterDescriptionInfo = errors.New(`'arg_parser.commandsSorted' element must have 'chapter_description_info' in source yaml file`)

	// ErrCommandAdditionalNames - 'arg_parser.commandsSorted' element has invalid element of 'additional_names'
	ErrCommandAdditionalNames = errors.New(`'arg_parser.commandsSorted' element has invalid element of 'additional_names' in source yaml file`)
)

// IsValid checks if Command is valid
func (c *Command) IsValid() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Command.IsValid: %w", err)
		}
	}()

	if err = c.GetMainName().IsValid(false); err != nil {
		return errors.Join(ErrCommandMainName, err)
	}

	if len(c.GetChapterDescriptionInfo()) == 0 {
		return fmt.Errorf(`%w: command main_name '%s'`, ErrCommandNoChapterDescriptionInfo, c.GetMainName())
	}

	for _, nameCommand := range c.GetAdditionalNames() {
		if err = nameCommand.IsValid(false); err != nil {
			return errors.Join(fmt.Errorf(`%w: command main_name '%s'`, ErrCommandAdditionalNames, c.GetMainName()), err)
		}
	}

	return nil
}

// CommandOpt - source for construct a description of a mainName line mainName
// (which can contain flags and arguments)
type CommandOpt struct {
	MainName               string `yaml:"main_name"`
	ChapterDescriptionInfo string `yaml:"chapter_description_info"`

	// optional
	UsingPlaceholders []string `yaml:"using_placeholders"`
	AdditionalNames   []string `yaml:"additional_names"`
}

// NewCommand converts opt to CommandDescription pointer
func NewCommand(opt *CommandOpt) *Command {
	if opt == nil {
		return nil
	}

	sort.Strings(opt.UsingPlaceholders)
	return &Command{
		mainName:                coty.NameCommand(opt.MainName),
		chapterDescriptionInfo:  coty.InfoChapterDESCRIPTION(opt.ChapterDescriptionInfo),
		usingPlaceholdersSorted: coty.ToSliceTypesSorted[coty.NamePlaceholder](opt.UsingPlaceholders),
		additionalNames:         coty.ToSliceTypesSorted[coty.NameCommand](opt.AdditionalNames),
	}
}

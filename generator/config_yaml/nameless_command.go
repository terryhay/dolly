package config_yaml

import (
	"errors"
	"fmt"
	"sort"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// NamelessCommand - pseudo mainName which doesn't have call name
type NamelessCommand struct {
	chapterDescriptionInfo coty.InfoChapterDESCRIPTION

	// optional
	usingPlaceholders []coty.NamePlaceholder
}

// GetChapterDescriptionInfo gets infoChapterDESCRIPTION field
func (nc *NamelessCommand) GetChapterDescriptionInfo() coty.InfoChapterDESCRIPTION {
	if nc == nil {
		return coty.InfoChapterDESCRIPTIONUndefined
	}
	return nc.chapterDescriptionInfo
}

// GetUsingPlaceholders gets placeholders field
func (nc *NamelessCommand) GetUsingPlaceholders() []coty.NamePlaceholder {
	if nc == nil {
		return nil
	}
	return nc.usingPlaceholders
}

var (
	// ErrNamelessCommandNoChapterDescriptionInfo - 'arg_parser.commandsSorted' element must have 'chapter_description_info' in source yaml file
	ErrNamelessCommandNoChapterDescriptionInfo = errors.New(`'arg_parser.nameless_command' element must have 'chapter_description_info' in source yaml file`)
)

// IsValid checks if NamelessCommand is valid
func (nc *NamelessCommand) IsValid() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("NamelessCommand.IsValid: %w", err)
		}
	}()

	if len(nc.GetChapterDescriptionInfo()) == 0 {
		return ErrNamelessCommandNoChapterDescriptionInfo
	}

	return nil
}

// NamelessCommandOpt - source for construct a pseudo mainName which doesn't have call name
type NamelessCommandOpt struct {
	ChapterDescriptionInfo string `yaml:"chapter_description_info"`

	// optional
	UsingPlaceholders []string `yaml:"using_placeholders"`
}

// NewNamelessCommand converts opt to NamelessCommand pointer
func NewNamelessCommand(opt *NamelessCommandOpt) *NamelessCommand {
	if opt == nil {
		return nil
	}

	sort.Strings(opt.UsingPlaceholders)
	return &NamelessCommand{
		chapterDescriptionInfo: coty.InfoChapterDESCRIPTION(opt.ChapterDescriptionInfo),
		usingPlaceholders:      coty.ToSliceTypesSorted[coty.NamePlaceholder](opt.UsingPlaceholders),
	}
}

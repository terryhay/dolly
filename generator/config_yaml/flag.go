package config_yaml

import (
	"errors"
	"fmt"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Flag - description of a mainName line flag
type Flag struct {
	mainName               coty.NameFlag
	chapterDescriptionInfo coty.InfoChapterDESCRIPTION

	// optional
	isOptional      bool
	additionalNames []coty.NameFlag
}

// GetMainName gets mainName filed
func (f *Flag) GetMainName() coty.NameFlag {
	if f == nil {
		return coty.NameFlagUndefined
	}
	return f.mainName
}

// GetDescriptionHelpInfo gets infoChapterDESCRIPTION field
func (f *Flag) GetDescriptionHelpInfo() coty.InfoChapterDESCRIPTION {
	if f == nil {
		return coty.InfoChapterDESCRIPTIONUndefined
	}
	return f.chapterDescriptionInfo
}

// GetIsOptional gets isOptional field
func (f *Flag) GetIsOptional() bool {
	if f == nil {
		return false
	}
	return f.isOptional
}

// GetAdditionalNames gets additionalNames filed
func (f *Flag) GetAdditionalNames() []coty.NameFlag {
	if f == nil {
		return nil
	}
	return f.additionalNames
}

var (
	// ErrFlagMainName - one of NameFlag fields is invalid
	ErrFlagMainName = errors.New(`Flag.IsValid: 'arg_parser.placeholders.flags.main_name' is invalid`)

	// ErrFlagNoChapterDescriptionInfo - no 'chapter_description_info' field in source yaml file
	ErrFlagNoChapterDescriptionInfo = errors.New(`Flag.IsValid: 'arg_parser.placeholders.flags.chapter_description_info' must be set`)

	// ErrFlagAdditionalName - one of elements of 'arg_parser.placeholders.flags.additional_names' is invalid
	ErrFlagAdditionalName = errors.New(`Flag.IsValid: one of elements of 'arg_parser.placeholders.flags.additional_names' is invalid`)
)

// IsValid checks if Flag is valid
func (f *Flag) IsValid() error {
	if err := f.GetMainName().IsValid(); err != nil {
		return errors.Join(ErrFlagMainName, err)
	}

	if len(f.chapterDescriptionInfo) == 0 {
		return fmt.Errorf(`flag main_name '%s': %w`, f.GetMainName(), ErrFlagNoChapterDescriptionInfo)
	}

	for _, name := range f.GetAdditionalNames() {
		if err := name.IsValid(); err != nil {
			return errors.Join(fmt.Errorf(`flag main_name '%s': %w`, f.GetMainName(), ErrFlagAdditionalName), err)
		}
	}

	return nil
}

// NamesSorted collects all flag call names in new slice and return it
func (f *Flag) NamesSorted() []coty.NameFlag {
	if f == nil {
		return nil
	}

	sorted := make([]coty.NameFlag, 0, 1+len(f.additionalNames))
	sorted = append(sorted, f.mainName)
	sorted = append(sorted, f.additionalNames...)

	coty.SortSlice(sorted)
	return sorted
}

// FlagOpt - source description of a mainName line flag
type FlagOpt struct {
	MainName               string `yaml:"main_name"`
	ChapterDescriptionInfo string `yaml:"chapter_description_info"`

	// optional
	IsOptional      bool     `yaml:"is_optional"`
	AdditionalNames []string `yaml:"additional_names"`
}

// NewFlag converts opt to Flag
func NewFlag(opt *FlagOpt) *Flag {
	if opt == nil {
		return nil
	}

	return &Flag{
		mainName:               coty.NameFlag(opt.MainName),
		chapterDescriptionInfo: coty.InfoChapterDESCRIPTION(opt.ChapterDescriptionInfo),
		isOptional:             opt.IsOptional,

		additionalNames: coty.ToSliceTypesSorted[coty.NameFlag](opt.AdditionalNames),
	}
}

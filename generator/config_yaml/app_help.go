package config_yaml

import (
	"errors"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// AppHelp - data for application help info page
type AppHelp struct {
	appName                coty.NameApp
	chapterNameInfo        coty.InfoChapterNAME
	chapterDescriptionInfo []coty.InfoChapterDESCRIPTION
}

// GetApplicationName gets appName field
func (ah *AppHelp) GetApplicationName() coty.NameApp {
	if ah == nil {
		return coty.AppNameUndefined
	}
	return ah.appName
}

// GetHelpInfoChapterName gets chapterNameInfo field
func (ah *AppHelp) GetHelpInfoChapterName() coty.InfoChapterNAME {
	if ah == nil {
		return coty.InfoChapterNAMEUndefined
	}
	return ah.chapterNameInfo
}

// GetHelpInfoChapterDescription gets infoChapterDESCRIPTION field
//
//goland:noinspection ALL
func (ah *AppHelp) GetHelpInfoChapterDescription() []coty.InfoChapterDESCRIPTION {
	if ah == nil {
		return nil
	}
	return ah.chapterDescriptionInfo
}

// AppHelpOpt - source for construct AppHelp object
type AppHelpOpt struct {
	AppName                string   `yaml:"app_name"`
	ChapterNameInfo        string   `yaml:"chapter_name_info"`
	ChapterDescriptionInfo []string `yaml:"chapter_description_info"`
}

//goland:noinspection ALL
var (
	// ErrAppHelpNilPointer -  no required field 'arg_parser.app_help' in source yaml file
	ErrAppHelpNilPointer = errors.New(`AppHelp.IsValid: no required field 'arg_parser.app_help' in source yaml file`)

	// ErrAppHelpNoAppName - no required field 'arg_parser.app_help.app_name' in source yaml file
	ErrAppHelpNoAppName = errors.New(`AppHelp.IsValid: no required field 'arg_parser.app_help.app_name' in source yaml file`)

	// ErrAppHelpNoChapterNameInfo - no required field 'arg_parser.app_help.chapter_name_info' in source yaml file
	ErrAppHelpNoChapterNameInfo = errors.New(`AppHelp.IsValid: no required field 'arg_parser.app_help.chapter_name_info' in source yaml file`)
)

// IsValid checks if AppHelp is valid
//
//goland:noinspection ALL
func (ah *AppHelp) IsValid() error {
	if ah == nil {
		return ErrAppHelpNilPointer
	}

	if len(ah.appName) == 0 {
		return ErrAppHelpNoAppName
	}

	if len(ah.chapterNameInfo) == 0 {
		return ErrAppHelpNoChapterNameInfo
	}

	return nil
}

// NewAppHelp converts opt to AppHelp object
func NewAppHelp(opt *AppHelpOpt) *AppHelp {
	if opt == nil {
		return nil
	}

	return &AppHelp{
		appName:         coty.NameApp(opt.AppName),
		chapterNameInfo: coty.InfoChapterNAME(opt.ChapterNameInfo),

		chapterDescriptionInfo: coty.ToSliceTypesSorted[coty.InfoChapterDESCRIPTION](opt.ChapterDescriptionInfo),
	}
}

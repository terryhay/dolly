package arg_parser_config

import (
	"unsafe"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Application contains specification of application for help output
type Application struct {
	appName         coty.NameApp
	infoChapterNAME coty.InfoChapterNAME
}

// ApplicationOpt contains source data for cast to Application
type ApplicationOpt struct {
	AppName         coty.NameApp
	InfoChapterNAME coty.InfoChapterNAME
}

// MakeApplication constructs Application object
func MakeApplication(opt ApplicationOpt) Application {
	return *(*Application)(unsafe.Pointer(&opt))
}

// GetAppName gets appName field
func (i Application) GetAppName() coty.NameApp {
	return i.appName
}

// GetNameHelpInfo gets infoChapterNAME field
func (i Application) GetNameHelpInfo() coty.InfoChapterNAME {
	return i.infoChapterNAME
}

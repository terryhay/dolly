package arg_parser_config

import "unsafe"

// ApplicationDescription contains specification of application for help output
type ApplicationDescription struct {
	appName             string
	nameHelpInfo        string
	descriptionHelpInfo []string
}

// ApplicationDescriptionSrc contains source data for cast to ApplicationDescription
type ApplicationDescriptionSrc struct {
	AppName             string
	NameHelpInfo        string
	DescriptionHelpInfo []string
}

// ToConst converts src to ApplicationDescription object
func (src ApplicationDescriptionSrc) ToConst() ApplicationDescription {
	return *(*ApplicationDescription)(unsafe.Pointer(&src))
}

// GetAppName - appName field getter
func (i ApplicationDescription) GetAppName() string {
	return i.appName
}

// GetNameHelpInfo - nameHelpInfo field getter
func (i ApplicationDescription) GetNameHelpInfo() string {
	return i.nameHelpInfo
}

// GetDescriptionHelpInfo - nameHelpInfo field getter
func (i ApplicationDescription) GetDescriptionHelpInfo() []string {
	return i.descriptionHelpInfo
}

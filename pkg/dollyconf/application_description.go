package dollyconf

// ApplicationDescription contains specification of application for help output
type ApplicationDescription struct {
	AppName             string
	NameHelpInfo        string
	DescriptionHelpInfo []string
}

// GetAppName - AppName field getter
func (i ApplicationDescription) GetAppName() string {
	return i.AppName
}

// GetNameHelpInfo - NameHelpInfo field getter
func (i ApplicationDescription) GetNameHelpInfo() string {
	return i.NameHelpInfo
}

// GetDescriptionHelpInfo - NameHelpInfo field getter
func (i ApplicationDescription) GetDescriptionHelpInfo() []string {
	return i.DescriptionHelpInfo
}

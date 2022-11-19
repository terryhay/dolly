package generate

import (
	"fmt"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	"strings"
)

// sectionAppDescription - string with application description
type sectionAppDescription string

// genAppDescriptionSection creates a paste section with application description
func genAppDescriptionSection(appDescription *confYML.AppHelpDescription) sectionAppDescription {
	descriptionHelpInfo := "nil"
	if len(appDescription.GetDescriptionHelpInfo()) > 0 {
		builder := strings.Builder{}
		builder.WriteString("[]string{")
		for _, paragraph := range appDescription.GetDescriptionHelpInfo() {
			builder.WriteString(fmt.Sprintf("\n\t\t\t\t\"%s\",", paragraph))
		}
		builder.WriteString("\n\t\t\t}")
		descriptionHelpInfo = builder.String()
	}

	return sectionAppDescription(fmt.Sprintf(`
		AppDescription: apConf.ApplicationDescriptionSrc{
			AppName: "%s",
			NameHelpInfo: "%s",
			DescriptionHelpInfo: %s,
		}.ToConst()`,
		appDescription.GetApplicationName(),
		appDescription.GetNameHelpInfo(),
		descriptionHelpInfo))
}

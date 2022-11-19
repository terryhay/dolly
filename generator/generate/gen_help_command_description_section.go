package generate

import (
	"fmt"
	confYML "github.com/terryhay/dolly/generator/config_yaml"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"sort"
	"strings"
)

const (
	helpCommandDescriptionNilPattern = `
		HelpCommandDescription: nil`
	helpCommandDescriptionPrefixPart = `
		HelpCommandDescription: apConf.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[apConf.Command]bool{`
	helpCommandDescriptionCommandMapElementPart = `
				%s: true,`
	helpCommandDescriptionPostfixPart = `
			},
		)`
)

// sectionHelpCommandDescription - string with help command description section
type sectionHelpCommandDescription string

// genHelpCommandDescriptionSection creates a paste section with help command description
func genHelpCommandDescriptionSection(
	helpCommandDescription *confYML.HelpCommandDescription,
	commandsIDTemplateData map[string]*creator.IDTemplateData,
) sectionHelpCommandDescription {

	if helpCommandDescription == nil {
		return helpCommandDescriptionNilPattern
	}

	sortedCommandNameIDs := make([]string, 0, len(helpCommandDescription.GetAdditionalCommands())+1)
	sortedCommandNameIDs = append(sortedCommandNameIDs,
		commandsIDTemplateData[helpCommandDescription.GetCommand()].GetNameID())

	for i := 0; i < len(helpCommandDescription.GetAdditionalCommands()); i++ {
		sortedCommandNameIDs = append(sortedCommandNameIDs,
			commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[i]].GetNameID())
	}

	sort.Strings(sortedCommandNameIDs)

	builder := strings.Builder{}
	builder.WriteString(helpCommandDescriptionPrefixPart)

	builder.WriteString(fmt.Sprintf(helpCommandDescriptionCommandMapElementPart, sortedCommandNameIDs[0]))
	for i := 1; i < len(sortedCommandNameIDs); i++ {
		builder.WriteString(fmt.Sprintf(helpCommandDescriptionCommandMapElementPart, sortedCommandNameIDs[i]))
	}

	builder.WriteString(helpCommandDescriptionPostfixPart)
	return sectionHelpCommandDescription(builder.String())
}

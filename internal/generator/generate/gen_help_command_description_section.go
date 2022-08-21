package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/config_yaml"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"sort"
	"strings"
)

const (
	helpCommandDescriptionNilPattern = `
		// helpCommandDescription
		nil`
	helpCommandDescriptionPrefixPart = `
		// helpCommandDescription
		dollyconf.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[dollyconf.Command]bool{`
	helpCommandDescriptionCommandMapElementPart = `
				%s: true,`
	helpCommandDescriptionPostfixPart = `
			},
		)`
)

// HelpCommandDescriptionSection - string with help command description section
type HelpCommandDescriptionSection string

// GenHelpCommandDescriptionSection creates a paste section with help command description
func GenHelpCommandDescriptionSection(
	helpCommandDescription *config_yaml.HelpCommandDescription,
	commandsIDTemplateData map[string]*id_template_data_creator.IDTemplateData,
) HelpCommandDescriptionSection {

	if helpCommandDescription == nil {
		return helpCommandDescriptionNilPattern
	}

	var i int

	sortedCommandNameIDs := make([]string, 0, len(helpCommandDescription.GetAdditionalCommands())+1)
	sortedCommandNameIDs = append(sortedCommandNameIDs,
		commandsIDTemplateData[helpCommandDescription.GetCommand()].GetNameID())

	for i = 0; i < len(helpCommandDescription.GetAdditionalCommands()); i++ {
		sortedCommandNameIDs = append(sortedCommandNameIDs,
			commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[i]].GetNameID())
	}

	sort.Strings(sortedCommandNameIDs)

	builder := strings.Builder{}
	builder.WriteString(helpCommandDescriptionPrefixPart)

	builder.WriteString(fmt.Sprintf(helpCommandDescriptionCommandMapElementPart, sortedCommandNameIDs[0]))
	for i = 1; i < len(sortedCommandNameIDs); i++ {
		builder.WriteString(fmt.Sprintf(helpCommandDescriptionCommandMapElementPart, sortedCommandNameIDs[i]))
	}

	builder.WriteString(helpCommandDescriptionPostfixPart)
	return HelpCommandDescriptionSection(builder.String())
}

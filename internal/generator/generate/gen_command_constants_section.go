package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	commandConstantsPrefixPart = `
const (`
	commandConstantsFirstLinePart = `
	// %s - %s
	%s dollyconf.Command = "%s"`
	commandConstantsLinePart = `
	// %s - %s
	%s = "%s"`
	commandConstantsPostfixPart = `
)`
)

// CommandListSection - string with command constant definitions list
type CommandListSection string

// GenCommandListSection creates a paste section with commands
func GenCommandListSection(
	commandsTemplateData map[string]*id_template_data_creator.IDTemplateData,
) CommandListSection {

	if len(commandsTemplateData) == 0 {
		return ""
	}

	sortedCommandsTemplateData := sortByNameID(commandsTemplateData)

	builder := strings.Builder{}
	templateData := sortedCommandsTemplateData[0]

	builder.WriteString(commandConstantsPrefixPart)

	builder.WriteString(fmt.Sprintf(commandConstantsFirstLinePart,
		templateData.GetNameID(), templateData.GetComment(), templateData.GetNameID(), templateData.GetCallName()))

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		builder.WriteString(fmt.Sprintf(commandConstantsLinePart,
			templateData.GetNameID(), templateData.GetComment(), templateData.GetNameID(), templateData.GetCallName()))
	}

	builder.WriteString(commandConstantsPostfixPart)

	return CommandListSection(builder.String())
}

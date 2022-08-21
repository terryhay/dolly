package generate

import (
	"fmt"
	"github.com/terryhay/dolly/internal/generator/id_template_data_creator"
	"strings"
)

const (
	commandIDConstantsPrefixPart = `
const (`
	commandIDConstantsFirstLinePart = `
	// %s - %s
	%s dollyconf.CommandID = iota + 1`
	commandIDConstantsLinePart = `
	// %s - %s
	%s`
	commandIDConstantsPostfixPart = `
)`
)

// CommandIDListSection - string with command id constant definitions list
type CommandIDListSection string

// GenCommandIDListSection creates a paste section with command ids
func GenCommandIDListSection(
	commandsTemplateData map[string]*id_template_data_creator.IDTemplateData,
	nullCommandIDTemplateData *id_template_data_creator.IDTemplateData,
) CommandIDListSection {

	sortedCommandsTemplateData := sortCommandsTemplateData(commandsTemplateData, nullCommandIDTemplateData)
	if len(sortedCommandsTemplateData) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(commandIDConstantsPrefixPart)

	templateData := sortedCommandsTemplateData[0]
	builder.WriteString(fmt.Sprintf(commandIDConstantsFirstLinePart,
		templateData.GetID(), templateData.GetComment(), templateData.GetID()))

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		builder.WriteString(fmt.Sprintf(commandIDConstantsLinePart,
			templateData.GetID(), templateData.GetComment(), templateData.GetID()))
	}
	builder.WriteString(commandIDConstantsPostfixPart)

	return CommandIDListSection(builder.String())
}

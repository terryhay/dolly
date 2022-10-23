package generate

import (
	"fmt"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	commandConstantsPrefixPart = `
const (`
	commandConstantsFirstLinePart = `
	// %s - %s
	%s apConf.Command = "%s"`
	commandConstantsLinePart = `
	// %s - %s
	%s = "%s"`
	commandConstantsPostfixPart = `
)`
)

// sectionCommandList - string with command constant definitions list
type sectionCommandList string

// genCommandListSection creates a paste section with commands
func genCommandListSection(commandsTemplateData map[string]*creator.IDTemplateData) sectionCommandList {

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

	return sectionCommandList(builder.String())
}

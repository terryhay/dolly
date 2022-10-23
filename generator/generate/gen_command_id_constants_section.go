package generate

import (
	"fmt"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	commandIDConstantsPrefixPart = `
const (`
	commandIDConstantsFirstLinePart = `
	// %s - %s
	%s apConf.CommandID = iota + 1`
	commandIDConstantsLinePart = `
	// %s - %s
	%s`
	commandIDConstantsPostfixPart = `
)`
)

// sectionCommandIDList - string with command id constant definitions list
type sectionCommandIDList string

// genCommandIDListSection creates a paste section with command ids
func genCommandIDListSection(
	commandsTemplateData map[string]*creator.IDTemplateData,
	nullCommandIDTemplateData *creator.IDTemplateData,
) sectionCommandIDList {

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

	return sectionCommandIDList(builder.String())
}

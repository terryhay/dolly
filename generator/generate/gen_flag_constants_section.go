package generate

import (
	"fmt"
	creator "github.com/terryhay/dolly/generator/id_template_data_creator"
	"strings"
)

const (
	flagConstantsPrefixPart = `
const (`
	flagConstantsFirstLinePart = `
	// %s - %s
	%s apConf.Flag = "%s"`
	flagConstantsLinePart = `
	// %s - %s
	%s = "%s"`
	flagConstantsPostfixPart = `
)`
)

// sectionFlagStringIDList - string with flag constant definitions list
type sectionFlagStringIDList string

// genFlagIDConstantsSection - creates a paste section flag constants
func genFlagIDConstantsSection(
	flagsTemplateData map[string]*creator.IDTemplateData,
) sectionFlagStringIDList {

	if len(flagsTemplateData) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(flagConstantsPrefixPart)

	sortedFlagsTemplateData := sortByNameID(flagsTemplateData)

	templateData := sortedFlagsTemplateData[0]
	builder.WriteString(fmt.Sprintf(flagConstantsFirstLinePart,
		templateData.GetNameID(),
		templateData.GetComment(),
		templateData.GetNameID(),
		templateData.GetCallName()))

	for i := 1; i < len(sortedFlagsTemplateData); i++ {
		templateData = sortedFlagsTemplateData[i]
		builder.WriteString(fmt.Sprintf(flagConstantsLinePart,
			templateData.GetNameID(),
			templateData.GetComment(),
			templateData.GetNameID(),
			templateData.GetCallName()))
	}

	builder.WriteString(flagConstantsPostfixPart)
	return sectionFlagStringIDList(builder.String())
}
